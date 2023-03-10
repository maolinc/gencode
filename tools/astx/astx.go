package astx

import (
	"fmt"
	"github.com/maolinc/gencode/tools/filex"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"golang.org/x/tools/go/ast/astutil"
	"os"
	"strings"
)

func MergeSource(srcFile, destFile string, pkgName string) error {
	fset := token.NewFileSet()
	srcAst, err := parseFile(srcFile)
	if err != nil {
		return err
	}
	destAst, err := parseFile(destFile)
	if err != nil {
		return err
	}

	mergeImport(srcAst, destAst)
	pkg := ast.Package{
		Name:    pkgName,
		Scope:   nil,
		Imports: nil,
		Files:   map[string]*ast.File{"a.go": srcAst, "b.go": destAst},
	}
	mergeFile := ast.MergePackageFiles(&pkg, 7)

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

func mergeImport(src *ast.File, dest *ast.File) {
	fset := token.NewFileSet()
	destImports := astutil.Imports(fset, dest)
	for _, destImport := range destImports {
		for _, spec := range destImport {
			name := strings.Trim(spec.Path.Value, "\"")
			astutil.DeleteImport(fset, dest, name)
			astutil.AddImport(fset, src, name)
		}
	}
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
			return fn.Body, nil
		}
	}
	return nil, nil
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
	var (
		err  error
		err1 error
	)
	file, err := parseString(filePath)
	if err != nil {
		filePath = filex.GetAbs(filePath)
		if file, err1 = parser.ParseFile(fset, filePath, nil, parser.ParseComments); err1 != nil {
			return nil, err1
		}
		return file, nil
	}
	return file, nil
}

func parseString(txt string) (*ast.File, error) {
	fset := token.NewFileSet()
	return parser.ParseFile(fset, "", txt, parser.ParseComments)
}
