package main

import "fmt"

type Person struct {

	Name string
	Age int
}

func (p *Person) String() string {
	return fmt.Sprintf("%s is %d", p.Name, p.Age)
}

func main() {
	p := &Person{"Ardeshir" , 44}

	fmt.Println(p)
}

