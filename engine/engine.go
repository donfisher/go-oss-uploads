package engine

import "fmt"

type Engine struct {
	Analyzes Analyzes
	Uploads Uploads
}
type Analyzes interface {
	Check(path string) (bool bool,err error)
	IsFile(path string) (bool bool)
}
type Uploads interface {
	UploadOne(path string)
	UploadMany(path string)
}


func (e *Engine) Upload(path string)  {
	//分析输入是目录还是文件,以及文件的合法性
	b, err := e.Analyzes.Check(path)
	if err != nil || b == false{
		fmt.Println("目录或文件不存在，请核对后再试！")
	}
	//分别执行上传
	bool2 := e.Analyzes.IsFile(path)
	if bool2 {
		fmt.Printf("准备上传单个文件！%s",path)
		fmt.Println("")
		//上传单个
		e.Uploads.UploadOne(path)
	}else{
		fmt.Printf("准备批量上传文件！%s",path)
		fmt.Println("")
		//并发上传
		e.Uploads.UploadMany(path)
	}

}