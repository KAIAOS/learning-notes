### 一、 kubernetes简介

> 需要掌握
>
> 前世今生   kubernetes框架   kubernetes组件

<img src="C:\Users\hanka\AppData\Roaming\Typora\typora-user-images\image-20201222103851288.png" alt="image-20201222103851288" style="zoom:50%;" />

#### 1、 特点

1. 轻量级：消耗资源少
2. 弹性收缩
3. 负载均衡：IPVS

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
kubectl log pod  [pod名称]    #查看pod日志
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
   3. 稳定的网络标识，Pod重新调度后PodName和HostName都不变
   4. 有序部署、有序扩展，pod部署或者扩展时按照预定的顺序
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

Kubenetes`Service`定义了这样一种抽象：一个`Pod`的逻辑分组，一种可以访问它们的策略----通常称之为微服务/这一组Pod能够被`Service`访问到，就是通过`Lable Selector`

<img src="C:\Users\hanka\AppData\Roaming\Typora\typora-user-images\image-20201222212442122.png" alt="image-20201222212442122" style="zoom:50%;" />



### 六、存储

> 需要掌握
>
> 多种存储类型的特点    应用场景



### 七、调度器

> 需要掌握
>
> 调度器原理       根据pod定义到想要的节点运行



### 八、集群安全

>需要掌握
>
>集群的认证      鉴权     访问控制     原理及流程



### 九、HELM

>需要掌握
>
>相当于Linux yum   掌握HELM原理    HELM自定义模板    部署常用的插件

