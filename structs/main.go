package main

import (
	"fmt"
)

type person struct {
	firstName string
	lastName  string
	age       int
	working   bool
	contactInfo
}
type contactInfo struct {
	email string
	zip   int
}

func main() {
	sampath := person{
		firstName: "Sampathkumar",
		lastName:  "Subramaniam",
		age:       39,
		working:   true,
		contactInfo: contactInfo{
			email: "a@a.com",
			zip:   70139,
		},
	}
	sampathOriginalAddress := &sampath
	sampathOriginalAddress.updateName("Nadhira")
	fmt.Printf("%+v\n", sampath)
	tmp := sampath
	tmp.updateName("hello")
	fmt.Printf("%+v", sampath)
}

func (p *person) updateName(newFirstName string) {
	(*p).firstName = newFirstName
}
