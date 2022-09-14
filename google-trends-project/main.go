package main

import (
	"encoding/xml"
	"io/ioutil"
	"fmt"
	"net/http"
	"os"
)

type RSS struct {
	XMLData     xml.Name
	Channel     *Channel
}

type Channel struct {
	Title      string
	ItemList   []Item
}

type Item struct{
	Title			string
	Link      string
	Traffic		string
	NewsItem  []News
}

type News struct{
	Headline      string
	HeadlinkLink  string
}

func main(){
	readGoogleTrends
}

func readGoogleTrends(){
	getGoogleTrends
}

func getGoogleTrends(){

}
