package handlers

import (
	"errors"
	"fmt"
)

//FizzBuzz : returns all multiple of multint1, multint2, and both in the range of limit
func FizzBuzz(multint1, multint2, limit int, multstr1, multstr2 string) ([]string, error) {
	if limit <= 0 {
		err := errors.New("limit parameter must be positive")
		return nil, err
	}
	if multint1 <= 0 || multint2 <= 0 {
		err := errors.New("multiples must be positive")
		return nil, err
	}
	var result []string
	for cpt := 1; cpt <= limit; cpt++ {
		ismultint1 := (cpt%multint1 == 0)
		ismultint2 := (cpt%multint2 == 0)

		if ismultint1 && ismultint2 {
			result = append(result, multstr1+multstr2)
		} else if ismultint1 {
			result = append(result, multstr1)
		} else if ismultint2 {
			result = append(result, multstr2)
		} else {
			item := fmt.Sprintf("%d", cpt)
			result = append(result, item)
		}
	}
	return result, nil
}
