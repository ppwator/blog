sql优化: https://www.eversql.com/?utm_medium=organic&utm_source=ext_blog&utm_content=ftisiotwebsite
## 1. explain
### 1. type: 
system:表只有一行记录
const:通过索引一次就找到.用于比较 主键或唯一索引
eq_ref: 连接查询. 常见于主键或唯一索引
ref: 普通索引
### 2. ref
索引的哪一列使用了
### 3. rows
需要遍历的行数
### 4.extra
Using temporary: 分组时未用到索引.一般出现时Using filesort也会出现. 
## 2.索引
is null可以使用索引,is not null不使用索引
### 1.单表优化
1. sample
```sql
where cid=1 and comments>1 order by views desc limit 1;
```

可以建立cid和views的复合索引优化. 
```sql
alter table a add index idx_cv(cid,views)
```
### 2.关联查询优化
内连接时,MySQL会自动把小结果集的表选为驱动表,所以大表的字段最好加上索引.
左连接时,左表会全表扫描,所以右边大表字段最好加上索引.右连接同理
### 3. 查看约束constraint
SELECT * from information_schema.table_constraints WHERE TABLE_NAME='user';
### 4. 外键约束
1、Cascade方式。

当update/delete记录在主表上时，同步update/delete掉从表的匹配记录。

2、Set null方式。

当update/delete记录在主表上时，从表上匹配记录的列被设置为null，但需要注意的是，子表的外键列不能是notnull。

3、No action方式。

如果子表中有匹配记录，则不允许update/delete操作与父表相对应的候选键。

4、Restrict方法，与no action相同。

立即检查外键约束。

5、Set default模式，空白可能显示在可视化工具SQLyog中。

当父表发生变化时，子表将外键列设置为默认值，但Innodb无法识别。

推荐
on update cascade on delete restrict
### 5. 定义条件与处理程序
1. 定义条件
declare 错误 condition for 错误码/错误条件
2. 处理程序

1062: mysql_error_code
```sql
declare exit/continue/undo handler for 1062 set @sv=-1;
```
### 6. 流程控制
循环结构：loop ， while...do
repeat ... until... 相当于 do while循环
### 7. MySQL 8
1. check约束

2. DDL原子化

3. 计算列 int generated always as(a+b) virtual
重启后自增不重置，写在redo_log. mysql 5.7重启后，自增会重置，保存在内存。 （删除后重启，8显示为主键不连续）

4. 全局变量持久化persist

5. 内部临时表的默认存储引擎由memory改为tempTable

6. 窗口函数
作用类似于分组函数。不合并结果，数据多少条就有多少结果

select 窗口函数 over 条件(相当于where后条件)

7. 公用表表达式CTE(common table compressions)
可以多次调用.子查询不能多次调用
- 普通：  with 名称 as 表达式
- 递归：  with recursive  名称 as 表达式
### 8. 存储过程与存储函数
存储过程 procedure， 调用call
存储函数 function ， 调用select 
### 9. 游标cursor
cursor for 
open
fetch
close
### 10. 触发器trigger
条件insert/update时可以使用new，delete可以使用old表示表中已经存在的记录
```
create trigger ... 
before/after insert/update/delete on 表
for each row
条件
```
自定义错误
```
then signal sqlstate 'HY000' set message_text="高于a"
```

