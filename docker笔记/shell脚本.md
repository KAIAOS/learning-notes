### 一、shell中的变量

#### 1.1 基本语法

- 查看： echo $变量
- 定义变量: 变量=值------------等号前后不能有空格
- 撤销变量 unset  变量
- bash中，变量默认类型都是字符串类型，无法进行数值计算
- 变量值如果有空格，需要使用双引号或者单引号括起来
- 可把变量提升为全局变量，可供其他的Shell程序使用   export 变量名

#### 1.2 特殊变量

1. $n  (n为数字，$0代表该脚本名称，$1-$9代表第一到第九个参数)
2. $# （获取所有的输入参数的个数，常用于循环）
3. $*    （获取命令行中的所有参数，$*把所有参数当作一个整体）
4. $@    (也是获得命令行的所有参数 不过是每个参数区分对待)
5. $?    (最后一次执行命令的返回状态。如果值为0则是正常执行；非0则是上一个命令执行不正确，具体错误码是定义的)

#### 1.3 运算符

1. $((运算式子))   或者  $[运算式子]
2. expr    加减乘（\\*）除取余     运算符之间要有空格     expr 2 + 3

### 二、条件判断

#### 1 基本语法

[condition] (注意condition前后要有空格)

```bash
#中括号前后必须有空格
[ -w helloword.sh ] #判断是否有写权限
```

#### 2 常用判断条件

1. 两个整数之间的比较
   1. =字符串比较
   2. -lt、-gt、-le（小于等于）、-ge、-eq、-ne（不等 not equal）
2. 按照文件权限进行判断  -r -w -x
3. 按照文件类型进行判断   -f 文件存在并且是一个常规的文件 -e文件存在  -d文件存在并且是目录
4. 多条件判断
   1. （&& 表示前一条命令执行成功时，才执行后一条命令）
   2. ||表示上一条命令执行失败后，才执行下一条命令

#### 3 流程控制

1. ​	if判断

```bash
#[ 条件判断式 ]中括号前后必须有空格
#if 后要有空格
if [ 条件判断式 ];then
	命令
elif [ 条件判断 ];then
	命令
fi
#或者
if [ 条件判断式 ]
then
	命令
fi
```

2.  for循环

```bash
# 语法一
for(( 初始值;循环控制条件;变量变化 ))
	do
		命令
    done
#例子
s=0
for (( i=0;i<=100;i++ ))
do
	s=$[$s+$i]
done
echo $s
```

```bash
#语法二
for 变量 in 值1 值2 值3 ...
do
	命令
done

# 例子
#!/bin/bash
for i in $*   //此处若为 "$" 则是所有参数当作一个整体  "$@"仍是单独的参数
do
        echo "like $i"
done
```

3. while循环

```bash
while [ 判断式 ]
do
	命令
done

#例子
s=0
i=1
while [ $i -le 100 ]
do
	s=$[$S+$i]
	i=$[$i+1]
done
echo $s
```

### 三、read读取控制台输入

#### 1 基本语法

```bash
read (选项)(参数)
选项： -p指定读取值时的提示符；-t指定读取值时等待的时间
参数：指定读取值的变量名
#!/bin/bash
read -t 7 -p "Enter your name in 7 seconds" Name
echo $Name
```

### 四、函数

#### 1 系统函数

1. basename [path/filename] [suffix]
   1. basename /home/example.txt  输出 example.txt
   2. basename /home/example.txt .txt 输出example
2. dirname [文件绝对路径] 返回文件路径
   1. dirname /home/example.txt  输出 /home

#### 2 自定义函数

1. 基本用法

   1. 必须在使用之前声明函数
   2. 函数返回值，只能通过$?系统变量获得，可以显式加return返回，如果不加将以最后一条命令运行结果作为返回值。return之后跟数值n(0~255)

   ```bash
   function funcname(){
   	命令
   }
   funcname #调用
   ```

   ```bash
   #!/bin/bash
   #计算两数之和
   function sum(){
   	s=0;
   	s=$[$1+$2]
   	echo $s
   }
   
   read -p "input parameter1：" p1
   read -p "input parameter2：" p2
   sum $p1 $p2
   ```

### 五、Shell工具

#### 1 cut【选项参数】 filename

1. cut 【选项参数】 filename
   1. -f  列号，提取第几列   3- 表示第三列之后所有的
   2. -d 分隔符，按照指定分隔符分割列   默认是制表符

```bash
#cut.dong guan
shen zhen
wo  lai

cut -d " " -f 1,2 cut.txt #空格切割列并取1，2列
cat cut.txt|grep guan|cut -d " " -f2 #先以行过滤关键字
echo $PATH | cut -d " " -f 3-
```

#### 2 sed【选项】 filename

> sed是一种流编辑器，它一次处理一行内容。处理时，把当前处理的行储存在临时缓冲区中，称为“模式空间”，接着用sed命令处理缓冲区中的内容，处理完成后，把缓冲区的内容送往屏幕。接着处理下一行，这样不断重复，直到文件末尾。**文件内容并没有改变**。

1. 选项参数 -e 直接在列模式上进行sed的动作编辑
2. 命令功能描述
   1. a 	新增，a的后面接字符串，在下一行出现
   2. d     删除
   3. s      查找并替换

```bash
 sed "2a something" cut.txt #增加一行内容输出到屏幕
 sed "/wo/d" cut.txt #删除包含特定内容的行输出到屏幕 
 sed -e "2d" -e "s/wo/ni/g" cut.txt #删除第二行 并且全局替换wo为ni
```

#### 3 awk

1. awk [选项参数] ‘pattern1{action1} pattern2{action2}’ filename
   1. pattern:表示匹配模式 action：找到匹配内容后执行的一系列命令
   2. -F 指定输入文件的分隔符      -v赋值一个用户定义变量

```bash
awk -F : -v i=1 '{print $3+i}' passwd  #所有的第三列数字加一打印
```

#### 4 sort

1. sort  [选项参数] [待排序的文件列表]
   1. -n 依照数值大小排序
   2. -r以相反的顺序排序
   3. -t设置排序时所用的分割字符
   4. -k指定需要排序的列

```bash
sort -t : -nrk 2 sort.sh
```

#### 5 实例

```bash
#查询file1中空行所在的行号
awk '/^$/{pring NR}' file1

```

1197617594@qq.com