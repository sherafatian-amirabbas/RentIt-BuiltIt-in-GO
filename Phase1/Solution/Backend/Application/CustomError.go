package application

// CustomError will hold the information for error handling
type CustomError struct {
	Code    string
	Message string
}

// ToString returns CustomError as a string
func (cErr CustomError) ToString() string {

	return cErr.Code + ": " + cErr.Message
}
