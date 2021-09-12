// Code generated from Lang.g4 by ANTLR 4.9.2. DO NOT EDIT.

package parser // Lang

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 22, 124,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 3, 2, 3, 2, 3, 2, 7, 2, 32, 10, 2, 12, 2, 14, 2, 35,
	11, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 3, 4, 5, 4,
	47, 10, 4, 3, 4, 3, 4, 7, 4, 51, 10, 4, 12, 4, 14, 4, 54, 11, 4, 3, 4,
	3, 4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 6, 3, 6, 3, 7, 5, 7, 65, 10, 7, 3, 7,
	3, 7, 3, 7, 3, 7, 3, 7, 3, 8, 7, 8, 73, 10, 8, 12, 8, 14, 8, 76, 11, 8,
	3, 8, 5, 8, 79, 10, 8, 3, 8, 3, 8, 3, 8, 3, 8, 7, 8, 85, 10, 8, 12, 8,
	14, 8, 88, 11, 8, 3, 8, 3, 8, 3, 9, 3, 9, 5, 9, 94, 10, 9, 3, 10, 5, 10,
	97, 10, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 11, 7, 11, 104, 10, 11, 12,
	11, 14, 11, 107, 11, 11, 3, 11, 5, 11, 110, 10, 11, 3, 11, 3, 11, 3, 11,
	3, 11, 3, 11, 3, 11, 3, 12, 3, 12, 3, 13, 3, 13, 3, 14, 3, 14, 3, 14, 2,
	2, 15, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 2, 5, 3, 2, 10,
	11, 3, 2, 20, 21, 3, 2, 16, 18, 2, 122, 2, 28, 3, 2, 2, 2, 4, 38, 3, 2,
	2, 2, 6, 42, 3, 2, 2, 2, 8, 57, 3, 2, 2, 2, 10, 61, 3, 2, 2, 2, 12, 64,
	3, 2, 2, 2, 14, 74, 3, 2, 2, 2, 16, 93, 3, 2, 2, 2, 18, 96, 3, 2, 2, 2,
	20, 105, 3, 2, 2, 2, 22, 117, 3, 2, 2, 2, 24, 119, 3, 2, 2, 2, 26, 121,
	3, 2, 2, 2, 28, 33, 5, 4, 3, 2, 29, 32, 5, 14, 8, 2, 30, 32, 5, 12, 7,
	2, 31, 29, 3, 2, 2, 2, 31, 30, 3, 2, 2, 2, 32, 35, 3, 2, 2, 2, 33, 31,
	3, 2, 2, 2, 33, 34, 3, 2, 2, 2, 34, 36, 3, 2, 2, 2, 35, 33, 3, 2, 2, 2,
	36, 37, 7, 2, 2, 3, 37, 3, 3, 2, 2, 2, 38, 39, 7, 3, 2, 2, 39, 40, 7, 20,
	2, 2, 40, 41, 7, 4, 2, 2, 41, 5, 3, 2, 2, 2, 42, 43, 7, 5, 2, 2, 43, 44,
	5, 22, 12, 2, 44, 46, 7, 6, 2, 2, 45, 47, 5, 8, 5, 2, 46, 45, 3, 2, 2,
	2, 46, 47, 3, 2, 2, 2, 47, 52, 3, 2, 2, 2, 48, 49, 7, 7, 2, 2, 49, 51,
	5, 8, 5, 2, 50, 48, 3, 2, 2, 2, 51, 54, 3, 2, 2, 2, 52, 50, 3, 2, 2, 2,
	52, 53, 3, 2, 2, 2, 53, 55, 3, 2, 2, 2, 54, 52, 3, 2, 2, 2, 55, 56, 7,
	8, 2, 2, 56, 7, 3, 2, 2, 2, 57, 58, 5, 22, 12, 2, 58, 59, 7, 9, 2, 2, 59,
	60, 5, 26, 14, 2, 60, 9, 3, 2, 2, 2, 61, 62, 9, 2, 2, 2, 62, 11, 3, 2,
	2, 2, 63, 65, 5, 10, 6, 2, 64, 63, 3, 2, 2, 2, 64, 65, 3, 2, 2, 2, 65,
	66, 3, 2, 2, 2, 66, 67, 7, 12, 2, 2, 67, 68, 5, 22, 12, 2, 68, 69, 7, 13,
	2, 2, 69, 70, 7, 14, 2, 2, 70, 13, 3, 2, 2, 2, 71, 73, 5, 6, 4, 2, 72,
	71, 3, 2, 2, 2, 73, 76, 3, 2, 2, 2, 74, 72, 3, 2, 2, 2, 74, 75, 3, 2, 2,
	2, 75, 78, 3, 2, 2, 2, 76, 74, 3, 2, 2, 2, 77, 79, 5, 10, 6, 2, 78, 77,
	3, 2, 2, 2, 78, 79, 3, 2, 2, 2, 79, 80, 3, 2, 2, 2, 80, 81, 7, 15, 2, 2,
	81, 82, 5, 22, 12, 2, 82, 86, 7, 13, 2, 2, 83, 85, 5, 16, 9, 2, 84, 83,
	3, 2, 2, 2, 85, 88, 3, 2, 2, 2, 86, 84, 3, 2, 2, 2, 86, 87, 3, 2, 2, 2,
	87, 89, 3, 2, 2, 2, 88, 86, 3, 2, 2, 2, 89, 90, 7, 14, 2, 2, 90, 15, 3,
	2, 2, 2, 91, 94, 5, 18, 10, 2, 92, 94, 5, 20, 11, 2, 93, 91, 3, 2, 2, 2,
	93, 92, 3, 2, 2, 2, 94, 17, 3, 2, 2, 2, 95, 97, 5, 10, 6, 2, 96, 95, 3,
	2, 2, 2, 96, 97, 3, 2, 2, 2, 97, 98, 3, 2, 2, 2, 98, 99, 5, 24, 13, 2,
	99, 100, 5, 22, 12, 2, 100, 101, 7, 4, 2, 2, 101, 19, 3, 2, 2, 2, 102,
	104, 5, 6, 4, 2, 103, 102, 3, 2, 2, 2, 104, 107, 3, 2, 2, 2, 105, 103,
	3, 2, 2, 2, 105, 106, 3, 2, 2, 2, 106, 109, 3, 2, 2, 2, 107, 105, 3, 2,
	2, 2, 108, 110, 5, 10, 6, 2, 109, 108, 3, 2, 2, 2, 109, 110, 3, 2, 2, 2,
	110, 111, 3, 2, 2, 2, 111, 112, 5, 22, 12, 2, 112, 113, 7, 6, 2, 2, 113,
	114, 7, 8, 2, 2, 114, 115, 7, 13, 2, 2, 115, 116, 7, 14, 2, 2, 116, 21,
	3, 2, 2, 2, 117, 118, 9, 3, 2, 2, 118, 23, 3, 2, 2, 2, 119, 120, 9, 4,
	2, 2, 120, 25, 3, 2, 2, 2, 121, 122, 7, 19, 2, 2, 122, 27, 3, 2, 2, 2,
	14, 31, 33, 46, 52, 64, 74, 78, 86, 93, 96, 105, 109,
}
var literalNames = []string{
	"", "'package'", "';'", "'@'", "'('", "','", "')'", "'='", "'private'",
	"'public'", "'interface'", "'{'", "'}'", "'class'", "'String'", "'Int'",
	"'Bool'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "INTEGER",
	"PACKAGE_NAME", "TYPE_NAME", "WS",
}

