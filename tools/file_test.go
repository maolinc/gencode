package tools

import (
	"fmt"
	"github.com/maolinc/gencode/tools/astx"
	"github.com/maolinc/gencode/tools/mergex"
	"log"
	"os"
	"testing"
)

func TestReadFile(t *testing.T) {
	//println(filepath.Dir(`a/b/`))
	//fN := "go.mod"
	path := "C:\\Users\\maolin.chen\\Desktop\\me\\workspace\\xyz\\app\\artifact\\model"
	module, fixPath := GetModule(path)

	log.Println(module, fixPath)
	//abs, _ := filepath.Abs(path)
	//fileContent, err := FindFileToBack(path, fN)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Println(string(fileContent))
}

func TestParseGoFile(t *testing.T) {
	goF := "C:\\Users\\maolin.chen\\Desktop\\me\\workspace\\generatecode\\app\\api\\internal\\logic\\martifact\\create_m_artifact_logic.go"
	//dest := `
	//	package martifact
	//
	//	import (
	//		"github.com/maolinc/copier"
	//		"github.com/maolinc/gencode/app/model"
	//		"github.com/pkg/errors"
	//	)
	//
	//	func (l *CreateMArtifactLogic) CreateMArtifact(req *types.CreateMArtifactReq) (resp *types.CreateMArtifactResp, err error) {
	//
	//		data := &model.MArtifact{}
	//		_ = copier.Copiers(data, req)
	//		err = l.svcCtx.MArtifactModel.Insert(l.ctx, nil, data)
	//		if err != nil {
	//			return nil, errors.Wrapf(err, "req: %v", req)
	//		}
	//
	//		return &types.CreateMArtifactResp{}, nil
	//	}
	//`
	err := astx.MergeSource(goF, "C:\\Users\\maolin.chen\\Desktop\\me\\workspace\\generatecode\\tools\\create_m_artifact_logic_back.txt")
	if err != nil {
		log.Println(err)
		return
	}
	//fset := token.NewFileSet()
	//f, err := parser.ParseFile(fset, "src.go", src, parser.AllErrors)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//
	//for _, decl := range f.Decls {
	//	if fn, ok := decl.(*ast.FuncDecl); ok {
	//		fmt.Printf("Function %s\n", fn.Name.Name)
	//		ast.Print(fset, fn.Body)
	//	}
	//}
}

func TestMerge(t *testing.T) {
	f1 := "C:\\Users\\maolin.chen\\Desktop\\me\\workspace\\generatecode\\tools\\create_m_artifact_logic_back.txt"
	f2 := "C:\\Users\\maolin.chen\\Desktop\\me\\workspace\\generatecode\\tools\\create_m_artifact_logic_back.txt"
	f3 := "C:\\Users\\maolin.chen\\Desktop\\me\\workspace\\generatecode\\tools\\create_m_artifact_logic_back.go"
	err := mergex.MergeFiles(f1, f2, f3)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("Files merged successfully.")
}
