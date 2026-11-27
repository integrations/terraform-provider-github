package github

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
)

var _ statecheck.StateCheck = expectFunc{}

// expectFunc is a statecheck.StateCheck implementation that delegates to a closure. It's for assertions that
// can't be expressed against state alone, e.g. unmanaged attributes that must be verified via the API.
type expectFunc struct {
	fn func(ctx context.Context, req statecheck.CheckStateRequest) error
}

// CheckState runs the wrapped closure and reports any error it returns.
func (c expectFunc) CheckState(ctx context.Context, req statecheck.CheckStateRequest, resp *statecheck.CheckStateResponse) {
	resp.Error = c.fn(ctx, req)
}

// ExpectFunc returns a statecheck.StateCheck that fails if fn returns an error.
func ExpectFunc(fn func(ctx context.Context, req statecheck.CheckStateRequest) error) statecheck.StateCheck {
	return expectFunc{fn: fn}
}

var _ knownvalue.Check = setAbsent{}

// setAbsent is a knownvalue.Check implementation that asserts that a specific value is not present in a set.
type setAbsent struct {
	value []knownvalue.Check
}

// CheckValue determines whether the passed value of type []any, and does not contain the expected value.
func (v setAbsent) CheckValue(other any) error {
	otherVals, ok := other.([]any)
	if !ok {
		return fmt.Errorf("expected []any value for SetAbsent check, got: %T", other)
	}

	for _, otherVal := range otherVals {
		match := true
		for _, check := range v.value {
			if err := check.CheckValue(otherVal); err != nil {
				match = false
				break
			}
		}

		if match {
			return fmt.Errorf("found unexpected value %s for SetAbsent check", v.String())
		}
	}

	return nil
}

// String returns the string representation of the value.
func (v setAbsent) String() string {
	var setVals []string

	for _, val := range v.value {
		setVals = append(setVals, val.String())
	}

	return fmt.Sprintf("%s", setVals)
}

// SetAbsent returns a knownvalue.Check for asserting the value defined by the []knownvalue.Check slice is not present in the set.
func SetAbsent(value []knownvalue.Check) setAbsent {
	return setAbsent{
		value: value,
	}
}
