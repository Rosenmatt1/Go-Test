package main

import ( 
	"fmt"
 )

//  func main() {
// 	 x := 5
// 	 y := 7
// 	 sum := x + y
// 	 fmt.Println(sum)
// 	//  fmt.Println(y)
//  }

// func main() {
// 	x := 5
// 	if x > 6 {
// 		 fmt.Println("More than 6")
// 	} else if x < 2 {
// 		fmt.Println("Less than 2")
// 	} else {
// 		fmt.Println("In Between")
// 	}
//  }

// func main() {
// 	// var a [5]int
// 	// a := [5]int{5,4,3,2,1}
// 	a := []int{5,4,3,2,1}
// 	a = append(a,13)
// 	// a[2] = 7
// 	fmt.Println(a)
//  }

 func main() {
	vertices := make(map[string]int)
//Creates key value pairs
	vertices["triangle"] = 2
	vertices["square"] = 3
	vertices["dodecagon"] = 12

	delete(vertices, "square")

	fmt.Println(vertices)   //shows whole array
	// fmt.Println(vertices["triangle"])  //for a key

 }