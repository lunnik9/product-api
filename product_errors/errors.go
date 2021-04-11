package product_errors

import "fmt"

type ProductError struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
}

func (e ProductError) Error() string {
	return fmt.Sprintf("error occured: %v, code: %d", e.Message, e.StatusCode)
}

func New(statusCode int, message string) ProductError {
	return ProductError{statusCode, message}
}
