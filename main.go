package main

import (
  "fmt"
  "os"
  "os/user"
  "monkey/repl"
)

func main() {
  user, err := user.Current()
  if err != nil {
    panic(err)
  }
  fmt.Printf("Monkey REPL v1. User: %s \n", user.Username)
  repl.Start(os.Stdin, os.Stdout)
}


