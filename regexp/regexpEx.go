package regex

import (
	"fmt"
	"regexp"
)

func regexEx() {
	regexpStr := `^[0-9]\.[0-9]\.[0-9](\.[0-9])?$`
	reg, err := regexp.Compile(regexpStr)
	if err != err {
		fmt.Println(err)
	}
	fmt.Println(reg.Match([]byte("2.3.22")))
}
