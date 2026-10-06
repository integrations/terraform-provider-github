package github

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
)

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
