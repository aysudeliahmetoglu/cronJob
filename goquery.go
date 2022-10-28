package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/PuerkitoBio/goquery"
)
func main(){
	titles,err := getTitle("https://blog.golang.org")
	if err != nil{
		log.Println(err)
	}
	fmt.Println("Titles:")
	fmt.Printf(titles)


}

func getTitle(url string)(string,error){
	resp,err := http.Get(url)
	//Get HTML
	if err != nil{
		return "",err
	}

	//convert goquery doc
	doc,err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil{
		return "",err
	}

	//create list
	title_list := ""
	doc.Find(".title a").Each(func(i int, s *goquery.Selection){
		title_list += "- "  s.Text() + "\n"
	})

	return title_list,nil

}