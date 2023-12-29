package main
/*
Create a golang program that generates a number between 0 and 10. 
The starting number (0) is not given and thus 0 is assumed as lowest number.
*/

import (
    "fmt"
    "math/rand"
    "time"
)

func random(min int, max int) int {
    return rand.Intn(max-min) + min
}

func main() {
    rand.Seed(time.Now().UnixNano())
    randomNum := random(0, 10)
    fmt.Printf("Random number: %d\n", randomNum)
}