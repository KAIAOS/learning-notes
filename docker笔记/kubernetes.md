### 一、 kubernetes简介

> 需要掌握
>
> 前世今生   kubernetes框架   kubernetes组件
>
> kubeadm init --apiserver-advertise-address=192.168.117.4 --image-repository registry.aliyuncs.com/google_containers --service-cidr=10.96.0.0/12  --pod-network-cidr=10.244.0.0/16

<img src="C:\Users\hanka\AppData\Roaming\Typora\typora-user-images\image-20201222103851288.png" alt="image-20201222103851288" style="zoom:50%;" />

#### 1、 特点

1. 容器化集群管理系统
2. 轻量级：消耗资源少
3. 弹性收缩
4. 负载均衡：IPVS

#### 2、组件

1. APISERVER：所有服务访问的统一入口
2. ControllerManager：维持副本的期望数目
3. Scheduler：负责介绍任务，选择合适的节点进行分配
4. Kubelet：直接跟容器引擎交互，实现容器的生命周期
5. Kube-proxy：负责写入规则到IPTABLES、IPVS 实现服务的映射访问

#### 3、其他插件

1. COREDNS：可以为集群中的SVC创建一个域名IP的对应解析
2. DASHBOARD：给k8s集群提供一个B/S结构的访问体系
3. INGRESS CONTROLLER：本来是四成代理，INGRESS可以实现七层代理
4. fedetation：可以提供一个跨集群中心多k8s同意管理的功能
5. prometheus：提供k8s集群的监控能力
6. ELK：提供k8s集群日志统一分析平台



### 二、Pod

> 需要掌握
>
> 什么是Pod  k8s网络通讯模式

#### 1、pod分类

1. 自主式pod和控制器控制pod
2.  每个pod都有一个pause容器，一个pod内不同的容器会共享pause的网络栈和存储卷
3. 同一个pod中的容器端口不能冲突

#### 2、 网络通讯模式

> kubenetes的网络模型假定所有Pod都在一个可以直接连通的扁平网络空间，这里使用Flannel实现
>
> **Flannel**是CoreOS团队针对kubenetes设计的网络规划服务，它使得集群中不同节点主机创建的Docker容器都具有全集群唯一的虚拟IP地址。而且它还能在这些IP地址之间建立一个overlay network，通过这个覆盖网络，将数据包传递到目标容器内
>
> <img src="C:\Users\hanka\AppData\Roaming\Typora\typora-user-images\image-20201222135504010.png" alt="image-20201222135504010" style="zoom:50%;" />

1. 同一个Pod内多个容器：由于公用了pause容器的网络栈，所以可以直接lo回环网卡
2. 各个Pod之间的通讯：overlay network
   1. 在同一台主机。Pod的地址是与docker0在同一个网段的，由docke0网桥转发
   2. 不在同一台主机。因为docker0网段与宿主机网卡不在同一网段，所以将pod 的Ip与Node的Ip关联起来，通过这关联让Pod可以互相访问
3. Pod与Service之间的通讯：由iptables进行维护和转发
4. Pod到外网：iptables执行SNAT转换，把源Ip改为宿主机网卡ip
5. 外网访问Pod：通过service，NodePod暴露服务（ClusterIp是内网访问）
6. 其实一共有三层网络：节点网络、Pod网络、Service网络
   1. ![image-20201222140703990](C:\Users\hanka\AppData\Roaming\Typora\typora-user-images\image-20201222140703990.png)
7. ETCD之Flannel提供说明：
   1. 存储管理Flannel可分配的Ip地址段资源
   2. 监控每一个Pod的实际地址，并在内存中建立维护Pod节点路由表



### 三、资源清单

> 需要掌握
>
> k8s中所有的内容都抽象为资源，资源实例化以后称为对象
>
> 资源        资源清单的语法           编写pod       **掌握pod的生命周期**

####  1、集群资源分类

1. 名称空间级别   kubeadm会把组件放在kube-system下   kubecrl get pod -n kube-system
   1. 工作负载型资源（workload）：Pod、ReplicaSet、Deployment、StatefulSet、DaemonSet、Job、CronJob
   2. 服务发现及负载均衡资源（erviceDiscovery LoadBalance）:Service、Ingress
   3. 配置与存储型资源：Volume、CSI（容器存储接口、可扩展各种各样的第三方存储卷）
   4. 特殊类型的存储卷：ConfigMap（当配置中心使用的资源类型）、Secret、DownwardApi
