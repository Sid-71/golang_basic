package main

import "fmt"
 

 func add(value1 int ,value2 int ) int {
	return value1 + value2;
 }

  func proAdded(values ...int ) (int,string) {
	sum :=0;
	for _,value := range values {
		sum += value;
	}
	return sum , "mnesage from coder";
  }
func main(){
  result  := add(1,3);
  fmt.Printf(" %v ",result);


  // lets say we don't know how mny values i will enter 
  result2,messgae := proAdded(1,2,3,4,5);
  println(result2,messgae);



}