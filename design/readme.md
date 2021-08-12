# 说明

## 1. 概述

这是设计相关的资源。

## 2. 程序结构 Program Structure 

```plantuml
@startuml
ditaa
+------------------------------------------------------+
|                 interface / log / ...                | framework
+------------------------------------------------------+
| authn / config / object / persistence / script / ... | core module 
+------------------------------------------------------+
@enduml
```

## 3. 核心模块 Core Module 

- rejson
- object-service-core
- object-service-core-rust

## 4. 服务 Service

- http-object-redisjson-go
- http-object-redisjson-openresty
- http-object
- mqtt-object

## 5. 接口 interface

![design-map](image/design-map.png)

## 数据文件

## 项目布局

## 6. 关键依赖

- https://github.com/gin-gonic/gin
- https://github.com/spf13/viper
- https://github.com/yuin/gopher-lua
- https://github.com/protocolbuffers/protobuf