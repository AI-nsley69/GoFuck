package optimizer

import (
    // "trainsley69.me/GoFuck/tokenizer"
    "trainsley69.me/GoFuck/parser"
    "fmt"
    "reflect"
)

func Optimize(nodes []parser.Node) []parser.Node {
    i := 0;
    var newNodes []parser.Node
    current := nodes[i]
    for i < len(nodes) {
        current = nodes[i]
        switch current.(type) {
            case parser.PlusNode:
                delta := 0
                for true {
                    current := nodes[i]
                    if (fmt.Sprintf("%T", current) == "parser.PlusNode") {
                        delta += int(reflect.ValueOf(current).Int())
                    } else {
                        break
                    }
                    i++
                }
                newNodes = append(newNodes, parser.PlusNode(delta))
                
            case parser.PointerNode:
                delta := 0
                for true {
                    current := nodes[i]
                    if (fmt.Sprintf("%T", current) == "parser.PointerNode") {
                        delta += int(reflect.ValueOf(current).Int())
                    } else {
                        break
                    }
                    i++
                }
                newNodes = append(newNodes, parser.PointerNode(delta))

            case parser.PrintNode: newNodes = append(newNodes, current)
            case parser.InputNode: newNodes = append(newNodes, current)

            case parser.LoopNode: {
                var newLoop []parser.Node 
                for _, node := range reflect.ValueOf(current).Interface().(parser.LoopNode) {
                    newLoop = append(newLoop, node)
                }
                newLoop = Optimize(newLoop)
                newNodes = append(newNodes, parser.LoopNode(newLoop))
            }
        }
        
        i++
    }

    return newNodes
}
