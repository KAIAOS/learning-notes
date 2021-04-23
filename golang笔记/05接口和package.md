### 接口interface

不关心变量是什么类型 只关心能调用什么方法

接口是一个类型

type 接口名 interface{

​	方法名 （参数，参数）（返回值，返回值）

}

用来给变量、参数、返回值 设置类型

```go
type speaker interface{  //定义一个可以叫的类型
    speak()				//方法签名
}

type cat struct{}

type dog struct{}


func (c cat) speak(){
    fmt.Println("miao")
}

func (d dog) speak(){
    fmt.Println("wang")
}

func da(x speaker){
    x.speak()
}

func main(){
    var c1 cat
    var d1 dog
    da(c1)
    da(d1)
}
```

### 接口的实现

如果一个变量实现了接口中规定的所有方法，那么这个变量就实现了这个接口，可以成为这个接口类型的变量

接口保存的分为两部分 值类型和值本身，这样可以实现对不同值的存储

<img src="D:\markdown\golang笔记\05接口.assets\image-20210422150716852.png" alt="image-20210422150716852" style="zoom: 33%;" />![image-20210422150739027](D:\markdown\golang笔记\05接口.assets\image-20210422150739027.png)

<img src="D:\markdown\golang笔记\05接口.assets\image-20210422150755412.png" alt="image-20210422150755412" style="zoom: 33%;" />

### 指针接收者和值接收者区别

使用值接收者实现接口： 存结构体或者接口体指针都可

使用指针接收者实现接口：只能存结构体指针

![image-20210422151252133](D:\markdown\golang笔记\05接口.assets\image-20210422151252133.png)

### 接口和类型的关系

多个类型可以实现同一个接口

一个类型可以实现多个接口

接口也可以嵌套

```go
type animal interface { //接口也可以嵌套
    mover
    eater
}

type mover interface {
    move()
}

type eater interface {
    eat(string)
}

type cat struct {
    name string
    feet int8
}

func (c *cat) move(){
    fmt.Println("move")
}

func (c *cat) eat(food string){
    fmt.Println(food)
}

func main(){
	c1 := cat{
		name: "ccc",
		feet: 2,
	}
	var a1 animal
	a1 = &c1
	fmt.Println(a1)
}
```

### 空接口

空接口可以接受任何类型，解释： 因为实现了接口的所有方法就是实现了该接口，所以任何变量都实现了空接口，所以可以接受所有变量

```go
interface{}  //空接口  没必要起名字
var m1 map[string]interface{}
m1 = make(map[string]interface{}, 16)
```

### 类型断言

变量，ok ：= x.(T)

 x标识类型为空接口的变量   T标识可能的类型  返回第一个是转换后的变量，第二个是布尔值true断言成功

```go
// 第一种
func assign(a interface{}) {
    str, ok = a.(string)
    if !ok {
        fmt.Println("猜错了")
    }else{
        fmt.Println("猜对了")
    }    
}

//switch case 实现
func assign(a interface{}) {
	switch t := a.(type) {
	case string:
	fmt.Println("string",t)
	case int:
	fmt.Println("int",t)
	default:
	fmt.Println("unsupport type",t)
	}
}
func main(){
	assign(true)
}
```

### package概念

包名是从 $GOPATH/src/后开始计算的

想被别的包调用的标识符都要首字母大写

导入包不使用时 需要匿名导入（init（））

禁止循环导入

```go
package calc

func Add(a, b int)(res int){
   res = a + b
   return
}
```

```go
//main.go
import (
   zhou "./calc" //起别名
   "fmt"
) // 导入内置 fmt 包

func main(){
	fmt.Println(zhou.Add(1,2))
}
```

### init函数

go执行导入包会自动触发内部init()函数，init()函数没有参数和返回值，不能主动调用

<img src="D:\markdown\golang笔记\05接口.assets\image-20210422155324298.png" alt="image-20210422155324298" style="zoom:33%;" />

```go
func init(){
    fmt.Println("导入自动执行")
}
```

多个包时：

<img src="D:\markdown\golang笔记\05接口.assets\image-20210422155430250.png" alt="image-20210422155430250" style="zoom: 50%;" />