### 一、锁机制

> 锁分为 行级锁、页级锁、表级锁

#### 1 锁分类

1. 行级锁：
   1. 行级锁锁定颗粒度最小的,所以发生锁定资源争用的概率也最小，能够给予应用程序尽可能大的并发处理能力而提高一些需要高并发应用系统的整体性能
   2. 由于锁定资源的颗粒度很小，所以每次获取锁和释放锁消耗的资源也更多，带来的消耗自然也就更大了。此外，行级锁定也最容易发生死锁。
   3. 行级锁定的主要是Innodb存储引擎和NDBCluster存储引擎
2. 表级锁
   1. 表级锁最大颗粒度的锁定机制。由于表级锁一次会将整个表锁定，所以可以很好的避免死锁问题。  当然，锁定颗粒度大所带来最大的负面影响就是出现锁定资源争用的概率也会最高，致使并大度较低。
   2. MySQL数据库中，使用表级锁定的主要是MyISAM，Memory，CSV等一些非事务性存储引擎
3. 页级锁
   1. 页级锁定的特点是锁定颗粒度介于行级锁定与表级锁之间，所以获取锁定所需要的资源开销，以及所能提供的并发处理能力也同样是介于上面二者之间。另外，页级锁定和行级锁定一样，会发生死锁。 

```sql
-- 一些监控指标
一、慢查询
　　mysql> show variables like '%slow%';
　　+------------------+-------+
　　| Variable_name | Value |
　　+------------------+-------+
　　| log_slow_queries | ON |
　　| slow_launch_time | 2 |
　　+------------------+-------+
　　mysql> show global status like '%slow%';
　　+---------------------+-------+
　　| Variable_name | Value |
　　+---------------------+-------+
　　| Slow_launch_threads | 0 |
　　| Slow_queries | 4148 |
　　+---------------------+-------+　　
十、表锁情况
　　mysql> show global status like 'table_locks%';
　　+-----------------------+-----------+
　　| Variable_name | Value |
　　+-----------------------+-----------+
　　| Table_locks_immediate | 490206328 |
　　| Table_locks_waited | 2084912 |
　　+-----------------------+-----------+　　
-- [FOR MORE]  https://blog.csdn.net/iteye_11910/article/details/82366959
```



