package main

import (
	"fmt"
	"sort"
)


func main(){
	// bascially slices are nothing but just a dynamic arrays
	var fruits =[]string {"apples","fruites"};
	fmt.Println(fruits);
	// to add in slices , 
	fruits = append(fruits, "mango","banana");
	fmt.Printf(" %v ",fruits);


	// another  syntax for slices

	highscores := make([]int,2);
	fmt.Println(highscores);

	// to sort an slice 

     b  := []int{123,21,3,};
	 sort.Ints(b);
	 fmt.Println(b);
   


	 // to delete an an element from an array 
	 c := []string{"hello","hi","sdf"};
	 c = append(c, "world");
	  var index int =1;
	 c = append(c[ :index ],c[index+1 :]...);
	 fmt.Println(c);

}




