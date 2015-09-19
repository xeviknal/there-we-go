package main

import "fmt"

func main() {
  jenkins("I want you NOW")

  go jenkins("Wen you want, come")
  go func(message string) {
    for i := 0; i < 10; i ++ {
      fmt.Println(message, ", I am the unnamed emancipated children no. ", i)
    }
  }("Jenkins")

  go jenkins("I am lost in a thread")

  var input string
  fmt.Scanln(&input)
  fmt.Println(input)
}

func jenkins(message string) {
  fmt.Println("Did you say: ", message)
}
