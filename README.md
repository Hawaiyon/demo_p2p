
# 功能说明
金额默认都为int64（BIGINT），可以根据需要，按“分”为单位存储.
以下功能未实现：

- 用户权限验证
- 创建用户和创建交易加入token校验，避免重复提交、CSRF攻击

# 接口说明
要求实现的为以下接口，具体路径已经用swagger定义，可以查看在线文档
1. 创建用户接口，请求的参数支持设置初始金额、返回用户 ID
2. 创建一笔借款交易，参数是两个用户的 ID 和金额
3. 创建一笔还款交易，参数是两个用户的 ID 和金额
4. 查询一个用户的账户情况，参数是用户 ID，返回当前余额、借出总额和借入总额
5. 查询用户之间的债务情况，参数是两个用户的 ID，返回两者间当前的借入借出总额

# 部署说明
1. 安装依赖的go pkg
2. 安装 mysql 5.7/ mariadb, 按`data/p2p.sql`的表结构创建数据库，并且更改`models/dbmodels/db.go`中的数据库连接信息（暂时未做成配置项）
3. 编译运行，可以用`--port `指定监听端口

# 数据库表结构