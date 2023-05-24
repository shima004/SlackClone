package delivery

import (
	"strconv"

	"github.com/go-playground/validator"
)

func isRequestValid(m interface{}) (bool, error) {
	validate := validator.New()
	err := validate.Struct(m)
	if err != nil {
		return false, err
	}
	return true, nil
}

func StringToUint(s string) (uint, error) {
	u, err := strconv.ParseUint(s, 10, 32)
	return uint(u), err
}
