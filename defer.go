package main

// import "fmt"


/*
*imagination 
its like in javascirpt all synncrhousn code execute and after that asyncouns 



in an exection of funciton , all lines of code which have defer are pushed into stack of the 
and after that scope executes , then those defer commands are executed 

*/

func main(){
   
	 println("c");
	defer println("golang");
	defer println("rust");
	defer println("java");
	defer println("c++");

	println("c++");

}