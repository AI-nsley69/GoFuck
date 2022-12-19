package optimizer

import (
    // "trainsley69.me/GoFuck/tokenizer"
    "trainsley69.me/GoFuck/parser"
    "fmt"
    "reflect"
)

func Optimize(nodes []parser.Node) []parser.Node {
    i := 0;
    lim := len(nodes)
    var newNodes []parser.Node
    current := nodes[i]
    for i < lim {
        current = nodes[i]
        switch fmt.Sprintf("%T", current) {
            case "parser.PlusNode":
                delta := 0
                delta += int(reflect.ValueOf(current).Int())
                for fmt.Sprintf("T%", PeekNext(nodes, i)) == "parser.PlusNode" {
                    i++
                    current := nodes[i]
                    delta += int(reflect.ValueOf(current).Int())
                }

                newNodes = append(newNodes, parser.PlusNode(delta))
                
            case "parser.PointerNode":
                delta := 0
                delta += int(reflect.ValueOf(current).Int())
                for fmt.Sprintf("T%", PeekNext(nodes, i)) == "parser.PointerNode" {
                    i++
                    current := nodes[i]
                    delta += int(reflect.ValueOf(current).Int())
                }
               
                newNodes = append(newNodes, parser.PointerNode(delta))

            case "parser.PrintNode": newNodes = append(newNodes, parser.PrintNode{})
            case "parser.InputNode": newNodes = append(newNodes, parser.InputNode{})

            case "parser.LoopNode": {
                var newLoop []parser.Node 
                for _, node := range reflect.ValueOf(current).Interface().(parser.LoopNode) {
                    newLoop = append(newLoop, node)
                }
                value := Optimize(newLoop)
                newNodes = append(newNodes, parser.LoopNode(value))
            }
        }
        
        i++
    }

    return newNodes
}

func PeekNext(nodes []parser.Node, i int) parser.Node {
    next := i +1
    if (next >= len(nodes)) {
        return parser.InputNode{}
    }
    return nodes[next]
}

