package main 

import (
	"fmt"
	"bufio"
	"strconv"
	"os"
)

func main() {
	fmt.Println("Enter a number to start the Fizzbuzz at: ")
	scanner := bufio.NewScanner(os.Stdin)                         //scanner object
	scanner.Scan()                                                //read the user input
	user_input, _ := strconv.ParseInt(scanner.Text(), 10, 64)     //read the input as an integer and store in a variable
	fizzbuzz(user_input)                                          //call fizzbuzz()
}

func fizzbuzz(start_number int64) {                             //takes in the user input as an int parameter
	for i:=start_number; i<start_number+100; i++ {                //start from userinput, go until user input + 100
		if i%3 == 0 && i%5 == 0 {                                   //both
			fmt.Println(i,"Fizzbuzz")
			continue                                                  //back to the loop
		}else if i%3 == 0 {                                         //3 only
			fmt.Println(i,"Fizz")
			continue                                                  //back to the loop
		} else if i%5 == 0 {                                        //5 only
			fmt.Println(i,"Buzz")
			continue                                                  //back to the loop
		} else {                                                    //neither
			continue
		}
	}
}
