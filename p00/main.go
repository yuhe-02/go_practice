
//  definition of package
package main

import "fmt"

//  definition of function
func main() {
	var sum int

	for i := 1; i <= 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}