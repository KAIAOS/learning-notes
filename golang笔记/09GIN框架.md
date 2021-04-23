### 一、gin框架

- Gin是一个golang的为框架，封装比较优雅，API友好，源码注释比较明确具有快速灵活 容错方便等特点
- 对于golang来说 web框架的依赖远比python、java之类的小。自身的net/http足够简单，性能也非常不错
- 借助框架开发，不仅可以省去很多常用的封装带来的时间也有助于编码风格形成

hello world

#### 基本路由

gin框架中路由是基于httprouter做的

```go
func main(){
   // 创建路由
    //默认使用了2个中间件 Logger() Recovery()
    //也可以使用不带中间件的路由  r := gin.New()
   r := gin.Default()
    
   
   //绑定路由规则 执行的函数
   //gin.Context 封装了request和response
   r.GET("/", func(c *gin.Context) {
      c.String(http.StatusOK, "hello world!")
   })
   
    r.POST("/xxxPost/", getting)
	
   //监听端口 默认在8080
   r.Run(":8000")
}

func getting(c *gin.Context){
    
}
```

#### API参数

- 可以通过Context的Param方法获取API参数
- localhost:8000/user/zhangshan
- /user/:name/*action    action := c.Param("action")  

```go
func main(){
   // 创建路由
   r := gin.Default()
   //绑定路由规则 执行的函数
   //gin.Context 封装了request和response
   r.GET("/user/:name/", func(c *gin.Context) {
      name := c.Param("name")
      c.String(http.StatusOK, name)
   })

   //监听端口 默认在8080
   r.Run(":8000")
}
```

#### URL参数

- URL参数可以通过DefaultQuery() 或Query()方法获取
- DefaultQuery() 若参数不存在 则返回默认值  Query()若不存在返回空串
- API？name=zs

```go
func main(){
   // 创建路由
   r := gin.Default()
   //绑定路由规则 执行的函数
   //gin.Context 封装了request和response
   r.GET("/user", func(c *gin.Context) {
      name := c.DefaultQuery("name","ccc")
      c.String(http.StatusOK, name)
   })

   //监听端口 默认在8080
   r.Run(":8000")
}
```

#### 表单参数

DefaultPostForm（）

PostForm（）

PostFormArray()

```go
r.POST("/form",func(c *gin.Context) {
    	type1 := c.DefaultPostForm("type", "alert")
      c.String(http.StatusOK, name)
   })
```

#### 路由组

routes group 是为了管理一些相同的url

v1的submit应该是有问题的

![image-20210423155331706](D:\markdown\golang笔记\09GIN框架.assets\image-20210423155331706.png)

#### 路由原理

httprouter会将所有路由规则构造一棵前缀树

### 二、gin数据解析与绑定

#### 1、json数据解析与绑定

content-type:application/json

客户端传参 后端接收解析并绑定结构体

c.ShouldBind(&json)

```go
func main(){
   // 创建路由
   r := gin.Default()

   r.POST("/login", func(c *gin.Context) {
      var json Login
      if err := c.ShouldBind(&json);err !=nil{
         //返回错误信息
         c.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
         return
      }
      //将request的body中的数据 自动按照json格式解析到结构体
      if json.User != "root" || json.Password != "admin"{
         c.JSON(http.StatusBadRequest, gin.H{"status":"304"})
         return
      }
      c.JSON(http.StatusOK, gin.H{"status":"200"})
   })

   //监听端口 默认在8080
   r.Run(":8000")
}
```

#### 2、表单数据解析与绑定

 c.Bind(&json)

```go
func main(){
   // 创建路由
   r := gin.Default()

   r.POST("/login", func(c *gin.Context) {
      var json Login
      if err := c.Bind(&json);err !=nil{
         //返回错误信息
         c.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
         return
      }
      //将request的body中的数据 自动按照json格式解析到结构体
      if json.User != "root" || json.Password != "admin"{
         c.JSON(http.StatusBadRequest, gin.H{"status":"304"})
         return
      }
      c.JSON(http.StatusOK, gin.H{"status":"200"})
   })

   //监听端口 默认在8080
   r.Run(":8000")
}
```

#### 3、URI数据解析与绑定

http://localhost:8000/root/admin

```go
func main(){
   // 创建路由
   r := gin.Default()

   r.GET("/login/:user/:password", func(c *gin.Context) {
      var json Login
      if err := c.ShouldBindUri(&json);err !=nil{
         //返回错误信息
         c.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
         return
      }
      //将request的body中的数据 自动按照json格式解析到结构体
      if json.User != "root" || json.Password != "admin"{
         c.JSON(http.StatusBadRequest, gin.H{"status":"304"})
         return
      }
      c.JSON(http.StatusOK, gin.H{"status":"200"})
   })

   //监听端口 默认在8080
   r.Run(":8000")
}
```

### 三、多种响应方式

```go
func main(){
   // 创建路由
   r := gin.Default()
   //1 json
   r.GET("json", func(c *gin.Context) {
      c.JSON(200,gin.H{"message":"someJson","status":200})
   })

   //2 结构体
   r.GET("/struct", func(c *gin.Context) {
      var msg struct{
         Name string
         Message string
         Number int
      }
      msg.Name="aa"
      msg.Message="asda"
      msg.Number=123
      c.JSON(200, msg)
   })
    
    
   //监听端口 默认在8080
   r.Run(":8000")
}
```

#### 异步执行

```go
func main(){
   // 创建路由
   r := gin.Default()
   //异步 需要搞一个context副本
   r.GET("/long", func(c *gin.Context) {
      copyContext := c.Copy()
      go func(){
         time.Sleep(time.Second*3)
         log.Println("异步执行了"+ copyContext.Request.URL.Path)
      }()

   })

   //监听端口 默认在8080
   r.Run(":8000")
}
```

### 四、 cookie

 

![image-20210423171232480](D:\markdown\golang笔记\09GIN框架.assets\image-20210423171232480.png)

#### cookie 中间件

![image-20210423171920451](D:\markdown\golang笔记\09GIN框架.assets\image-20210423171920451.png)

#### session中间件  