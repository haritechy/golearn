package controller

import (
	"fmt"
	"strings"
)

func Arrays() {

	maping()
	var myname [5]string
	myname[0] = "hari"
	myname[1] = "haran"
	artrsy := []string{}
	araysName := append(artrsy, "kjdfgf", "jhj", "yty7gfd")
	for _, name := range araysName {
		// fmt.Printf("%s", name)

		fmt.Println(name)
	}

	firstName := myname[0]
	lastName := myname[1]
	fullName := firstName + " " + lastName
	// fmt.Printf("%s", araysName)
	fmt.Printf("FullName  is : %s\n", fullName)

	var a int = 100

	var n int = 100
	m := 10
	n++
	m++

	x := &a

	for i := 0; i <= 10; i++ {
		a++
		fmt.Println(a)

	}

	if a == n {

		fmt.Println("trus")
		return
	}
	fmt.Print(false)

	fmt.Println(*x)
	fmt.Printf("inremet operator ++ %v\n", (n == a))
	fmt.Printf("%v\n", (a == n) && (a > n))

	type company struct {
		manage    string
		software  string
		developer string
	}

	emlpoyee := company{manage: "hariharan", software: "welcome", developer: "hui"}
	slice := strings.Split(emlpoyee.manage[0:5], " ")

	fmt.Println(slice)
	fmt.Println(emlpoyee)

	/*mapping*/

	employeemapping := map[string]string{
		"manage":   emlpoyee.manage,
		"develope": emlpoyee.developer,

		"software": emlpoyee.software,
	}
	delete(employeemapping, "develope")
	fmt.Printf("enter your mapping  key\n")

	// switch city {
	// case "london":
	// 	fmt.Printf("Welcome to the city of %v in the UK\n", city)
	// case "paris":
	// 	fmt.Printf("Welcome to the city of %v in France\n", city)
	// case "sydney":
	// 	fmt.Printf("Welcome to the city of %v in Australia\n", city)
	// case "mumbai":
	// 	fmt.Printf("Welcome to the city of %v in India\n", city)
	// default:
	// 	fmt.Printf("Your entered city %v is not found in our data\n", city)
	// }

}

func maping() {

	var key int

	fmt.Print("Enter your choice")
	fmt.Scan(&key)
	fmt.Println("yoour number is ", key)
}
