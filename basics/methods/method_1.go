package main

import "fmt"

type User struct {
	Name  string
	Email string
}

type Player struct {
	User
	GamerId int
}

func (p *Player) Greetings() string {
	return fmt.Sprintf("GamerID: %d, Email: %s", p.GamerId, p.Email)
}

func (U *User) Greetings() string {
	return fmt.Sprintf("Name: %s, email %s", U.Name, U.Email)
}

func main() {
	p := Player{}
	p.Name = "Ram"
	p.Email = "ram@gm.com"
	p.GamerId = 123
	fmt.Println("resp:", p.Greetings())

	u := User{}
	u.Name = "saammm"
	u.Email = "sam@mm.com"
	fmt.Println("resp:", u.Greetings())
}