var ruleNames = []string{
	"program", "namespace", "annotation", "key_value", "modifier", "nInterface",
	"nClass", "classContent", "fieldDeclaration", "methodDeclaration", "identifier",
	"langType", "value",
}

type LangParser struct {
	*antlr.BaseParser
}

// NewLangParser produces a new parser instance for the optional input antlr.TokenStream.
//
// The *LangParser instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewLangParser(input antlr.TokenStream) *LangParser {
	this := new(LangParser)
	deserializer := antlr.NewATNDeserializer(nil)
	deserializedATN := deserializer.DeserializeFromUInt16(parserATN)
	decisionToDFA := make([]*antlr.DFA, len(deserializedATN.DecisionToState))
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "Lang.g4"

	return this
}

// LangParser tokens.
const (
	LangParserEOF          = antlr.TokenEOF
	LangParserT__0         = 1
	LangParserT__1         = 2
	LangParserT__2         = 3
	LangParserT__3         = 4
	LangParserT__4         = 5
	LangParserT__5         = 6
	LangParserT__6         = 7
	LangParserT__7         = 8
	LangParserT__8         = 9
	LangParserT__9         = 10
	LangParserT__10        = 11
	LangParserT__11        = 12
	LangParserT__12        = 13
	LangParserT__13        = 14
	LangParserT__14        = 15
	LangParserT__15        = 16
	LangParserINTEGER      = 17
	LangParserPACKAGE_NAME = 18
	LangParserTYPE_NAME    = 19
	LangParserWS           = 20
)

