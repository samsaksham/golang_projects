package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"time"
)

func main() {
	csvFileName := flag.String("csv", "problems.csv", "A csv file for questions")
	timeLimet := flag.Int("limit", 30, "Time Limit for QUIZ")
	flag.Parse()
	

	file, err := os.Open(*csvFileName)
	if err != nil {
		exit(fmt.Sprintf("Failed to open the csv file %s \n", *csvFileName))

	}

	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		exit(fmt.Sprintf("CSV IS Invalid "))
	}
	problems := readLines(lines)
	timer := time.NewTimer(time.Duration(*timeLimet) * time.Second)

	

	cnt := 0
	for i, p := range problems {
		fmt.Printf("The Question number %d is %s = ", i+1, p.q)
		anschan := make(chan string)

		go func ()  {
			var answer string

			fmt.Scanf("%s\n", &answer)

			anschan <- answer
		}()

		select {
		case <-timer.C:
			fmt.Println("You Scored ", cnt)
			return

		case answer := <- anschan:
			

			if answer == p.a {
				fmt.Printf("\nCorrent Answer \n")
				cnt++
			}


		}
		
		

	}

	

}

func readLines(lines [][]string) []problem {
	ret := make([]problem, len(lines))

	for i, line := range lines {
		ret[i] = problem{
			q: line[0],
			a: line[1],
		}
	}

	return ret
}

type problem struct {
	q string
	a string
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
