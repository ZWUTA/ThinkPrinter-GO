# 全新JWB打印系统

ThinkPrint-GO

## 简介

ThinkPrint-GO是一个基于Golang的全新JWB打印系统，它的前身是ThinkPrint-python

本项目采用纯Golang重写，前后端分离设计。

前端使用Vue3 + NaiveUI，后端使用SQLite、Gin、Gorm、golang-jwt、viper、slog

注意，本项目需要设备内安装``Microsoft Word 2016``或``Microsoft Office 365``以上版本，否则无法正常使用

## 功能

- [x] 用户注册
- [x] 用户登录
- [x] JWT鉴权
- [x] 批量文件上传
- [x] 多用户打印队列
- [x] 打印历史日志
- [x] Microsoft Word打印
- [x] 高准确率解析

支持的文件格式：
- [x] docx
- [x] doc
- [x] pdf
- [x] odt
- [x] rtf

## 部署

### 服务器端

1. 安装``Microsoft Word 2016``以上或``Microsoft Office 365``
2. 从GitHub下载最新的``ThinkPrint-GO``二进制文件和``thinkprint-vue``前端包
3. 将``ThinkPrint-GO``二进制文件和``thinkprint-vue``前端包放置在同一目录下。
4. 请确保``dist``文件夹与``ThinkPrint-GO``同级
5. 将工作目录切换到``ThinkPrint-GO``所在目录
6. 运行``ThinkPrint-GO``，首次启动程序将自动创建``config.toml``配置文件
7. 自定义修改``config.toml``配置文件
8. 再次运行``ThinkPrint-GO``，程序将自动创建数据库文件``sqlite.db``

### 配置文件注解

````toml
[core]
# 服务器端口号
Port = 5204
# 服务器绑定IP
Bind = '0.0.0.0'
# 日志等级，可选：debug, info, warn, error
LogLevel = 'info'

[print]
# Word程序路径，注意使用双斜杠，不要使用单斜杠。
# 请根据实际情况修改，如果你的Word安装在其他位置，请自行修改
# 请务必确保Word程序路径正确，否则无法正常打印
WordExePath = 'C:\\Program Files\\Microsoft Office\\root\\Office16\\WINWORD.EXE'

[security]
# 密码盐值，创建用户前请修改
PasswordSalt = 'ThinkPrinter'
# JWT密钥，使用UUID生成，请随意修改
JWTSecret = 'c487798f-0335-43c0-906f-4feb247132a4'
# JWT过期时间，单位为秒
JWTExpiration = 1800
````

### 客户端

1. 访问``http://<服务器IP>:<端口号>``，默认端口号为``5204``
2. 进行登录或注册
3. 上传文件（可多选）
4. 你将看到前方排队打印人数
5. 等待打印完成，取走你的文件

## git commit

✨`:sparkles:`feat：新功能（feature）

🐛`:bug:`fix：修补bug

📝`:memo:`docs：文档（documentation）

🎨`:art:`style： 格式（不影响代码运行的变动）

♻️`:recycle:`refactor：重构（即不是新增功能，也不是修改bug的代码变动）

🧪`:test_tube:`test：增加测试

🔨`:hammer:`chore：构建过程或辅助工具的变动
