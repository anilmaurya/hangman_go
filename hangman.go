package main

import "fmt"
import "strings"

func main(){
  word := "elephant"
  placeholder := make([]string, len(word))
  chances := 6
  fmt.Println("Guess ",  len(word), " char string")

  for chances > 0 && strings.Compare( strings.Join(placeholder, ""), word) == -1{
    var input string
    fmt.Scanln(&input)
    found := false
    for i, v := range strings.Split(word, ""){
      if input == v{
        placeholder[i] = v
        found = true
      }
    }
    fmt.Println(placeholder)
    if !found{
      chances  = chances - 1
    }
    found = false
  }

  if strings.Compare( strings.Join(placeholder, ""), word) == 0{
    fmt.Println("Won")
  }else{
    fmt.Println("Loss")
  }

}
