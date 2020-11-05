package main

import (
	"fmt"
	"time"

	"github.com/richardbertozzo/getting-started-go/person"
)

func main() {
	fmt.Print("Hello World")

	name := "Richard"
	p := person.Person{
		Name: name,
		Age:  26,
	}
	p.SetAge(30)
	fmt.Println(p)

	go ShowMsg("Oi pessoa: ", p)
	go ShowMsg("Oi teste: ", p)
	ShowMsg("Oi print: ", p)

	time.Sleep(5 * time.Second)
	s1, _ := CreateThings()
	fmt.Println(s1)

	db := person.NewMongo("dajndasjd")
	db.Get("uuid")
}

func ShowMsg(msg string, p person.Person) {
	fmt.Printf("%s %s \n", msg, p.Name)
}

func CreateThings() (string, string) {
	return "things", "picollo"
}
