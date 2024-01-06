package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)

func main() {
	csvFileName := flag.String("csv","problems.csv","A csv file for questions")
	flag.Parse()
	
	
	file,err := os.Open(*csvFileName)
	if err != nil{
		exit(fmt.Sprintf("Failed to open the csv file %s \n",*csvFileName))
		
	}

	r := csv.NewReader(file)
	lines,err := r.ReadAll()
	if err != nil{
		exit(fmt.Sprintf("CSV IS Invalid "))
	}
	problems := readLines(lines)
	
	cnt := 0
	for i,p := range problems{
		fmt.Printf("The Question number %d is %s = ",i+1,p.q)

		var ans string 
		
		fmt.Scanf("%s\n",&ans)

		if ans == p.a{
			fmt.Printf("Corrent Answer \n")
			cnt++
		}
	
	}

	fmt.Println("You Scored ",cnt)



}

func readLines(lines [][]string) []problem {
	ret := make([]problem, len(lines))

    for i, line := range lines{
	ret[i] = problem{
		q : line[0],
		a : line[1],
	}
}

	return ret
}

type problem struct {
	q string
	a string
}


func exit(msg string){
	fmt.Println(msg)
	os.Exit(1)
}