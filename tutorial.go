package main

import "fmt"

// func main() {
// 	fmt.Println("Welcome to my nightmare!")

// 	// var name string = "Cheryl"
// 	// The below syntax gets Go to guess the type.
// 	name := "Cheryl"
// 	age := 40
// 	// if it's a float a number must be assigned to it to show how much computer memory is being used ex float64
// 	// if a variable isn't used go will give you an error

// 	// notice Printf is differnt than Println (print line)
// 	// %v is a place holder
// 	// fmt.Printf("Hello %v %v, how are you doing???", name, name)
// 	fmt.Printf("Hello %v, You are %v", name, age)
// 	// https://gobyexample.com/string-formatting
// }

func main() {
	fmt.Println("Welcome to my game!")
	fmt.Printf("Enter your name: ")
	// \n will create a line break
	var name string

// Scan allows you to enter a variable see below
// & allows you to set reference to that variable.
	fmt.Scan(&name)
	
	fmt.Printf("Hello %v, Welcome to my game!", name)

	fmt.Printf("Enter your age: ")
	// uint will not allow a negative value.
	var age uint
	fmt.Scan(&age)
	fmt.Println(age >= 10) //This will return a boolean 
}