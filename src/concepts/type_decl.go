package main

import "fmt"

type Farenheit float64
type Celsius float64
//define a String() method for the Celsius type
func (c Celsius) String() string { return fmt.Sprintf("%g°C", c) }

func main() {

	var f Farenheit = 99.0
	var c Celsius = 21.0

	//fmt.Println(f + c)

	// error when building
	// go:13: invalid operation: f + c (mismatched types Farenheit and Celsius)

	// You can cast the type so that it works, even though it is semantically wrong
	fmt.Println(f + Farenheit(c))
	fmt.Println(c + Celsius(f))
}
