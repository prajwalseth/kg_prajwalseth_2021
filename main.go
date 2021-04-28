package main
import "os"
import "fmt"

func main() {
  englishWords := []string{"Zero", "One", "Two", "Three", 
"Four", "Five", "Six", "Seven", "Eight", "Nine"} // create a string array for each digit
  final := ""
  for _, i := range os.Args[1:] { // iterate over inputs
    for j := 0; j < len(i); j++ { // iterate over characters in each input integer
      if i[j] >= 48 && i[j] < 58 { // check if character is an integer
		final += (englishWords[i[j] - 48]) // append appropriate string from englishWords for each digit
      } else {
        fmt.Println("There is a non-integer in the input array") // error message
		return
        }
    }
    final += "," // concatenate strings with a comma
  }
  fmt.Println(final[:len(final)-1]) // print out the concatenated string minus the final comma
}