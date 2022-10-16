package main

import (
	"fmt"
	"strings"
)

func main() {
	myStr := " Azie Melza Pratama   "
	fmt.Println("Original : ", myStr)

	trimMyStr := strings.Trim(myStr, " ")
	fmt.Println(trimMyStr, "==> Trimmed")

	upperMyStr := strings.ToUpper(trimMyStr)
	fmt.Println(upperMyStr, "==> Upper Cased")

	lowerMyStr := strings.ToLower(trimMyStr)
	fmt.Println(lowerMyStr, "==> Lower Cased")

	str1 := "umar berjalan di pagar ular"
	titledMyStr := strings.Title(str1)
	fmt.Println(titledMyStr, "==> Titled from ==> ", str1)

	var splitMyStr []string = strings.Split(trimMyStr, " ")
	fmt.Println(splitMyStr)

	ctnMyStr1 := strings.Contains(trimMyStr, "Melza")
	ctnMyStr2 := strings.Contains(trimMyStr, "Charles")
	fmt.Println("Does This String Contains Melza ? : ", ctnMyStr1)
	fmt.Println("Does This String Contains Charles ? : ", ctnMyStr2)

	replaceMyStr := strings.ReplaceAll(trimMyStr, "Pratama", "Vanhautten")
	fmt.Println(replaceMyStr)
}
