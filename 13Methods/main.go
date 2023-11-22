package main

import "fmt"

func main() {
	// no inheritence in golang;  No Super No Parent
	rohit := User{"Rohit", "rohitvarshney87@gmail.com", true, 23}
	fmt.Println(rohit)
	fmt.Printf("Rohit details are: %+v \n", rohit)
	fmt.Printf("Name : %v and Email is: %v\n", rohit.Name, rohit.Email)
	rohit.getStatus()
	rohit.NewEmail()
	fmt.Println(rohit.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) getStatus() {
	fmt.Println("Status of user : ", u.Status)
}

func (u User) NewEmail() {
	u.Email = "test@go.dev"
	fmt.Println("Email of this user is: ", u.Email)
}
