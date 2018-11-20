package Uploads

import (
	"fmt"
	"github.com/aliyun-oss-go-sdk/oss"
	"io/ioutil"
	"os"
	"strings"
	"errors"
)
type Worker struct {
	in chan string
	done chan bool
}

type Uploads struct {
}

const endpoint = "oss-cn-hangzhou.aliyuncs.com" // oss endpoint
const accessKeyId = "xxx" // oss key
const accessKeySecret = "xxx" //oss secret
const bucketName = "xxx" //oss bucket名称
const objectName = "" //oss远程目标地址
const workerCount = 100 //设置最大并发数
const suffix = "" //筛选目录下需要上传的格式



//上传单个文件
func (u *Uploads) UploadOne(path string) {
	// 创建OSSClient实例。
	localFileName := path
	split := strings.Split(localFileName, `\`)
	length := len(split)
	fileName := split[length-1:length]



	client, err := oss.New(endpoint, accessKeyId, accessKeySecret)
	if err != nil {
		handleError(err)
	}
	// 获取存储空间。
	bucket, err := client.Bucket(bucketName)
	if err != nil {
		handleError(err)
	}
	up(bucket,objectName+fileName[0],localFileName)
}
func createWorker(bucket *oss.Bucket,objectName string, id int) Worker {
	w := Worker{
		in : make(chan string),
		done: make(chan bool),
	}
	go doWork(id , w.in , w.done,bucket,objectName)
	return w
}

//并发上传目录下所有文件
func (u *Uploads) UploadMany(path string) {
	files, err := getAllFiles(path)
	if err != nil{
		handleError(err)
	}
	fileCount := len(files)
	if fileCount == 0{
		handleError(errors.New("目录下没有指定文件，请重试！"))
	}
	client, err := oss.New(endpoint, accessKeyId, accessKeySecret)
	if err != nil {
		handleError(err)
	}
	bucket, err := client.Bucket(bucketName)
	if err != nil {
		handleError(err)
	}
	var workers [workerCount] Worker
	for i:=0;i<fileCount;i++{
		workers[i] = createWorker(bucket,objectName,i)
	}
	for i:=0;i<fileCount;i++{
		workers[i].in <- files[i]
	}
	for i:=0;i<fileCount;i++{
		<-workers[i].done
	}
}

func handleError(err error) {
	fmt.Println("Error:", err)
	os.Exit(-1)
}

func up(bucket *oss.Bucket,objectName string ,localFileName string)  {
	// 上传文件。
	err := bucket.PutObjectFromFile(objectName, localFileName)
	if err != nil {
		handleError(err)
	}
}

func doWork(id int ,c chan string,done chan bool,bucket *oss.Bucket,objectName string){
	for n := range c {
		split := strings.Split(n, `\`)
		length := len(split)
		fileName := split[length-1:length]
		fmt.Printf("worker : %d, object: %s, uploading file %v \n",id,objectName,n)
		up(bucket,objectName+fileName[0],n)
		go func() {done <- true}()
	}
}

//获取指定目录下的所有文件,包含子目录下的文件
func getAllFiles(dirPth string) (files []string, err error) {
	var dirs []string
	dir, err := ioutil.ReadDir(dirPth)
	if err != nil {
		return nil, err
	}
	PthSep := string(os.PathSeparator)
	for _, fi := range dir {
		if fi.IsDir() { // 目录, 递归遍历
			dirs = append(dirs, dirPth+PthSep+fi.Name())
			getAllFiles(dirPth + PthSep + fi.Name())
		} else {
			// 过滤指定格式
			if suffix != ""{
				ok := strings.HasSuffix(fi.Name(), suffix)
				if ok {
					files = append(files, dirPth+PthSep+fi.Name())
				}
			}else{
				files = append(files, dirPth+PthSep+fi.Name())
			}
		}
	}

	return files, nil
}
