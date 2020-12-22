### #{} 与 ${} 的效果是不一样的

在动态sql解析过程，#{} 与 ${} 的效果是不一样的：  

1. #{ } 解析为一个 JDBC 预编译语句（prepared statement）的参数标记符 ？。 　 

select * from user where  = #{id}; 　　
// 会被解析为： select * from user where id = ?; 
#{} 被解析为一个参数占位符？

2. ${ } 仅仅是String替换，在动态 SQL 解析阶段将会进行变量替换 　

select * from user where id = ${id};
// 若传入 id="12"   
// 变成 select * from user where id = "12"
虽然id是int类型，但是此时我们仍要传String类型，new Integer(id).toString(), 预编译之前的sql语句已经不包含变量name。

总结：

1. ${ } 的变量的替换阶段是在动态 SQL 解析阶段，而 #{ }的变量的替换是在DBMS中（在解析阶段变成 ‘？’）

2. 如果使用了${}， 那么传递来的参数必须是字符串。

3. 声明了statementType="STATEMENT",那么都使用${}, 传字符串

id传来的是字符串，不然报错
