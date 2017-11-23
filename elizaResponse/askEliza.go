package askEliza

import (
	"math/rand"
	"regexp"
	"strings"
	"time"
)

// When naming our function, we ensure it begins with an upppercase character
// so that it can be exported correctly to our Eliza.go file
func ElizaResponseFunc(inputString string) string {

	// Set a seed for our RNG using the current time in Nanoseconds
	rand.Seed(time.Now().UTC().UnixNano())

	// Assign our userInput to a string variable for comparison
	userInput := inputString

	// Create an array of strings that
	matches := []string{
		"(.*)hello(.*)",
	}

	outputs := []string{
		"Hi, how are things!?=Well, how are you my friend?",
	}

	//Regular Expressions
	for counter, _ := range matches {

		// Assign
		patternToMatch := regexp.MustCompile(matches[counter])

		if patternToMatch.MatchString(userInput) {

			newSolutions := strings.Split(outputs[counter], "=")
			print("Line Matched")
			return newSolutions[rand.Intn(len(newSolutions))]
		}
	}

	//array of random answers
	answers := []string{
		"Alright. Interesting....",
		"How does that make you feel?",
		"Why do you say that?",
		"Huh??",
		"I'm not even going to pretend that I know what you're trying to say...",
	}

	//return answer if input does not match expressions
	return answers[rand.Intn(len(answers))]

}
