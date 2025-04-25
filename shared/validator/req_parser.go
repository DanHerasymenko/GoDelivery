package req_parser

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"regexp"
	"github.com/gin-gonic/gin"
)

// доволі масивна штука - там багато оптимізацій, кешування і т.д.
// правильно не створювати на кожен чих новий валідатор

var v = validator.New()

var seasonRegex = regexp.MustCompile(`^(Winter|Spring|Summer|Autumn)[0-9]{4}$`)

func ParseReqBody(ctx *gin.Context, reqBody interface{}) error {
	
	if err := ctx.ShouldBindJSON(&reqBody); err != nil {
		return fmt.Errorf("failed to parse request body: %w", err)
	}

	if err := v.Struct(reqBody); err != nil {
		return fmt.Errorf("failed to validate request body: %w", err)
	}

	return nil
}