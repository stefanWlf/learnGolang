package main

import "fmt"

func main() {
	// adding all kind of stuff inside a slice
	xs := []string{"Almond Biscotti Café", "Banana Pudding ", "Balsamic Strawberry (GF)", "Bittersweet Chocolate (GF)", "Blueberry Pancake (GF)", "Bourbon Turtle (GF)", "Browned Butter Cookie Dough", "Chocolate Covered Black Cherry (GF)", "Chocolate Fudge Brownie", "Chocolate Peanut Butter (GF)", "Coffee (GF)", "Cookies & Cream", "Fresh Basil (GF)", "Garden Mint Chip (GF)", "Lavender Lemon Honey (GF)", "Lemon Bar", "Madagascar Vanilla (GF)", "Matcha (GF)", "Midnight Chocolate Sorbet (GF, V)", "Non-Dairy Chocolate Peanut Butter (GF, V)", "Non-Dairy Coconut Matcha (GF, V)", "Non-Dairy Sunbutter Cinnamon (GF, V)", "Orange Cream (GF) ", "Peanut Butter Cookie Dough", "Raspberry Sorbet (GF, V)", "Salty Caramel (GF)", "Slate Slate Different", "Strawberry Lemonade Sorbet (GF, V)", "Vanilla Caramel Blondie", "Vietnamese Cinnamon (GF)", "Wolverine Tracks (GF)"}

	// printing out the index and slice content
	for i, v := range xs {
		fmt.Printf("Iteration %v \t Ice Cream flavour: %v\n", i, v)
	}

	fmt.Println("-----------")
	// slicing something off a slice

	// [inclusive:exclusive]
	fmt.Println(xs[0:5])

	fmt.Println("-----------")

	// [:exclusive]
	fmt.Println(xs[:2])

	fmt.Println("-----------")

	// [inclusive:]
	fmt.Println(xs[28:])

	fmt.Println("-----------")

	// delete something from a slice
	test := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	test = append(test[:4], test[5:]...)

	fmt.Println(test) // expected outout: [0 1 2 3 5 6 7 8 9]

	//works like this: You append the slice with a certain cap to it, this cap is exclusive and WILL NOT get moved over, then just add the other wanted values as an argument
}
