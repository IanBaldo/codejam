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
	file, err := os.Open("B-small-practice.in")
	checkErr(err)
	
	// Closes the file once the program finnishes 
	defer file.Close()


	scanner := bufio.NewScanner(file)
	
	caseCounter := 1
	for scanner.Scan() {
		n, err := strconv.Atoi(scanner.Text())
		checkErr(err)
		fmt.Printf("Case #%d: %d\n", caseCounter, calcLastTidy(n))
		
		caseCounter++
	}


}


func calcLastTidy (num int) int {
	for i:= num; i>0; i-- {
		if isTidy(i) {
			return i
		}
	}
	return 0
}


func isTidy (n int) bool {
	
	exp := int(math.Log10(float64(n)))
	
	// 1 digit numbers are tidy by definition!
	if exp == 0 {
		return true
	}

	for i:= 1; i<= exp; i++ {
		var low, high int

		if i == 1 {
			high =  n % int(math.Pow10(i))
			low =  n % int(math.Pow10(i+1)) / int(math.Pow10(i))
		} else {
			high =  n % int(math.Pow10(i)) / int(math.Pow10(i-1))
			low =  n % int(math.Pow10(i+1)) / int(math.Pow10(i))
		}
		
		if high < low {
			return false
		}
	}

	return true
}



