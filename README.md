<p align="center">
<h2>
go本地上传文件到阿里云OSS插件
</h2>
</p>

## About go-oss-uploads
1.支持单个文件上传到OSS <br>
2.支持目录下所有文件并发上传到OSS  <br>

## Usage
1、确保GO已经安装<br>
2、确保GO环境变量和GOPATH环境变量正确设置<br>
3、将本项目clone到GOPATH下的src目录下，并安装aliyun-oss-go-sdk (git clone https://github.com/aliyun/aliyun-oss-go-sdk.git)<br>
4、修改uploads目录下upload配置<br><br>

const endpoint = "oss-cn-hangzhou.aliyuncs.com" // oss endpoint <br>
const accessKeyId = "xxx" // oss key <br>
const accessKeySecret = "xxx" //oss secret<br>
const bucketName = "xxx" //oss bucket名称<br>
const objectName = "" //oss远程目标地址<br>
const workerCount = 100 //设置最大并发数<br>
const suffix = "" //筛选目录下需要上传的格式<br><br>

5、运行main.go。选择目录或文件，上传至阿里云OSS。<br>

## License

licensed under the [MIT license](https://opensource.org/licenses/MIT).