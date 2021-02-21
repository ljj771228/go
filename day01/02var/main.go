package main //package是关键字，用来声明这个文件属于那个包（main包说明这个代码会编译成可执行文件）

//导入"fmt"标准库（go里的标准库）
import "fmt"

//上边的代码首先去GOROOT目录下加载fmt包，然后在main函数里调用fmt包的Println函数来打印

//go语言中声明变量推荐使用小驼峰式命名
//例如：var studentName string

//声明变量
//var name string
//var age int
//var isOk bool

//在函数体外声明的变量是全局变量，如下：
var (
	name string //“空字符串”
	age  int    // 0
	isOk bool   // fales
)

func main() { // func函数声明的意思（就是构造一个main函数） 是程序的入口，从main函数开始执行
	name = "李柚潼"
	age = 7
	isOk = true
	//var hahaha string
	//如上，GO语言中，非全局变量声明后必须要使用，否则编译不过去
	fmt.Print(isOk)               //在终端中输出要打印的内容
	fmt.Println()                 //打印一个换行空行
	fmt.Printf("name:%s\n", name) //%s是个占位符，\n是换行符号，使用后面name变量这个值，替换%s占位符
	fmt.Println(age, "岁")         //打印完指定内容后，会在后面增加一个换行符

	//声明变量并同时赋值
	var s1 string = "李佳锦"
	fmt.Println(s1)
	//类型推导，根据值判断该变量是什么类型
}
