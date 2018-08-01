# 短链接服务

这是一个提供短链接和短地址的服务，可以用来给应用生成短地址。

[English Document](https://github.com/WindGreen/ifth/blob/master/README.md)

演示地址：http://www.fith.net

![example](https://raw.githubusercontent.com/WindGreen/ifth/master/ifth-www.png)

## 功能

- 提供生成地址的页面
  - 可选http/https
  - 可选长度（开发中）
  - 申请个性化地址（规划中）
- 提供生成地址的接口（开发中）
  - 可指定长度



## 安装

### 编译

本程序依赖MongoDB，请先安装mongoDB

- 网站服务

  ```shell
  cd www && go build && ./www
  ```

  

- 短链接服务

  ```shell
  cd url && go build && ./url
  ```

  

### Docker

- 网站服务

  ```shell
  docker network create ifth
  docker run -d --name mongo --network ifth mongo
  docker run -d -p 80:80 --network ifth windgreen/ifth-www:1.0.0
  ```

- 短链接服务

  ```shell
  docker network create ifth
  docker run -d --name mongo --network ifth mongo
  docker run -d -p 80:80 --network ifth windgreen/ifth-url:1.0.0
  ```

  

## 配置

DNS配置

服务器端口

数据库配置

地址ID长度

是否允许给相同地址生成多个短地址

