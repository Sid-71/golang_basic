package main

import "fmt"

func main(){
	a := []string{"monday","tuesday"};
	// normal for loop
	for  i:=0 ; i<len(a); i++{
		fmt.Println(a[i]);
	}
   // 2 method
	for j := range a{
		fmt.Println(a[j]);
	}
  

	// 3 method 
   // you are going to use this one alot of times 
	for index , value := range a{
		fmt.Println(index,value);
	}
}