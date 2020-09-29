package main

import "fmt"

func main() {
	// Declaration of variables
	var (
		day = 0
		month = 0
	)

	// Input data
	fmt.Scan(&day)
	fmt.Scan(&month)

	zodiac := ""
	// Conditions
	if  day >= 21 && month == 1 || day <= 18 && month == 2 {
		zodiac = "📝 acuario"
	} else if day >= 19 && month == 2 || day <= 20 && month == 3 {
		zodiac = "📝 piscis"
	} else if day >= 21 && month == 3 || day <= 20 && month == 4 {
		zodiac = "📝 aries"
	} else if day >= 21 && month == 4 || day <= 20 && month == 5 {
		zodiac = "📝 tauro"
	} else if day >= 21 && month == 5 || day <= 21 && month == 6 {
		zodiac = "📝 geminis"
	} else if day >= 22 && month == 6 || day <= 22 && month == 7 {
		zodiac = "📝 cancer"
	} else if day >= 23 && month == 7 || day <= 22 && month == 8 {
		zodiac = "📝 leo"
	} else if day >= 23 && month == 8 || day <= 22 && month == 9 {
		zodiac = "📝 virgo"
	} else if day >= 23 && month == 9 || day <= 22 && month == 10 {
		zodiac = "📝 libra"
	} else if day >= 23 && month == 10 || day <= 22 && month == 11 {
		zodiac = "📝 escorpio"
	} else if day >= 23 && month == 11 || day <= 21 && month == 12 {
		zodiac = "📝 sagitario"
	} else { // day >= 22 && month == 12 || day <= 20 && month == 1
		zodiac = "📝 capricornio"
	}

	// Print result
	println(zodiac)
}