// LangParser rules.
const (
	LangParserRULE_program           = 0
	LangParserRULE_namespace         = 1
	LangParserRULE_annotation        = 2
	LangParserRULE_key_value         = 3
	LangParserRULE_modifier          = 4
	LangParserRULE_nInterface        = 5
	LangParserRULE_nClass            = 6
	LangParserRULE_classContent      = 7
	LangParserRULE_fieldDeclaration  = 8
	LangParserRULE_methodDeclaration = 9
	LangParserRULE_identifier        = 10
	LangParserRULE_langType          = 11
	LangParserRULE_value             = 12
)

// IProgramContext is an interface to support dynamic dispatch.
type IProgramContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsProgramContext differentiates from other interfaces.
	IsProgramContext()
}

type ProgramContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgramContext() *ProgramContext {
	var p = new(ProgramContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LangParserRULE_program
	return p
}

func (*ProgramContext) IsProgramContext() {}

func NewProgramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramContext {
	var p = new(ProgramContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LangParserRULE_program

	return p
}

func (s *ProgramContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgramContext) Namespace() INamespaceContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INamespaceContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INamespaceContext)
}

func (s *ProgramContext) EOF() antlr.TerminalNode {
	return s.GetToken(LangParserEOF, 0)
}

func (s *ProgramContext) AllNClass() []INClassContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*INClassContext)(nil)).Elem())
	var tst = make([]INClassContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(INClassContext)
		}
	}

	return tst
}

func (s *ProgramContext) NClass(i int) INClassContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INClassContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(INClassContext)
}

func (s *ProgramContext) AllNInterface() []INInterfaceContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*INInterfaceContext)(nil)).Elem())
	var tst = make([]INInterfaceContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(INInterfaceContext)
		}
	}

	return tst
}

func (s *ProgramContext) NInterface(i int) INInterfaceContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INInterfaceContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(INInterfaceContext)
}

func (s *ProgramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgramContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LangListener); ok {
		listenerT.EnterProgram(s)
	}
}

func (s *ProgramContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LangListener); ok {
		listenerT.ExitProgram(s)
	}
}

func (p *LangParser) Program() (localctx IProgramContext) {
	localctx = NewProgramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, LangParserRULE_program)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(26)
		p.Namespace()
	}
	p.SetState(31)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<LangParserT__2)|(1<<LangParserT__7)|(1<<LangParserT__8)|(1<<LangParserT__9)|(1<<LangParserT__12))) != 0 {
		p.SetState(29)
		p.GetErrorHandler().Sync(p)
		switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(27)
				p.NClass()
			}

		case 2:
			{
				p.SetState(28)
				p.NInterface()
			}

		}

		p.SetState(33)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(34)
		p.Match(LangParserEOF)
	}

	return localctx
}

// INamespaceContext is an interface to support dynamic dispatch.
type INamespaceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNamespaceContext differentiates from other interfaces.
	IsNamespaceContext()
}

type NamespaceContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNamespaceContext() *NamespaceContext {
	var p = new(NamespaceContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LangParserRULE_namespace
	return p
}

func (*NamespaceContext) IsNamespaceContext() {}

func NewNamespaceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NamespaceContext {
	var p = new(NamespaceContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LangParserRULE_namespace

	return p
}

func (s *NamespaceContext) GetParser() antlr.Parser { return s.parser }

func (s *NamespaceContext) PACKAGE_NAME() antlr.TerminalNode {
	return s.GetToken(LangParserPACKAGE_NAME, 0)
}

func (s *NamespaceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NamespaceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NamespaceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LangListener); ok {
		listenerT.EnterNamespace(s)
	}
}

func (s *NamespaceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LangListener); ok {
		listenerT.ExitNamespace(s)
	}
}

