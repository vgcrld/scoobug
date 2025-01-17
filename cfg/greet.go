package cfg

import "fmt"

// create a person struct
type Person struct {
	Name string
	Age  int
}

// create a method for the person struct to return the name
func (p Person) GetName() string {
	return p.Name
}

// create a set name
func (p *Person) SetName(name string) {
	p.Name = name
}

// print this struct
func (p Person) Print() {
	fmt.Printf("Name: %s\n", p.Name)
	fmt.Printf("Age: %d\n", p.Age)
}
