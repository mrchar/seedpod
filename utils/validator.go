package utils

import "github.com/go-playground/validator/v10"

// 用于验证结构体约束
var validate = validator.New()

func ValidateStruct(v interface{}) error {
	return validate.Struct(v)
}
