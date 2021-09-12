// Code generated from Lang.g4 by ANTLR 4.9.2. DO NOT EDIT.

package parser

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 22, 138,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4,
	18, 9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 3, 2, 3, 2, 3, 2,
	3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 4, 3, 4, 3, 5, 3, 5, 3, 6,
	3, 6, 3, 7, 3, 7, 3, 8, 3, 8, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9,
	3, 9, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 11, 3, 11, 3,
	11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 12, 3, 12, 3, 13,
	3, 13, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 15, 3, 15, 3, 15, 3,
	15, 3, 15, 3, 15, 3, 15, 3, 16, 3, 16, 3, 16, 3, 16, 3, 17, 3, 17, 3, 17,
	3, 17, 3, 17, 3, 18, 6, 18, 116, 10, 18, 13, 18, 14, 18, 117, 3, 19, 6,
	19, 121, 10, 19, 13, 19, 14, 19, 122, 3, 20, 3, 20, 7, 20, 127, 10, 20,
	12, 20, 14, 20, 130, 11, 20, 3, 21, 6, 21, 133, 10, 21, 13, 21, 14, 21,
	134, 3, 21, 3, 21, 2, 2, 22, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13, 8, 15,
	9, 17, 10, 19, 11, 21, 12, 23, 13, 25, 14, 27, 15, 29, 16, 31, 17, 33,
	18, 35, 19, 37, 20, 39, 21, 41, 22, 3, 2, 7, 3, 2, 50, 59, 3, 2, 99, 124,
	4, 2, 67, 92, 99, 124, 5, 2, 50, 59, 67, 92, 99, 124, 5, 2, 11, 12, 15,
	15, 34, 34, 2, 141, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2,
	2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13, 3, 2, 2, 2, 2, 15, 3, 2,
	2, 2, 2, 17, 3, 2, 2, 2, 2, 19, 3, 2, 2, 2, 2, 21, 3, 2, 2, 2, 2, 23, 3,
	2, 2, 2, 2, 25, 3, 2, 2, 2, 2, 27, 3, 2, 2, 2, 2, 29, 3, 2, 2, 2, 2, 31,
	3, 2, 2, 2, 2, 33, 3, 2, 2, 2, 2, 35, 3, 2, 2, 2, 2, 37, 3, 2, 2, 2, 2,
	39, 3, 2, 2, 2, 2, 41, 3, 2, 2, 2, 3, 43, 3, 2, 2, 2, 5, 51, 3, 2, 2, 2,
	7, 53, 3, 2, 2, 2, 9, 55, 3, 2, 2, 2, 11, 57, 3, 2, 2, 2, 13, 59, 3, 2,
	2, 2, 15, 61, 3, 2, 2, 2, 17, 63, 3, 2, 2, 2, 19, 71, 3, 2, 2, 2, 21, 78,
	3, 2, 2, 2, 23, 88, 3, 2, 2, 2, 25, 90, 3, 2, 2, 2, 27, 92, 3, 2, 2, 2,
	29, 98, 3, 2, 2, 2, 31, 105, 3, 2, 2, 2, 33, 109, 3, 2, 2, 2, 35, 115,
	3, 2, 2, 2, 37, 120, 3, 2, 2, 2, 39, 124, 3, 2, 2, 2, 41, 132, 3, 2, 2,
	2, 43, 44, 7, 114, 2, 2, 44, 45, 7, 99, 2, 2, 45, 46, 7, 101, 2, 2, 46,
	47, 7, 109, 2, 2, 47, 48, 7, 99, 2, 2, 48, 49, 7, 105, 2, 2, 49, 50, 7,
	103, 2, 2, 50, 4, 3, 2, 2, 2, 51, 52, 7, 61, 2, 2, 52, 6, 3, 2, 2, 2, 53,
	54, 7, 66, 2, 2, 54, 8, 3, 2, 2, 2, 55, 56, 7, 42, 2, 2, 56, 10, 3, 2,
	2, 2, 57, 58, 7, 46, 2, 2, 58, 12, 3, 2, 2, 2, 59, 60, 7, 43, 2, 2, 60,
	14, 3, 2, 2, 2, 61, 62, 7, 63, 2, 2, 62, 16, 3, 2, 2, 2, 63, 64, 7, 114,
	2, 2, 64, 65, 7, 116, 2, 2, 65, 66, 7, 107, 2, 2, 66, 67, 7, 120, 2, 2,
	67, 68, 7, 99, 2, 2, 68, 69, 7, 118, 2, 2, 69, 70, 7, 103, 2, 2, 70, 18,
	3, 2, 2, 2, 71, 72, 7, 114, 2, 2, 72, 73, 7, 119, 2, 2, 73, 74, 7, 100,
	2, 2, 74, 75, 7, 110, 2, 2, 75, 76, 7, 107, 2, 2, 76, 77, 7, 101, 2, 2,
	77, 20, 3, 2, 2, 2, 78, 79, 7, 107, 2, 2, 79, 80, 7, 112, 2, 2, 80, 81,
	7, 118, 2, 2, 81, 82, 7, 103, 2, 2, 82, 83, 7, 116, 2, 2, 83, 84, 7, 104,
	2, 2, 84, 85, 7, 99, 2, 2, 85, 86, 7, 101, 2, 2, 86, 87, 7, 103, 2, 2,
	87, 22, 3, 2, 2, 2, 88, 89, 7, 125, 2, 2, 89, 24, 3, 2, 2, 2, 90, 91, 7,
	127, 2, 2, 91, 26, 3, 2, 2, 2, 92, 93, 7, 101, 2, 2, 93, 94, 7, 110, 2,
	2, 94, 95, 7, 99, 2, 2, 95, 96, 7, 117, 2, 2, 96, 97, 7, 117, 2, 2, 97,
	28, 3, 2, 2, 2, 98, 99, 7, 85, 2, 2, 99, 100, 7, 118, 2, 2, 100, 101, 7,
	116, 2, 2, 101, 102, 7, 107, 2, 2, 102, 103, 7, 112, 2, 2, 103, 104, 7,
	105, 2, 2, 104, 30, 3, 2, 2, 2, 105, 106, 7, 75, 2, 2, 106, 107, 7, 112,
	2, 2, 107, 108, 7, 118, 2, 2, 108, 32, 3, 2, 2, 2, 109, 110, 7, 68, 2,
	2, 110, 111, 7, 113, 2, 2, 111, 112, 7, 113, 2, 2, 112, 113, 7, 110, 2,
	2, 113, 34, 3, 2, 2, 2, 114, 116, 9, 2, 2, 2, 115, 114, 3, 2, 2, 2, 116,
	117, 3, 2, 2, 2, 117, 115, 3, 2, 2, 2, 117, 118, 3, 2, 2, 2, 118, 36, 3,
	2, 2, 2, 119, 121, 9, 3, 2, 2, 120, 119, 3, 2, 2, 2, 121, 122, 3, 2, 2,
	2, 122, 120, 3, 2, 2, 2, 122, 123, 3, 2, 2, 2, 123, 38, 3, 2, 2, 2, 124,
	128, 9, 4, 2, 2, 125, 127, 9, 5, 2, 2, 126, 125, 3, 2, 2, 2, 127, 130,
	3, 2, 2, 2, 128, 126, 3, 2, 2, 2, 128, 129, 3, 2, 2, 2, 129, 40, 3, 2,
	2, 2, 130, 128, 3, 2, 2, 2, 131, 133, 9, 6, 2, 2, 132, 131, 3, 2, 2, 2,
	133, 134, 3, 2, 2, 2, 134, 132, 3, 2, 2, 2, 134, 135, 3, 2, 2, 2, 135,
	136, 3, 2, 2, 2, 136, 137, 8, 21, 2, 2, 137, 42, 3, 2, 2, 2, 7, 2, 117,
	122, 128, 134, 3, 8, 2, 2,
}

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'package'", "';'", "'@'", "'('", "','", "')'", "'='", "'private'",
	"'public'", "'interface'", "'{'", "'}'", "'class'", "'String'", "'Int'",
	"'Bool'",
}

var lexerSymbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "INTEGER",
	"PACKAGE_NAME", "TYPE_NAME", "WS",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
	"T__9", "T__10", "T__11", "T__12", "T__13", "T__14", "T__15", "INTEGER",
	"PACKAGE_NAME", "TYPE_NAME", "WS",
}

type LangLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

// NewLangLexer produces a new lexer instance for the optional input antlr.CharStream.
//
// The *LangLexer instance produced may be reused by calling the SetInputStream method.
// The initial lexer configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewLangLexer(input antlr.CharStream) *LangLexer {
	l := new(LangLexer)
	lexerDeserializer := antlr.NewATNDeserializer(nil)
	lexerAtn := lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)
	lexerDecisionToDFA := make([]*antlr.DFA, len(lexerAtn.DecisionToState))
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "Lang.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// LangLexer tokens.
const (
	LangLexerT__0         = 1
	LangLexerT__1         = 2
	LangLexerT__2         = 3
	LangLexerT__3         = 4
	LangLexerT__4         = 5
	LangLexerT__5         = 6
	LangLexerT__6         = 7
	LangLexerT__7         = 8
	LangLexerT__8         = 9
	LangLexerT__9         = 10
	LangLexerT__10        = 11
	LangLexerT__11        = 12
	LangLexerT__12        = 13
	LangLexerT__13        = 14
	LangLexerT__14        = 15
	LangLexerT__15        = 16
	LangLexerINTEGER      = 17
	LangLexerPACKAGE_NAME = 18
	LangLexerTYPE_NAME    = 19
	LangLexerWS           = 20
)
