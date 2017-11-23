

// Shortened version of the document.ready jquery function
$(function(){

// Function that executes dependant on certain keypresses
 $("#userInput").keypress(function(key) {
    
// Assign the input taken by the user to jquery variable called bigInput
const bigList = $("#list-group");
const bigInput = $("#userInput");

// Assign the value from htmlUserInput to a jQuery variable called userText
const userText = bigInput.val();

            // If the user presses the enter key(keycode(13)), execute the following code
            if(key.keyCode == 13) 
            {
                
                        // Stop the page from refreshing
                        event.preventDefault();
            
                        // Test
                        console.log(userText);
                      
                        // Clear the text box
                        bigInput.val(" ");

                        // Append the unordered list in our HTML with OUR response containted with a HTML list element
                      bigList.append("<li class='list-group-item list-group-item-success text-left'>" + "User : " + userText + "</li>");
                       
            }
            
            // If the user enters another key that's not enter, let them try again
            else
            {
                return;
            }


            // Passing 
            const queryParams = {
                "userInput" : userText
            }
            $.get("/chat", queryParams)
            .done(function(resp) {
                const nextListItem = "<li class='list-group-item list-group-item-warning text-right'>" + "Eliza : " + resp + "</li>";
                bigInput.append(nextListItem)                
            }
            
            
           
       
    });
   
});


