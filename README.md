# 【一】前言

## 【1】为什么选择GVA

### （1）介绍

- GIN-VUE-ADMIN：是一款基于GIN+VUE+ElementUI开发的全栈基础开发平台。
- GIN-VUE-ADMIN是一个基于vue和gin开发的全栈前后端分离的开发基础平台，拥有jwt鉴权，动态路由，动态菜单，casbin鉴权，表单生成器，代码生成器等功能，提供了多种示例文件，让大家把更多时间专注在业务开发上。

### （2）官网

- https://www.gin-vue-admin.com/guide/introduce/project.html

## 【2】技术选型

### （1）相关官网

- Gin文档：https://gin-gonic.com/zh-cn/docs/

- gorm数据操作：https://gorm.io/
- 数据：myql >=5.7+ mysql8+ 
- Vue3 文档：https://cn.vuejs.org/
- Vue-router: https://router.vuejs.org/zh/guide/essentials/nested-routes.html
- Pinia: https://pinia.vuejs.org/zh/
- ElementUI文档：https://element-plus.gitee.io
- Redis 缓存：https://redis.io/
- Git版本控制: https://git-scm.com/

### （2）Docker部署脚本

#### [1] MySQL8+

- 本地项目结构

```ini
├── docker-compose.yml # docker-compose.yml文件
├── conf
│   └── my.cnf  # MySQL配置文件
├── data # 数据库数据文件目录
├── log # 日志存放目录
└── bak # 手动备份数据的目录
```

- `docker-compose` 配置文件

```ini
services:
  mysqldb:
    # 启动方式
    #restart: unless-stopped
    
    # 镜像
    image: mysql:8.0
    # 容器名字
    container_name: mysql8
    privileged: true
    environment:
      # 时区
      TZ: Asia/Shanghai
      # root用户的密码
      MYSQL_ROOT_PASSWORD: 222333
      # 用户(不能是root, 后续需要给此用户赋予权限)
      MYSQL_USER: adminUser
      # 用户密码，建议不要和root一样。
      MYSQL_PASSWORD: 333444
    command:
      --character-set-server=utf8mb4
      --collation-server=utf8mb4_general_ci
      --explicit_defaults_for_timestamp=true
    ports:
      - 3306:3306
    volumes:
      - ./log/:/var/log/mysql/
      - ./data/:/var/lib/mysql/
      - ./conf/:/etc/mysql/conf.d/
      - ./bak/:/bak/
```

- `my.cnf` 配置文件

```ini
[client]
default-character-set=utf8mb4
 
[mysql]
default-character-set=utf8mb4
 
[mysqld]
#服务端口号 默认3306
port=3306
 
# 数据路径，默认是/var/lib/mysql/
#datadir = /app/data/
 
init_connect='SET NAMES utf8mb4'
character-set-server=utf8mb4
collation-server=utf8mb4_unicode_ci
 
# 最大连接数
max_connections=50
 
# 连接失败的最大次数。防止有人从该主机试图攻击数据库系统
max_connect_errors=20
 
# 创建新表时将使用的默认存储引擎
default-storage-engine=INNODB
 
skip-name-resolve
 
# 最大binlog文件大小。达到设定大小后，会创建新的bin log日志。
max_binlog_size=500M
# binlog过期天数。默认是0（永远存在）
# expire_logs_days=5
 
# 若不配置log-error默认输出到控制台。
# 若只配置但不指定文件，即：log-error，，则为datadir目录下的 `hostname`.err，hostname为主机名。
log-error = /var/log/mysql/error.log
 
# 是否启用慢查询日志，默认不启用
# slow_query_log={1|ON|0|OFF} 
slow_query_log=ON
# 默认路径为库文件目录下主机名加上-slow.log
slow_query_log_file=/var/log/mysql/slow.log 
 
# 指定慢查询超时时长(默认10秒)，超出此时长的属于慢查询
long_query_time=10 
# 定义一般查询日志和慢查询日志的输出格式，默认为file
# log_output={TABLE|FILE|NONE} 
# 查询没有使用索引的时候是否也记入慢查询日志，默认OFF
# log_queries_not_using_indexes=OFF
 
# 开启一般查询日志，默认off关闭
# general_log=ON
# 指定一般查询日志路径
# general_log_file=/var/log/mysql/general.log 
 
# 用于监控MySQL服务在运行过程中的资源消耗、资源等待等情况，默认ON
# performance_schema = OFF
```

