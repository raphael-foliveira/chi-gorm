package validate

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

func Required(key string, value any) error {
	t := reflect.TypeOf(value)
	noop := reflect.Zero(t).Interface()
	if reflect.DeepEqual(value, noop) {
		return fmt.Errorf("%v: is required", key)
	}
	return nil
}

func NumberString(key string, value string) error {
	if _, err := strconv.Atoi(value); err != nil {
		return fmt.Errorf("%v: must be a number", key)
	}
	return nil
}

func Email(key string, value string) error {
	if !strings.Contains(value, "@") {
		return fmt.Errorf("%v: must be an email", key)
	}
	return nil
}

func Min(key string, value int, target int) error {
	if value < target {
		return fmt.Errorf("%v: must be greater than %v", key, target)
	}
	return nil
}

func Max(key string, value int, target int) error {
	if value > target {
		return fmt.Errorf("%v: must be less than %v", key, target)
	}
	return nil
}

func MinLength(key, value string, target int) error {
	if len(value) < target {
		return fmt.Errorf("%v: must be at least %v characters", key, target)
	}
	return nil
}

func MaxLength(key, value string, target int) error {
	if len(value) > target {
		return fmt.Errorf("%v: must be at most %v characters", key, target)
	}
	return nil
}