2. 集群级别    role  一旦定义了全集群都可见
   1. NameSpace、Node、Role、ClusterRole、RoleBinding、ClusterRoleBinding
3. 元数据类型  例如HPA、PodTemplate、LimitRange

#### 2、资源清单

​	1、必须存在的字段属性

<img src="C:\Users\hanka\AppData\Roaming\Typora\typora-user-images\image-20201222143015131.png" alt="image-20201222143015131" style="zoom:50%;" />

2. 比较重要的字段属性

   <img src="C:\Users\hanka\AppData\Roaming\Typora\typora-user-images\image-20201222143544849.png" alt="image-20201222143544849" style="zoom:50%;" />

*******

![image-20201222143834801](C:\Users\hanka\AppData\Roaming\Typora\typora-user-images\image-20201222143834801.png)

***

![image-20201222143914217](C:\Users\hanka\AppData\Roaming\Typora\typora-user-images\image-20201222143914217.png)

<img src="C:\Users\hanka\AppData\Roaming\Typora\typora-user-images\image-20201222144018972.png" alt="image-20201222144018972" style="zoom:50%;" />

#### 3、编写资源清单

```yaml
# pod.yml
apiVersion: v1
kind: Pod
metedata:
  name: myapp-pod
  labels:
    app: myapp
    version: v1
spec:
  containers:
  -name: app
   image: myapp:v1 
```

```bash
kubectl apply -f pod.yml    #根据yml文件部署pod
kubectl replace -f  pod.yml #根据yml文件重启pod

kubectl get pod         #查询当前pod
kubectl describe pod  [pod名称]    #查看pod详细状态
kubectl logs pod  [pod名称]    #查看pod日志
kubectl log [pod名称] -c [容器名称]    #查看pod内容器日志
kubectl edit pod [pod名称]     #编辑这个pod的资源清单

kubectl delete deployment --all #删除所有控制器
kubectl delete pod --all   #删除所有Pod
kubectl delete svc [service name]   #删除所有service

kubectl exec [pod名称] (-c [容器名称])  -it -- /bin/sh
```

#### 4、Pod生命周期

![image-20201222162921609](C:\Users\hanka\AppData\Roaming\Typora\typora-user-images\image-20201222162921609.png)

1. 过程介绍
   1. 先创建pause容器，在初始化容器 init C
   2. 进入MainC，执行START 命令；MainC结束后，执行STOP命令
   3. readiness 就绪检查程序，可以设置延时启动
   4. Liveness 会一直检查容器是否正常，否则可能要重启（根据指定的规则）
2. init C可以是一个或多个先于容器启动的init容器，init C总是运行到成功完成为止，每个Init容器必须在下一个Init容器启动之前完成。如果init c失败，kubernetes会不断的重启该Pod，直到Init C成功为止，如果restartPolicy为Never就不会重启

#### 5、探针

1. 探针是由kubelet对容器执行的定期诊断。kubelet调用由容器实现的Handler，共三种：
   1. ExecAction: 在容器内执行指定命令。如果命令退出返回码为0，则诊断成功
   2. TCPSocketAction：对指定端口的容器的IP进行TCP检查。如果端口打开则诊断成功
   3. HTTPGetAction：对指定端口的容器Ip地址执行HTTP Get请求。如果响应状态码大于等于200且小于400，则诊断成功
2. 每次诊断只有三种结果之一：
   1. 成功：通过诊断
   2. 失败：未通过诊断
   3. 未知：诊断失败，不会采取任何行动
3. 探测方式
   1. livenessProbe：指示容器是否正在运行。如果存活探测失败，则kubelet会杀死容器，并容器受到重启策略的影响。如果不提供存活探针，则默认Success
   2. readinessProbe：指示容器是否准备好服务请求。如果就绪探测失败，端点控制器将从与Pod匹配的所有Service的端点中删除该Pod的IP地址。初始延迟之前的就绪状态默认为Failure。如果不提供就绪探针，则默认Success



### 四、Pod控制器

> 需要掌握
>
> 各种控制器特点和使用方式

#### 1、控制器类型

