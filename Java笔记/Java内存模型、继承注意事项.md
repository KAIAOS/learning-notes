继承中方法重写

@override

1、必须保证父子类的方法名称相同参数列表也相同

2、子类方法返回值必须小于等于父类方法的返回值范围

3、java.lang.Object是所有类的公共最高父类

4、子类的方法权限必须大于等于父类方法的权限修饰符

​	public>protected>(default)>private 	(default)是什么都不写

>  只有子类的构造方法才能调用父类的构造方法super()
>
>  super的父类构造调用必须是子类构造方法的第一个语句唯一一个

<img src="C:\Users\hanka\AppData\Roaming\Typora\typora-user-images\image-20200620000159888.png" alt="image-20200620000159888" style="zoom:67%;" />

### 内存模型

``` 
jvisualvm
```



![image-20200620000523836](C:\Users\hanka\AppData\Roaming\Typora\typora-user-images\image-20200620000523836.png)

![image-20200715180940934](C:\Users\hanka\AppData\Roaming\Typora\typora-user-images\image-20200715180940934.png)

#### * Java常用的内存区域

1、栈内存空间：Java栈的区域很小，只有1M，特点是存取速度很快，所以在stack中存放的都是快速执行的任务，每个方法执行时都会创建一个栈帧（Stack Frame），描述的是java方法执行的内存模型，用于存储局部变量，操作数栈，方法出口,基本数据类型的数据，和对象的引用（reference）等。每个方法的调用都对应的出栈和入栈。

2、堆内存空间：jvm只有一个堆区(heap)被所有线程共享，堆中不存放基本类型和对象引用，只存放对象本身　。

3、方法区（又叫静态区）：存放所有的①类（class），②静态变量（static变量），③静态方法，④常量和⑤成员方法(就是普通方法，由访问修饰符，返回值类型，类名和方法体组成)。

4、程序计数器寄存器（PC register）：每个线程启动的时候，都会创建一个PC（Program Counter，程序计数器）寄存器。PC寄存器里保存有当前正在执行的JVM指令的地址。

5、常量池：是方法区的一部分。Class文件中除了有类的版本、字段、方法、接口等描述信息外，还有一项信息是常量池，用于存放编译器生成的各种字面量和符号引用，这部分内容将在类加载后进入方法区的运行时常量池中存放。

6、本地方法栈：本地方法栈与Java虚拟机栈发挥的作用非常相似,它们之间的区别在于虚拟机栈为虚拟机执行java方法(也就是字节码文件)服务,而本地方法栈则为使用到Native方法服务