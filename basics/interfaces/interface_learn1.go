package main

import "fmt"

type User struct {
	FirstName string
	LastName  string
}

type Customer struct {
	CustomerId   int
	CustomerName string
}

func (u *User) Name() string {
	return fmt.Sprintf("%s, %s", u.FirstName, u.LastName)
}

func (c *Customer) Name() string {
	return fmt.Sprintf("%s", c.CustomerName)
}

type Namer interface {
	Name() string
}

func Wish(n Namer) string {
	return fmt.Sprintf("Hello, %s Good Morning", n.Name())
}

func main() {
	usr := &User{
		FirstName: "Sanu",
		LastName:  "Yadav",
	}
	fmt.Println(Wish(usr))

	cust := &Customer{
		CustomerId:   123,
		CustomerName: "Sanuky",
	}
	fmt.Println(Wish(cust))
}
