package listener

import (
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
func (tree *TreeLangListener) EnterProgram(c *parser.ProgramContext) {

	namespace := c.Namespace().(*parser.NamespaceContext)

	pkg := new(lang_types.NPackage)
	pkg.Name = namespace.PACKAGE_NAME().GetText()

	nclass := c.AllNClass()
	for _, value := range nclass {
		ParseClass(&value, pkg)
	}
}

func ParseClass(ctx *parser.INClassContext, pkg *lang_types.NPackage) {

}
