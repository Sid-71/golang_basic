package main

import(
	"fmt"
)

var doo = 23;
func main(){
	// three types of ways to declare a variables
	// 1 
	var j int = 2;
	var goo = 234;
	i :=2
	
     // if you used colon for assignment it will always set its value by itself.
	 
	 
	 
	// shadowing the variables
	doo = 234;

	var nameTobeGiven string = "Uzumaki naruto"
	fmt.Printf("%v \n",nameTobeGiven);

   // if you used pascal casing . it can be used as an export , if you used camel casing , it will 
   // local to its scope






	fmt.Printf("%v\n",i);
	fmt.Printf("%v\n",j);
	fmt.Printf("%v\n",goo);
	fmt.Printf("%v\n",doo);

}