package main

import (
	"flag"
	"fmt"
	"password-generator/util"
)

var length int
var useLower bool
var useUpper bool
var useDigit bool
var specialChars string

func init() {
	flag.IntVar(&length, "length", 16, "Set the length of the password with this data.")
	flag.IntVar(&length, "L", 16, "Set the length of the password with this data.")

	flag.BoolVar(&useLower, "lower", false, "Let the password use lower case letters.")
	flag.BoolVar(&useLower, "l", false, "Let the password use lower case letters.")

	flag.BoolVar(&useUpper, "upper", false, "Let the password use upper case letters.")
	flag.BoolVar(&useUpper, "u", false, "Let the password use upper case letters.")

	flag.BoolVar(&useDigit, "digit", false, "Let the password use digits.")
	flag.BoolVar(&useDigit, "d", false, "Let the password use digits.")

	flag.StringVar(&specialChars, "specialChars", "", "Let the password use these characters.")
	flag.StringVar(&specialChars, "s", "", "Let the password use these characters.")
}

func main() {
	flag.Parse()

	// log.SetPrefix("[INFO]")
	// log.Printf("useLower = %v\n", useLower)
	// log.Printf("useUpper = %v\n", useUpper)
	// log.Printf("useDigit = %v\n", useDigit)
	// log.Printf("special = [%v]\n", specialChars)
	// log.Printf("length = %v\n", length)

	defer func() {
		err := recover()

		// Handle panic
		if err != nil {
			fmt.Println("Error!", err)
		}
	}()

	password := util.GeneratePassword(length, useDigit, useLower, useUpper, specialChars)

	fmt.Printf("%v\n", password)
}
