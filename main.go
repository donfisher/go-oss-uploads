package main

import (
	"fmt"
	"uploads/Analyzes"
	"uploads/Uploads"
	"uploads/engine"
)
var (
	path            string
)
func main() {
	fmt.Println("此程序上传指定目录文件到阿里云")
	fmt.Println("请输入目录或者文件完整地址")
	fmt.Scanln(&path)
	e := engine.Engine{
		Analyzes:&Analyzes.Analyzes{},
		Uploads:&Uploads.Uploads{},
	}
	e.Upload(path)
}
