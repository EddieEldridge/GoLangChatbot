


$(function(){
    
 $("#userInput").keypress(function(key) {
    
// Assign the input taken by the user to jquery variable called htmlUserInput
const bigList = $("#list-group")
const bigInput = $("#userInput");

// Assign the value from htmlUserInput to userText
const userText = bigInput.val();

            if(key.keyCode == 13) 
            {
                
                        // Stop the page from refreshing
                        event.preventDefault();
            
                        // Test
                        console.log(userText);
                      
                        // Clear the text box
                        bigInput.val("");

                        // Append the unordered list in our HTML with OUR response containted with a HTML list element
                        $("ul").append("<li class='list-group-item list-group-item-success text-left'>" + "User : " + userText + "</li>");
            }
            
           
       
    });
   
});


