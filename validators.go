// Author: Alexandre Wilhelm
// See LICENSE file for full LICENSE
// Copyright 2016 Aporeto.

package elemental

import (
	"fmt"
	"regexp"
)

const (
	maximumIntFailFormat             = `Data '%d' of attribute '%s' should be less than %d`
	maximumIntExclusiveFailFormat    = `Data '%d' of attribute '%s' should be less or equal than %d`
	minimumIntFailFormat             = `Data '%d' of attribute '%s' should be greater than %d`
	minimumIntExclusiveFailFormat    = `Data '%d' of attribute '%s' should be greater or equal than %d`
	maximumFloatFailFormat           = `Data '%g' of attribute '%s' should be less than %g`
	maximumFloatExclusiveFailFormat  = `Data '%g' of attribute '%s' should be less or equal than %g`
	minimumFloatFailFormat           = `Data '%g' of attribute '%s' should be greater than %g`
	minimumFloatExclusiveFailFormat  = `Data '%g' of attribute '%s' should be greater or equal than %g`
	maximumLengthFailFormat          = `Data '%s' of attribute '%s' should be less than %d chars long`
	maximumLengthExclusiveFailFormat = `Data '%s' of attribute '%s' should be less or equal than %d chars long`
	minimumLengthFailFormat          = `Data '%s' of attribute '%s' should be greater than %d chars long`
	minimumLengthExclusiveFailFormat = `Data '%s' of attribute '%s' should be greater or equal than %d chars long`
	patternFailFormat                = `Data '%s' of attribute '%s' should match '%s'`
	requiredFloatFailFormat          = `Attribute '%s' is required`
	requiredStringFailFormat         = `Attribute '%s' is required`
	requiredIntFailFormat            = `Attribute '%s' is required`
	floatInListFormat                = `Data '%g' of attribute '%s' is not in list '%g'`
	stringInListFormat               = `Data '%s' of attribute '%s' is not in list '%s'`
	intInListFormat                  = `Data '%d' of attribute '%s' is not in list '%d'`
)

// Validatable is the interface for objects that can be validated.
type Validatable interface {
	Validate() Errors
}

// ValidateStringInList validates if the string is in the list
func ValidateStringInList(attribute string, value string, enums []string) *Error {

	for _, v := range enums {
		if v == value {
			return nil
		}
	}

	return NewError("Validation Error", fmt.Sprintf(stringInListFormat, value, attribute, enums), attribute, 409)
}

// ValidateFloatInList validates if the string is in the list
func ValidateFloatInList(attribute string, value float64, enums []float64) *Error {

	for _, v := range enums {
		if v == value {
			return nil
		}
	}

	return NewError("Validation Error", fmt.Sprintf(floatInListFormat, value, attribute, enums), attribute, 409)
}

// ValidateIntInList validates if the string is in the list
func ValidateIntInList(attribute string, value int, enums []int) *Error {

	for _, v := range enums {
		if v == value {
			return nil
		}
	}

	return NewError("Validation Error", fmt.Sprintf(intInListFormat, value, attribute, enums), attribute, 409)
}

// ValidateRequiredInt validates is the int is set to 0
func ValidateRequiredInt(attribute string, value int) *Error {

	if value == 0 {
		return NewError("Validation Error", fmt.Sprintf(requiredIntFailFormat, attribute), attribute, 409)
	}

	return nil
}

// ValidateRequiredFloat validates is the int is set to 0
func ValidateRequiredFloat(attribute string, value float64) *Error {

	if value == 0.0 {
		return NewError("Validation Error", fmt.Sprintf(requiredFloatFailFormat, attribute), attribute, 409)
	}

	return nil
}

// ValidateMaximumFloat validates a float against a maximum value
func ValidateMaximumFloat(attribute string, value float64, max float64, exclusive bool) *Error {

	if !exclusive && value > max {
		return NewError("Validation Error", fmt.Sprintf(maximumFloatFailFormat, value, attribute, max), attribute, 409)
	} else if exclusive && value >= max {
		return NewError("Validation Error", fmt.Sprintf(maximumFloatExclusiveFailFormat, value, attribute, max), attribute, 409)
	}

	return nil
}

// ValidateMinimumFloat validates a float against a maximum value
func ValidateMinimumFloat(attribute string, value float64, min float64, exclusive bool) *Error {

	if !exclusive && value < min {
		return NewError("Validation Error", fmt.Sprintf(minimumFloatFailFormat, value, attribute, min), attribute, 409)
	} else if exclusive && value <= min {
		return NewError("Validation Error", fmt.Sprintf(minimumFloatExclusiveFailFormat, value, attribute, min), attribute, 409)
	}

	return nil
}

// ValidateMaximumInt validates a integer against a maximum value
func ValidateMaximumInt(attribute string, value int, max int, exclusive bool) *Error {

	if !exclusive && value > max {
		return NewError("Validation Error", fmt.Sprintf(maximumIntFailFormat, value, attribute, max), attribute, 409)
	} else if exclusive && value >= max {
		return NewError("Validation Error", fmt.Sprintf(maximumIntExclusiveFailFormat, value, attribute, max), attribute, 409)
	}

	return nil
}

// ValidateMinimumInt validates a integer against a maximum value
func ValidateMinimumInt(attribute string, value int, min int, exclusive bool) *Error {

	if !exclusive && value < min {
		return NewError("Validation Error", fmt.Sprintf(minimumIntFailFormat, value, attribute, min), attribute, 409)
	} else if exclusive && value <= min {
		return NewError("Validation Error", fmt.Sprintf(minimumIntExclusiveFailFormat, value, attribute, min), attribute, 409)
	}

	return nil
}

// ValidateRequiredString validates is the string is empty or not
func ValidateRequiredString(attribute string, value string) *Error {

	if value == "" {
		return NewError("Validation Error", fmt.Sprintf(requiredStringFailFormat, attribute), attribute, 409)
	}

	return nil
}

// ValidatePattern validates a string against a regular expression
func ValidatePattern(attribute string, value string, pattern string) *Error {

	re := regexp.MustCompile(pattern)

	if !re.MatchString(value) {
		return NewError("Validation Error", fmt.Sprintf(patternFailFormat, value, attribute, pattern), attribute, 409)
	}

	return nil
}

// ValidateMinimumLength validates the minimum length of a string
func ValidateMinimumLength(attribute string, value string, min int, exclusive bool) *Error {

	length := len([]rune(value))

	if !exclusive && length < min {
		return NewError("Validation Error", fmt.Sprintf(minimumLengthFailFormat, value, attribute, min), attribute, 409)
	} else if exclusive && length <= min {
		return NewError("Validation Error", fmt.Sprintf(minimumLengthExclusiveFailFormat, value, attribute, min), attribute, 409)
	}

	return nil
}

// ValidateMaximumLength validates the maximum length of a string
func ValidateMaximumLength(attribute string, value string, max int, exclusive bool) *Error {

	length := len([]rune(value))

	if !exclusive && length > max {
		return NewError("Validation Error", fmt.Sprintf(maximumLengthFailFormat, value, attribute, max), attribute, 409)
	} else if exclusive && length >= max {
		return NewError("Validation Error", fmt.Sprintf(maximumLengthExclusiveFailFormat, value, attribute, max), attribute, 409)
	}

	return nil
}
