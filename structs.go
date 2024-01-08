package main

import "fmt"

type  todo struct {
	title  string
	desc  string
}
func main(){
    todo1 := todo{"eating pizza","will do it before night"};


	// here if i don't use printf and used println it will print only values not the keys
	fmt.Printf("%+v\n",todo1);
	fmt.Println(todo1);

	// you don't usere inheritence here 


	// if else training 
	var number int = 212;
	if  (number > 231) {
		fmt.Printf("hellow wordl");
	}else if(number == 212){
		fmt.Printf("tomorow  wordl");
	}else {
	  fmt.Print("sdfsd");
	}
}