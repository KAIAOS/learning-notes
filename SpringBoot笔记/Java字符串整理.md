### 内存示意图

8个基本类型 对应有包装类 （引用类型）

<img src="C:\Users\hanka\AppData\Roaming\Typora\typora-user-images\image-20200619221021844.png" alt="image-20200619221021844" style="zoom: 67%;" />

### 1、字符串常量池

​	对于直接写上双引号的字符串就在常量池中

​	对于基本类型来说 ==是进行数值比较

​	对于引用类型来说，==是进行地址值比较

<img src="C:\Users\hanka\AppData\Roaming\Typora\typora-user-images\image-20200619221327697.png" alt="image-20200619221327697" style="zoom:80%;" />

### 2、字符串常用比较方法

public boolean equals(Object, obj) 参数是字符串且内容相同才会给出true

``` java
//public boolean equals(Object, obj) 
//public boolean equalsIgnoreCase(Object, obj) 不区分大小写
String str = "hello"；
"abc".equals(str)； //推荐
str.equals("abc")； //不推荐
```

### 3、字符串获取方法

> 字符串是不可变类型，每次获取都是新的字符串
>
> 变化的是地址值而不是字符串

``` java
public int length()
public String concat(String str)
public char charAt(int index)
public int indexOf(String str)//查找字符串在本字符串中首次出现的索引位置 如果没有返回-1
```



### 4、字符串的截取方法

``` java
public String substring(int index);//截取从参数位置一直到字符串末尾，返回新字符串
public String substring(int begin, int end);//截取从begin还是到end结束[begin,end)
```

### 5、字符串的转换相关方法

``` java
public char[] toCharArray();//返回字符数组
public byte[] getBates();//返回当前字符串底层字节数组
public String replace(CharSequence oldString, Charsequence newString);//字符串全局替换
```

### 6、字符串分割方法

``` java
public String[] split(String regex);//将字符串切分成若干部分
//split参数是一个正则表达式 例如. 必须写\\.
```



### java中用于处理字符串常用的有三个类:

### 1、java.lang.String

### 2、java.lang.StringBuffer

### 3、java.lang.StrungBuilder

三者共同之处:都是final类,不允许被继承，主要是从性能和安全性上考虑的，因为这几个类都是经常被使用着，且考虑到防止其中的参数被参数修改影响到其他的应用。

StringBuffer是线程安全，可以不需要额外的同步用于多线程中;

StringBuilder是非同步,运行于多线程中就需要使用着单独同步处理，但是速度就比StringBuffer快多了;

StringBuffer与StringBuilder两者共同之处:可以通过append、indert进行字符串的操作。