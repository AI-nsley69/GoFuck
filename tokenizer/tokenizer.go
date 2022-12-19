package tokenizer

import (
    "fmt"
)

// Create a new enum-like type for tokentypes
type Type int16

const (
    PointerInc Type = 0
    PointerDec      = 1
    IncValue        = 2
    DecValue        = 3
    Print           = 4
    Input           = 5
    WhileOpen       = 6
    WhileClose      = 7
)

type Token struct {
    TknType Type
    Pos int
    Char rune
}

var tokens []Token

func Tokenize(source string) []Token {
    for pos, char := range source {
        switch char {
    	    case '>': addToken(PointerInc, pos, char)
    	    case '<': addToken(PointerDec, pos, char)
    	    case '+': addToken(IncValue, pos, char)
    	    case '-': addToken(DecValue, pos, char)
    	    case '.': addToken(Print, pos, char)
    	    case ',': addToken(Input, pos, char)
    	    case '[': addToken(WhileOpen, pos, char)
    	    case ']': addToken(WhileClose, pos, char)
    	    case '\n': continue
    	    default: fmt.Println(fmt.Sprintf("Unrecognized token %c at position %v", char, pos))
        }
    }

    return tokens;
}

func addToken(token Type, pos int, char rune) {
    var newToken = Token{ TknType: token, Pos: pos, Char: char }
    tokens = append(tokens, newToken)
}

func TokenToString(token Type) string {
    switch token {
        case PointerInc: return "PointerInc"
        case PointerDec: return "PointerDec"
        case IncValue: return "IncValue"
        case DecValue: return "DecValue"
        case Print: return "Print"
        case Input: return "Input"
        case WhileOpen: return "WhileOpen"
        case WhileClose: return "WhileClose"
        default: return ""
    }
}

