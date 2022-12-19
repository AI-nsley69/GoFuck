package interpreter

import (
    "trainsley69.me/GoFuck/parser"
    
    "fmt"
    "reflect"
)

var stack = [65665]int{}
var stackIdx int = 0
var output string = ""

func Run(nodes []parser.Node) string {
    Interpret(nodes)
    return output
}

func Interpret(nodes []parser.Node) {
    for _, node := range nodes {
        if len(output) > 0 && rune(output[len(output) - 1]) == '\n' {
            fmt.Printf(output)
            output = ""
        }
        switch node.(type) {
            case parser.PlusNode: stack[stackIdx] += int(reflect.ValueOf(node).Int())
            case parser.PointerNode: stackIdx += int(reflect.ValueOf(node).Int())

            case parser.PrintNode: output += string(rune(stack[stackIdx]))
            case parser.InputNode: continue;

            case parser.LoopNode:
                var newLoop []parser.Node
                for _, n := range reflect.ValueOf(node).Interface().(parser.LoopNode) {
                    newLoop = append(newLoop, n)
                }
                for stack[stackIdx] > 0 {
                    Interpret(newLoop)
                }
           default: continue
        }
    }
}

