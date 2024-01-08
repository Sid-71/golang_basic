package main

import "fmt"


func main(){
   frequence := make(map[string]string);
   frequence["hello"] = "world";
   frequence["iit"] = "dhawan";
   frequence["mudit"] = "dhawan";
   frequence["dishu"] = "dhawan";


   fmt.Println(frequence["hello"]);
   fmt.Println(frequence);
 
   // to delete particular key 
   delete(frequence,"iit");

   fmt.Println(frequence);


   // to loop in a map
   for key,value := range frequence {
	fmt.Println(key,value);
   }

   

   
}