func (p *LangParser) Namespace() (localctx INamespaceContext) {
	localctx = NewNamespaceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, LangParserRULE_namespace)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(36)
		p.Match(LangParserT__0)
	}
	{
		p.SetState(37)
		p.Match(LangParserPACKAGE_NAME)
	}
	{
		p.SetState(38)
		p.Match(LangParserT__1)
	}

	return localctx
}

// IAnnotationContext is an interface to support dynamic dispatch.
type IAnnotationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAnnotationContext differentiates from other interfaces.
	IsAnnotationContext()
}

type AnnotationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAnnotationContext() *AnnotationContext {
	var p = new(AnnotationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LangParserRULE_annotation
	return p
}

func (*AnnotationContext) IsAnnotationContext() {}

func NewAnnotationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AnnotationContext {
	var p = new(AnnotationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LangParserRULE_annotation

	return p
}

func (s *AnnotationContext) GetParser() antlr.Parser { return s.parser }

func (s *AnnotationContext) Identifier() IIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *AnnotationContext) AllKey_value() []IKey_valueContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IKey_valueContext)(nil)).Elem())
	var tst = make([]IKey_valueContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IKey_valueContext)
		}
	}

	return tst
}

func (s *AnnotationContext) Key_value(i int) IKey_valueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IKey_valueContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IKey_valueContext)
}

func (s *AnnotationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AnnotationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AnnotationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LangListener); ok {
		listenerT.EnterAnnotation(s)
	}
}

func (s *AnnotationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LangListener); ok {
		listenerT.ExitAnnotation(s)
	}
}

func (p *LangParser) Annotation() (localctx IAnnotationContext) {
	localctx = NewAnnotationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, LangParserRULE_annotation)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(40)
		p.Match(LangParserT__2)
	}
	{
		p.SetState(41)
		p.Identifier()
	}
	{
		p.SetState(42)
		p.Match(LangParserT__3)
	}
	p.SetState(44)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == LangParserPACKAGE_NAME || _la == LangParserTYPE_NAME {
		{
			p.SetState(43)
			p.Key_value()
		}

	}
	p.SetState(50)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == LangParserT__4 {
		{
			p.SetState(46)
			p.Match(LangParserT__4)
		}
		{
			p.SetState(47)
			p.Key_value()
		}

		p.SetState(52)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(53)
		p.Match(LangParserT__5)
	}

	return localctx
}

// IKey_valueContext is an interface to support dynamic dispatch.
type IKey_valueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsKey_valueContext differentiates from other interfaces.
	IsKey_valueContext()
}

type Key_valueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyKey_valueContext() *Key_valueContext {
	var p = new(Key_valueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LangParserRULE_key_value
	return p
}

func (*Key_valueContext) IsKey_valueContext() {}

func NewKey_valueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Key_valueContext {
	var p = new(Key_valueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LangParserRULE_key_value

	return p
}

func (s *Key_valueContext) GetParser() antlr.Parser { return s.parser }

func (s *Key_valueContext) Identifier() IIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *Key_valueContext) Value() IValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *Key_valueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Key_valueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Key_valueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LangListener); ok {
		listenerT.EnterKey_value(s)
	}
}

func (s *Key_valueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LangListener); ok {
		listenerT.ExitKey_value(s)
	}
}

func (p *LangParser) Key_value() (localctx IKey_valueContext) {
	localctx = NewKey_valueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, LangParserRULE_key_value)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(55)
		p.Identifier()
	}
	{
		p.SetState(56)
		p.Match(LangParserT__6)
	}
	{
		p.SetState(57)
		p.Value()
	}

	return localctx
}

// IModifierContext is an interface to support dynamic dispatch.
type IModifierContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsModifierContext differentiates from other interfaces.
	IsModifierContext()
}

type ModifierContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyModifierContext() *ModifierContext {
	var p = new(ModifierContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LangParserRULE_modifier
	return p
}

func (*ModifierContext) IsModifierContext() {}

func NewModifierContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ModifierContext {
	var p = new(ModifierContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LangParserRULE_modifier

	return p
}

func (s *ModifierContext) GetParser() antlr.Parser { return s.parser }
func (s *ModifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ModifierContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ModifierContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LangListener); ok {
		listenerT.EnterModifier(s)
	}
}

func (s *ModifierContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LangListener); ok {
		listenerT.ExitModifier(s)
	}
}

