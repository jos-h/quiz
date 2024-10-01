# Go Quiz App

A simple command-line quiz application written in Go, which reads questions and answers from a CSV file, asks questions to the user one at a time and enforces time limit for each question and the entire quiz. The program randomly picks the question from the CSV file. The same question is not being repeated again.

# Features
- Read questions and answers from the CSV file
- User has been give a configurable amout of time to answer the question
- Time limit has been enforced on the entire quiz
- Final score has been displayed at the end of the quiz / the quiz limit has been expired

# Installation

To use the quiz application, ensure you must have installed [GO](https://go.dev/) on your system
1. Clone the repository using
   - **git clone https://github.com/jos-h/quiz.git**
2. Navigate the directory
   - **cd quiz**
3. Build the project
   - **go build**

# Usage
** Command-Line Options **
- ```-csv```: This flag allows you to specify the CSV file containing the quiz questions and answers. The CSV file should have two columns: the first column should be the question, and the second column is the answer
   ```
     Example:
       5+5, 10
       4+3, 7
   ```
- Example command
    ``` ./quiz -csv=problems.csv```

# Quiz Structure
- After starting the quiz, the app will display each question one at a time.
- Users will have a limited time (e.g., 10 seconds per question) to answer each question.
- The quiz will also have a total time limit (e.g., 25 seconds) for all questions.
- If the user exceeds the time limit for a question, the app will move on to the next one.
- At the end of the quiz, the app displays the user score.

# Example

![image](https://github.com/user-attachments/assets/b37d9467-cb6c-42e0-ae14-8d01409d410e)

If time limit is exceeded for a question or for the whole quiz

  ![image](https://github.com/user-attachments/assets/d91cf00c-9e4c-4811-a0d5-530046f1ad20)

# Configuration
You can adjust the time limits by modifying constants in the Go code:

```
TIME_TO_ANS_IN_SECONDS: Time allowed for each question.
TOTAL_QUIZ_TIME_IN_SECONDS: Total time allowed for the entire quiz.
```

# Contributing
Feel free to submit issues or pull requests to improve the app. Contributions are welcome!

# Enhancement
- The configuration can be done via command-line arguments by making use of flag
