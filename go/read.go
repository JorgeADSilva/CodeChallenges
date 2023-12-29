package main

import (
	"strings"
	"fmt"
  "bufio"
	"os"
)
/*
Write a program which reads information from a file 
and represents it in a slice of structs. Assume that 
there is a text file which contains a series of names. 
Each line of the text file has a first name and a last 
name, in that order, separated by a single space on the 
line. 

Your program will define a name struct which has two 
fields, fname for the first name, and lname for the 
last name. Each field will be a string of size 20 
(characters).

Your program should prompt the user for the name of the 
text file. Your program will successively read each line 
of the text file and create a struct which contains the 
first and last names found in the file. Each struct created 
will be added to a slice, and after all lines have been read 
from the file, your program will have a slice containing one 
struct for each line in the file. After reading all lines from 
the file, your program should iterate through your slice of 
structs and print the first and last names found in each struct.
*/

type Name struct {
  FisrtName string
  LastName string
}

func main() {
  // Prompt the user for the name of the text file
	fmt.Print("Enter the file name with the names: ")
	var fileName string
	fmt.Scanln(&fileName)
	
  file, err := os.Open(fileName)
  if err != nil{
    fmt.Println("Error opening the file: ", err)
    os.Exit(1)
  }

  defer file.Close()

  var nameArray []Name

  scanner := bufio.NewScanner(file)

  for scanner.Scan(){
    line := scanner.Text()
    
    //Split the line into first and last name fields
    nameParts := strings.Fields(line)

    //Create a the Name struct and add it to the slice
    if(len(nameParts)>=2){
      name := Name{
        FisrtName: nameParts[0],
        LastName: nameParts[1],
      }

      nameArray = append(nameArray, name)
    }
  }
  //Check for scanner errors
  if err := scanner.Err(); err != nil{
    fmt.Println("Error reading file: ", err)
    os.Exit(1)
  }

  fmt.Println("Names inside the file provided:")
  for _, n:= range nameArray{
    fmt.Printf("First Name: %-20s Last Name: %-20s\n", n.FisrtName, n.LastName)
  }
}