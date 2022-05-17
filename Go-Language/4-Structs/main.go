package main

import "fmt"

type contactInfo struct{
	email string
	zipCode int
}

type person struct {
	firstname string
	lastName  string
	contact contactInfo
}

func main() {
	fmt.Println("------Struct lesson 1------")
	kaan := person{firstname: "Kaan",lastName:  "Akg"}
	josuke := person{lastName: "Higashikata",firstname: "Josuke"}
	fmt.Println(kaan.firstname,kaan.lastName)
	fmt.Println(josuke.firstname,josuke.lastName)

	var jonathan person
	fmt.Printf("%+v\n",jonathan)
	jonathan.firstname="Jonathan"
	jonathan.lastName="Joestar"
	fmt.Printf("%+v\n",jonathan)
	
	fmt.Println("------Struct lesson 2------")

	jim:= person{
		firstname: "Jim",
		lastName: "Mor",
		contact: contactInfo{
			email: "jim@gmail.com",
			zipCode: 40000,
		},
	}

	fmt.Printf("%+v\n",jim)

}