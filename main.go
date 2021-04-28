package main
import "os"
import "fmt"

func main() {
  englishWords := []string{"Zero", "One", "Two", "Three", 
"Four", "Five", "Six", "Seven", "Eight", "Nine"}
  final := ""
  for _, i := range os.Args[1:] {
    for j := 0; j < len(i); j++ {
      if i[j] >= 48 && i[j] < 58 {
		final = final + (englishWords[i[j] - 48])
      } else {
        fmt.Println("There is a non-integer in the input array")
		return
        }
    }
    final = final + ","
  }
  fmt.Println(final[:len(final)-1])
}