- 启动容器

```shell
docker-compose up -d 
```

- 赋予用户操作权限

```shell
# 进到MySQL这个docker中
docker exec -it mysql8 bash

# 连接到MySQL服务
mysql -uroot -p

# 授予root用户所有权限
GRANT ALL ON *.* TO 'root'@'%';

# 授予新用户所有权限
GRANT ALL ON *.* TO 'adminUser'@'%';

#  刷新权限
FLUSH PRIVILEGES;
```

#### [2] Redis

- 项目结构

```ini
myProject                       ----项目名
├── compose      
│    └── docker-compose.yml     ----compose配置文件
└── redis     
   ├── data                     ----redis数据存储目录     
   ├── logs                     ----redis日志文件目录(redis.conf中logfile设置相对路径则不需要)     
   └── redis.conf               ----redis配置文件
```

- `redis.conf` 配置文件内容

```ini
#开启远程可连接
#bind 127.0.0.1
#自定义密码
requirepass 12345678
#指定 Redis 监听端口(默认:6379)
port 6379
#客户端闲置指定时长后关闭连接(单位:秒。0:关闭该功能)
timeout 0
# 900s内如果至少一次写操作则执行bgsave进行RDB持久化操作
save 900 1
# 在300s内，如果至少有10个key进行了修改，则进行持久化操作
save 300 10
#在60s内，如果至少有10000个key进行了修改，则进行持久化操作
save 60 10000
#是否压缩数据存储(默认:yes。Redis采用LZ 压缩，如果为了节省 CPU 时间，可以关闭该选项，但会导致数据库文件变的巨大)
rdbcompression yes
#指定本地数据文件名(默认:dump.rdb)
dbfilename dump.rdb
#指定本地数据文件存放目录
dir /data
#指定日志文件位置(如果是相对路径，redis会将日志存放到指定的dir目录下)
logfile "redis.log"
```

- `docker-compose.yml` 配置文件内容

```ini
version: "3.8"
services:
  redis:
    # 镜像及版本      
    image: redis:6.2.6
    # 自定义容器名
    container_name: my-redis
    # docker启动时,自动启动该容器
    restart: always
    # 挂载映射，可以让数据或配置持久化
    volumes:
      # <本地配置文件> : <docker中的配置文件> : <ro:docker容器对该文件只读,默认是rw可读可写>
      - ../redis/redis.conf:/etc/redis/redis.conf:ro
      # <本地数据目录> : <docker中的数据目录>
      - ../redis/data:/data
      # <本地日志目录> : <docker中的日志目录>
      # redis不具有自动创建/logs的权限，如果redis.conf中指定的相对位置,则数据目录已经可以映射出日志文件
      #- ../redis/logs:/logs
    # docker执行的启动命令
    command: redis-server /etc/redis/redis.conf
    ports:
     # <本地端口> : <docker容器端口>
     - 6378:6379
```

- 部署启动

```shell
#  cd compose
#  docker-compose up
 
// 或以后台方式启动
#  docker-compose up -d
 
// 关闭docker-compose
#  docker-compose down
```

## 【3】GVA现有开发的模块

- 插件中心 NEW :基于 Gva自己的一套设计风格，独创 go的插件中心，现已支持 ：微信支付、登录等，K8s相关操作 ，第三方登录 等等插件
- 权限管理：基于jwt和casbin实现的权限管理
- 文件上传下载：实现基于七牛云的文件上传操作（为了方便大家测试，我公开了自己的七牛测试号的各种重要token，恳请大家不要乱传东西）
- 用户管理：系统管理员分配用户角色和角色权限。
- 角色管理：创建权限控制的主要对象，可以给角色分配不同api权限和菜单权限。
- 菜单管理：实现用户动态菜单配置，实现不同角色不同菜单。
- api管理：不同用户可调用的api接口的权限不同。
- 配置管理：配置文件可前台修改（测试环境不开放此功能）。
- 富文本编辑器：MarkDown编辑器功能嵌入。
- 条件搜索：增加条件搜索示例。
- restful示例：可以参考用户管理模块中的示例API。

## 【4】任务

