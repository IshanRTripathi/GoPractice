package main

import "fmt"

func main() {
	fmt.Println("Yeaaaaahhhhhhh!")

	// initialising variables
	initialiseVariables()

	// initialising complex data structures
	var response = initialiseDataStructures()
	fmt.Println(response, " is the response from initialiseDataStructures() function")

	// conditional statements
	performConditionalStatements()

	// looping statements
	performLoops()

	// pass params in a function
	res := getReceivedValue("ishan")
	fmt.Println("res = ", res)
}

func getReceivedValue(input string) string {
	return input+" r tripathi"
}

func performLoops() {
	for i:=1; i<10; i++ {
		if i%2 == 0 && i > 2 {
			fmt.Println(i, " is the current even value of i > 2")
		}
	}
	fmt.Println("New loop => ")
	for i:=1; i<10; i++ {
		if i%2 == 0 && i > 4 {
			fmt.Println(i, " is the current even value of i > 4, hence breaking")
			break
		}
	}
}

func performConditionalStatements() {
	x := 6
	var y = 7

	if x > y {
		fmt.Println(x, " > ", y)
	} else if x == y {
		fmt.Println(x, " == ", y)
	} else {
		fmt.Println(x, " <= ", y)
	}

	var day = 65656
	switch day {
	case (1):
		fmt.Println("Saturday")
	case (2):
		fmt.Println("Sunday")
	case 2222:
		fmt.Println("ajshfvajfvakday")
	default:
		fmt.Println("default day")
	}
}

func initialiseDataStructures() bool {

	return true
}

func initialiseVariables() {
	var num = 5
	fmt.Println(num)

	numsArr := [3]int{1,2,3}
	fmt.Println(numsArr)

	var numsArr2 = [5]int{1,2,3,4,5}
	fmt.Println(numsArr2)

	numsArr3 := [...]int{1,2,3,4,5,6,7}
	fmt.Println(numsArr3)

	var numsArr4 = [...]int{1,2,3,4,5,6,7,8,9}
	fmt.Println(numsArr4)
	fmt.Println(numsArr4[2], " is the 2nd index value ")
	numsArr4[0] = 100
	fmt.Println(numsArr4, " modified 0th index")

	var myString = "ishan"
	fmt.Println(myString)

	myString += "a"
	fmt.Println(myString)

}
