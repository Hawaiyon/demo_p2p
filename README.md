# 接口文档
接口已经用swagger定义，可以查看[在线文档](http://demo.hawaiyon.ml/swagger-ui/)

# 演示地址
[http://demo.hawaiyon.ml/p2p/v1](http://demo.hawaiyon.ml/p2p/v1
)

# 设计思路
金额默认都为int64（BIGINT），可以根据需要，按“分”为单位存储.
以下功能未实现：
- 用户权限验证
- 创建用户和创建交易加入token校验，避免重复提交、CSRF攻击

因为涉及到事务，所以使用了MySQL来存储数据。

## 数据库设计
数据库有2个表，
1. 用户信息表，记录了用户名，总接入借出金额；
2. 交易信息表，记录交易对象（两个用户id）、金额、时间、是否为还款的信息。
在提交借款、还款时用了事务，避免并发可能带来的问题
表结构如下

user表


| 字段名称 | 类型 | 说明 |
|---|---|---|
| id | BIGINT | id, pk, auto_inc|
|username| char(20)| 用户名 |
|email| char(30) |邮箱|
|balance| bigint| 余额 |
|total_borrow|bigint|当前借入未还的金额|
|total_lend| bigint|当前借出未还的金额|


transaction 表


| 字段名称 | 类型 | 说明 |
|---|---|---|
|id | bigint(20)| pk, auto_inc |
|from_user_id |bigint | 金额转出用户id|
|to_user_id |bigint| 金额转入用户id|
|amount| bigint| 交易金额|
|created_date| datetime | 交易时间|
|status| char | 交易状态|
|is_repay| tinyint(1)|是否为还款交易|

## 开发框架
1. 使用swagger的格式定义接口，使用[go-swagger](https://github.com/go-swagger/go-swagger)生成api server的代码框架
2. 使用xorm 来操作MysQL
3. 自定义的实现大部分在`models/dbmodel/`下，`restapi/configure_p2_p.go`下调用了这些函数。
# 部署说明
1. 安装依赖的go pkg
2. 安装 mysql 5.7/ mariadb, 按`data/p2p.sql`的表结构创建数据库，并且更改`models/dbmodels/db.go`中的数据库连接信息（暂时未做成配置项）
3. 编译运行，可以用`--port `指定监听端口

