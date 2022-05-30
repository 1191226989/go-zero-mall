# user service

| api 服务 | 端口：8000 | rpc 服务 | 端口：9000 |
| ----  | ---- | ---- | ---- |
| login | 用户登录接口 | login | 用户登录接口 |
| register | 用户注册接口 | register | 用户注册接口 |
| userinfo | 用户信息接口 | userinfo | 用户信息接口 |


### 1. 生成 user model 模型

- user/model/user.sql
```sql
CREATE TABLE `user` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(255)  NOT NULL DEFAULT '' COMMENT '用户姓名',
  `gender` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '用户性别',
  `mobile` varchar(255)  NOT NULL DEFAULT '' COMMENT '用户电话',
  `password` varchar(255)  NOT NULL DEFAULT '' COMMENT '用户密码',
  `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_mobile_unique` (`mobile`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4;

```

- 运行model模板生成命令
```shell
goctl model mysql ddl -src ./model/user.sql -dir ./model -c
```

### 2. 生成 user rpc 服务

- 编写 `user/rpc/user.proto` 文件

- 运行rpc模板生成命令
```shell
goctl rpc protoc user.proto --go_out=. --go-grpc_out=. --zrpc_out=.
```

- 回到 mall 项目根目录执行
```shell
go mod tidy
```

### 3. 编写 user rpc 服务

- 修改 `rpc/etc/user.yaml`
修改服务监听地址，端口号为0.0.0.0:9000，Etcd 服务配置，Mysql 服务配置，CacheRedis 服务配置

- 添加 user model 依赖，在`user/rpc/internal/config/config.go `添加 Mysql 服务配置，CacheRedis 服务配置的实例化

- 在 `user/rpc/internal/svc/servicecontext.go` 注册服务上下文 user model 的依赖

- 在 `user/rpc/internal/logic/registerlogic.go` 添加用户注册逻辑 Register

- 在 `user/rpc/internal/logic/loginlogic.go` 添加用户登录逻辑 Login

- 在 `rpc/internal/logic/userinfologic.go` 添加用户信息逻辑 UserInfo

### 4. 生成 user api 服务

- 编写 `user/api/user.api`
```
type (
  // 用户登录
  LoginRequest {
    Mobile   string `json:"mobile"`
    Password string `json:"password"`
  }
  LoginResponse {
    AccessToken  string `json:"accessToken"`
    AccessExpire int64  `json:"accessExpire"`
  }
  // 用户登录

  // 用户注册
  RegisterRequest {
    Name     string `json:"name"`
    Gender   int64  `json:"gender"`
    Mobile   string `json:"mobile"`
    Password string `json:"password"`
  }
  RegisterResponse {
    Id     int64  `json:"id"`
    Name   string `json:"name"`
    Gender int64  `json:"gender"`
    Mobile string `json:"mobile"`
  }
  // 用户注册

  // 用户信息
  UserInfoResponse {
    Id     int64  `json:"id"`
    Name   string `json:"name"`
    Gender int64  `json:"gender"`
    Mobile string `json:"mobile"`
  }
  // 用户信息
)

service User {
  @handler Login
  post /api/user/login(LoginRequest) returns (LoginResponse)
  
  @handler Register
  post /api/user/register(RegisterRequest) returns (RegisterResponse)
}

@server(
  jwt: Auth
)
service User {
  @handler UserInfo
  post /api/user/userinfo returns (UserInfoResponse)
}
```

- 运行api模板生成命令
```shell
goctl api go -api ./api/user.api -dir ./api
```

### 5. 编写 user api 服务

- 修改 `user/api/etc/user.yaml` 配置文件，修改服务地址，端口号为0.0.0.0:8000，Mysql 服务配置，CacheRedis 服务配置，Auth 验证配置

- 在 `user/api/internal/config/config.go` 添加 user rpc 服务配置的实例化

- 在 `user/api/internal/svc/servicecontext.go` 注册服务上下文 user rpc 的依赖

- 在 `user/api/internal/logic/registerlogic.go` 添加用户注册逻辑 Register

- 在 `user/api/internal/logic/loginlogic.go` 添加用户登录逻辑

- 在 `user/api/internal/logic/userinfologic.go` 添加用户信息逻辑 UserInfo

### 6. 本地调试运行

- 启动 user rpc 服务

- 启动 user api 服务

### 7. 开发完成，复制api和rpc各自的etc配置文件到 `src/etc` 目录，替换配置文件与 `.env` 文件相关的公共环境变量（配置文件统一复制到 `src/etc` 目录是为了运维修改方便）

- src/etc/user-api.yaml
- src/etc/user-rpc.yaml

### 8. 增加user-rpc user-api的服务容器配置
```yaml
# 服务容器配置
services:
  user-rpc:                                # 自定义容器名称
    build:
      context: ./src                # 指定构建使用的 Dockerfile 文件
      # dockerfile: ./src/Dockerfile
      args:                      # 指定构建使用的 go/etc 文件
        - BUILD_FILE_MAIN=service/user/rpc/user.go
        - BUILD_FILE_ETC=etc/user-rpc.yaml
    image: user-rpc:1.0 
    # environment:                         # 设置环境变量
    #   - TZ=${TZ}
    privileged: true
    ports:                               # 设置端口映射
      - "9000:9000"
    stdin_open: true                     # 打开标准输入，可以接受外部输入
    tty: true
    networks:
      - net-mall
    restart: always                      # 指定容器退出后的重启策略为始终重启

  user-api:
    build:
      context: ./src
      args:
        - BUILD_FILE_MAIN=service/user/api/user.go
        - BUILD_FILE_ETC=etc/user-api.yaml
    image: user-api:1.0 
    privileged: true
    ports:
      - "8000:8000"
    stdin_open: true
    tty: true
    networks:
      - net-mall
    restart: always
    depends_on:
      - user-rpc
```
### 8. 在docker-compose.yml所在的目录启动容器服务
```shell
docker-compose up -d
```