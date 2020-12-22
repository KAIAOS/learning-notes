![img](https://upload-images.jianshu.io/upload_images/845143-a8f0bc0ad45017f8.png?imageMogr2/auto-orient/strip|imageView2/2/w/498/format/webp)

```git
#关联远程仓库
git init 
git remote add origin https://github.com/KAIAOS/learning-notes.git
git add .
git commit -m 'init repository'
git push -u origin master
```



``` git
#建新分支
git checkout -b dev #创建新分支 = git branch dev and git checkout dev
git push origin dev #本地分支push到远端
git branch --set-upstream-to=origin/dev #本地分支关联远端分支
```

> git合并流程
>
> https://www.jianshu.com/p/4142210eb2b2

``` git
//不适用快进式合并
git checkout master
git merge --no-ff develop

//合并commit
git checkout master 
git merge --squash develop
git commit -m "branch功能完成，合并到主干" 

//撤销merge用来撤销还没commit 的merge,其实原理就是放弃index和工作区的改动。
1、git merge --abort
2、git reset --hard HEAD

//查看分支信息
git branch -a
//删除本地分支
git branch -d <BranchName>
//删除远程分支
git push origin --delete <BranchName>

//取消本地的commit
git log //找到commit ca31e8528fd92e9d174c87c067d663497b3333e1 (origin/dev)
git reset --hard commit_id

关于git reset命令,包括 --mixed，--soft --hard等，其中--mixed为默认方式，他们之间的区别如下
git reset –mixed：此为默认方式，不带任何参数的git reset，即时这种方式，它回退到某个版本，只保留源码，回退commit和index信息
git reset –soft：回退到某个版本，只回退了commit的信息，不会恢复到index file一级。如果还要提交，直接commit即可
git reset –hard：彻底回退到某个版本，本地的源码也会变为上一个版本的内容

git reset -soft :取消了commit  
git reset -mixed（默认） :取消了commit ，取消了add
git reset -hard :取消了commit ，取消了add，取消源文件修改
```



![image-20200723065722123](C:\Users\hanka\AppData\Roaming\Typora\typora-user-images\image-20200723065722123.png)

| Overview                                                     | Messages | Message rates | +/-   |       |         |       |          |               |      |
| :----------------------------------------------------------- | :------- | :------------ | :---- | :---- | :------ | :---- | :------- | :------------ | :--- |
| Name                                                         | Type     | Features      | State | Ready | Unacked | Total | incoming | deliver / get | ack  |
| [ExcelImportFail](http://localhost:15672/#/queues/%2F/ExcelImportFail) | classic  | D Args        | idle  | 0     | 0       | 0     |          |               |      |
| [ExcelImportFormatInvalid](http://localhost:15672/#/queues/%2F/ExcelImportFormatInvalid) | classic  | D Args        | idle  | 0     | 0       | 0     |          |               |      |
| [ExcelImportSuccess](http://localhost:15672/#/queues/%2F/ExcelImportSuccess) | classic  | D Args        | idle  | 0     | 0       | 0     |          |               |      |
| [ExcelImportTask](http://localhost:15672/#/queues/%2F/ExcelImportTask) | classic  | D Args        | idle  | 0     | 0       | 0     |          |               |      |
| [ExcelImportWarning](http://localhost:15672/#/queues/%2F/ExcelImportWarning) | classic  | D Args        | idle  | 0     | 0       | 0     |          |               |      |
| [GenerateBill](http://localhost:15672/#/queues/%2F/GenerateBill) |          |               |       |       |         |       |          |               |      |