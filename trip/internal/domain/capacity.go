package domain

import "errors"

type Capacity map[string]int

func ValidateCapacity(c Capacity) error {
	required := []string{
		"male", "female", "base_person", "extra_person",
	}
	for _, k := range required {
		if _, ok := c[k]; !ok {
			return errors.New("er1012")
		}
	}
	return nil
}
