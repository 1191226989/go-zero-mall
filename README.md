# go-zero实战：让微服务Go起来
参考 `go-zero` 入门学习教程的示例代码，教程地址：[go-zero实战：让微服务Go起来](https://juejin.cn/post/7036011047391592485)。

**本项目由以上项目改进为kubesphere部署k8s集群启动所有的微服务，使用了更新版本的protoc工具生成代码**

## 使用说明

### 1. `kubesphere` 安装k8s集群环境 [kubesphere办公网站](https://kubesphere.com.cn/)

### 2. 数据库创建
创建数据库 `mall`

创建数据表 `user`、`product`、`order`、`pay`

`SQL`语句在 `service/[user,product,order,pay]/model` 目录

创建数据库 `dtm_barrier`

`SQL`语句在 `sql/dtm_barrier.sql` 分布式事务

> 提示：相关 mysql 配置，请使用你修改的端口号，账号，密码连接访问数据库。

### 3. 项目中间件搭建

#### 3.1 etcd
kubesphere > 项目 > 应用负载 > 服务 > 有状态服务，创建etcd服务
~~~bash
# 容器环境变量
TZ Asia/Shanghai
ALLOW_NONE_AUTHENTICATION yes
ETCD_ADVERTISE_CLIENT_URLS http://[etcd服务名]:2379
~~~

#### 3.2 redis
kubesphere > 项目 > 应用负载 > 服务 > 有状态服务，创建redis服务
~~~bash
# 配置字典 redis-conf
# redis.conf
appendonly yes
port 6379
bind 0.0.0.0

# 启动命令
redis-server
# 启动参数
/etc/redis/redis.conf

# 只读挂载配置存储卷 redis-conf
/etc/redis/
~~~

#### 3.3 dtm分布式事务
kubesphere > 项目 > 应用负载 > 服务 > 无状态服务，创建dtm服务
~~~bash
# 配置字典 dtm-config
# config.yaml
# 微服务
MicroService:
  Driver: dtm-driver-gozero        # 要处理注册/发现的驱动程序的名称
  Target: etcd://[集群内部etcd服务的dns地址]:2379/dtmservice # 注册 dtm 服务的 etcd 地址
  EndPoint: [集群内部dtm服务的dns地址]:36790

# 容器挂载端口
http-36789
grpc-36790

# 启动命令
/app/dtm/dtm
# 启动参数
-c=/app/dtm/configs/config.yaml

~~~

### 4. 业务微服务部署
- 制作每个微服务的Dockerfile文件 > 打包镜像并push镜像到hub.docker.com
- kubesphere > 项目 > 应用负载 > 应用 > 自制应用 > 创建项目应用
- 创建应用 > 创建服务 > 无状态服务 > 容器配置
~~~bash
# 容器镜像地址

# 容器端口
api TCP
rpc GRPC

# 每个微服务的etc配置文件添加为配置字典，各个中间件服务的地址替换为集群内部该服务的dns地址
# 只读挂载配置存储卷，挂载目录为Dockerfile文件规定的打包以后配置文件所在的目录
~~~

### 5. 项目目录说明
```bash
├── docker # Dockerfile文件目录
├── LICENSE
├── README.md
├── scripts # 工具剧本
│   └── gomall.postman_collection.json
├── sql # 数据库文件
│   ├── dtm_barrier.sql
│   └── mall.sql
└── src # 项目业务代码
    ├── common
    │   ├── cryptx
    │   ├── jwtx
    │   └── randx
    ├── go.mod
    ├── go.sum
    └── service
        ├── order
        ├── pay
        ├── product
        └── user
```
- mysql/redis/etcd 中间件服务部署到公用服务节点（多项目复用）

## 感谢

- [go-zero](https://github.com/zeromicro/go-zero)
- [DTM](https://github.com/dtm-labs/dtm)
- [kubusphere](https://kubesphere.com.cn)