package helper

import "errors"

func CheckEmpty(input ...interface{}) error {
	for _, value := range input {
		switch value {
		case "":
			return errors.New("input cannot be empty")
		case 0:
			return errors.New("input cannot be zero")
		case nil:
			return errors.New("input cannot be empty")
		}
	}

	return nil
}