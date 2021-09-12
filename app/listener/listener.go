package listener

import (
	"fmt"
	"strconv"

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
	c := lang_types.NewClass(ctx.Identifier().GetText(), expose, p)
	
	for _, v := range ctx.AllAnnotation() {
		a := ParseAnnotation(v.(*parser.AnnotationContext))
		c.AnnotationList = append(c.AnnotationList, a)
	}
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

/**
 * modifier interface {}
 */
func ParseAnnotation(ctx *parser.AnnotationContext) *lang_types.NAnnotation {
	a := lang_types.NewAnnotation(ctx.Identifier().GetText())

	for _, v := range ctx.AllKey_value() {
		a.Values = append(a.Values, ParseKeyValue(v.(*parser.Key_valueContext)))
	}
	return a
}

func ParseKeyValue(ctx *parser.Key_valueContext) *lang_types.KeyValue {

	k := ctx.Identifier().GetText()
	v := ctx.Value().(*parser.ValueContext)

	var nv interface{}
	if t := v.INTEGER(); t != nil {
		integer, err := strconv.Atoi(t.GetText())
		if err != nil {
			panic(fmt.Errorf("%s invalid INTEGER", k))
		}
		nv = integer
	}

	a := lang_types.NewKeyValue(k, nv)
	return a
}
