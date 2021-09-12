// Code generated from Lang.g4 by ANTLR 4.9.2. DO NOT EDIT.

package parser // Lang

import "github.com/antlr/antlr4/runtime/Go/antlr"

// LangListener is a complete listener for a parse tree produced by LangParser.
type LangListener interface {
	antlr.ParseTreeListener

	// EnterProgram is called when entering the program production.
	EnterProgram(c *ProgramContext)

	// EnterNamespace is called when entering the namespace production.
	EnterNamespace(c *NamespaceContext)

	// EnterAnnotation is called when entering the annotation production.
	EnterAnnotation(c *AnnotationContext)

	// EnterKey_value is called when entering the key_value production.
	EnterKey_value(c *Key_valueContext)

	// EnterModifier is called when entering the modifier production.
	EnterModifier(c *ModifierContext)

	// EnterNInterface is called when entering the nInterface production.
	EnterNInterface(c *NInterfaceContext)

	// EnterNClass is called when entering the nClass production.
	EnterNClass(c *NClassContext)

	// EnterClassContent is called when entering the classContent production.
	EnterClassContent(c *ClassContentContext)

	// EnterFieldDeclaration is called when entering the fieldDeclaration production.
	EnterFieldDeclaration(c *FieldDeclarationContext)

	// EnterMethodDeclaration is called when entering the methodDeclaration production.
	EnterMethodDeclaration(c *MethodDeclarationContext)

	// EnterIdentifier is called when entering the identifier production.
	EnterIdentifier(c *IdentifierContext)

	// EnterLangType is called when entering the langType production.
	EnterLangType(c *LangTypeContext)

	// EnterValue is called when entering the value production.
	EnterValue(c *ValueContext)

	// ExitProgram is called when exiting the program production.
	ExitProgram(c *ProgramContext)

	// ExitNamespace is called when exiting the namespace production.
	ExitNamespace(c *NamespaceContext)

	// ExitAnnotation is called when exiting the annotation production.
	ExitAnnotation(c *AnnotationContext)

	// ExitKey_value is called when exiting the key_value production.
	ExitKey_value(c *Key_valueContext)

	// ExitModifier is called when exiting the modifier production.
	ExitModifier(c *ModifierContext)

	// ExitNInterface is called when exiting the nInterface production.
	ExitNInterface(c *NInterfaceContext)

	// ExitNClass is called when exiting the nClass production.
	ExitNClass(c *NClassContext)

	// ExitClassContent is called when exiting the classContent production.
	ExitClassContent(c *ClassContentContext)

	// ExitFieldDeclaration is called when exiting the fieldDeclaration production.
	ExitFieldDeclaration(c *FieldDeclarationContext)

	// ExitMethodDeclaration is called when exiting the methodDeclaration production.
	ExitMethodDeclaration(c *MethodDeclarationContext)

	// ExitIdentifier is called when exiting the identifier production.
	ExitIdentifier(c *IdentifierContext)

	// ExitLangType is called when exiting the langType production.
	ExitLangType(c *LangTypeContext)

	// ExitValue is called when exiting the value production.
	ExitValue(c *ValueContext)
}
