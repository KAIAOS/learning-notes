* # 爬虫系统交接指南

  ## 日常维护

  ### 代码更新

  目前是从爬虫分支上先提交更新,然后再listfetcher更新对应分支爬虫,最后merge到master分支上.我一般的操作是下面这样,可以根据自己习惯.多用git status查看情况.

  ``` shell
  初始在listfetcher或master分支,如果现有代码有改动,先git stash保存到暂存区,最后更新完再git stash pop恢复
  git checkout listfetcher_hankai
  git pull
  git checkout listfetcher
  # 更新hankai分支负责的所有爬虫,checkout更新目标分支文件到当前分支
  git checkout listfetcher_by_hankai crawler/fetcher/chinesemooc/list_fetcher.py crawler/fetcher/haodaxue/list_fetcher.py crawler/fetcher/pmphmooc/list_fetcher.py crawler/fetcher/weizhiku/list_fetcher.py crawler/fetcher/zhihuishu/list_fetcher.py crawler/fetcher/wanke/list_fetcher.py crawler/fetcher/umooc/list_fetcher.py`1
  ( git checkout listfetcher_by_fuchen crawler/fetcher/cqooc/list_fetcher.py crawler/fetcher/icourse163/list_fetcher.py crawler/fetcher/xueyinonline/list_fetcher.py crawler/fetcher/ulearning/list_fetcher.py crawler/fetcher/gaoxiaobang/list_fetcher.py crawler/fetcher/erya/list_fetcher.py crawler/fetcher/livedu/list_fetcher.py crawler/fetcher/ehuixue/list_fetcher.py
  git checkout listfetcher_by_wzx crawler/fetcher/huixuexi/list_fetcher.py crawler/fetcher/gaoxiaowaiyumuke/list_fetcher.py crawler/fetcher/xuetangx/list_fetcher.py crawler/fetcher/youkelianmeng/list_fetcher.py crawler/fetcher/zhihuizhijiao/list_fetcher.py crawler/fetcher/zhihuizhijiao/proxy_ip.py crawler/fetcher/zhongkeyun/list_fetcher.py crawler/fetcher/zhejiangmooc/list_fetcher.py
  )
  # 更新完毕
  git commit -m "update listfetcher --xxxxxxx"
  git push origin listfetcher
  git checkout master
  git merge  --squash listfetcher
  git commit -m "update listfetcher ---xxxx"
  git push origin master
  (git stash pop)
  ```

  ### 服务器程序启动

  目前在gpu-node1 192.168.232.1节点上的crawl下运行,分为三个程序启动,目前同时起了6个crawler,1个manager,1个pipeline.

  账户:crawl 密码： crawl
  
  使用tmux进行窗口管理,使用tmux进行窗口管理,登入账号之后,`tmux a -t 0`进入窗口 操作方式prefix:Ctrl+B
  常用操作:
  	`Ctrl+B "` 竖直分屏 | | |
	`Ctrl+B %` 水平分屏 =
  	`Ctrl+B x` 关闭窗口

  启动程序
  
  ```shell
  python3 start.py -r Manager -c config/manager.json
  python3 start.py -r Crawler -c config/crawler.json
  python3 start.py -r Pipeline -c config/pipeline.json
  ```

## 项目结构

