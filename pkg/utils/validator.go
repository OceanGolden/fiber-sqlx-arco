package utils

import "fiber-sqlx-arco/pkg/configs"

func ValidateStruct(s interface{}) []string {
	err := configs.Validate.Struct(s)
	if err != nil {
		return configs.ValidatorErrors(err)
	}
	return nil

}
