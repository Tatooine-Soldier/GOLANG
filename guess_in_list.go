package main                      //You try to guess if your number is in this randomly generated array

import (
	"fmt"
	"strconv"
	"bufio"
	"os"
	"math/rand"
)

func main() {
	fmt.Println("Enter your Number(0-100): ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	number,_ := strconv.ParseInt(scanner.Text(), 10, 64)                         //take in as an integer(64bit)
	get_pattern(number)                                                          //pass user input into method below
}

func get_pattern(dnumber int64) {                              
	var count int64 = 0
	darray := [3]int64{rand.Int63n(100), rand.Int63n(100), rand.Int63n(100)}     //array of 3 random numbers
	for i:=0;i<len(darray);i++ { 
		if darray[i] == dnumber {                                                //if number is in the array
			fmt.Printf("Number %v found at positon %v\n",dnumber,i)              //say when its found 
			count = count + 1                                                    //counter for number of appearances
			continue 
		} 
	} 
	fmt.Printf("%v appeared in the array %v time(s)...",dnumber,count)            //display how many times the entered number appears
} 