func (p *LangParser) Modifier() (localctx IModifierContext) {
	localctx = NewModifierContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, LangParserRULE_modifier)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(59)
		_la = p.GetTokenStream().LA(1)

		if !(_la == LangParserT__7 || _la == LangParserT__8) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// INInterfaceContext is an interface to support dynamic dispatch.
type INInterfaceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNInterfaceContext differentiates from other interfaces.
	IsNInterfaceContext()
}

type NInterfaceContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNInterfaceContext() *NInterfaceContext {
	var p = new(NInterfaceContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LangParserRULE_nInterface
	return p
}

func (*NInterfaceContext) IsNInterfaceContext() {}

func NewNInterfaceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NInterfaceContext {
	var p = new(NInterfaceContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LangParserRULE_nInterface

	return p
}

func (s *NInterfaceContext) GetParser() antlr.Parser { return s.parser }

func (s *NInterfaceContext) Identifier() IIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *NInterfaceContext) Modifier() IModifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IModifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IModifierContext)
}

func (s *NInterfaceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NInterfaceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NInterfaceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LangListener); ok {
		listenerT.EnterNInterface(s)
	}
}

func (s *NInterfaceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LangListener); ok {
		listenerT.ExitNInterface(s)
	}
}

func (p *LangParser) NInterface() (localctx INInterfaceContext) {
	localctx = NewNInterfaceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, LangParserRULE_nInterface)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(62)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == LangParserT__7 || _la == LangParserT__8 {
		{
			p.SetState(61)
			p.Modifier()
		}

	}
	{
		p.SetState(64)
		p.Match(LangParserT__9)
	}
	{
		p.SetState(65)
		p.Identifier()
	}
	{
		p.SetState(66)
		p.Match(LangParserT__10)
	}
	{
		p.SetState(67)
		p.Match(LangParserT__11)
	}

	return localctx
}

// INClassContext is an interface to support dynamic dispatch.
type INClassContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNClassContext differentiates from other interfaces.
	IsNClassContext()
}

type NClassContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNClassContext() *NClassContext {
	var p = new(NClassContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LangParserRULE_nClass
	return p
}

func (*NClassContext) IsNClassContext() {}

func NewNClassContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NClassContext {
	var p = new(NClassContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LangParserRULE_nClass

	return p
}

func (s *NClassContext) GetParser() antlr.Parser { return s.parser }

func (s *NClassContext) Identifier() IIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *NClassContext) AllAnnotation() []IAnnotationContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAnnotationContext)(nil)).Elem())
	var tst = make([]IAnnotationContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAnnotationContext)
		}
	}

	return tst
}

func (s *NClassContext) Annotation(i int) IAnnotationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAnnotationContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAnnotationContext)
}

func (s *NClassContext) Modifier() IModifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IModifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IModifierContext)
}

func (s *NClassContext) AllClassContent() []IClassContentContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IClassContentContext)(nil)).Elem())
	var tst = make([]IClassContentContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IClassContentContext)
		}
	}

	return tst
}

func (s *NClassContext) ClassContent(i int) IClassContentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IClassContentContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IClassContentContext)
}

func (s *NClassContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NClassContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NClassContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LangListener); ok {
		listenerT.EnterNClass(s)
	}
}

func (s *NClassContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LangListener); ok {
		listenerT.ExitNClass(s)
	}
}

func (p *LangParser) NClass() (localctx INClassContext) {
	localctx = NewNClassContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, LangParserRULE_nClass)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(72)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == LangParserT__2 {
		{
			p.SetState(69)
			p.Annotation()
		}

		p.SetState(74)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(76)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == LangParserT__7 || _la == LangParserT__8 {
		{
			p.SetState(75)
			p.Modifier()
		}

	}
	{
		p.SetState(78)
		p.Match(LangParserT__12)
	}
	{
		p.SetState(79)
		p.Identifier()
	}
	{
		p.SetState(80)
		p.Match(LangParserT__10)
	}
	p.SetState(84)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<LangParserT__2)|(1<<LangParserT__7)|(1<<LangParserT__8)|(1<<LangParserT__13)|(1<<LangParserT__14)|(1<<LangParserT__15)|(1<<LangParserPACKAGE_NAME)|(1<<LangParserTYPE_NAME))) != 0 {
		{
			p.SetState(81)
			p.ClassContent()
		}

		p.SetState(86)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(87)
		p.Match(LangParserT__11)
	}

	return localctx
}

