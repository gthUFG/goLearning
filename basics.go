// ~ Basic syntax

// Programs in GO are part of packages. You have to declare the package, generally main, then import packages that give you functions.

package main

import (
	"fmt"
	"strconv"
)

// It will have a main function, which will be always executed.

func main() {
	fmt.Println("hello")
	// meetingVars()
	// meetingArrays()
	// conds()
	// fmt.Println(showYourself("Taçarolo", 18))
	fmt.Println(returnMultiple("Maluco", 3))
}

// ~ Variables

func meetingVars() {

	// You can define variables with two ways:

	// With the var keyword, defining the type of the variable
	var var1 = 4
	var var2 int = 5
	var var3 int
	// Note that with var, you may declare variables inside and outside functions. Moreover, you may or may not give the type and value of the variable

	// And with the := sign
	var4 := "hello mademoiselle"
	// This syntax can only be used inside functions and requires a value, because the ":=" tells the reader to infer the type of the variable.

	// You can also declare multiple variables in one line:
	var x, y, z int = 2, 3, 4
	a, b, c := "What's", "Your", "Name"

	// Or in a block:
	var (
		xx int
		yy = 32
		zz = true
	)

	// Const keyword works similar to var. But once it declares constants, you have to declare its value.

	const PI = 3.14

	fmt.Println(var1, var2, var3, var4, x, y, z, a, b, c, xx, yy, zz)

	// ^ Types

	// Boolean
	const c1 bool = true
	// Int
	const c2 int = -412
	// Float32
	const c3 float32 = 2.3
	// String
	const c4 string = "dynamite"

}

// ~ Arrays

func meetingArrays() {
	// Syntax: [length ("..." for inferred length)]datatype{values}
	var array1 = [3]int{1, 2, 3}
	var array2 = [...]bool{true, false}
	array3 := [...]string{"Bom", "dia", "lindeza"}
	fmt.Println(len(array1), array2[1], array3)

	// Note that you can't make the datatype of arrays stay inferred.

	// ^ Slices
	// Slices are like arrays, but flexible. They can change its length. They have, moreover than length, capacity (returned by cap() method), which is the number of elements the slice can grow or shrink to.

	// Syntax:
	// []datatype{values}
	// array[start:end]
	// make([]type, length, capacity)
	var slice1 = []float32{1.2, 3.2, 4.2}
	var slice2 = array3[1:]
	var slice3 = make([]int, 3, 8)
	fmt.Println(slice1, slice2, slice3)

	// ^ Functions

	//? append (add elements to slice)
	// Syntax: append(slice, element) or
	// 		   append(slice1, slice2...)

	fmt.Println(append([]int{12, 34}, 23))

	//? copy (copy elements to other slice)
	// Syntax: append(destination, source)

	// mySlice := make([]int, 5, 5)
	mySlice := make([]int, 2, 4)
	copy(mySlice, array1[:])
	fmt.Println(mySlice, array1)
}

// ~ Conditions

func conds() {
	// Syntax:
	// if condition {
	// } else if condition2 {
	// } else {}

	// Switch syntax
	// switch expression {
	// 		case x: code
	// 		case y: code
	// 		default: code
	// }
}

// ~ Conditions

func showYourself(name string, age int) string {
	// You cannot create functions inside functions.
	// If you want your function to return a value, you must declare its type when creating it.
	return "Hi, my name is " + name + " and i'm " + strconv.Itoa(age) + "."
}

func returnMultiple(x string, y int) (z int, w string) {
	z = y + 4
	w = "Hello " + x
	return
}

// ~ Struct

type Person struct {
	name    string
	age     int
	job     string
	salary  float32
	species string
}

func createPerson(n string, a int, j string, s float32) {
	var person1 Person
	person1.name = n
	person1.age = a
	person1.job = j
	person1.salary = s
	person1.species = "human"
}
