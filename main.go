package main

import (
	"fmt"
	"sort"
)

func main() {
	//s := "gopher"
	//fmt.Printf("Hello and welcome, %s!\n", s)
	//
	//for i := 1; i <= 5; i++ {
	//	fmt.Println("i =", 100/i)
	//}

	//fmt.Println("Enter the degrees in Farhneit")
	//var input float64
	//
	//fmt.Scanf("%f", &input)
	//
	//var output float64 = (input - 32) * 5 / 9
	//fmt.Println("The temperature is ", output, "C")

	//numbers := []int{10, 20, 30, 40, 50}
	//for _, number := range numbers {
	//	fmt.Println(number)
	//}

	//// Rethinking the modulus operator
	//for i := 1; i <= 100; i++ {
	//	if i%3 == 0 {
	//		fmt.Printf("%d ", i)
	//	}
	//}

	// The FizzBuzz question
	//for i := 1; i <= 100; i++ {
	//	if i%3 == 0 {
	//		fmt.Printf("Fizz \n")
	//	} else if i%5 == 0 {
	//		fmt.Printf("Buzz \n")
	//	} else if i%3 == 0 && i%5 == 0 {
	//		fmt.Printf("FizzBuzz \n")
	//	} else {
	//		fmt.Printf("%d \n", i)
	//	}
	//}

	// Arrays and Maps
	//Find the smallest number in this array
	x := []int{
		48, 96, 86, 68,
		57, 82, 63, 70,
		37, 34, 83, 27,
		19, 97, 9, 17,
	}

	minVal := x[0]

	for _, v := range x {
		if v < minVal {
			minVal = v
		}
	}
	fmt.Println("The minimum value using linear scan is:", minVal)

	sortedX := make([]int, len(x))
	copy(sortedX, x)
	sort.Ints(sortedX)
	fmt.Printf("This is the sorted array %d\n", sortedX)
	fmt.Println("The minimum value using go's sorting is:", minVal)

	// Traverse the array comparing the two nearest values to each other,
	//find the minimum and arrange it in ascending order then print the first number in the array,
	//or possibly the  whole arranged array first

	//x := make(map[string]int)
	//x["key"] = 10
	//fmt.Println(x["key"])

	//	 Maps as dictionary
	//elements := map[string]map[string]string{
	//	"H": {
	//		"name":  "Hydrogen",
	//		"state": "gas",
	//	},
	//	"He": {
	//		"name":  "Helium",
	//		"state": "gas",
	//	},
	//	"Li": {
	//		"name":  "Lithium",
	//		"state": "solid",
	//	},
	//	"Be": {
	//		"name":  "Beryllium",
	//		"state": "solid",
	//	},
	//	"B": {
	//		"name":  "Boron",
	//		"state": "solid",
	//	},
	//	"C": {
	//		"name":  "Carbon",
	//		"state": "solid",
	//	},
	//	"N": {
	//		"name":  "Nitrogen",
	//		"state": "gas",
	//	},
	//	"O": {
	//		"name":  "Oxygen",
	//		"state": "gas",
	//	},
	//	"F": {
	//		"name":  "Fluorine",
	//		"state": "gas",
	//	},
	//	"Ne": {
	//		"name":  "Neon",
	//		"state": "gas",
	//	},
	//}
	//if el, ok := elements["Li"]; ok {
	//	fmt.Println(el["name"], el["state"])
	//}

}
