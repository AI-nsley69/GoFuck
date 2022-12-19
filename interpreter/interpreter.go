package interpreter

import (
    "trainsley69.me/GoFuck/parser"
    
    "fmt"
    "reflect"
)

var stack = [65665]int{}
var stackIdx int = 0;

func Interpret(nodes []parser.Node) {
    for _, node := range nodes {
        switch node.(type) {
            case parser.PlusNode: stack[stackIdx] += int(reflect.ValueOf(node).Int())
            case parser.PointerNode: stackIdx += int(reflect.ValueOf(node).Int())

            case parser.PrintNode: fmt.Printf("%c", stack[stackIdx])
            case parser.InputNode: continue;

            case parser.LoopNode:
                var newLoop []parser.Node
                for _, n := range reflect.ValueOf(node).Interface().(parser.LoopNode) {
                    newLoop = append(newLoop, n)
                }
                for stack[stackIdx] > 0 {
                    Interpret(newLoop)
                }
        }
    }
}
