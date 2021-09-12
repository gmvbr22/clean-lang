package listener

import (
	"fmt"

	"github.com/gmvbr/clean-lang/app/lang_types"
	"github.com/gmvbr/clean-lang/app/parser"
)

type TreeLangListener struct {
	*parser.BaseLangListener
}

func NewTreeLangListener() *TreeLangListener {
	return new(TreeLangListener)
}

/**
 * package name;
 *
 * => class
 * => interface
 */
func (tree *TreeLangListener) EnterProgram(ctx *parser.ProgramContext) {
	namespace := ctx.Namespace().(*parser.NamespaceContext)

	p := lang_types.NewPackage(namespace.PACKAGE_NAME().GetText())

	for _, v := range ctx.AllNClass() {
		ParseClass(v.(*parser.NClassContext), p)
	}
	for _, v := range ctx.AllNInterface() {
		ParseInterface(v.(*parser.NInterfaceContext), p)
	}
}

/**
 * @annotation(arg1=value, arg2=value) modifier class {}
 */
func ParseClass(ctx *parser.NClassContext, p *lang_types.NPackage) {
	modifier := ctx.Modifier()
	expose := true
	if modifier != nil {
		expose = modifier.GetText() == "public"
	}
	fmt.Println(expose)
	lang_types.NewClass(ctx.Identifier().GetText(), expose, p)
}

/**
 * modifier interface {}
 */
func ParseInterface(ctx *parser.NInterfaceContext, p *lang_types.NPackage) {
	modifier := ctx.Modifier()
	expose := true
	if modifier != nil {
		expose = modifier.GetText() == "public"
	}
	lang_types.NewInterface(ctx.Identifier().GetText(), expose, p)
}
