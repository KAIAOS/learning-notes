![image-20200619231044900](C:\Users\hanka\AppData\Roaming\Typora\typora-user-images\image-20200619231044900.png)

``` java
/*
static 修饰成员变量时，变量不再属于对象而是属于类，被所有对象共享

static 修饰成员方法时，就是静态方法，同样属于类，类方法
如果没有static关键字那么必须首先创建对象，然后通过对象才能使用他 类名称.静态方法名 
1、静态方法不能访问直接访问非静态变量，因为在内存中先有静态后有非静态
2、静态不能使用this关键字
*/
```

### 在方法区中有一个独立的空间存储静态区

![image-20200619232135069](C:\Users\hanka\AppData\Roaming\Typora\typora-user-images\image-20200619232135069.png)

### 静态代码块

``` java
//特点时当第一次使用到本类时，静态代码块执行唯一的一次，静态代码块比构造方法先执行
public class 类名称{
     static{
         //静态代码块内容
     }
 }
```

