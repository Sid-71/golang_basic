package main

import "fmt"

type todo struct {
	title string
	desc  string
}

func (t todo) setData(desc string) {
	t.desc = desc
	println("calling the function ")
	fmt.Println(" %v ", t)
}

func main() {
	todo1 := todo{"eating pizza", "will do it before night"}
	fmt.Println(" %v ", todo1)
	todo1.setData("no more")
	println("after making changfes ")
	fmt.Println(" %v ", todo1)

}



/*
these are the simple methods inside the class , which  you have already seen in java
and they are specificlaly called by call by value as you can see the results.
*/
