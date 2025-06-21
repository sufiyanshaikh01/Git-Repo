package main

import (
	"bufio"
	"fmt"
	"os"
)

func SimpleFuncation() {
	fmt.Println("We are Starting funcation for Go language")

}

func add(a, b int) (result int) {

	result = a + b
	return

}

func multiplay(c, d int) (result int) {

	result = c * d
	return
}

func divide(x, y float64) (float64, error) {
	if y == 0 {
		return 0, fmt.Errorf("denometer is not zero")
	}
	return x / y, nil

}

func main() {

	fmt.Println("Hello world!!")

	fmt.Println("Hey,Whaht's Your name is:")
	reader := bufio.NewReader(os.Stdin)
	name, _ := reader.ReadString('\n')
	fmt.Println("HELLO MR:", name)

	var Name string = "sufiyan"
	fmt.Println("NAME: ", Name)

	age := 19
	height := 6.2
	person := "MUmbai"

	fmt.Printf("Age %d\nHeight %f\nPerson %s\n", age, height, person)

	ans1 := add(3, 4)
	ans2 := multiplay(5, 4)

	fmt.Println("add of two number is:", ans1)
	fmt.Println("multiply of two number is: ", ans2)

	fmt.Println("starting Error Handing for Go language")

	ans3, _ := divide(10, 4)
	fmt.Println("division of two number is:", ans3)

	fmt.Println("We are Starting of Array in Go language")

	var namee = [3]string{"sufiyan", "shaikh", "saeed"}

	fmt.Println("Name of the person:", namee)
	fmt.Println("Length of the array is:", namee)
	fmt.Println("value of namee at 0 index is: ", namee[0])
	fmt.Println("value of name at 1 index is:", namee[1])
	fmt.Println("value of name at 2 index is:", namee[2])

	var price [5]int
	fmt.Println("price is:", price)

	var price1 [5]bool
	fmt.Println("price is:", price1)

	var price2 [5]string
	fmt.Printf("price is %q\n", price2)

	number := []int{1, 2, 3, 4, 5}
	number = append(number, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15)
	fmt.Println("Number :", number)
	fmt.Printf("Number has data type : %T\n", number)
	fmt.Println("length of number: ", len(number))

	num := make([]int, 3, 5)

	fmt.Println("slice: ", num)
	fmt.Println("length: ", len(num))
	fmt.Println("cpacity: ", cap(num))

}
