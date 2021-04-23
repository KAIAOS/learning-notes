### go module

go module是Go语言默认的依赖管理工具 使用go module管理依赖后会在项目根目录下生成两个文件`go.mod` `go.sum`

go mod命令

```shell
go mod download 下载依赖module到本地cache（默认为$GOPATH/pkg/mod目录）
go mod edit		编辑go.mod文件
go mod graph	打印模块依赖图
go mod init		初始化当前文件夹创建go.mod文件
go mod tidy		将增加缺少的module文件 删除无用的module
go mod vendor	将依赖复制到vendor下
go mod verify	校验依赖
go mod why		解释为什么需要依赖
```

go.mod 结构

```shell
module miaosha

go 1.15

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/gin-gonic/gin v1.6.3
	github.com/go-redis/redis/v8 v8.4.8
	github.com/jinzhu/gorm v1.9.16
	github.com/streadway/amqp v1.0.0
)
```

### go get

go get下载依赖包  go get -u 升级到最新的版本    go get -u=patch 将会升级到最新的修订版本

go get package@version升级到指定版本

下载所有依赖 go mod download

### Go标准库Context 非常重要

如何通知子goroutine退出

```go
var (
	wg sync.WaitGroup
	notify bool
)

func f() {
	defer wg.Done()
	for  {
		fmt.Println("aa")
		time.Sleep(time.Millisecond*500)
		if notify{
			break
		}
	}
}
func main() {
	wg.Add(1)
	go f()
	time.Sleep(time.Second*3)
	notify = true
	wg.Wait()
}
```

在go http包的server中，每一个请求都有一个对应的goroutine取处理。请求处理函数通常会启动额外的goroutine用来访问后端，比如数据库和RPC服务。用来处理一个请求的goroutine通常要访问一些与请求特定的数据，比如终端用户的身份认证信息，验证相关的token、请求的截止时间。当一个请求被取消或超时时，所有用来处理该请求的goroutine都应迅速退出，然后才能释放这些goroutine占的资源

```go
//使用context的版本
var (
	wg sync.WaitGroup
)

func f(ctx context.Context) {
	defer wg.Done()
	flag := false
	for  {
		if flag {
			break
		}
		fmt.Println("aa")
		time.Sleep(time.Millisecond*500)
		select {
		case <- ctx.Done():   // ctx.Done()返回一个只读的通道
			flag = true
		default:
		}
	}
}
func main() {
	ctx, cancel := context.WithCancel(context.Background())
	wg.Add(1)
	go f(ctx)
	time.Sleep(time.Second*3)
	//如何通知子Goroutine退出  只需要调用cancel()
	cancel()
	wg.Wait()
}
```

![image-20210423140620841](D:\markdown\golang笔记\08gomodule和context.assets\image-20210423140620841.png)

![image-20210423140827955](D:\markdown\golang笔记\08gomodule和context.assets\image-20210423140827955.png)

![image-20210423140904004](D:\markdown\golang笔记\08gomodule和context.assets\image-20210423140904004.png) 

![image-20210423141242756](D:\markdown\golang笔记\08gomodule和context.assets\image-20210423141242756.png)

```go
func main() {
	d := time.Now().Add(50* time.Millisecond)
	ctx, cancel := context.WithDeadline(context.Background(), d)

	//尽管ctx会过期，但是在任何情况下调用它的cancel函数都是很好的实践
	//如果不这样做，可能会使得上下文及其父类存活时间超过必要时间
	defer cancel()

	select {
		case <-time.After(1 * time.Second):
			fmt.Println("zzz")
		case <-ctx.Done():
			fmt.Println(ctx.Err())
	}
}
```

![image-20210423142129520](D:\markdown\golang笔记\08gomodule和context.assets\image-20210423142129520.png)

```go
var wg sync.WaitGroup

func worker(ctx context.Context){
   LOOP:
      for  {
         fmt.Println("db connecting")
         time.Sleep(time.Millisecond*10)
         select {
         case <- ctx.Done():
            break LOOP
         default:

         }
      }
      fmt.Println("worker done")
      wg.Done()
}

func main() {

   ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*50)

   wg.Add(1)
   go worker(ctx)
   time.Sleep(time.Second * 5)
   cancel()
   wg.Wait()
   fmt.Println("over")
}
```

![image-20210423142947494](D:\markdown\golang笔记\08gomodule和context.assets\image-20210423142947494.png)

**仅对API和进程之间传递请求域数据使用上下文**，而不是传递参数给函数

所提供的键必须是可比较的，并且不是string类型或任何其他内置内省，以避免使用上下文在包之间发生冲突，

WithValue的用户应该为键定义自己的类型，为了避免在分配给interface{}时进行分配，上下文键通常具有具体类型struct{}或者导出上下文关键变量的静态类型应该时指针或接口

```go
type TraceCode string

var wg sync.WaitGroup

func worker(ctx context.Context){
	key:= TraceCode("TRACE_CODE")
	traceCode, ok := ctx.Value(key).(string)
	if !ok {
		fmt.Println("invalid trace code")
	}

LOOP:
	for {
		fmt.Printf("worker,traceCode : %s\n", traceCode)
		time.Sleep(time.Millisecond*10) //假设正常连接数据库耗时10毫秒
		select {
		case <- ctx.Done():
			break LOOP
		default:

		}
	}
	fmt.Println("worker done!")
	wg.Done()
}

func main() {

	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*50)
	ctx = context.WithValue(ctx, TraceCode("TRACE_CODE"), "123456")
	wg.Add(1)
	go worker(ctx)
	time.Sleep(time.Second * 5)
	cancel()
	wg.Wait()
	fmt.Println("over")
}
```

```
func TestGetAll(t *testing.T) {
   all, err := GetAll()
   if err != nil{
      panic(err)
   }
   for _,v := range all{
      t.Logf("is:  %#v",v)
   }
}
```