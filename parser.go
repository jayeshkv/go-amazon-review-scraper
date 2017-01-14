package main

import (
"bytes" //for appending url
"fmt"
)
var(
prdUrl bytes.Buffer //holds the temp url
prdId string //gets product id from user
)
func main() {
	fmt.Print("Enter the product ID: ")
	fmt.Scanln(&prdId)
	prdUrl.WriteString(prdId)
	fmt.Println(prdUrl.String())
	
}
func init() {
prdUrl.WriteString("http://amazon.com/review/")
}