1. ReplicationController和ReplicaSet 
   1. 用来确保容器应用副本数满足期望值，容器异常退出九自动创建新的；多出来的会自动回收
   2. ReplicaSet 跟ReplicationController没有本质不同，只是它支持集合式的selector
   3. R eplicaSet支持：创建pod的时候会打标签，例如app=apache、version=v1，当操作pod时可以按照标签操作，例如副本数目监控就是基于标签
2. Deployment
   1. 虽然ReplicaSet可以独立使用，但是一般还是建议使用Deployment来自动管理ReplicaSet
   2. 支持扩容和缩容
   3. ReplicaSet不支持rolling-update，但是deployment支持
3. Horizontal Pod Autoscaling 
   1. 仅适用deployment 和ReplicaSet，可根据CPU利用率扩容
   2. 会监控pod的cpu利用率达到设定值时，会自动新建副本
4. StatefulSet
   1. 解决有状态服务对应Deployment 和 ReplicaSet是无状态服务
   2. 稳定的持久化存储，即Pod重新调度可以访问到相同的持久化数据
   3. 稳定的网络标识，Pod重新调度后PodName和HostName都不变 $(podname).$(headless server name)
   4. 有序部署、有序扩展，pod部署或者扩展时按照预定的顺序
      1. 如果有很多Pod副本，它们会被顺序的创建（0---N-1）并且下一个Pod运行之前所有Pod必须是Running和Ready状态
5. DaemonSet
   1. DaemonSet确全部（或者一些）Node上仅运行一个Pod副本，当有新的Node加入集群时，会增加一个Pod，当Node删除时候这些Pod被回收
   2. 一些典型用法
      1. 运行集群存储daemon，在每一个node上运行日志收集daemon，在每一个node上运行监控daemon
6. Job，CronJob
   1. Job负责批处理任务，即仅执行一次的任务，它可以保证批处理任务的一个或多个Pod成功结束
   2. CronJob 管理基于时间的Job，可以给定时间运行一次也可以周期性的运行

### 五、服务发现

>需要掌握
>
>SVC原理       构建方式

#### 1、介绍

Kubenetes`Service`定义了这样一种抽象：一个`Pod`的逻辑分组，一种可以访问它们的策略----通常称之为微服务/这一组Pod能够被`Service`访问到，就是通过`Lable Selector`

<img src="C:\Users\hanka\AppData\Roaming\Typora\typora-user-images\image-20201222212442122.png" alt="image-20201222212442122" style="zoom:50%;" />

#### 2、Service类型

1. ClusterIp：默认类型，自动分配一个仅Cluster内部可以访问的虚拟Ip
   1. 还有一种HeadService：有时候不需要负载均衡和单独的ServiceIp，通过指定ClusterIP的 值为‘None’来创建Headless Service
2. NodePort：在ClusterIp基础上为每台机器上绑定一个端口，这样就可以通过Node IP：NodePort 来访问服务
3. LoadBalancer：在NodePort的基础上，借助Cloud Provider创建一个外部的负载均衡器，并将请求转发到  Node IP：NodePort，，云供应商需要额外收费的
4. ExternalName：把集群外部的服务引入到集群内部来，在集群内部直接使用。没有任何类型的代理被创建（是集群内部访问外部服务会用到） 

<img src="C:\Users\hanka\AppData\Roaming\Typora\typora-user-images\image-20201223143158124.png" alt="image-20201223143158124" style="zoom:50%;" />

​		kube-proxy负责将相应标签的Pod访问写入iptables，当客户端访问时会经过iptables访问到相应的Pod，apiserver负责监控，现版本使用IPVS

#### 3、Port、TargetPort、NodePort

1. 这里的port表示：service暴露在cluster ip上的端口，**<cluster ip>:port** 是提供给集群内部客户访问service的入口。
2. targetPort很好理解，targetPort是pod上的端口，从port和nodePort上到来的数据最终经过kube-proxy流入到后端pod的targetPort上进入容器。
3. nodePort是kubernetes提供给集群外部客户访问service入口的一种方式（另一种方式是LoadBalancer），所以，<nodeIP>:nodePort 是提供给集群外部客户访问service的入口。

### 六、存储

> 需要掌握
>
> 多种存储类型的特点    应用场景

#### 1、ConfigMap

#### 2、Secret

#### 3、Volume