```
[~]$: mooc_crawler
.
├── config         						# 配置文件,包括各角色配置,对应start.py里启动的参数,以及各平台的相关配置.
│   ├── crawler.json  				
│   ├── manager.json
│   ├── pipeline.json
│   └── platform
├── crawler										
│   ├── crawler.py						# crawler主程序,加载任务对应的fetcher
│   ├── fetcher
│   ├── __init__.py
│   ├── loader.py							# 加载到对应爬虫类的加载器
│   └── README.md
├── exceptions
│   ├── __init__.py
│   └── offlineError.py				
├── manager
│   ├── handler.py					 # 用来处理client发来的rpc请求
│   ├── launchers						 # 各平台的登录器
│   ├── login_info_mgr.py		 # 获取登录信息
│   ├── rpc_client.py				 # 约定rpc客户端的请求方法及参数
│   ├── rpc_server.py				 # Manager主进程
│   ├── scheduler.py				 # 实现调度策略,接收到每个任务要怎么处理
│   └── timer.py
├── message_queue						 
│   ├── base_queue.py				
│   ├── mq_consumer.py
│   └── mq_producer.py
├── persistence							
│   ├── config							# 配置文件对应类
│   ├── db									# 一个数据表对应一个类,封装了操作数据库的接口
│   └── model								# 处理对象类
├── pipeline
│   ├── handler.py						
│   ├── pipelines						 # 数据的处理流程,主要写在data_pipeline里,继承了base_pipeline
│   ├── rpc_client.py
│   └── rpc_server.py
├── Pipfile
├── README.md
├── requirments.txt						
├── rpc											 # 封装的rpc模块,实现不同角色间的通信
│   └── http
├── settings
│   ├── __init__.py
│   └── settings.py
├── start.py									# 启动程序
├── tests
└── utils
    ├── __init__.py
    └── utils.py
```

目前维护主要改动的地方就是fetcher下面会更新,更新之后只需要把crawler重启就行了,manager和pipeline没有改动就不用重启;persistence里主要先看一下schedule_task和task_info这两个表,熟悉一下这两个表结构,结合manager/scheduler理解一下调度逻辑;

schedule_task: status: 1在队列中，2在爬取中，3爬取完成，4，失败

## 数据库数据导出

可使用该命令查看当天任务结果,写个脚本自动导出报表以及发送邮件,通过crontab工具每天定时执行.每天还是把报表更新一下在线版本,然后本地留存一份,保证每天各平台负责人要核查数据.

``` sql
select course_name,crawl_num,course_num,crawled_next_time,start_time,crawl_finish_time,start_handle_time,finish_time ,TIMESTAMPDIFF(SECOND,start_time,crawl_finish_time)/course_num as crawl_per_course_time,TIMESTAMPDIFF(SECOND,create_time,start_time) as wait_start,TIMESTAMPDIFF(SECOND,start_time,crawl_finish_time) as crawl_time,TIMESTAMPDIFF(SECOND,crawl_finish_time,start_handle_time) as wait_time,TIMESTAMPDIFF(SECOND,start_handle_time,finish_time) as handle_time,crawled_failed_num,status
from (select a.course_name,b.status,b.create_time,b.start_time,b.crawl_finish_time,b.finish_time,b.crawled_failed_num,b.start_handle_time,b.crawled_next_time,b.crawl_num
from (SELECT * from schedule_task where crawled_next_time >= curdate()) b 
JOIN (select course_name,id from task_info) a on a.id = b.task_id) c 
left join (SELECT DISTINCT(platform),count(*)as course_num from course_crowd_weekly where update_time >= curdate() GROUP BY platform) d on c.course_name = d.platform 
ORDER BY course_num desc
```

task_info

schedule_task

![image-20200315195510718](C:\Users\hanka\AppData\Roaming\Typora\typora-user-images\image-20200315195510718.png)

## 其他组件

1. 消息队列:[网页可视化查看](http://192.168.232.154:15672),队列名是[task_queue](http://192.168.232.254:15672/#/queues/%2F/task_queue)

2. 账号`admin`密码`admin`

3. Unacked 是正在进行爬取的任务,爬取完成即从队列移出,Ready是在等待的任务,本身有2个Unacked的消息,可以忽略.

4. graylog日志管理系统:[网页入口](http://192.168.232.254:9001) 需要通过代理访问

   看爬虫的话就是 faciility:mooc_crawler_test 另外两个就是mooc_pipelinetest/mooc_manager_test
   级别看错误日志的就level:3

## Linux相关

* 服务器没有开放管理员权限,可以自己租台服务器熟悉一下,实验室都是centos7系统;docker部署应用,nginx代理配置,防火墙配置,这些在平时部署服务的时候会经常用到.
* 可以在自己服务器上把环境配一下,让系统可以在你的服务器上跑起来,这样你基本就熟悉了,可以多跑跑做些测试,更容易理解整个系统结构.

### 爬虫异常重爬：

1. task_info对应平台crawled_next_time提前一天
2. schedule_task对应平台的status改成4
3. 发送post



