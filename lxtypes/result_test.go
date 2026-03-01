package lxtypes_test

import (
	"errors"
	"strconv"
	"testing"

	"github.com/nthanhhai2909/lx/lxtypes"
)

func TestResultSuccess(t *testing.T) {
	result := lxtypes.Success(42)

	if !result.IsSuccess() {
		t.Error("Expected Success to return true for IsSuccess()")
	}
	if result.IsFailure() {
		t.Error("Expected Success to return false for IsFailure()")
	}
	if got := result.Value(); got != 42 {
		t.Errorf("Value() = %v, want 42", got)
	}
}

func TestResultFailure(t *testing.T) {
	err := errors.New("test error")
	result := lxtypes.Failure[int](err)

	if result.IsSuccess() {
		t.Error("Expected Failure to return false for IsSuccess()")
	}
	if !result.IsFailure() {
		t.Error("Expected Failure to return true for IsFailure()")
	}
	if got := result.Error(); got.Error() != err.Error() {
		t.Errorf("Error() = %v, want %v", got, err)
	}
}

func TestResultValuePanics(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("Expected Value on Failure to panic")
		}
	}()
	result := lxtypes.Failure[int](errors.New("error"))
	result.Value()
}

func TestResultErrorPanics(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("Expected Error on Success to panic")
		}
	}()
	result := lxtypes.Success(42)
	result.Error()
}

func TestResultValueOr(t *testing.T) {
	success := lxtypes.Success(42)
	failure := lxtypes.Failure[int](errors.New("error"))

	if got := success.ValueOr(0); got != 42 {
		t.Errorf("Success.ValueOr(0) = %v, want 42", got)
	}
	if got := failure.ValueOr(99); got != 99 {
		t.Errorf("Failure.ValueOr(99) = %v, want 99", got)
	}
}

func TestResultValueOrElse(t *testing.T) {
	success := lxtypes.Success(42)
	failure := lxtypes.Failure[int](errors.New("error"))

	if got := success.ValueOrElse(func(e error) int { return 0 }); got != 42 {
		t.Errorf("Success.ValueOrElse(...) = %v, want 42", got)
	}
	if got := failure.ValueOrElse(func(e error) int { return 99 }); got != 99 {
		t.Errorf("Failure.ValueOrElse(...) = %v, want 99", got)
	}
}

func TestResultMap(t *testing.T) {
	double := func(n int) int { return n * 2 }

	success := lxtypes.Success(21)
	mapped := lxtypes.ResultMap(success, double)

	if !mapped.IsSuccess() {
		t.Error("Expected mapped Success to be Success")
	}
	if got := mapped.Value(); got != 42 {
		t.Errorf("Mapped value = %v, want 42", got)
	}

	failure := lxtypes.Failure[int](errors.New("error"))
	mappedFailure := lxtypes.ResultMap(failure, double)

	if !mappedFailure.IsFailure() {
		t.Error("Expected mapped Failure to be Failure")
	}
}

func TestResultAndThen(t *testing.T) {
	safeDivide := func(a, b int) lxtypes.Result[int] {
		if b == 0 {
			return lxtypes.Failure[int](errors.New("division by zero"))
		}
		return lxtypes.Success(a / b)
	}

	success := lxtypes.Success(10)
	result := lxtypes.ResultAndThen(success, func(n int) lxtypes.Result[int] {
		return safeDivide(100, n)
	})

	if !result.IsSuccess() {
		t.Error("Expected Success after AndThen")
	}
	if got := result.Value(); got != 10 {
		t.Errorf("Result = %v, want 10", got)
	}

	// Test chaining that fails
	successZero := lxtypes.Success(0)
	resultFailure := lxtypes.ResultAndThen(successZero, func(n int) lxtypes.Result[int] {
		return safeDivide(100, n)
	})

	if !resultFailure.IsFailure() {
		t.Error("Expected Failure after AndThen with zero")
	}

	// Test Failure propagation
	failure := lxtypes.Failure[int](errors.New("initial error"))
	resultProp := lxtypes.ResultAndThen(failure, func(n int) lxtypes.Result[int] {
		return safeDivide(100, n)
	})

	if !resultProp.IsFailure() {
		t.Error("Expected Failure to propagate through AndThen")
	}
	if got := resultProp.Error().Error(); got != "initial error" {
		t.Errorf("Error = %v, want 'initial error'", got)
	}
}

func TestResultOrElse(t *testing.T) {
	success := lxtypes.Success(42)
	failure := lxtypes.Failure[int](errors.New("error"))

	fallback := func(e error) lxtypes.Result[int] {
		return lxtypes.Success(99)
	}

	// Success should ignore OrElse
	resultSuccess := success.OrElse(fallback)
	if !resultSuccess.IsSuccess() {
		t.Error("Expected Success.OrElse to return Success")
	}
	if got := resultSuccess.Value(); got != 42 {
		t.Errorf("Success.OrElse value = %v, want 42", got)
	}

	// Failure should call fallback
	resultFailure := failure.OrElse(fallback)
	if !resultFailure.IsSuccess() {
		t.Error("Expected Failure.OrElse to return fallback Success")
	}
	if got := resultFailure.Value(); got != 99 {
		t.Errorf("Failure.OrElse value = %v, want 99", got)
	}
}

func TestResultRecover(t *testing.T) {
	failure := lxtypes.Failure[int](errors.New("error"))

	recovered := lxtypes.ResultRecover(failure, func(e error) lxtypes.Result[int] {
		return lxtypes.Success(99)
	})

	if !recovered.IsSuccess() {
		t.Error("Expected Recover to return Success")
	}
	if got := recovered.Value(); got != 99 {
		t.Errorf("Recovered value = %v, want 99", got)
	}
}

func TestFromError(t *testing.T) {
	// Success case
	value, err := strconv.Atoi("42")
	result1 := lxtypes.FromError(value, err)

	if !result1.IsSuccess() {
		t.Error("Expected FromError with nil error to be Success")
	}
	if got := result1.Value(); got != 42 {
		t.Errorf("Value = %v, want 42", got)
	}

	// Failure case
	value2, err2 := strconv.Atoi("invalid")
	result2 := lxtypes.FromError(value2, err2)

	if !result2.IsFailure() {
		t.Error("Expected FromError with error to be Failure")
	}
}

func TestResultChaining(t *testing.T) {
	// Complex chaining scenario
	parseNum := func(s string) lxtypes.Result[int] {
		n, err := strconv.Atoi(s)
		return lxtypes.FromError(n, err)
	}

	double := func(n int) int { return n * 2 }

	validate := func(n int) lxtypes.Result[int] {
		if n > 15 {
			return lxtypes.Success(n)
		}
		return lxtypes.Failure[int](errors.New("too small"))
	}

	result := parseNum("10")
	doubled := lxtypes.ResultMap(result, double)
	validated := lxtypes.ResultAndThen(doubled, validate)
	final := lxtypes.ResultMap(validated, func(n int) int { return n + 5 })

	if !final.IsSuccess() {
		t.Error("Expected chained result to be Success")
	}
	if got := final.Value(); got != 25 {
		t.Errorf("Chained result = %v, want 25", got)
	}

	// Chaining that fails
	result2 := parseNum("invalid")
	if !result2.IsFailure() {
		t.Error("Expected parse error")
	}
	if got := result2.ValueOr(99); got != 99 {
		t.Errorf("Failed parse result = %v, want 99", got)
	}
}
