# Short Url Service

This is a service to generate short url for long url.

[中文文档](./Readme_zh-cn.md)

Demo：http://www.fith.net

![example](/home/yqfwind/go/src/ifth/ifth-www.png)

## Features

- Web page to generate short url
  - Choosing http/https
  - Choosing length (developing)
  - Custom ID (Scheduling)
- API to generate short url(developing)
  - Length input



## Installation

### Build

Service is depended on MongoDB, please install first. [Mongo Installation](https://docs.mongodb.com/manual/installation/)

- Web page

  ```shell
  cd www && go build && ./www
  ```

  

- Url Service

  ```shell
  cd url && go build && ./url
  ```

  

### Docker

- Web page

  ```shell
  docker network create ifth
  docker run -d --name mongo --network ifth mongo
  docker run -d -p 80:80 --network ifth windgreen/ifth-www:1.0.0
  ```

- Url Service

  ```shell
  docker network create ifth
  docker run -d --name mongo --network ifth mongo
  docker run -d -p 80:80 --network ifth windgreen/ifth-url:1.0.0
  ```

  

## Configration

DNS tips

Server Port

MongoDB Connection

Length ID for default generation

Difference ID for one url