- 提供运行环境
- 安装和下载gva
- 运行gva
- 开发研究gva开发流程和步骤
- 添加用户
- 添加管理员
- 添加角色
- 给角色授权
- 添加菜单
- 分配路由
- 定义业务
- 开始添加你真正意义的模块，比如：社区问答，课程，数据管理等等开始放入gva项目中。
- ==自动构建（通过数据库表自动生成对应的模块）==

## 【5】目标需求

- 登录流程+验证码
- JWT接口的安全性
- 权限的开发
- 权限控制（vue指令）
- 国际化处理
- 路由管理（动态路由）
- 学会elementplus调用
- 尝试学会自定定义组件
- gin的路由管理（路由组）
- gin的中间的处理(权限拦截和token拦截)
- redis的处理
- 配置yaml或properties解析
- 日志的保存和zap
- 统计和分析
- 发布和部署项目

# 【二】GVA的准备工作

## 【1】**Go环境的安装**

- 版本：1.19.2

## 【2】MySQL8

- https://dev.mysql.com/downloads/mysql/

## 【3】**MYSQL 图形化界面工具**

- navicat161_premium_cs_x64.exe
- syslog

## 【4】Redis

- Linux版本：https://github.com/redis/redis/archive/7.0.11.tar.gz
- Windows版本：Redis-x64-5.0.14.1.msi
  - windiow版本的下载：https://github.com/tporadowski/redis/tags

## 【5】Redis 图形化界面工具

- redis-desktop-manager-0.8.8.384

## 【6】NodeJS

- https://nodejs.org/download/release/latest-v16.x/

## 【7】Git

- https://git-scm.com/
- 或者安装图形化界面工具：TortoiseGit-2.13.0.1-64bit.msi

## 【8】数据库创建

```ini
# 创建数据库
create database social_project
```

# 【三】GVA项目初始化

## 【1】GitHub地址

- 克隆 GVA 到本地 

```shell
git clone https://github.com/flipped-aurora/gin-vue-admin
```

## 【2】后端启动

- 打开 `server` 地址 `gin-vue-admin/server`

- 配置 Go 代理 `GOPROXY=https://goproxy.io,direct`

- 启动 Server 

  ```shell
  go mod tidy
  
  # 把go的模块，清理一下，看看还有哪些没有下载或者没有同步。就全部进行重新下载和同步一次。
  #
  ```

- 然后开始启动项目，找到项目下的 `main.go` 文件开始运行和启动即可

  - 在 `main.go` 得代码可以看到，如果你第一次运行和安装go可以尝试执行如下得命令

  ```shell
  go:generate go env -w GO111MODULE=on
  go:generate go env -w GOPROXY=https://goproxy.cn,direct
  go:generate go mod tidy
  go:generate go mod download
  ```

## 【3】前端启动

- 打开 `web` 项目地址 `gin-vue-admin/web`

- 使用 `pnpm` 来下载和管理

  ```shell
  # 先在全局安全pnpm
  npm install -g pnpm
  
  # 然后在项目web命令下执行
  pnpm install
  ```

- 使用 `pnpm` 或者 `npm` 来启动项目

  ```shell
  npm run serve
  pnpm run serve
  ```

> 不要纠结 `pnpm` 启动失败得问题，失败了就直接使用`npm run serve`来启动即可。

## 【4】安装数据库和脚本

- 保证 `gin server`是启动状态
- web也是正确启动状态
- 然后访问： http://localhost:8080

- 这里点击立即初始化，会把对应得数据库 `gva` 创建好，并且把对应得数据库表也会创建好。

- 同时也会把项目中得 `config.yaml`文件中关于 MySQL 得配置也会自动配置好

  ```ini
  mysql:
    path: 127.0.0.1
    port: "3306"
    config: charset=utf8mb4&parseTime=True&loc=Local
    db-name: ksd-social-db
    username: root
    password: mkxiaoer
    prefix: ""
    singular: false
    engine: ""
    max-idle-conns: 10
    max-open-conns: 100
    log-mode: error
    log-zap: false
  ```

- 然后把 `go server` 服务重启即可。然后开始体验 `gva` 吧！！

- 然后输入：http://localhost:8080，然后输入账号和密码如下：

  - 用户名：admin

  - 密码：123456