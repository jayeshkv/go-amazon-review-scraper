package main

import (
"github.com/PuerkitoBio/goquery"
"bytes" //for appending url
"fmt"
"log"
)
var(
prdUrl bytes.Buffer //holds the temp url
prdId string //gets product id from user
)
func main() {
	fmt.Print("Enter the product ID: ")
	fmt.Scanln(&prdId)
	prdUrl.WriteString(prdId)
	doc, err := goquery.NewDocument(prdUrl.String())
	if err != nil {
    	log.Fatal(err)
	}else {


	doc.Find(".a-section .review").Each(func(i int, s *goquery.Selection){
	author:=s.Find(".author").Text()
	rating:=s.Find(".a-icon-alt").Text()
	review:=s.Find(".review-text").Text()
	fmt.Printf("Review# %d \n Author: %s \n Rating %s \n Review: %s \n",i, author, rating, review)
	})
	}
}

func init() {
prdUrl.WriteString("http://amazon.com/review/")
}