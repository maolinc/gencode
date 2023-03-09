package astx

import (
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/printer"
	"go/token"
	"os"
	"strings"
)

func MergeSource(srcFile, destText string) error {
	// 解析ast
	srcAst, err := parseFile(srcFile)
	if err != nil {
		return err
	}
	destAst, err := parseFile(destText)
	if err != nil {
		return err
	}
	pkg := ast.Package{
		Name:    "martifact",
		Scope:   nil,
		Imports: nil,
		Files:   map[string]*ast.File{"a.go": srcAst, "b.go": destAst},
	}
	mergeFile := ast.MergePackageFiles(&pkg, 7)
	fset := token.NewFileSet()
	open, err := os.Create(srcFile)
	if err != nil {
		return err
	}
	defer open.Close()
	err = printer.Fprint(open, fset, mergeFile)
	if err != nil {
		return err
	}
	return nil
}

func hasFunc(aFile *ast.File, funcDecl *ast.FuncDecl, replace bool) bool {
	name, methodName := getMethodId(funcDecl)
	decls := make([]ast.Decl, 0)
	for _, aDecl := range aFile.Decls {
		if aFunc, ok := aDecl.(*ast.FuncDecl); ok {
			aN, aM := getMethodId(aFunc)
			if aN == name && aM == methodName {
				// Same function name, check if signature matches
				if aFunc.Type.Params.NumFields() == funcDecl.Type.Params.NumFields() &&
					aFunc.Type.Results.NumFields() == funcDecl.Type.Results.NumFields() {
					if replace {
						aFunc = funcDecl

					}
					return true
				}
			}
		}
		decls = append(decls, aDecl)
	}
	return false
}

func parseBlockStmt(text string) (*ast.BlockStmt, error) {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "", text, parser.AllErrors)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	for _, decl := range f.Decls {
		if fn, ok := decl.(*ast.FuncDecl); ok {
			//fmt.Printf("Function %s\n", fn.Name.Name)
			//ast.Print(fset, fn.Body)
			return fn.Body, nil
		}
	}
	return nil, nil
}

func generateCode(node *ast.File) string {
	fset := token.NewFileSet()
	var buf strings.Builder
	if err := format.Node(&buf, fset, node); err != nil {
		panic(err)
	}
	return buf.String()
}

func getMethodId(fdecl *ast.FuncDecl) (structName, methodName string) {
	recv := fdecl.Recv
	if recv == nil || len(recv.List) != 1 {
		return "", fdecl.Name.Name
	}
	field, ok := recv.List[0].Type.(*ast.StarExpr)
	if !ok {
		return "", fdecl.Name.Name
	}
	ident, ok := field.X.(*ast.Ident)
	if !ok {
		return "", fdecl.Name.Name
	}
	return ident.Name, fdecl.Name.Name
}

func findMethod(node *ast.File, structName, methodName string) *ast.FuncDecl {
	for _, decl := range node.Decls {
		if fdecl, ok := decl.(*ast.FuncDecl); ok {
			recv := fdecl.Recv
			if recv == nil || len(recv.List) != 1 {
				continue
			}
			field, ok := recv.List[0].Type.(*ast.StarExpr)
			if !ok {
				continue
			}
			ident, ok := field.X.(*ast.Ident)
			if !ok || ident.Name != structName || fdecl.Name.Name != methodName {
				continue
			}
			return fdecl
		}
	}
	return nil
}

func parseFile(filePath string) (*ast.File, error) {
	fset := token.NewFileSet()
	return parser.ParseFile(fset, filePath, nil, parser.ParseComments)
}

func parseString(txt string) (*ast.File, error) {
	fset := token.NewFileSet()
	return parser.ParseFile(fset, "", txt, parser.ParseComments)
}

func modifyMethod(node *ast.FuncDecl, newBody *ast.BlockStmt) *ast.FuncDecl {
	node.Body = newBody
	return node
}
