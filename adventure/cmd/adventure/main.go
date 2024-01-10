package main

import (
	"fmt"
	"flag"
	"os"
	"D:\Projects\Go-Learning\golang_mini_projects\adventure\"
	
	
)

func main() {
	filename := flag.String("file","gophers.json","The Adenture Book Json File")
	flag.Parse()
	fmt.Println(*filename)

	f, err := os.Open(*filename)
	if err != nil {
		panic(err)
	}

	d := json.Decoder(f)
	var story adventure.Story
	if err := d.decode(&story); err != nil{
		panic(err)
	}

	fmt.Printf("%v\n",story)
}