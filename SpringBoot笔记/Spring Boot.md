# 一、Spring Boot 入门

## 1、Spring Boot 简介

> 简化Spring应用开发的一个框架；
>
> 整个Spring技术栈的一个大整合；
>
> J2EE开发的一站式解决方案；

## 2、微服务

> In short, the microservice architectural style is an approach to developing a single application as a **suite of small services**, each **running in its own process** and communicating with lightweight mechanisms, often an HTTP resource API. These services are **built around business capabilities** and **independently deployable** by fully automated deployment machinery. There is a **bare minimum of centralized management** of these services, which may be written in different programming languages and use different data storage technologies.
>
> -- [James Lewis and Martin Fowler (2014)](https://martinfowler.com/articles/microservices.html)

2014 Martin Fowler

微服务：架构风格

一个应用应该是一组小型服务；可以通过http的方式进行互通

每一个功能元素最终都是一个可独立替换和独立升级的软件单元

## 3、helloworld探究

### 1、POM文件

#### 1、父项目

```xml
<parent>
        <groupId>org.springframework.boot</groupId>
        <artifactId>spring-boot-starter-parent</artifactId>
        <version>2.2.2.RELEASE</version>
        <relativePath/> <!-- lookup parent from repository -->
</parent>

它的父项目是
<parent>
    <groupId>org.springframework.boot</groupId>
    <artifactId>spring-boot-dependencies</artifactId>
    <version>2.2.2.RELEASE</version>
    <relativePath>../../spring-boot-dependencies</relativePath>
</parent>
它的properties标签管理Spring Boot应用里面的所有依赖版本
```

Spring Boot的版本仲裁中心；

以后导入依赖默认是不需要写版本；（没有在dependencies里面管理的依赖则需要声明版本号）

#### 2、启动器

```xml
<dependency>
            <groupId>org.springframework.boot</groupId>
            <artifactId>spring-boot-starter-web</artifactId>
</dependency>
```

**spring-boot-starter**-web：

​	spring-boot-starter：spring-boot场景启动器；帮我们导入了web模块正常运行所依赖的组件

Spring Boot把每个功能场景打包成一个个的starter，只需要在项目中引入这些starter，相关场景的所有依赖就都会导入进来。

### 2、主程序类，入口类

```java
@SpringBootApplication
public class HelloWorldMainApplication {
    public static void main(String[] args){
        SpringApplication.run(HelloWorldMainApplication.class, args);
    }
}
```

**@SpringBootApplication**: Spring Boot 应用标注在某个类上说明这个类是SpringBoot的主配置类，SpringBoot就应该运行这个类的main方法来启动SpringBoot应用

```java
@Target({ElementType.TYPE})
@Retention(RetentionPolicy.RUNTIME)
@Documented
@Inherited
@SpringBootConfiguration
@EnableAutoConfiguration
@ComponentScan(
    excludeFilters = {@Filter(
    type = FilterType.CUSTOM,
    classes = {TypeExcludeFilter.class}
), @Filter(
    type = FilterType.CUSTOM,
    classes = {AutoConfigurationExcludeFilter.class}
)}
)
public @interface SpringBootApplication {
```

**@SpringBootConfiguration**：标注某个类表示SpringBoot的配置类

​		**@Configuration**  Spring中的注解 表示配置类

​			配置类就是配置文件；配置类也是容器中的一个组件；@Component

**@EnableAutoConfiguration**：开启自动配置

​		以前需要配置的东西，Spring Boot自动配置

```java
@AutoConfigurationPackage
@Import({AutoConfigurationImportSelector.class})
public @interface EnableAutoConfiguration {
```

​	@AutoConfigurationPackage自动配置包

>    @Import({Registrar.class}) Spring的底层注解，给容器中导入一个组件；导入的组件由Registrar.class将**主配置类（@SpringBootApplication标注的类）**的所在包及下面所有子包里面的所有组件扫描到Spring容器

​	@Import({AutoConfigurationImportSelector.class}) @Import作用是给容器中导入组件

​		AutoConfigurationImportSelector导入哪些组件的选择器

​		将需要导入的组件以全类名的方式返回；这些组件就被添加到容器中

​		会给容器中导入非常多的自动配置类；就是给容器导入这个场景需要的所有组件，并配置好这些组件

​		有了自动配置类，就免去了我们手动编写配置注入功能组件的工作

## 4、使用Spring initializer快速创建Spring Boot项目

IDE支持快速创建一个SpringBoot项目，选择需要的模块；向导会联网创建Spring Boot项目

默认生成的SpringBoot项目

- 主程序已经生成好了，只需要编写自己的逻辑
- resource文件夹的目录结构
  - static：保存所有静态资源；js、css、images
  - templates：保存所有的模板页面；（SpringBoot默认jar包使用嵌入式的Tomcat，默认不支持jsp页面）但是可以使用模板引擎（freemarker，thymeleaf）
  - application.properties : Spring Boot应用的配置文件例如：server.port =8081可以改变端口

# 二、配置文件

SpringBoot使用一个全局的配置文件，配置文件名是固定的；

- application.properties 例如 server.servlet.context-path=/boot
- application.yml

配置文件作用：修改SpringBoot自动配置的默认值；SpringBoot在底层自动配置

yaml：以数据为中心的配置文件

# 三、数据访问

## 1、整合jpa

jpa：orm（object relational map）；

1）编写一个实体类（bean）和数据表映射，并配置好映射关系

```java
@Entity//告诉JPA这是一个实体类
@Table(name = "tbl_user")//指定和哪个数据表对应；如果省略默认表名就是类名小写
public class User {

    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    private Integer id;
    @Column(name = "last_name", length = 50)
    private String lastName;
    @Column//省略默认列名就是属性名
    private String email;
```



2）编写一个Dao接口来操作实体类对应的数据表（repository）

```java
//继承Jparepository完成对数据库的操作
public interface UserRepository extends JpaRepository<User, Integer> {
}
```



3）基本的配置

```yaml
spring:
jpa:
    hibernate:
      ddl-auto: update
    show-sql: true
```

- springboot 与 docker

- springboot 启动配置原理   七

- springboot 自定义starters

- springboot 与安全

- springboot 与分布式

  ![image-20201218150914107](C:\Users\hanka\AppData\Roaming\Typora\typora-user-images\image-20201218150914107.png)

<img src="C:\Users\hanka\AppData\Roaming\Typora\typora-user-images\image-20201220154401991.png" alt="image-20201220154401991" style="zoom:50%;" />分割人人人人人都是是否缺少维生素的完全

<img src="C:\Users\hanka\AppData\Roaming\Typora\typora-user-images\image-20201220163716215.png" alt="image-20201220163716215" style="zoom:67%;" />

![image-20201220210256800](C:\Users\hanka\AppData\Roaming\Typora\typora-user-images\image-20201220210256800.png)

