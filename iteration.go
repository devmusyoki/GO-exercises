package main

import "fmt"

func main() {
	i := 1
	for i <= 10 {
		fmt.Println(i)
		i = i + 1
	}
}

// First we create a variable i, initialize it to 1 
// create an expression such that i equals or is less than 10
// Iterate and print i only when it is less than 10
//  i should increment by 1

//ALTERNATIVE

package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
         fmt.Println(i)
	}
}
