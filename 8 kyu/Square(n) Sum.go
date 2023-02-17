//problem 			: https://www.codewars.com/kata/515e271a311df0350d00000f/train/go

//Best pratices  	: https://www.codewars.com/kata/reviews/5d82854b15a9070001e404df/groups/5e077bf78c89df0001d86657
package main

import "fmt"

func main() {
	//fmt.Println(math.Pow(2, 2))
	var balance = []int{1, 2, 2}
	fmt.Println(SquareSum(balance))
}

func SquareSum(numbers []int) int {
	// your code here

	var totalSum int
	for i := 0; i < len(numbers); i++ {
		totalSum = totalSum + (numbers[i] * numbers[i])

	}
	return totalSum
}