> 当容器崩溃时，kubelet会重启，容器会以干净的状态重新启动。其次在Pod同时运行多个容器时，这些容器之间通常需要共享文件。kubenetes中Volume解决了这个问题
>
> Volume有明确的寿命-- 与封装它的Pod相同，当Pod删除时，它也不存在

1. emptyDir：当Pod被分配给节点时，首先创建emptyDir卷，并且只要该Pod在该节点上运行，该卷就会存在，最初时空的。Pod中的容器可以读取和写入emptyDir卷中的相同文件，尽管改卷可以挂载到每个容器的相同或不同的路径上。当出于任何原因从节点删除Pod时，empthDir中的数据被永久删除
2. hostPath：将主机节点的文件系统中的文件或者目录挂载到集群中，用途如下
   1. 运行需要访问Docker内部容器的文件或容器使用主机的文

<img src="C:\Users\hanka\AppData\Roaming\Typora\typora-user-images\image-20201223161118089.png" alt="image-20201223161118089" style="zoom:50%;" />

**使用这种卷时需要注意：**

- 由于每个节点的文件不同，具有相同配置的Pod在不同节点上的行为会有所不同
- 当kubenetes按照计划添加资源感知的调度时，将无法考虑hostPath的资源
- 在底层主机上创建的文件或目录只能由root写入，需要在特权容器中以root运行，或者修改主机的文件权限以便写入hostPath

```yml
# 一个使用hostPath的Pod资源清单
kind: Pod
metedata:
  name: test-pd
spec:
  containers:
    - image: k8s.gcr.io/test-webserver
      name: test-container
      volumeMounts:
      - mountPath: /test-pd
        name: test-volume
  volumes:
  - name: test-volume
    hostPath:
     # directory name on host
     path: /data
     # this field is optional,Directory means the directory path must exist
     type: Directory
```

#### 4、PV and PVC

1. PersistentVolume（PV）是由管理员设置的存储，它是集群的一部分。就像是集群中的资源一样，PV也是集群中的资源，此API对象包含存储实现的细节，即NFS、ISCSI或特定于云供应商的存储系统
2. PersistentVolumeClaim（PVC）是用户存储的请求。它与Pod类似，Pod消耗节点资源，PVC消耗PV资源，声明可以请求特定的大小和访问模式



### 七、调度器

> 需要掌握
>
> 调度器原理       根据pod定义到想要的节点运行

#### 1、调度说明

1. Scheduler是Kubenetes的调度器，任务是把定义的Pod分配到集群的节点上，保证公平、资源高效利用、效率和灵活，Scheduler是作为单独的程序运行的，启动之后一直监听API server，获取`PodSpec。NodeName`为空的pod，为其创建一个binding，表明该pod放在哪个节点上
2. 调度过程：首先过滤掉不满足条件的节点，这个过程为predicate，然后通过节点按照优先级排序priority

#### 2、调度亲和性

1. `pod.spec.nodeAffinity` 节点亲和性
   1. preferedDuringSchedulingIgnoredDuringExecution 软策略 （满足就运行，不满足就算了）
   2. requiredDuringSchedulingIgnoredDuringExecution  硬策略（不满足就不运行，pod会pending）
2. `pod.spec.podAffinity/podAntiAffinity` Pod亲和性
   1. preferedDuringSchedulingIgnoredDuringExecution 软策略
   2. requiredDuringSchedulingIgnoredDuringExecution  硬策略

#### 3、污点 taint和toleration

1. 节点亲和性是pod的一种属性，它使得pod被调度到一类特定的节点，而taint则相反，它使得节点能够排斥一类特定的Pod。taint和toleration相互配合，避免pod被分配到不合适的节点，每个节点都可以应用多个taint，不能容忍这些taint的pod是不会运行的，如果可以容忍则有可能会运行在该节点
2. `kubectl taint` 可以给某个节点设置污点，每个污点由`key=value:effect`组成，value可以为空，effect描述污点的作用   effetc:
   1.  NoSchedule: 表示k8s将不会把Pod调度到具有该污点的Node上
   2. PreferNoSchedule: 表示k8s尽量避免将pod调度到该污点的Node上
   3. NoExecute: 表示k8s将不会把Pod调度到该node上，同时会将Node上已经存在的Pod驱逐出去
