package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(10, 31))
}

// When two or more consecutive named function parameters share a type, you can omit the type from all but the last.
package main

import "fmt"

func add(x, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(11, 4))
}

// multiple results
package main

import "fmt"

func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	a, b := swap("boy", "girl")
	fmt.Println(a, b)
}

// Named return values
package main

import "fmt"

func split(sum int) (x, y int) { // Go's return values may be named. If so, they are treated as variables defined at the top of the function.
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	fmt.Println(split(17))
}