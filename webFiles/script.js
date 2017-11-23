
// Assign the input taken by the user to jquery variable called htmlUserInput
const htmlUserInput = $("#userInput");
const htmlUserList = $("#conversation-list")

 $("#userInput").keypress(function(e) {
    
            console.log("test");

            // If keypress of the enter key(key 13) isn't done, let them enter another key
            if(e.keyCode == 13) {
            
            alert('You pressed enter!');  

            // Stop the page from refreshing
            event.preventDefault();

            // Assign the value from htmlUserInput to userText
            const userText = htmlUserInput.val();
            
            // Clear the text box
            htmlUserInput.val(" ");

            htmlUserList.append("<li class='list-group-item list-group-item-primary text-right'>" + "User : " + userText + "</li>");
        }
    });
   
    
    




