### 一、指针

Go语言中不存在指针操作，只需要记住两个符号

&取地址   *根据地址取值

```go
func main() {
    n := 18
    p := &n
    fmt.Println(p)
    fmt.Printf("%T\n",p)
    
    m := *p
    fmt.Println(m)
    fmt.Printf("%T\n",m)
}
```

### 二、new

```go
func main() {
    var a1 *int // nil pointer
    fmt.Println(a1)
    var a2 = new(int)
    fmt.Println(a2)
    fmt.Println(*a2) //0
    *a2 = 100
    fmt.Println(*a2) //100
}
```

实际上 make也是分配内存的，与new不同的是，make作用于slice、map、chan的内存创建，而且它返回的类型就是这三个基本类型本身，而不是他们的指针类型，因为这三种类型就是引用类型，所以没必要返回他们的指针

### 三、map

> map是一种无序的基于k-v的数据结构 map是引用类型 必须初始化才能使用
>
> 如果不存在该某个key  则返回map对应值类型的零值
>
> var m1 map[key类型] 值类型

```go
func main(){
    var m1 map[string]int //此时map为空 还是nil
    m1 = make(map[string]int, 10)//要估算好map容量，避免在程序运行时动态扩容
    m1["aa"] = 1
    
    v, ok := m1["cc"]
    if !ok{
		fmt.Println("没有key")
	}else{
		fmt.Println(v)
	}
    fmt.Println(m1)
}
```

#### map遍历
```go
func main(){
    var m1 map[string]int //此时map为空 还是nil
    m1 = make(map[string]int, 10)//要估算好map容量，避免在程序运行时动态扩容
    m1["aa"] = 1
    m1["bb"] = 2
    v, ok := m1["cc"]
    for k,v := range m1 {
    	fmt.Println(k,v)
    }
    for k := range m1 {
    	fmt.Println(k)
    }
    for _,v := range m1 {
    	fmt.Println(v)
    }
    fmt.Println(m1)
}
```

#### map 删除

```go
var m1 map[string]int //此时map为空 还是nil
m1 = make(map[string]int, 10)//要估算好map容量，避免在程序运行时动态扩容
m1["aa"] = 1
m1["bb"] = 2
delete(m1, "aa")
fmt.Println(m1)
delete(m1, "cc") //删除一个不存在的key   什么都不做 不报错
```

#### 按照指定顺序遍历map

```go
func main(){
	rand.Seed(time.Now().UnixNano())

	var scoreMap = make(map[string] int, 200)

	for i := 0; i < 100; i++ {
		key := fmt.Sprintf("stu%02d",i)
		value := rand.Intn(100)
		scoreMap[key] = value
	}

	var keys = make([]string, 0, 200)
	for key := range scoreMap {
		keys = append(keys, key)
	}

	sort.Strings(keys)

	for _, key := range keys{
		fmt.Println(key, scoreMap[key])
	}
}
```

#### map和slice组合

```go
//值为map的切片
func main() {
    var s1 = make([]map[int]string, 10, 10) //这里仅仅对切片初始化 而里面的map还没有初始化
    s1[0] = make(map[int]string, 1) //对map初始化
    s1[0][10] = "sasdasdad"
    fmt.Println(s1)
}

//值为切片的map
func main() {
   var m = make(map[string][]int, 10)
    m["beijing"] = []int{10,20,30}
    fmt.Println(m)
}
```

