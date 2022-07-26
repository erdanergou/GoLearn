package main

/*
net/http包的用法，如何发请求
当需要频繁发送请求的时候（每五秒从阿里云同步接口数据）：定义一个全局的client，后续发请求的操作都使用这个全局的client


单元测试
单元测试文件名必须是xxx/ccc_test.go

go内置的测试工具
	go test
单元测试函数的格式
	import "testing"
	func  TestSplit(t *testing.T){

	}


性能测试/基准测试
函数格式
	func BenchmarkSplit(b *testing.B){
		// b.N被测试函数执行的次数
	}
执行命令：
	go test -bench -v

pprof
	记录cpu快照信息
	记录内存的快照信息

flag
 os.Args
	./xxx a b c
	osArgs:["./xxx""a""b""c"]
 flag标准库
	1.声明变量  flag.Type()   flag.TypeVar()
	nameval1 := flag.String("name","张三","姓名") // 返回的是指针变量

	var nameVal2 string
	flag.StringVal(&nameVal2,"name","张三","姓名")  // 把一个已经存在的变量绑定到命令行
	必须要调用
	flag.Parse() //解析命令传入的参数，赋值给对应的变量

	其他方法:
	flag.Args() //返回命令行参数后的其他参数，以[]string类型
	flag.NArg() // 返回命令行参数后的其他参数类型
	flag.NFlag() //返回使用的命令行参数
*/

func main() {

}
