package main

import (
	"fmt"
  	"bufio"
	"os"
  	"strings"
)
/*
Write a program which prompts the user to enter a string. 
The program searches through the entered string for the 
characters ‘i’, ‘a’, and ‘n’. The program should print 
“Found!” if the entered string starts with the character 
‘i’, ends with the character ‘n’, and contains the character 
‘a’. The program should print “Not Found!” otherwise. 
The program should not be case-sensitive, so it does not 
matter if the characters are upper-case or lower-case.
*/

func main() {
	var ian string
	var ianFound bool = true
    scanner := bufio.NewScanner(os.Stdin)
	
	fmt.Print("Verify if it's an ian ->")
    scanner.Scan()
    ian = scanner.Text()
	
  	ianLowerStr := strings.ToLower(ian)

	if( ianLowerStr[0] != 'i' || 
		ianLowerStr[len(ianLowerStr)-1] != 'n' || 
		!strings.Contains(ianLowerStr, "a")){
		ianFound = false
	}

	if ianFound{
		fmt.Println("Found!")
	}else{
		fmt.Println("Not Found!")
	}
}