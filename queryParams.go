package main

import (
	"fmt"
	"net/url"

)

func main() {
	const myurl = "http://localhost:3000?name=sidharth"
	parsedData ,_ := url.Parse(myurl);
	data := parsedData.Query()
	fmt.Println(data["name"]);
}