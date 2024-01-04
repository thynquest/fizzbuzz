package service

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var stats map[string]int

type FizzBuzz interface {
	FizzBuzz(int1 int, int2 int, limit int, str1, str2 string) ([]string, error)
	GetStats() map[string]int
}

func init() {
	stats = make(map[string]int)
}

type fizzBuzz struct{}

func NewFizzBuzz() FizzBuzz {
	return &fizzBuzz{}
}

func (f *fizzBuzz) updateStats(multint1, multint2, limit, multstr1, multstr2 string) {
	key := strings.Join([]string{multint1, multint2, limit, multstr1, multstr2}, ",")
	stats[key] = stats[key] + 1
}

func (f *fizzBuzz) GetStats() map[string]int {
	return stats
}

func (f *fizzBuzz) FizzBuzz(int1 int, int2 int, limit int, str1, str2 string) ([]string, error) {
	if limit <= 0 {
		err := errors.New("limit parameter must be positive")
		return nil, err
	}
	if int1 <= 0 || int2 <= 0 {
		err := errors.New("multiples must be positive")
		return nil, err
	}
	var result []string
	for cpt := 1; cpt <= limit; cpt++ {
		ismultint1 := (cpt%int1 == 0)
		ismultint2 := (cpt%int2 == 0)

		if ismultint1 && ismultint2 {
			result = append(result, str1+str2)
		} else if ismultint1 {
			result = append(result, str1)
		} else if ismultint2 {
			result = append(result, str2)
		} else {
			item := fmt.Sprintf("%d", cpt)
			result = append(result, item)
		}
	}
	multint1 := strconv.Itoa(int1)
	multint2 := strconv.Itoa(int2)
	l := strconv.Itoa(limit)
	f.updateStats(multint1, multint2, l, str1, str2)
	return result, nil
}
