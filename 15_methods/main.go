package main

import "fmt"

func main() {
	fmt.Println("welcome to structs in golang")

	// no inheritance in golang

	user1 := User{Name: "Alex",Email: "alex@iitj.ac.in", IsVerified: false, Phone: 7894561231}
	fmt.Println(user1)

	fmt.Printf("the details of user1 are %+v\n", user1)
	fmt.Printf("the emailId of user1 is %v\n", user1.Email)

	user1.GetStatus()

	user1.NewMail("rishabh@go.dev")
	fmt.Println(user1)
}

type User struct {
	Name       string
	Email      string
	IsVerified bool
	Phone      int
}


func (u User) GetStatus (){
	fmt.Println("Is the email of user is verified ", u.IsVerified)
}

func (u User) NewMail(email string){
	u.Email = email
	// note this will no update the actual value in the object
}