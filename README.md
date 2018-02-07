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
    - 创建bucket
    - 删除bucket
    - 列出bucket内容
    - 列出正在上传的object

- objcet
    - 普通上传
    - 多线程上传
    - 删除
    - 下载
    - 复制