// IClassContentContext is an interface to support dynamic dispatch.
type IClassContentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsClassContentContext differentiates from other interfaces.
	IsClassContentContext()
}

type ClassContentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyClassContentContext() *ClassContentContext {
	var p = new(ClassContentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LangParserRULE_classContent
	return p
}

func (*ClassContentContext) IsClassContentContext() {}

func NewClassContentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ClassContentContext {
	var p = new(ClassContentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LangParserRULE_classContent

	return p
}

func (s *ClassContentContext) GetParser() antlr.Parser { return s.parser }

func (s *ClassContentContext) FieldDeclaration() IFieldDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFieldDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFieldDeclarationContext)
}

func (s *ClassContentContext) MethodDeclaration() IMethodDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMethodDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMethodDeclarationContext)
}

func (s *ClassContentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ClassContentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ClassContentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LangListener); ok {
		listenerT.EnterClassContent(s)
	}
}

func (s *ClassContentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LangListener); ok {
		listenerT.ExitClassContent(s)
	}
}

func (p *LangParser) ClassContent() (localctx IClassContentContext) {
	localctx = NewClassContentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, LangParserRULE_classContent)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(91)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(89)
			p.FieldDeclaration()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(90)
			p.MethodDeclaration()
		}

	}

	return localctx
}

// IFieldDeclarationContext is an interface to support dynamic dispatch.
type IFieldDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFieldDeclarationContext differentiates from other interfaces.
	IsFieldDeclarationContext()
}

type FieldDeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldDeclarationContext() *FieldDeclarationContext {
	var p = new(FieldDeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LangParserRULE_fieldDeclaration
	return p
}

func (*FieldDeclarationContext) IsFieldDeclarationContext() {}

func NewFieldDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldDeclarationContext {
	var p = new(FieldDeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LangParserRULE_fieldDeclaration

	return p
}

func (s *FieldDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldDeclarationContext) LangType() ILangTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILangTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILangTypeContext)
}

func (s *FieldDeclarationContext) Identifier() IIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *FieldDeclarationContext) Modifier() IModifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IModifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IModifierContext)
}

func (s *FieldDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LangListener); ok {
		listenerT.EnterFieldDeclaration(s)
	}
}

func (s *FieldDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LangListener); ok {
		listenerT.ExitFieldDeclaration(s)
	}
}

func (p *LangParser) FieldDeclaration() (localctx IFieldDeclarationContext) {
	localctx = NewFieldDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, LangParserRULE_fieldDeclaration)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(94)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == LangParserT__7 || _la == LangParserT__8 {
		{
			p.SetState(93)
			p.Modifier()
		}

	}
	{
		p.SetState(96)
		p.LangType()
	}
	{
		p.SetState(97)
		p.Identifier()
	}
	{
		p.SetState(98)
		p.Match(LangParserT__1)
	}

	return localctx
}

// IMethodDeclarationContext is an interface to support dynamic dispatch.
type IMethodDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMethodDeclarationContext differentiates from other interfaces.
	IsMethodDeclarationContext()
}

type MethodDeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMethodDeclarationContext() *MethodDeclarationContext {
	var p = new(MethodDeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LangParserRULE_methodDeclaration
	return p
}

func (*MethodDeclarationContext) IsMethodDeclarationContext() {}

func NewMethodDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MethodDeclarationContext {
	var p = new(MethodDeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LangParserRULE_methodDeclaration

	return p
}

func (s *MethodDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *MethodDeclarationContext) Identifier() IIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *MethodDeclarationContext) AllAnnotation() []IAnnotationContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAnnotationContext)(nil)).Elem())
	var tst = make([]IAnnotationContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAnnotationContext)
		}
	}

	return tst
}

func (s *MethodDeclarationContext) Annotation(i int) IAnnotationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAnnotationContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAnnotationContext)
}

func (s *MethodDeclarationContext) Modifier() IModifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IModifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IModifierContext)
}

func (s *MethodDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MethodDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MethodDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LangListener); ok {
		listenerT.EnterMethodDeclaration(s)
	}
}

func (s *MethodDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LangListener); ok {
		listenerT.ExitMethodDeclaration(s)
	}
}

