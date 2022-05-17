package main

import "fmt"

type contactInfo struct{
	email string
	zipCode int
}

type person struct {
	firstname string
	lastName  string
	contactInfo // you dont have to specify field name here
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
		contactInfo: contactInfo{
			email: "jim@gmail.com",
			zipCode: 40000,
		},
	}

	fmt.Printf("%+v\n",jim)

	fmt.Println("------Struct with receiver function------")
	//go is passby language if you call a type it make copy of it in memory
	// jimPointer := &jim //but tehre is a shortcut for this process
	// jimPointer.updateName("jimmy")

	jim.updateName("Jimmy")//jim is person type but its still work
	//if you make receiver type of pointer go automaticly takes it pointer
	//You dont have to do pointer thing for slices
	jim.printPerson()
	
}


func (pointerToPerson *person) updateName(newFirstName string){
	(*pointerToPerson).firstname=newFirstName
}

func (p person) printPerson()  {
	fmt.Printf("%+v\n",p)
}