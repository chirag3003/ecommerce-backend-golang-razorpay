package helpers

import "strings"

type Errors map[string][]string

func (e Errors) Add(field string, err string) {
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

func (e Errors) CheckLen(str string, minLen int) bool {
	if len(strings.TrimSpace(str)) < minLen {
		return false
	}
	return true
}

func (e Errors) CheckMaxLen(str string, maxLen int) bool {
	if len(strings.TrimSpace(str)) > maxLen {
		return false
	}
	return true
}

func (e Errors) CheckMinValue(val float64, minVal float64) bool {
	if val < minVal {
		return false
	}
	return true
}

func (e Errors) CheckMaxValue(val float64, maxVal float64) bool {
	if val > maxVal {
		return false
	}
	return true
}