3. `pod.spec.toleration`设置容忍。key、value、effect、要与node上的taint一致，当不指定key值时，表示容忍所有的污点key，当不指定effect时，表示容忍所有的污点作用

#### 4、固定节点

1. Pod.spec.nodeName 将Pod直接调度到指定的Node节点上，会跳过Scheduler的调度策略，该匹配规则是强制匹配
2. Pod.spec.nodeSelector 通过kubenetes的label-selector机制选择节点，由调度器调度策略匹配label<kbd>kubectl label node k8s-node01 disk=ssd</kbd>可以打标签<kbd>Ctrl</kbd>+<

### 八、集群安全

>需要掌握
>
>集群的认证      鉴权     访问控制     原理及流程

**机制说明：Kubenetes作为一个分布式集群的管理工具，保证集群的安全性是重要任务。API Server是集群内部各个组件通信的中介，也是外部控制的入口。所以kubenetes的安全机制就是围绕保护API Server来设计的。Kubernetes使用了认证（Authentication）、鉴权（Authorization）、准入控制（Admission Controll）**

#### 1、 Authtication

- HTTP Token认证：通过一个Token来识别合法用户
- HTTP Base认证： 通过用户名+密码的方式认证
  - 用户名+密码用BASE64算法进行编码后的字符串放在HTTP Request的Header中
- 最严格的HTTPS证书认证，基于CA根证书签名的客户端身份认证方式

**安全性说明：Cotroller Manager、Scheduler与API Server在同一台机器，所以直接用API Server的非安全端口访问；kubectl、kubelet、kube-proxy访问API Server则需要HTTPS证书双向认证**

#### 2、Authorization

- AlwaysDeny:  拒绝所有请求
- AlwaysAllow：允许所有请求
- ABAC(Attribute-Based Access Controll)：基于属性的访问控制，表示使用用户的授权规则对用户请求进行匹配和控制 
- Webbook：调用外部服务对用户授权
- **RBAC(Role-Based  Access Controll)**: 基于角色的访问控制，先行默认规则

### 九、HELM

>需要掌握
>
>相当于Linux yum   掌握HELM原理    HELM自定义模板    部署常用的插件

**在没使用helm之前，向kubernetes部署应用，我们依次部署deployment、svc等，不走较繁琐，helm通过打包的方式，支持发布的版本管理和控制，很大程度上简化了kubernetes应用的部署和管理**

**Helm的本质就是让k8s的应用管理可配置，能动态生成，通过动态生成资源清单文件（deployment.yml,service.yml）调用kubectl自动执行k8s资源部署**

**Helm是官方提供的类似YUM的包管理器，是部署环境的流程封装。Helm有两个重要的概念：chart和release**

- chart是创建一个应用的信息集合，包括各种kubernetes对象的配置模板、参数定义、依赖关系、文档说明。chart是应用部署的逻辑单元。类似与yum中的软件安装包
- release是chart的运行实例，代表了一个正在运行的应用。当chart被安装到kubernetes集群，就生成了一个release。chart能够多次安装到同一个集群，每次安装都是一个release

 ```shell
helm repo add [名称]
helm install [安装之后名称] [搜索名称]
helm list
helm status  [安装之后名称]

#修改service的yaml文件，type改为nodePort
kubectl edit svc ui-weave-scope
 ```

#### 创建chart

`helm create mychart`  创建chart

### 十、部署项目流程

#### 1、容器交付流程

1. 开发代码，测试，编写Dockerfile
2. 持续集成Continuous Integration(*CI*)和持续交付Continuous Delivery(*CD*)
   1. 制作镜像 上传镜像仓库
3. 应用部署 （环境准备  Pod Service Ingress）
4. 运维  （监控  故障排除   升级版本）

#### 2、 k8s部署项目流程

1. 制作镜像
2. 上传镜像仓库
   1. docker tag [imageId] 
3. 通过控制器部署镜像
   1. kubectl create deployment java-deloyment --image=registry.us-west-1.aliyuncs.com/kaidev/java:1.0.0 --dry-run -o yaml>java-deployment.yaml
   2. kubectl scale deployment  [dp名称] --replicas=3
4. Service或ingress对外暴露应用
   1. kubectl expose deployment java-deloyment  --port=8080 --target-port=8111 --type=NodePort --dry-run -o yaml>java-deployment-service.yaml
5. 监控升级