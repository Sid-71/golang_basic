package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url  = "https://catfact.ninja/fact?name=sidharth"
func main() {
	response ,err := http.Get(url);
	if(err != nil){
    panic(err);
	}
	// we got the response 

	dataInBytes , err := ioutil.ReadAll(response.Body);
	if(err != nil){
		panic(err);
	}
	wholeData := string(dataInBytes);
	fmt.Printf(wholeData);

   defer response.Body.Close();
   


}