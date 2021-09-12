// Code generated from Lang.g4 by ANTLR 4.9.2. DO NOT EDIT.

package parser // Lang

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseLangListener is a complete listener for a parse tree produced by LangParser.
type BaseLangListener struct{}

var _ LangListener = &BaseLangListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseLangListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseLangListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseLangListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseLangListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProgram is called when production program is entered.
func (s *BaseLangListener) EnterProgram(ctx *ProgramContext) {}

// ExitProgram is called when production program is exited.
func (s *BaseLangListener) ExitProgram(ctx *ProgramContext) {}

// EnterNamespace is called when production namespace is entered.
func (s *BaseLangListener) EnterNamespace(ctx *NamespaceContext) {}

// ExitNamespace is called when production namespace is exited.
func (s *BaseLangListener) ExitNamespace(ctx *NamespaceContext) {}

// EnterAnnotation is called when production annotation is entered.
func (s *BaseLangListener) EnterAnnotation(ctx *AnnotationContext) {}

// ExitAnnotation is called when production annotation is exited.
func (s *BaseLangListener) ExitAnnotation(ctx *AnnotationContext) {}

// EnterKey_value is called when production key_value is entered.
func (s *BaseLangListener) EnterKey_value(ctx *Key_valueContext) {}

// ExitKey_value is called when production key_value is exited.
func (s *BaseLangListener) ExitKey_value(ctx *Key_valueContext) {}

// EnterModifier is called when production modifier is entered.
func (s *BaseLangListener) EnterModifier(ctx *ModifierContext) {}

// ExitModifier is called when production modifier is exited.
func (s *BaseLangListener) ExitModifier(ctx *ModifierContext) {}

// EnterNInterface is called when production nInterface is entered.
func (s *BaseLangListener) EnterNInterface(ctx *NInterfaceContext) {}

// ExitNInterface is called when production nInterface is exited.
func (s *BaseLangListener) ExitNInterface(ctx *NInterfaceContext) {}

// EnterNClass is called when production nClass is entered.
func (s *BaseLangListener) EnterNClass(ctx *NClassContext) {}

// ExitNClass is called when production nClass is exited.
func (s *BaseLangListener) ExitNClass(ctx *NClassContext) {}

// EnterClassContent is called when production classContent is entered.
func (s *BaseLangListener) EnterClassContent(ctx *ClassContentContext) {}

// ExitClassContent is called when production classContent is exited.
func (s *BaseLangListener) ExitClassContent(ctx *ClassContentContext) {}

// EnterFieldDeclaration is called when production fieldDeclaration is entered.
func (s *BaseLangListener) EnterFieldDeclaration(ctx *FieldDeclarationContext) {}

// ExitFieldDeclaration is called when production fieldDeclaration is exited.
func (s *BaseLangListener) ExitFieldDeclaration(ctx *FieldDeclarationContext) {}

// EnterMethodDeclaration is called when production methodDeclaration is entered.
func (s *BaseLangListener) EnterMethodDeclaration(ctx *MethodDeclarationContext) {}

// ExitMethodDeclaration is called when production methodDeclaration is exited.
func (s *BaseLangListener) ExitMethodDeclaration(ctx *MethodDeclarationContext) {}

// EnterIdentifier is called when production identifier is entered.
func (s *BaseLangListener) EnterIdentifier(ctx *IdentifierContext) {}

// ExitIdentifier is called when production identifier is exited.
func (s *BaseLangListener) ExitIdentifier(ctx *IdentifierContext) {}

// EnterLangType is called when production langType is entered.
func (s *BaseLangListener) EnterLangType(ctx *LangTypeContext) {}

// ExitLangType is called when production langType is exited.
func (s *BaseLangListener) ExitLangType(ctx *LangTypeContext) {}

// EnterValue is called when production value is entered.
func (s *BaseLangListener) EnterValue(ctx *ValueContext) {}

// ExitValue is called when production value is exited.
func (s *BaseLangListener) ExitValue(ctx *ValueContext) {}
