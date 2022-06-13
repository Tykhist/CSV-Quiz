package main

import (
	"fmt"
	"csv"
)

/*
func quesAns(ques string, ans string) map[string]string{
	quiz := map[string]string{}

}
*/

func quiz(csv file, ans bool) {
	/*
		Function will take a 2-column CSV file input as a quiz, the
		columns being for a question and the answer respectively.
		Then it will use the question in a string, asking for user
		input. If the input matches the answer, a counter will be
		incremented for correct. Otherwise, the counter for incorrect
		will go up. At the end of the question-asking loop, the amount
		of correct and incorrect answers will be printed in a string.
		There is also an argument for deciding if the answer is
		displayed after every question.
	*/

	/*
		csvLines := []string{csv.readlines()} // Breaks CSV into lines and puts those into an array

		quiz := map[string]string{} // Creates a map for question/answer pairs

		for i := 0; i < len(csvLines); i++ {
			// Splits columns into question/answer pairs and adds them to quiz map
			question, answer := csvLines[i].split()
			quiz[question] = answer
		}

		// Counters to display test score
		correct := 0
		incorrect := 0
		score := (correct / (correct + incorrect)) * 100

		for i := 0; i < len(quiz); i++ {
			// Displays every question
			fmt.Printf("Question %d: %s\n", i + 1 , quiz[i])
			fmt.Scan(&userAns)

			// Adds to counters, with option to display after every question
			if userAns == quiz[quiz[i]] {
				correct += 1
				if ans {
					fmt.Println("Correct!")
				}
			} else {
				incorrect += 1
				if ans {
					fmt.Println("Incorrect...")
				}
			}

			fmt.Println()
		}

		// Final results
		fmt.Printf("Correct: %d  Incorrect: %d\n", correct, incorrect)
		fmt.Printf("You scored a %d%%\n", score)

		switch score {
			case score >= 99:
				fmt.Println("⁂!WOOOOOWWW!⁂")
			case score >= 90:
				fmt.Println("Wow!!!")
			case score >= 80:
				fmt.Println("Wow")
			case score >= 70:
				fmt.Println("Wow?")
			default:
				fmt.Println("Wow....")
		fmt.Println()
		}

	*/
}

func main() {
	/*
		fmt.Scan(&quizFile)
		fmt.Scan(&answerDisplay)
		quiz(quizFile, answerDisplay)

		fmt.Println("Goodbyeee (ﾉ◕ヮ◕)ﾉ*:･ﾟ✧")
	*/
}
