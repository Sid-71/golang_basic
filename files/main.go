package main

import (
	"fmt"
	"io"
	"os"
)


// panic will stop the exection and will show you the error

func main() {
	content := "mangekyo sharingan utlimate genjustu";

	// this iwll create the file
    file,err := os.Create("./data.txt");

	// this is the way to handle the error
	if(err != nil){
		panic(err);
	}
	// this is the way to add content on the file you have created
	len , err := io.WriteString(file,content);
	if(err != nil){
		panic(err);
	}
	fmt.Println(len);
	defer file.Close();
	readFile("./data.txt");
}





func readFile(fileName string) {
  dataInByte,err :=	os.ReadFile(fileName);
  if(err != nil){
	return;
  }
  // but the data which will be coming will be in bytes
  fmt.Printf("the content of string will be :%v",string(dataInByte));

}