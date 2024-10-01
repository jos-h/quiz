package demo

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"
	"time"
	"math/rand"
)
const (
	TIME_TO_ANS_IN_SECONDS		= 10
	TOTAL_QUIZ_TIME_IN_SECONDS	= 25
)

type problem struct {
	ques	string
	ans		string
}
func exit(msg string){
	fmt.Println(msg)
}

func openFile(file string) *os.File {
	f, err := os.Open(file)
	if err != nil {
		exit(fmt.Sprintf("failed to open csv file: %s\n", file))
	}
	return f
}
func readFile(file *os.File) [][]string {
	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil{
		exit(fmt.Sprintf("Unable to parse csv file: %s\n", file.Name()))
	}
	return lines
}


func parseLines(lines [][]string) []problem {
	problems := make([]problem, 0, len(lines))
	for _, line := range lines {
		if len(line) != 2{
			continue // skip invalid lines
		}
		problems = append(problems, problem{
			ques: line[0],
			ans:  strings.TrimSpace(line[1]),
		})
	}
	return problems
}

func startQuiz(problems []problem) {
	quiztimer := time.NewTimer(TOTAL_QUIZ_TIME_IN_SECONDS * time.Second)
	correct := 0
	askedQues:= make(map[string]int)

	for i, _ := range problems {
		rand.New(rand.NewSource(time.Now().UnixNano()))
		indx := rand.Intn(len(problems))
		prb := problems[indx]
		_, ok := askedQues[prb.ques]
		if !ok {
			fmt.Printf("Problem #%d: %s\n", i + 1, prb.ques)
			quesTimer := time.NewTimer(TIME_TO_ANS_IN_SECONDS * time.Second)

			askedQues[prb.ques]++
			answerCh := make(chan string)

			go func() {
				var answer string
				fmt.Scanln(&answer)
				answerCh <- answer
			}()

			select {
			case answer := <-answerCh:
				if answer == prb.ans {
					correct ++
				}
			case <-quesTimer.C:
				fmt.Println("Time to answer the question is over!!")
			case <-quiztimer.C:
				fmt.Println("Your total quiz time is up!")
				fmt.Printf("You scored %d out of %d \n", correct, len(problems))
				return
			}
		}
	}
	fmt.Printf("You scored %d out of %d \n", correct, len(problems))
}

func Quiz(file string) {
	f := openFile(file)
	lines := readFile(f)
	problems := parseLines(lines)
	startQuiz(problems)
}