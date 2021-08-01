package kulina

import (
	"fmt"
	"strings"
)

//asume giving a number with type variable string
func number(number string) {
	var newNumber string
	if strings.Contains(number, ".") {
		fmt.Println("masuk")
		newNumber = strings.ReplaceAll(number, ".", "")
	}

	for i := 0; i < len(newNumber); i++ {
		index := strings.Index(newNumber, string(newNumber[i]))
		if index < len(newNumber) {
			ix := len(newNumber) - index
			var zeros string
			for i := 1; i <= ix-1; i++ {
				if ix == 1 {
					break
				}
				zeros += "0"
			}
			fmt.Println(string(newNumber[i]) + zeros)
		}
	}
}