func (p *LangParser) MethodDeclaration() (localctx IMethodDeclarationContext) {
	localctx = NewMethodDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, LangParserRULE_methodDeclaration)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(103)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == LangParserT__2 {
		{
			p.SetState(100)
			p.Annotation()
		}

		p.SetState(105)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(107)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == LangParserT__7 || _la == LangParserT__8 {
		{
			p.SetState(106)
			p.Modifier()
		}

	}
	{
		p.SetState(109)
		p.Identifier()
	}
	{
		p.SetState(110)
		p.Match(LangParserT__3)
	}
	{
		p.SetState(111)
		p.Match(LangParserT__5)
	}
	{
		p.SetState(112)
		p.Match(LangParserT__10)
	}
	{
		p.SetState(113)
		p.Match(LangParserT__11)
	}

	return localctx
}

// IIdentifierContext is an interface to support dynamic dispatch.
type IIdentifierContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIdentifierContext differentiates from other interfaces.
	IsIdentifierContext()
}

type IdentifierContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIdentifierContext() *IdentifierContext {
	var p = new(IdentifierContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LangParserRULE_identifier
	return p
}

func (*IdentifierContext) IsIdentifierContext() {}

func NewIdentifierContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IdentifierContext {
	var p = new(IdentifierContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LangParserRULE_identifier

	return p
}

func (s *IdentifierContext) GetParser() antlr.Parser { return s.parser }

func (s *IdentifierContext) TYPE_NAME() antlr.TerminalNode {
	return s.GetToken(LangParserTYPE_NAME, 0)
}

func (s *IdentifierContext) PACKAGE_NAME() antlr.TerminalNode {
	return s.GetToken(LangParserPACKAGE_NAME, 0)
}

func (s *IdentifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentifierContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IdentifierContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LangListener); ok {
		listenerT.EnterIdentifier(s)
	}
}

func (s *IdentifierContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LangListener); ok {
		listenerT.ExitIdentifier(s)
	}
}

func (p *LangParser) Identifier() (localctx IIdentifierContext) {
	localctx = NewIdentifierContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, LangParserRULE_identifier)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(115)
		_la = p.GetTokenStream().LA(1)

		if !(_la == LangParserPACKAGE_NAME || _la == LangParserTYPE_NAME) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// ILangTypeContext is an interface to support dynamic dispatch.
type ILangTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLangTypeContext differentiates from other interfaces.
	IsLangTypeContext()
}

type LangTypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLangTypeContext() *LangTypeContext {
	var p = new(LangTypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LangParserRULE_langType
	return p
}

func (*LangTypeContext) IsLangTypeContext() {}

func NewLangTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LangTypeContext {
	var p = new(LangTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LangParserRULE_langType

	return p
}

func (s *LangTypeContext) GetParser() antlr.Parser { return s.parser }
func (s *LangTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LangTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LangTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LangListener); ok {
		listenerT.EnterLangType(s)
	}
}

func (s *LangTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LangListener); ok {
		listenerT.ExitLangType(s)
	}
}

func (p *LangParser) LangType() (localctx ILangTypeContext) {
	localctx = NewLangTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, LangParserRULE_langType)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(117)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<LangParserT__13)|(1<<LangParserT__14)|(1<<LangParserT__15))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IValueContext is an interface to support dynamic dispatch.
type IValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsValueContext differentiates from other interfaces.
	IsValueContext()
}

type ValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValueContext() *ValueContext {
	var p = new(ValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LangParserRULE_value
	return p
}

func (*ValueContext) IsValueContext() {}

func NewValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValueContext {
	var p = new(ValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LangParserRULE_value

	return p
}

func (s *ValueContext) GetParser() antlr.Parser { return s.parser }

func (s *ValueContext) INTEGER() antlr.TerminalNode {
	return s.GetToken(LangParserINTEGER, 0)
}

func (s *ValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LangListener); ok {
		listenerT.EnterValue(s)
	}
}

func (s *ValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LangListener); ok {
		listenerT.ExitValue(s)
	}
}

func (p *LangParser) Value() (localctx IValueContext) {
	localctx = NewValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, LangParserRULE_value)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(119)
		p.Match(LangParserINTEGER)
	}

	return localctx
}
