//Problem : https://www.codewars.com/kata/5a3fe3dde1ce0e8ed6000097/train/go

//clever solution : https://www.codewars.com/kata/reviews/5a4ecce7577aadb893002a57/groups/5a4ecce844acf77ff80032ca

package main

import "fmt"

func main() {
	fmt.Println(century(1800))
}

func century(year int) int {
	// Finish this :)
	if year%100 == 0 {
		return year / 100
	}
	return year/100 + 1
}
