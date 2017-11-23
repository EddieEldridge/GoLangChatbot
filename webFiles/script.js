


$(function(){
    
 $("#userInput").keypress(function(key) {
    
// Assign the input taken by the user to jquery variable called htmlUserInput
const htmlUserInput = $("#userInput");
const htmlUserList = $("#conversation-list")

// Assign the value from htmlUserInput to userText
const userText = htmlUserInput.val();

            if(key.keyCode == 13) 
            {
                
                        // Stop the page from refreshing
                        event.preventDefault();
            
                    
                        
                        // Clear the text box
                        htmlUserInput.val("");
            
                        htmlUserList.append("<li class='list-group-item list-group-item-primary text-right'>" + "User : " + userText + "</li>");
            }
            
           
       
    });
   
});


