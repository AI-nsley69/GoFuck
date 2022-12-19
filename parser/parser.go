package parser

import (
    "trainsley69.me/GoFuck/tokenizer"
    "log"
)

type Node interface {
    isAstNode()
}

type PlusNode int
type PointerNode int
type PrintNode struct {}
type InputNode struct {}
type LoopNode []Node

func (PlusNode) isAstNode() {}
func (PointerNode) isAstNode() {}
func (PrintNode) isAstNode() {}
func (InputNode) isAstNode() {}
func (LoopNode) isAstNode() {}

func looper(tokens []tokenizer.Token, ch chan tokenizer.Token) {
    for _, token := range tokens {
        ch <- token
    }
    close(ch)
}

func Parse(tokens []tokenizer.Token) []Node {
    var tknCh chan tokenizer.Token = make(chan tokenizer.Token)
    go looper(tokens, tknCh)
    return internalParse(tknCh, false)
}

func internalParse(ch chan tokenizer.Token, inLoop bool) []Node {
    var nodes []Node
    for token := range ch {
        switch token.TknType {
            case tokenizer.PointerInc: nodes = append(nodes, PointerNode(1))
            case tokenizer.PointerDec: nodes = append(nodes, PointerNode(-1))
            case tokenizer.IncValue: nodes = append(nodes, PlusNode(1))
            case tokenizer.DecValue: nodes = append(nodes, PlusNode(-1))
            case tokenizer.Print: nodes = append(nodes, PrintNode{})
            case tokenizer.Input: nodes = append(nodes, InputNode{})
            case tokenizer.WhileOpen: nodes = append(nodes, LoopNode(internalParse(ch, true)))
            case tokenizer.WhileClose:
                if inLoop {
                    return nodes
                } else {
                    log.Fatal("Unmatched ']'")
                }
           default: continue
        }
    }
    if inLoop {
        log.Fatal("Unmatched '['")
    }

    return nodes
}
