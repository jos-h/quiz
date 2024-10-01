package main

import (
	"flag"
	q "github.com/Demos/CSVFiles/demo"
)

func main(){
	csvFilename := flag.String("csv", "problems.csv", "A csv file in the format of question and answers")
	flag.Parse()
	file := *csvFilename
	q.Quiz(file)
}