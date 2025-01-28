package main

import "fmt"

/*
var (
	a, b, c *string
	d       *int
)

func init() {
	p := "Drop by drop, the bucket gets filled."
	q := "Persistently, patiently, you are bound to succeed."
	r := "The meaning of life is ..."
	n := 42
	a = &p
	b = &q
	c = &r
	d = &n
}

type dog struct {
	first string
}

func (d *dog) walk() {
	fmt.Println("My name is", d.first, "and I'm walking.")
}

func (d *dog) run() {
	fmt.Println("My name is", d.first, "and I'm running.")
}

type youngin interface {
	walk()
	run()
}

func youngRun(y youngin) {
	y.walk()
	y.run()
}

// change the name of the dog with though the adress of the pointer
func (d *dog) rename(n string) {
	d.first = n
}
*/

type person struct {
	first string
}

func main() {
	/*
		fmt.Printf("%T \t %T \t %T \t %T \n\n", a, b, c, d)
		fmt.Printf("%v \t %v \t %v \t %v \n", *a, *b, *c, *d)

		d1 := dog{"Henry"}
		d2 := dog{"Padget"}
		youngRun(&d1)
		youngRun(&d2)

		d1.rename("peter")

		youngRun(&d1)
	*/
	p1 := person{
		"peter",
	}

	p2 := person{
		"Harald",
	}

	fmt.Println(p1)
	fmt.Println(p2)

	p1 = p1.changeNameValue("Uwe")
	p2.changeNamePointer("Hans-Peter")

	fmt.Println("--------------")

	fmt.Println(p1)
	fmt.Println(p2)

}

func (p person) changeNameValue(s string) person {
	p.first = s
	return p
}

func (p *person) changeNamePointer(s string) {
	p.first = s
}
