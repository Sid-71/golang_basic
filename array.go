package main

import (
	"fmt"
)

func main() {

	var amounts [3]int = [3]int{10, 20, 30}
	fmt.Printf("amounts : %v\n", amounts)

	amounts2 := [3]int{20, 204, 24}
	fmt.Printf("%v\n", amounts2);

   // for printing length of an array
   fmt.Printf("%v\n",len(amounts2));


   // suppose if you don't know what is the size of an array
   // amounts2 := [...]int{20, 204, 24,.....}
   //in above comment you can easily know the size by this ... will give you dynamic size 


    amounts3 := &amounts2;
	fmt.Printf("%v the same refrenceing \n",amounts3);


	// suppose  you want to copy all elements of an array
	a := [...]int{100,200,300};
	b := a[:] // this will add all elemtns of array a into b 
	fmt.Printf("%v\n",b);


    // understading the different kinds of having an array 
	c  := a[2:]; // staring from 2 index and reaching to the end
	d :=a[ : 2]; // starting from 0 index and going to end
	e := a[2 : 3]; // starting from index 2 and going to 3 index but inlcuding that 

   fmt.Print(c,d,e);



   
}
