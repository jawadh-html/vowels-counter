package main


import "fmt"



func main() {

    var input string
    
    fmt.Print("Enter a string: ")

    fmt.Scanln(&input)
    
    count := 0

    for _, char := range input {

        if char == 'a' || char == 'e' || char == 'i' || char == 'o' || char == 'u' ||
           char == 'A' || char == 'E' || char == 'I' || char == 'O' || char == 'U' {
            count++
        }
    }
    
    fmt.Printf("'%s' has %d vowels\n", input, count)
}