package badpkg

import (
	"fmt"
	"go/ast"
	"go/token"
)

// Foo is whatever
func Foo() {
	fmt.Println(ast.Bad)
	fmt.Println(token.CASE)
}
