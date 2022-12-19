package main

import (
    "trainsley69.me/GoFuck/tokenizer"
    "trainsley69.me/GoFuck/parser"
    "trainsley69.me/GoFuck/optimizer"
    "trainsley69.me/GoFuck/interpreter"

    "fmt"
    "os"
    "log"
)

func main() {
    // Logging setup
    log.SetPrefix("GoFuck: ")
    log.SetFlags(0)
    // Get the source from specified file
    if len(os.Args) < 2 {
        log.Fatal("Missing file!")
    }
    filePath := os.Args[1]
    source := getSource(filePath)
    // Tokenize the source
    tokens := tokenizer.Tokenize(source)
    // Parse the tokens into nodes
    nodes := parser.Parse(tokens)
    // Optimize the nodes
    nodes = optimizer.Optimize(nodes)
    // Interpret the nodes
    interpreter.Interpret(nodes)
    fmt.Println("")
}

func getSource(source string) string {
    content, err := os.ReadFile(source)
    if err != nil {
        log.Fatal(err)
    }

    return string(content)
}
