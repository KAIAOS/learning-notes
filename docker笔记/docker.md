### 一、Docker简介

​	docker是基于go语言实现的云开源项目，docker的使用目标是“build,Ship  and run any app,anywhere”

在任何地方构建、发布并运行任何应用

![image-20200927172822405](C:\Users\hanka\AppData\Roaming\Typora\typora-user-images\image-20200927172822405.png)<img src="C:\Users\hanka\AppData\Roaming\Typora\typora-user-images\image-20200927172852010.png" alt="image-20200927172852010" style="zoom:50%;" />

docker架构图

docker deamon执行客户端的命令 

<img src="C:\Users\hanka\AppData\Roaming\Typora\typora-user-images\image-20200927194330533.png" alt="image-20200927194330533" style="zoom:67%;" />

### 二、Docker常用命令

- docker images [-a]:
  - -a:laughing:列出本地所有的镜像（含中间映象层）
  - -q:只显示镜像id
  - --digests: 显示摘要信息
  - --no-trunc：显示完整镜像名字
- docker search xxx
  - dockerhub查找镜像 -s 30 点赞数
- docker rmi -f hello-world
  - 删除镜像 -f 强制删除 
  - 删除全部 docker rmi -f $(docker images -qa)
- docker run [options]  image [command] [arg]
  - docker exec -t 容器id  命令 不进入容器直接运行命令
  - docker exec -it 容器id /bin/bash    在容器中打开新的终端并启动新进程
  - docker attach 容器id 重新进入容器
  - options 说明 
    - --name=""  为容器指定一个名称
    - -d 后台运行容器 并返回容器id
    - -i 以交互模式运行容器，通常与-t同时使用
    - -t为容器重新分配一个伪输入终端，通常与-i同时使用  docker run -it centos /bin/bash
    - -P 随机端口映射
    - -p 指定端口映射 有四种格式
      - ip:hostPort:containerPort
      - ip::containerPort
      - hostPort:containerPort
      - containerPort
- docker ps -a 
  - 查看在运行的所有容器 -a查看所有历史运行的容器
  - -l上一次运行的 -n 3 上三次运行的
- docker start/restart/stop/kill/rm  容器id或者容器名 启动容器/重启/停止/强制停止/删除已经停止的容器 
- ctrl + P + Q  不停止退出镜像的终端
- docker  logs -f -t --tail 容器id
  - -t加入时间戳 -f跟随最新的日志打印  --tail显示最后多少条


- docker top 容器id 查看容器的运行进程
- docker inspect  容器id 查看容器内部细节 返回json
- docker cp 容器id:/tpm/yum.log /root   从容器拷贝到主机

### 三、镜像相关

​	镜像是一种轻量级、可执行的独立软件包，用来打包软件运行环境和基于运行环境开发的软件，它包含摸个软件所需的所有内容，包括代码、运行时、库、环境变量和配置文件

#### unionFS（联合文件系统）

union文件系统是一种分层、轻量级并且高性能的文件系统，它支持对文件系统的修改作为一次提交来一层层的叠加，同时可以将不同目录挂载到同一个虚拟文件系统。union文件系统时docker镜像的基础。镜像可以通过分层来进行继承，基于基础镜像，可以制作出各种具体的应用镜像

特性 ：：一次同时加载多个文件系统，但从外面开起来，只能看到一个文件系统，联合加载会把各层文件叠加起来，这样最终的文件系统会包含所有底层的文件和目录。

#### bootfs（boot file system）

bootfs主要包含bootloader和kernel，bootloader主要时引导加载kernel。linux刚启动时会加载bootfs文件系统。在docker镜像最底层的就是bootfs。这一层与经典的linux操作系统是一样的，包含boot加载器和内核。当boot加载完成之后整个内核就都在内存之中了 ，此时内存的使用权已由bootfs转交给内核，此时会卸载bootfs

#### rootfs（root file system）

在bootfs之上包含典型linux文件系统的标准目录。rootfs就是各种不同的操作系统发行版，对于镜像来说，共用host的kernel所以只需要rootfs，所以centos几百兆大小不同发行版可以公用bootfs

#### docker commit

docker commit -a="author"  -m="commit messages" 容器id namespace/tomcat2:1.2

之后在 docker images 就有新的镜像

### 四、docker数据卷

- 直接命令添加
  - docker run -it -v /宿主机的绝对路径目录:/容器内目录 镜像名
  - docker run -it -v /宿主机的绝对路径目录:/容器内目录**:ro**  镜像名  容器内只读
- DockerFile添加
  - docker build -f /mydockerfile/dockerfile -t  命名空间/镜像名 .
- 数据卷容器
  - docker run -it --name dc02 --volumes-from dc01 镜像名
  - 回到dc01可以看到02/03各自添加的都能共享了
  - 删除01 dc02仍然能访问
  - 结论：容器之间配置信息的传递，数据卷的生命周期一直持续到没有容器使用它为止

