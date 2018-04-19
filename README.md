# cos-go-sdk-v5 

*腾讯云 对象储存（cos） API5.0 golang sdk*

---

## 安装使用

---

#### 安装

```shell
go get -u github.com/nelsonken/cos-go-sdk-v5/cos
```

#### 使用

```go
package main

import "github.com/nelsonken/cos-go-sdk-v5/cos"

client := cos.New(cos.Option{})
client.Bucket(name).PutObject(...)

```

## 功能概述

*bucket所有功能完备（生命周期，跨域除外）；object的操作（完备）：增、删、查、改、下载、复制；*

---

### bucket

---

- [x] 列出bucket列表
- [x] 创建bucket
- [x] 删除bucket
- [x] 设置bucket ACL
- [x] 列出bucket内容
- [x] bucket是否存在
- [x] 列出正在上传的obj
- [x] ACL设置
- [ ] CORS 跨域
- [ ] lifcycle设置 


### objcet

---

- [x] 普通上传
- [x] 多线程分片上传（整合分片上传）
- [x] 删除
- [x] 下载
- [x] 复制
- [x] 初始化分片上传
- [x] 上传分片
- [x] 列出正在上传分片
- [x] 完成分片上传
- [x] 放弃上传


    


