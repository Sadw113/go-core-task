package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"reflect"
	conv "strconv"
)

var numDecimal int = 42           // Десятичная система
var numOctal int = 052            // Восьмеричная система
var numHexadecimal int = 0x2A     // Шестнадцатиричная система
var pi float64 = 3.14             // Тип float64
var name string = "Golang"        // Тип string
var isActive bool = true          // Тип bool
var complexNum complex64 = 1 + 2i // Тип complex64

func typePrinter() {
	fmt.Printf("Type of numDecimal %d is %s\n", numDecimal, reflect.TypeOf(numDecimal))
	fmt.Printf("Type of numOctal %o is %s\n", numOctal, reflect.TypeOf(numOctal))
	fmt.Printf("Type of numHexadecimal %x is %s\n", numHexadecimal, reflect.TypeOf(numHexadecimal))
	fmt.Printf("Type of pi %f is %s\n", pi, reflect.TypeOf(pi))
	fmt.Printf("Type of name %s is %s\n", name, reflect.TypeOf(name))
	fmt.Printf("Type of isActive %t is %s\n", isActive, reflect.TypeOf(isActive))
	fmt.Printf("Type of complexNum %f is %s\n", complexNum, reflect.TypeOf(complexNum))
}

func mergeStringVars() string {
	var res string

	res = conv.Itoa(numDecimal) + conv.Itoa(numOctal) +
		conv.Itoa(numHexadecimal) + conv.FormatFloat(pi, 'e', -1, 64) +
		name + conv.FormatBool(isActive) + conv.FormatComplex(complex128(complexNum), 'e', -1, 64)

	return res
}

func main() {
	typePrinter()
	mergedStr := mergeStringVars()
	mergedRune := []rune(mergedStr)
	hasher := sha256.New()
	left := mergedRune[:len(mergedRune)/2]
	right := mergedRune[len(mergedRune)/2:]
	res := hasher.Sum([]byte(string(left) + "go-2024" + string(right)))
	fmt.Print(hex.EncodeToString(res))
}
