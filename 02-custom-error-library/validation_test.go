package validation_test

import (
	"errors"
	"testing"

	validation "github.com/amanycodes/golang-exercises/02-custom-error-library"
)

func TestValidationInput(t *testing.T) {
	err := validation.ValidateInput("")
	if err == nil {
		t.Fatal("expected error for empty input, got nil")
	}
	var vErr *validation.ValidationError
	if !errors.As(err, &vErr) {
		t.Fatalf("expected validaton error, got %v", err)
	}
	if vErr.Code != validation.ErrCodeValidation {
		t.Fatalf("unexpected error code: %v", vErr.Code)
	}
}

func TestCheckAuthorization(t *testing.T) {
	err := validation.CheckAuthorization("cde")
	if err == nil {
		t.Fatal("expected auth error, got nil")
	}
	var aErr *validation.AuthorizationError
	if !errors.As(err, &aErr) {
		t.Fatalf("expected Authorization Error, got %v", aErr)
	}
	if aErr.Code != validation.ErrCodeAuthorization {
		t.Fatalf("unexpected code, got %v", aErr.Code)
	}
}

func TestFindResource(t *testing.T) {
	_, err := validation.FindResource(12)
	if err == nil {
		t.Fatal("expected Resource error, got nil")
	}
	var fErr *validation.NotFoundError
	if !errors.As(err, &fErr) {
		t.Fatalf("expected not found error, got %v", err)
	}
	if fErr.Code != validation.ErrCodeNotFound {
		t.Fatalf("expected not found error, got %v", fErr.Code)
	}
}
