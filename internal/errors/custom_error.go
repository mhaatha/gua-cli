package errors

import "errors"

var (
	ErrCannotLoadEnv              = errors.New("cannot load .env file")
	ErrCallingNewRequest          = errors.New("failed when calling http.NewRequest")
	ErrNetworkConnectivityProblem = errors.New("network connectivity problem")
	ErrUsernameNotFound           = errors.New("username not found")
	ErrInvalidGithubPAT           = errors.New("provide your *valid* Github Personal Access Token in .env file")
	ErrReadResponseBody           = errors.New("failed to read the response body")
	ErrUnmarshalResponseBody      = errors.New("failed to unmarshal the response body")
)

type ErrMsg error

type AppError struct {
	Err error
}

func (e AppError) Error() string {
	return e.Err.Error()
}
