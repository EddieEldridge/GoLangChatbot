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

	// Create an array of strings that looks for matches of certain keywords using regexp syntax
	matches := []string{
		"(.*)hello(.*)",
		"(.*)good(.*)",
		"(.*)test(.*)",
	}

	// Create an array of strings that compares the matches and outputs results accordingly
	outputs := []string{
		"Hi, how are things!? // How are you my friend? // What would you like to talk about?",
		"My name is Eliza, what's yours?",
		"testing is correct",
	}

	//Regular Expressions
	for counter, _ := range matches {

		// Assign
		patternToMatch := regexp.MustCompile(matches[counter])

		if patternToMatch.MatchString(userInput) {

			newSolutions := strings.Split(outputs[counter], " // ")

			return newSolutions[rand.Intn(len(newSolutions))]
		}
	}

	//array of random answers
	answers := []string{
		"I'll be honest, I have no idea what you're saying.",
		"I wish I was smart enough to understand what you just said..",
		"English, user. Do you speak it?",
	}

	//return answer if input does not match expressions
	return answers[rand.Intn(len(answers))]

}
