package main

import (
  "bufio"
  "fmt"
  "os"
)

type cliCommand struct {
  name string
  description string
  callback func() error
}

var commandSet = map[string]cliCommand{
  "help": {
    name:        "help",
    description: "Displays a help message",
    callback:    commandHelp,
  },
  "version": {
    name:        "version",
    description: "Pokedex version",
    callback:    commandVersion,
  },
  "exit": {
    name:        "exit",
    description: "exit",
    callback:    func() error{return nil},
  },
}

func printPrompt() error {
  fmt.Print("pokedex > ")
  return nil
}

func commandVersion() error {
  fmt.Println("version 0.1")
  return nil
}

func commandHelp() error {
  fmt.Println("available commands: [version, help, exit]")
  return nil
}

func main() {
    

    reader := bufio.NewScanner(os.Stdin)
    printPrompt()
replloop:
    for reader.Scan() {
      command := reader.Text()
      switch commandSet[command].name {
      case "exit":
        break replloop
      case "help":
        commandHelp()
      case "version":
        commandVersion()
      default:
        commandHelp()
      }
      printPrompt()
    }
}

