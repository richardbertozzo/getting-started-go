package person

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%s - %d", p.Name, p.Age)
}

func (p Person) SetAge(age int) Person {
	p.Age = age
	return p
}

type Database interface {
	Save(Person) error
	Get(id string) (Person, error)
}

type Mysql struct {
	URL string
}

func NewMysql(url string) Database {
	return &Mysql{URL: url}
}

func NewMongo(url string) Database {
	return &Mongo{URL: url}
}

func (m *Mysql) Save(p Person) error {
	return nil
}

func (m *Mysql) Get(id string) (Person, error) {
	return Person{}, nil
}

type Mongo struct {
	URL string
}

func (m *Mongo) Save(p Person) error {
	return nil
}

func (m *Mongo) Get(id string) (Person, error) {
	return Person{}, nil
}
