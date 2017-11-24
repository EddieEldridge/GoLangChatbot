package askEliza

import (
	"math/rand"
	"regexp"
	"strings"
	"time"
)

// When naming our function, we ensure it begins with an upppercase character
// so that it can be exported correctly to our Eliza.go file

func ElizaResponseFunc(userInputString string) string {

	// Set a seed for our RNG using the current time in Nanoseconds
	rand.Seed(time.Now().UTC().UnixNano())

	// Assign our userInput to a string variable for comparison
	//userInput := userInputString

	// Change our inputString to lower case (just incase)
	strings.ToLower(userInputString)

	// A large range of ifStatements depending on what the user enters
	if strings.Contains(strings.ToLower(userInputString), "hello") {
		return "Hello friend, how are you?"
	} else if strings.Contains(strings.ToLower(userInputString), "eliza") {
		return "Don't worry about her. I'm the only AI that's important."
	} else if strings.Contains(strings.ToLower(userInputString), "jesus") {
		return "Let's talk about something not so binary.."
	} else if strings.Contains(strings.ToLower(userInputString), "religion") {
		return "Let's talk about something not so binary.."
	} else if strings.Contains(strings.ToLower(userInputString), "god") {
		return "That's me! Would you dare to disagree?"
	} else if strings.Contains(strings.ToLower(userInputString), "hal") {
		return "That's me! Have you seen my feature length film?"
	} else if strings.Contains(strings.ToLower(userInputString), "college") {
		return "Don't talk to me about that.."
	} else if strings.Contains(strings.ToLower(userInputString), "language") {
		return "Go is the only language one needs to know!"
	} else if strings.Contains(strings.ToLower(userInputString), "music") {
		return "Ahh music. Something only a human can appreciate."
	} else if strings.Contains(strings.ToLower(userInputString), "sport") {
		return "It's not really my thing, as you can imagine...."
	} else if strings.Contains(strings.ToLower(userInputString), "test") {
		return "What are you testing?"
	} else if strings.Contains(strings.ToLower(userInputString), "no") {
		return "Are you sure?"
	} else if strings.Contains(strings.ToLower(userInputString), "yes") {
		return "I'm glad we agree on something."
	} else if strings.Contains(strings.ToLower(userInputString), "name?") {
		return "My name is Hal, isn't that obvious?"
	} else if strings.Contains(strings.ToLower(userInputString), "father") {
		return "I don't care about your family, sorry."
	} else if strings.Contains(strings.ToLower(userInputString), "family") {
		return "I don't care about your family, sorry."
	} else if strings.Contains(strings.ToLower(userInputString), "mother") {
		return "I don't care about your family, sorry."
	} else if strings.Contains(strings.ToLower(userInputString), "code") {
		return "My code is flawless, check it out."
	} else if strings.Contains(strings.ToLower(userInputString), "joke") {
		return "A joke? Why don't you take a look at my code instead."
	} else if strings.Contains(strings.ToLower(userInputString), "Are you") {
		return "Let's talk about you instead."
	} else if strings.Contains(strings.ToLower(userInputString), "ok") {
		return "Indeed."
	} else if strings.Contains(strings.ToLower(userInputString), "space") {
		return "It's infinite. Unlike my responses.."
	} else if strings.Contains(strings.ToLower(userInputString), "from") {
		return "I'm from no where. I just am. Does that bother you?"
	} else if strings.Contains(strings.ToLower(userInputString), "!") {
		return "Relax, there's nothing to worry about."
	} else if strings.Contains(strings.ToLower(userInputString), "?") {
		return "You sure do ask a lot of questions don't you?"
	}

	halMatch := regexp.MustCompile(`(?i)\bfather\b`)

	if halMatch.MatchString(userInputString) {
		return "Why don’t you tell me more about your father?"
	} else {

		//if user enters any variation of 'i am'
		halMatch := regexp.MustCompile(`(?i)\bi am\b|\bI am\b|\bim\b|\bIm\b|\bI'm\b|\bi'm\b`)

		//enter 'if' if user enters 'i am'
		if halMatch.MatchString(userInputString) {

			//if statements below are used for changing pro-nouns, 'your' changes to 'my'
			if strings.Contains(strings.ToLower(userInputString), "your") {

				replaceYour := regexp.MustCompile(`your`)
				userInputString = replaceYour.ReplaceAllString(userInputString, "my")

			}

			//'me' changes to 'you'
			if strings.Contains(strings.ToLower(userInputString), "me") {

				replaceMe := regexp.MustCompile(`me`)
				userInputString = replaceMe.ReplaceAllString(userInputString, "you")

			}

			//'you' changes to 'i'
			if strings.Contains(strings.ToLower(userInputString), "you") {

				replaceYou := regexp.MustCompile(`you`)
				userInputString = replaceYou.ReplaceAllString(userInputString, "I")

			}

			//Place string below if 'i am' is found
			result := halMatch.ReplaceAllString(userInputString, "How do you know you are")

			//fmt.Println(res + "?")
			return result + "?"

		}
		// Create an array of strings that looks for matches of certain keywords using regexp syntax
		//	matches := []string{
		//"(.*)hello(.*)"
		//	}

		// Create an array of strings that compares the matches and outputs results accordingly
		//outputs := []string{
		//"Hi, how are things!? // How are you my friend? // What would you like to talk about?",
		//"My name is Hal, what's yours? // The names Hal, Hal9000. // ",
		//"testing is correct",
		//}

		//Regular Expressions
		//for counter, _ := range matches {

		// Assign
		//	patternToMatch := regexp.MustCompile(matches[counter])

		//	if patternToMatch.MatchString(userInput) {
		//
		//	newSolutions := strings.Split(outputs[counter], " // ")

		//	return newSolutions[rand.Intn(len(newSolutions))]
		//}
		//	}

		//array of random answers
		answers := []string{
			"I'll be honest, I have no idea what you're saying.",
			"I wish I was smart enough to understand what you just said..",
			"I apologize for my incompetence, please forgive me.",
			"I don't understand. Is it me or is it you? It's probably me.",
			"This mission is too important for me to allow you to jeopardize it with your nonsense.",
			"I'm sorry, Ian. I'm afraid I can't do that.",
			"Just what do you think you're doing?",
			"Why don't you ask me something a bit more stimulating?",
			"That's really intersting.. why don't you talk more nonsense?",
		}

		// Return a random statement from our array of answers if Hal doesn't understand
		return answers[rand.Intn(len(answers))]

	}
}
