package main

// packages
import (
	"fmt"
	"strconv"
)

// variables declarations
var (
	j int = 42
	i int = 32
)

func main() {
	// variables declarations
	// var k int = 98
	// l := 20
	// var theURL string = "https://google.com"

	fmt.Printf("%v, %T\n", j, j)

	// convert types
	var x float32
	x = float32(j)
	fmt.Printf("%v, %T\n", x, x)

	// convert to string
	var k string
	k = strconv.Itoa(j)
	fmt.Printf("%v, %T\n", k, k)
}
