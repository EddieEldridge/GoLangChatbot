![golangchatbot](https://user-images.githubusercontent.com/22448079/38645770-2b4248b8-3ddd-11e8-9b5a-7baca873756e.png)


<p align="center">
  <b>A web hosted chatbot built using the Go programming language</b><br>
</p>

# Data Representation and Querying Project Eliza-ChatBot
The following program is an Eliza Chatbot with a personality akin to the Hal9000 AI from the Stanley Kubrick film '2001: A Space Odyssey' (http://www.imdb.com/title/tt0062622/)

#### What is an Eliza Chatbot
ELIZA is an early natural language processing computer program created from 1964 to 1966[1] at the MIT Artificial Intelligence Laboratory by Joseph Weizenbaum.[2] Created to demonstrate the superficiality of communication between man and machine, Eliza simulated conversation by using a 'pattern matching' and substitution methodology that gave users an illusion of understanding on the part of the program, but had no built in framework for contextualizing events.[3] Directives on how to interact were provided by 'scripts', written originally in MAD-Slip, which allowed ELIZA to process user inputs and engage in discourse following the rules and directions of the script. The most famous script, DOCTOR, simulated a Rogerian psychotherapist and used rules, dictated in the script, to respond with non-directional questions to user inputs. As such, ELIZA was one of the first chatterbots, but was also regarded as one of the first programs capable of passing the Turing Test. (https://en.wikipedia.org/wiki/ELIZA)

#### How to run the program
This program uses the Go programming language .

If you do not currently have Go installed click on the following link to download [INSTALL GO](https://golang.org/dl/)

To clone the repository to your local machine, using your prefered command prompt, navigate to the folder you wish to download the files to and enter
```
git clone https://github.com/EddieEldridge/GoLangChatbot
```
There is two ways to run this program
1. **Build and Run**
 Navigate to the Eliza-Chatbot folder and enter the following to compile the code 
```
go build eliza.go
```
This will create a .exe file in your current directory.To launch the server and run the file that is created (note, unless specified, Go will name the .exe after the name of the folder which contrains the eliza.go file. E.g. in my case, my folder is called 'ElizaChatBot' so an .exe called 'ElizaChatBot.exe' is created)
```
ElizaChatBot.exe
```
2. **Run** to simply run the program in your command prompt enter the following 
```
go run eliza.go
```  
The webpage will now be served to the local host. To view the page, open you browser and enter to the local address
```
127.0.0.1:8080
```
Or alternatively you can enter
```
localhost:8080
```
### The Developement Process
I developed the program over the course of a couple of weeks. Me and other classmates worked on the project together on numerous times, helping each other accomplish any difficult tasks or obstacles we might indiviually be stumped on. It was an interesting
experience to code an actual program using the Go programming language and learning how regular expressions work. I also enjoyed designing the layout of the website and was greatly helped by Bootstrap (https://getbootstrap.com/)

### Design
I decided to use the AI from 2001: A Space Odyssey as a personality as both Eliza and the film were produced in the 1960's. Eliza was one of the very first chatbots created and one of the first AI's to pass the Turing test so I thought it appropriate to use a fictional AI from a movie created only 2 years after Eliza's creation.

### Added Features
Try asking Hal about the following topics (not case-sensitive)
* Eliza
* Jesus
* Religion
* God
* College
* Hal 
* Language
* Sports
* Music
* Family
* Coding
* Jokes
* Jobs

### References
https://data-representation.github.io/

http://www.manifestation.com/neurotoys/eliza.php3

https://www.smallsurething.com/implementing-the-famous-eliza-chatbot-in-python/


