package main

import "fmt"

func grading () {
	fmt.Println("this is a grading system using struck")
	type Person struct{
		Name string
		Age int 
	}
	person := Person{
		Name : "Pranto",
		Age : 25 ,
	}
	person1 := Person{
		Name : "shovon",
		Age : 26,

	}

	

	person.Age = 25
	person1.Age = 24

	fmt.Println(person)
	fmt.Println(person1)

}
