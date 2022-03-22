package helpers

import (
	"fmt"
	"strings"
)

type Errors map[string][]string

func (e Errors) Add(field, err string) {
	e[field] = append(e[field], err)
}

func (e Errors) Get(field string) []string {
	return e[field]
}

func (e Errors) IsValid() bool {
	if len(e) == 0 {
		return true
	}
	return false
}

func (e Errors) CheckLen(str string, minLen int, field ...string) bool {
	if len(strings.TrimSpace(str)) < minLen {
		if len(field) > 0 {
			e.Add(field[0], fmt.Sprintf("min length required is %d", minLen))
		}
		return false
	}
	return true
}

func (e Errors) CheckMaxLen(str string, maxLen int, field ...string) bool {
	if len(strings.TrimSpace(str)) > maxLen {
		if len(field) > 0 {
			e.Add(field[0], fmt.Sprintf("max length is %d", maxLen))
		}
		return false
	}
	return true
}

func (e Errors) CheckMinValue(val, minVal float64, field ...string) bool {
	if val < minVal {
		if len(field) > 0 {
			e.Add(field[0], fmt.Sprintf("min value required is %f", minVal))
		}
		return false
	}
	return true
}

func (e Errors) CheckMaxValue(val, maxVal float64, field ...string) bool {
	if val > maxVal {
		if len(field) > 0 {
			e.Add(field[0], fmt.Sprintf("max value required is %f", maxVal))
		}
		return false
	}
	return true
}
