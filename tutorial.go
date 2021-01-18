package main    //package

import "fmt"    //module

func main() {
	var name string = "tom"     //variable name can't start with a number
	var a uint8 = 200           //uint8 is integers up to 255((2^8)-1)
	var b uint8 = 40            //int8 is signed integers between -128 and 127
	fmt.Println("Hello World!", name, "\n", a+b)
	what_type()
}

func what_type() {
	number := 6                 //golang will assume the type type of the value(int) and := avoids using var
	fmt.Printf("%T", number)    //"%t" is the type of value
}
