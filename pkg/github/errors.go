package github

import "fmt"

type MissingRequiredParameterError struct {
	Parameter string
}

// Error implements the error interface for MissingRequiredParameterError.
func (e MissingRequiredParameterError) Error() string {
	return fmt.Sprintf("missing required parameter: %s", e.Parameter)
}

type InvalidParameterTypeError struct {
	Parameter string
	Expected  string
	Actual    string
}

// Error implements the error interface for InvalidParameterTypeError.
func (e InvalidParameterTypeError) Error() string {
	return fmt.Sprintf("parameter %s is not of type %s, is %s", e.Parameter, e.Expected, e.Actual)
}
