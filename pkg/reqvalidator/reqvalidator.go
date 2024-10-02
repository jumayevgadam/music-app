package reqvalidator

import (
	"context"
	"github.com/go-playground/validator/v10"
)

// Use a single instance of Validate, it caches struct info
var validate *validator.Validate

func init() {
	validate = validator.New()
}

// ValidateStruct fields for models
func ValidateStruct(ctx context.Context, s interface{}) error {
	return validate.StructCtx(ctx, s)
}
