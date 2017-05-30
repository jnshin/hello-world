package jnshin

import (
	"fmt"
	"strconv"
	"strings"
)

//
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
