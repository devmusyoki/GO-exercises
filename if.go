package main


import "fmt"

func main() {
	for i := 1; i <= 10; i++{
		if i % 2 == 0 {
			fmt.Println(i, "even")
		}else {
			fmt.Println(i, "odd")
		}
	}
}

Creating a variable called i, that should be less than or equal to 10, and i should be incremental.
create an if function, that would print even if there is no remainder when divided by 2.
and print odd when it's else.
