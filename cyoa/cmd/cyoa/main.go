package main

import (
	"fmt"
	"flag"
	"os"
	"github.com/samsaksham/golang_projects/cyoa/story"
	"encoding/json"
	
	
	
)

func main() {
	filename := flag.String("file","gopher.json","The Adenture Book Json File")
	flag.Parse()
	fmt.Println(*filename)

	f, err := os.Open(*filename)
	if err != nil {
		panic(err)
	}

	d := json.NewDecoder(f)
	var story cyoa.Story
	if err := d.Decode(&story); err != nil{
		panic(err)
	}

	fmt.Printf("%v\n",story)
}