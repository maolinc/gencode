package tools

import (
	"embed"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/printer"
	"go/token"
	"golang.org/x/mod/modfile"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func GetHomeDir() string {
	homeDir, _ := os.UserHomeDir()
	return homeDir
}

func CopyDirEm(dir embed.FS, toDir string, file string) (err error) {
	readDir, err := dir.ReadDir(file)
	if err != nil {
		return err
	}

	err = os.MkdirAll(toDir, 777)
	if err != nil {
		return err
	}

	for _, entry := range readDir {
		fileTo, err := os.OpenFile(toDir+"/"+entry.Name(), os.O_WRONLY|os.O_CREATE, os.ModePerm)
		if err != nil {
			return err
		}
		readFile, err := dir.ReadFile(file + "/" + entry.Name())
		_, err = io.WriteString(fileTo, string(readFile))
		if err != nil {
			return err
		}
		fileTo.Close()
	}
	return nil
}

func CopyDir(fromDir, toDir string) (err error) {
	fileFrom, err := os.Open(fromDir)
	if err != nil {
		return err
	}
	dirs, err := fileFrom.ReadDir(-1)
	if err != nil {
		return err
	}

	err = os.MkdirAll(toDir, 777)
	if err != nil {
		return err
	}

	for _, file := range dirs {
		openFile, err := os.OpenFile(fromDir+"/"+file.Name(), os.O_RDONLY, os.ModePerm)
		if err != nil {
			return err
		}
		fileTo, err := os.OpenFile(toDir+"/"+file.Name(), os.O_WRONLY|os.O_CREATE, os.ModePerm)
		if err != nil {
			return err
		}
		_, err = io.Copy(fileTo, openFile)
		if err != nil {
			return err
		}
		fileTo.Close()
		openFile.Close()
	}

	defer fileFrom.Close()

	return err
}

// PathExists 判断一个文件或文件夹是否存在
// 输入文件路径，根据返回的bool值来判断文件或文件夹是否存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func FindFileToBack(path, fileName string) (context []byte, destPath string, err error) {
	if filepath.IsAbs(path) {
		return findFile(path, fileName)
	}
	absPath, err := filepath.Abs(path)
	if err != nil {
		return nil, "", err
	}
	return findFile(absPath, fileName)
}

func findFile(path string, filename string) (context []byte, destPath string, err error) {
	fullPath := filepath.Join(path, filename)
	exist, err := PathExists(fullPath)
	if err != nil {
		return nil, "", err
	}
	if exist {
		if file, err := os.ReadFile(fullPath); err == nil {
			return file, path, nil
		}
		return nil, "", err
	}
	parentDir := filepath.Dir(path)
	if path == parentDir {
		return nil, parentDir, nil
	}
	return findFile(parentDir, filename)
}

func GetModule(path string) (module, fixPath string) {
	var (
		modName  = "go.mod11"
		context  []byte
		destPath string
		err      error
	)
	if !filepath.IsAbs(path) {
		if path, err = filepath.Abs(path); err != nil {
			return "", ""
		}
	}
	context, destPath, err = findFile(path, modName)
	if err != nil {
		return "", ""
	}
	modFile := modfile.ModulePath(context)
	return modFile, strings.TrimPrefix(path, destPath)
}

func ParseGoFile(path string, stcName, metName string) error {
	node, err := parseFile(path)
	if err != nil {
		panic(err)
	}
	node.Imports = append(node.Imports, &ast.ImportSpec{Path: &ast.BasicLit{Value: "github.com/maolinc/gencode/app/moddddddddddddel"}})

	methodNode := findMethod(node, stcName, metName)
	if methodNode == nil {
		panic("method not found")
	}
	src := `
package martifact




func (l *CreateMArtifactLogic) CreateMArtifact(req *types.CreateMArtifactReq) (resp *types.CreateMArtifactResp, err error) {

	//data := &model.MArtifact{}
	//_ = copier.Copiers(data, req)
	//err = l.svcCtx.MArtifactModel.Insert(l.ctx, nil, data)
	//if err != nil {
	//	return nil, errors.Wrapf(err, "req: %v", req)
	//}

	return &types.CreateMArtifactResp{}, nil
}
	`
	stmt, err := parseBlockStmt(src)
	if err != nil {
		return err
	}

	modifyMethod(methodNode, stmt)
	//code := generateCode(node)
	// 将修改后的 AST 转换为源代码
	f2, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	defer f2.Close()
	fset := token.NewFileSet()
	if err := printer.Fprint(f2, fset, node); err != nil {
		panic(err)
	}
	return err
}

func parseFile(filePath string) (*ast.File, error) {
	fset := token.NewFileSet()
	return parser.ParseFile(fset, filePath, nil, parser.ParseComments)
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

func modifyMethod(node *ast.FuncDecl, newBody *ast.BlockStmt) *ast.FuncDecl {
	node.Body = newBody
	return node
}

func generateCode(node *ast.File) string {
	fset := token.NewFileSet()
	var buf strings.Builder
	if err := format.Node(&buf, fset, node); err != nil {
		panic(err)
	}
	return buf.String()
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
