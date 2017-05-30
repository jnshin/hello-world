package jnshin

import (
	"fmt"
	"strconv"
	"strings"
)

// Cntoi : comma 표기가 포함된 숫자를 일반 int로 변환한다.
func Cntoi(str string) (int, error) {
	if len(str) == 0 {
		return 0, nil
	}
	if tmpVal, err := strconv.Atoi(strings.Replace(str, ",", "", -1)); err != nil {
		return 0, fmt.Errorf("Conversion failed. org %v. %v", str, err.Error())
	} else {
		return tmpVal, nil
	}
}

// Cntof : comma 표기가 포함된 숫자를 일반 float로 변환한다.
func Cntof(str string) (float64, error) {
	if len(str) == 0 {
		return 0, nil
	}
	if tmpVal, err := strconv.ParseFloat(strings.Replace(str, ",", "", -1), 32); err != nil {
		return 0, fmt.Errorf("Conversion failed. org %v. %v", str, err.Error())
	} else {
		return tmpVal, nil
	}
}
