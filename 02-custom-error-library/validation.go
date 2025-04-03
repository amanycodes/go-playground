package validation

import "fmt"

const (
	ErrCodeValidation    = "VAL_ERR"
	ErrCodeAuthorization = "AUTH_ERR"
	ErrCodeNotFound      = "NF_ERR"
)

type ValidationError struct {
	Code    string
	Message string
}

func (v *ValidationError) Error() string {
	return fmt.Sprintf("[%s] %s", v.Code, v.Message)
}

type AuthorizationError struct {
	Code    string
	Message string
}

func (a *AuthorizationError) Error() string {
	return fmt.Sprintf("[%s] %s", a.Code, a.Message)
}

type NotFoundError struct {
	Code    string
	Message string
}

func (n *NotFoundError) Error() string {
	return fmt.Sprintf("[%s] %s", n.Code, n.Message)
}

func ValidateInput(input string) error {
	if input == "" {
		return &ValidationError{
			Code:    ErrCodeValidation,
			Message: "Input cannot be empty",
		}
	}
	return nil
}

func FindResource(id int) (string, error) {
	if id != 42 {
		return "", &NotFoundError{
			Code:    ErrCodeNotFound,
			Message: "Could not find the resource",
		}
	}
	return "Resource Data", nil
}

func CheckAuthorization(token string) error {
	if token != "abc" {
		return &AuthorizationError{
			Code:    ErrCodeAuthorization,
			Message: "Invalid User",
		}
	}
	return nil
}

func Process(id int, input string, token string) error {
	if err := ValidateInput(input); err != nil {
		return fmt.Errorf("processing failed: %w", err)
	}

	_, err := FindResource(id)
	if err != nil {
		return fmt.Errorf("processing failed: %w", err)
	}

	if err = CheckAuthorization(token); err != nil {
		return fmt.Errorf("processing failed: %w", err)
	}
	return nil
}
