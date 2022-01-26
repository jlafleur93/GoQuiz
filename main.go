package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)
var correctPoints int = 0;

func main() {
	csvFile := flag.String("csv", "problems.csv", "a csv file in format of ans,que")
	flag.Parse()
	_ = csvFile

	file, err := os.Open(*csvFile)
	if err != nil {
		exitPrint(fmt.Sprintf("Failed to open the CSV file: %s\n", *csvFile ))
	}
	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		exitPrint("failed to parse the csv file")
	}
	problems := parseLines(lines)

	for i, p := range problems {
		fmt.Printf("Problem #%d: %s = \n", i+1, p.q)
		var answer string
		fmt.Scanf("%s \n", &answer)
		if answer == p.a{
			fmt.Println("Correct!")
			correctPoints++;

		}
		if i == 11{
		var wrongAnswers = 12 - correctPoints
		fmt.Println("You have gotten", correctPoints, "Questions correct")
		fmt.Println("You have gotten",wrongAnswers, "questions wrong" )
	}
	}
	
}
func parseLines(lines [][] string)[]problem {
	ret := make([]problem, len(lines))
	for i, line := range lines{
		ret[i] = problem{
			q: line[0],
			a: line[1],
		}

	}
	return ret

}
type problem struct{
	q string
	a string
}

func exitPrint(msg string){
	fmt.Println(msg)
	os.Exit(1)
}