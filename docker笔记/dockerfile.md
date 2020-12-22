### 一、简介

​	dockerfile是用来构建Docker镜像的构建文件，是由一系列命令和参数构成的脚本

- 手动编写dockerfile文件，符合其规范

- docker build -f Dockerfile -t [自定义镜像名称] .

- docker build 命令执行 获得自定义的镜像

- docker run -d -p 9080:8080 --name mycontainer 

  -v /home/tomcat9/test:/usr/local/apache-tomcat/webapps/test

  -v /home/tomcat9/logs:/usr/local/apache-tomcat/logs

  mytomcat_imag

``` dockerfile
FROM scratch //scratch所有镜像的元镜像
ADD centos-8-x86_64.tar.xz /
LABEL org.label-schema.schema-version="1.0"     org.label-schema.name="CentOS Base Image"     org.label-schema.vendor="CentOS"     org.label-schema.license="GPLv2"     org.label-schema.build-date="20200809"
CMD ["/bin/bash"]
```

```dockerfile
FROM centos
COPY a.txt /usr/local/a.txt
ADD jdk-8u171.tar.gz /usr/local  #复制且自动解压缩
ADD apache-tomcat.tar.gz /usr/local 
RUN yum -y install vim #容器构建时需要运行的命令
ENV MYPATH /usr/local
WORKDIR $MYPATH

ENV JAVA_HOME /usr/local
ENV CLASSPATH $JAVA_HOME/lib/dt.jar:$JAVA_HOME/lib/tools.jar
ENV CATALINA_HOME /usr/local/apache-tomcat
ENV CATALINA_BASE /usr/local/apache-tomcat
ENV PATH $PATH:$DAVA_HOME/bin:$CATALINA_HOME/lib:$CATALINA_HOME/bin

EXPOSE 8080 #暴露端口

#启动时运行tomcat
#ENTRYPOINT ["/usr/local/apache-tomcat/bin/start.sh"] 
CMD /usr/local/apache-tomcat/bin/start.sh && tail -F /usr/local/apache-tomcat/bin/logs/catalina.out
```



### 二、dockerfile构建过程解析

#### 1.dockefile内容基础知识

- 每条保留字指令必须为大写字母且后面至少跟一个参数
- 指令顺序执行，#表示注释
- 每条指令都会创建一个新的镜像层，并对镜像进行提交

#### 2. 执行流程

- docker从基础镜像运行一个容器
- 执行一条指令并对容器做出修改
- 执行类似docker commit的操作提交一个新的镜像层
- docker再基于刚提交的镜像运行一个新的容器
- 执行dockerfile中的下一条指令直到所有指令都执行完成

#### 3. docker保留字指令

 - FROM 基础镜像，当前镜像是基于哪一个镜像的
 - MAINTAINER 镜像维护者的姓名和邮箱地址
 - RUN 容器构建时需要运行的命令
 - EXPOSE 暴露端口
 - WORKDIR 登陆目录
 - ENV 构建镜像过程中设置环境变量   $引用变量
 - ADD 拷贝到镜像且自动处理URL和解压tar压缩包
 - COPY  拷贝到镜像。 将从构建上下文目录中<源路径>的文件/目录 复制到新的一层的镜像内的<目标路径>
 - VOLUMES 数据卷
 - CMD 指定一个容器运行时要运行的命令，dockerfile可以有多个CMD指令，但是最后一个生效，CMD会被docker run 之后的参数替换掉
 - ENTRYPOINT 指定一个容器运行时要运行的命令，不会被命令行覆盖而是追加
     - shell命令格式 CMD <命令>
     - exec格式 CMD [“可执行文件”,参数1,参数2]
- ONBUILD 当一个被继承的dockerfile时运行命令，