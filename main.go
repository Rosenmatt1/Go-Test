package main

import ( 
	"fmt"
	// "errors"
	// "math"
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

//  func main() {
// 	vertices := make(map[string]int)
// //Creates key value pairs
// 	vertices["triangle"] = 2
// 	vertices["square"] = 3
// 	vertices["dodecagon"] = 12

// 	delete(vertices, "square")

// 	fmt.Println(vertices)   //shows whole array
// 	// fmt.Println(vertices["triangle"])  //for a key
//  }

//for loop
// func main() {
// 	// a := []int{5,4,3,2,1}
// 	for i := 0; i < 5; i++ {
// 		fmt.Println(i)
// 	}
//  }

//while loop
// func main() {
// 	i := 0
// 	// a := []int{5,4,3,2,1}
// 	for i < 5 {
// 		fmt.Println(i)
// 		i++
// 	}
//  }

// loop though array
// func main() {
// 	arr := []string{"a", "b", "c"}

// 	for index, value := range arr {
// 		fmt.Println("index", index, "value", value)
// 	}
//  }

//loop through map
// func main() {
// 	m := make(map[string]string)
// 	m["a"] = "alpha"
// 	m["b"] = "beta"
// 	fmt.Println(m)

// 	for key, value := range m {
// 		fmt.Println("key", key, "value", value)
// 	}
//  }

//function with multiple return options
// func main() {
// 	result, err := sqrt(16)

// 	if err != nil {
// 		fmt.Println(err)
// 	} else {
// 		fmt.Println(result)
// 	}
// }
// //new function called sum
// func sqrt(x float64) (float64, error) {
// 	if x < 0 {
// 		return 0, errors.New("Undefined for negative numbers")
// 	}

// 	return math.Sqrt(x), nil
// }

//Creating a struct

// type person struct {
// 	name string
// 	age int
// }

// func main() {
// 	p := person{name: "Jake", age: 23}
// 	fmt.Println(p.age)
// }

//Pointer to i
func main() {
	i := 7
	inc(&i)
	fmt.Println(i)
}
func inc(x *int) {
	*x++
}
