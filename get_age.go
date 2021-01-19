package main 

import (
	"fmt"
	"bufio"           //buffered input/output
	"os"              //operating system
	"strconv"
)

func main() {
	fmt.Printf("Enter birth year: ")
	scanner := bufio.NewScanner(os.Stdin)                        //scanner object
	scanner.Scan()                                               //read input
	user_input, _ := strconv.ParseInt(scanner.Text(), 10, 64)    //store input in a variable as int(would automatically be string)
	get_age(user_input)                                          //call get_age 
}

func get_age(birth_year int64) {                               //takes in the birth year as an int parameter
	age := 2020-birth_year                                       //calculate the age
	if age > 0 {
		fmt.Printf("You are %d years old by 2021", age)
	} else {
		fmt.Println("Invalid date")
	}
	
}
