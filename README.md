# cos-go-sdk-v5 腾讯云 云储存API5.0 golang实现

## 安装使用
---

- 安装
```shell
go get -u github.com/nelsonken/cos-go-sdk-v5/cos
```

- 使用
```go
package main

import "github.com/nelsonken/cos-go-sdk-v5/cos"

client := cos.New(appid, secretid, secretkey, region)
client.Bucket(name).PutObject(...)

```

## 功能概述
---

- bucket
    - [x] 列出bucket列表
    - [x] 创建bucket
    - [x] 删除bucket
    - [x] 设置bucket ACL
    - [x] 列出bucket内容
    - [x] bucket是否存在
    - [x] 列出正在上传的obj


- objcet
    - [x] 普通上传
    - [x] 多线程分片上传
    - [x] 删除
    - [x] 下载
    - [x] 复制
    - [x] 上传分片
    


