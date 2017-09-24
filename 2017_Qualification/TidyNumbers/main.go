package main

import (
	"strconv"
	"math"
	"bufio"
	"os"
	"fmt"
)

func checkErr (err error) {
	if err != nil {
		panic(err)
	}
}

func main () {

	// Open File and checks for error
	file, err := os.Open("B-large-practice.in")
	checkErr(err)
	
	// Closes the file once the program finnishes 
	defer file.Close()


	scanner := bufio.NewScanner(file)
	
	caseCounter := 1
	for scanner.Scan() {
		n, err := strconv.Atoi(scanner.Text())
		checkErr(err)
		fmt.Printf("[%d]: Case #%d: %d\n", n, caseCounter, calcLastTidy(n))
		
		caseCounter++
	}


}


func calcLastTidy (num int) int {
	curr := num
	for curr > 0 {
		nComp := int(math.Log10(float64(curr)))
		original := curr
		
		// 1 digit numbers are tidy by definition!
		if nComp == 0 {
			return curr
		}
	
		numberModified := false
		for i:= nComp; i > 0 ; i-- {
			var low, high int
			
			if i == 1 {
				high =  curr % int(math.Pow10(i))
				low =  curr % int(math.Pow10(i+1)) / int(math.Pow10(i))
			} else {
				high =  curr % int(math.Pow10(i)) / int(math.Pow10(i-1))
				low =  curr % int(math.Pow10(i+1)) / int(math.Pow10(i))
			}
			
			if high < low && numberModified == false {
				numberModified = true
				curr -= (high + 1) * int(math.Pow10(i-1))
			} else if high < low && numberModified == true {
				curr += (9-high)*int(math.Pow10(i-1))
			}
		}

		if original == curr {
			return curr
		}

	}
	return 0
}
