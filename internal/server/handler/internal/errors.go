package internal

import (
	"errors"
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

type ValidationError struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func NewValidationError(message string, errs ...error) *ValidationError {
	return &ValidationError{
		Status:  "error",
		Message: validationErrorMessage(message, errs...),
	}
}

func (e *ValidationError) Error() string {
	return e.Message
}

func validationErrorMessage(message string, errs ...error) string {
	if len(errs) == 0 || errs[0] == nil {
		return message
	}

	err := errs[0]
	var validationErrs validator.ValidationErrors
	if !errors.As(err, &validationErrs) {
		return joinValidationMessage(message, err.Error())
	}
	if len(validationErrs) == 0 {
		return message
	}

	messages := make([]string, 0, len(validationErrs))
	for _, fieldErr := range validationErrs {
		messages = append(messages, validationFieldErrorMessage(fieldErr))
	}
	return joinValidationMessage(message, strings.Join(messages, " "))
}

func joinValidationMessage(message, detail string) string {
	if message == "" {
		return detail
	}
	if detail == "" {
		return message
	}
	return fmt.Sprintf("%s: %s", message, detail)
}

func validationFieldErrorMessage(err validator.FieldError) string {
	field := validationFieldName(err)

	switch err.Tag() {
	case "required":
		return fmt.Sprintf("The %s field is required.", field)
	case "email":
		return fmt.Sprintf("The %s field must be a valid email address.", field)
	case "numeric":
		return fmt.Sprintf("The %s field must be numeric.", field)
	case "len":
		return fmt.Sprintf("The %s field must be %s characters.", field, err.Param())
	case "min":
		return fmt.Sprintf("The %s field must be at least %s characters.", field, err.Param())
	case "max":
		return fmt.Sprintf("The %s field must not be greater than %s characters.", field, err.Param())
	case "pixelfed_username":
		return fmt.Sprintf("The %s field must be a valid username.", field)
	default:
		return fmt.Sprintf("The %s field is invalid.", field)
	}
}

func validationFieldName(err validator.FieldError) string {
	field := err.Field()
	if field == "" {
		return "value"
	}
	return snakeCase(field)
}

func snakeCase(s string) string {
	var b strings.Builder
	b.Grow(len(s))

	for i := 0; i < len(s); i++ {
		c := s[i]
		if c >= 'A' && c <= 'Z' {
			if i > 0 {
				b.WriteByte('_')
			}
			b.WriteByte(c + ('a' - 'A'))
			continue
		}
		b.WriteByte(c)
	}

	return b.String()
}
