package main

import "fmt"

func main() {

	createArray(1, 2, 3, 4, 6)
	createSlice(42, 43, 44, 45, 46, 47, 48, 49, 50, 51)
	states := []string{` Alabama`, ` Alaska`, ` Arizona`, ` Arkansas`, ` California`, ` Colorado`, ` Connecticut`, ` Delaware`, ` Florida`, ` Georgia`, ` Hawaii`, ` Idaho`, ` Illinois`, ` Indiana`, ` Iowa`, ` Kansas`, ` Kentucky`, ` Louisiana`, ` Maine`, ` Maryland`, ` Massachusetts`, ` Michigan`, ` Minnesota`, ` Mississippi`, ` Missouri`, ` Montana`, ` Nebraska`, ` Nevada`, ` New Hampshire`, ` New Jersey`, ` New Mexico`, ` New York`, ` North Carolina`, ` North Dakota`, ` Ohio`, ` Oklahoma`, ` Oregon`, ` Pennsylvania`, ` Rhode Island`, ` South Carolina`, ` South Dakota`, ` Tennessee`, ` Texas`, ` Utah`, ` Vermont`, ` Virginia`, ` Washington`, ` West Virginia`, ` Wisconsin`, ` Wyoming`}
	stateSliceAppend(states)
	multidimensionalSlice()
}

func createArray(a, b, c, d, e int) {
	array := [5]int{a, b, c, d, e}

	for i, v := range array {
		fmt.Printf("index: %v \t value: %v \t type %T\n", i, v, v)
	}
	fmt.Println("--------------------")
}

func createSlice(a, b, c, d, e, f, g, h, i, j int) {
	xi := []int{a, b, c, d, e, f, g, h, i, j}
	// create a copy of the slice without pointer pointing at the same background array, for later deletion
	xic := make([]int, 10)
	copy(xic, xi)

	xs := xi[:5]
	ys := xi[5:]
	zs := xi[2:7]
	ts := xi[1:6]

	fmt.Println(xs)
	fmt.Println(ys)
	fmt.Println(zs)
	fmt.Println(ts)

	fmt.Println("--------------")

	xi = append(xi, 52)
	xi = append(xi, 53, 54, 55)
	y := []int{56, 57, 58, 59, 60}
	xi = append(xi, y...)

	fmt.Println(xi)

	fmt.Println("--------------")

	for i, v := range xi {

		fmt.Printf("index: %b \t value: %v \t type %T\n", i, v, v)
	}

	fmt.Println("--------------")

	// print slice bfore deleting a part of it
	fmt.Println(xic)

	// deleting a part of the slice
	xic = append(xic[:3], xic[6:]...)

	// print new slice
	fmt.Println(xic)
}

func stateSliceAppend(states []string) {
	x := make([]string, 0, 50)
	x = append(x, states...)

	fmt.Println(len(x))
	fmt.Println(cap(x))

	for i := 0; i < len(x); i++ {
		fmt.Printf("%v - %v \n", i, x[i])
	}
}

func multidimensionalSlice() {
	x1 := []string{"James", "Bond", "Shaken nit stirred"}
	x2 := []string{"Miss", "Moneypenny", "I'm 008"}

	x := [][]string{x1, x2}

	fmt.Println(x)
	fmt.Println("---------------")

	for i, v := range x {
		fmt.Printf("%v - %v \n", i, v)
		for a, b := range v {
			fmt.Printf("%v - %v \n", a, b)
		}
	}
}
