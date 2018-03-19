package main

import (
	"bufio"
	"bytes"
	"fmt"
	"strings"
	"time"

	"github.com/nelsonken/cos-go-sdk-v5/cos"
)

var Client *cos.Client
var AppID string = ""
var SecretID string = ""
var SecretKey string = ""
var Region string = ""
var Domain string = ""
var Bucket string = ""

func init() {
	Client = cos.New(&cos.Option{AppID, SecretID, SecretKey, Region, Domain, Bucket})
}

//获取BUCKETList列表
func GetBucketlist() (*cos.ListAllMyBucketsResult, error) {
	ctx := cos.GetTimeoutCtx(time.Second * 5)
	resp, err := Client.GetBucketList(ctx)
	if err != nil {
		return nil, err
	}
	return resp, err
}

//创建一个BUCKET
func CreateBucket(BucketName string) error {
	ctx := cos.GetTimeoutCtx(time.Second * 5)
	err := Client.CreateBucket(ctx, BucketName, &cos.AccessControl{})
	return err
}

//是否存在Bucket
func ExitBucket(BucketName string) (bool, error) {
	ctx := cos.GetTimeoutCtx(time.Second * 5)
	err := Client.BucketExists(ctx, BucketName)
	if err != nil && !strings.Contains(err.Error(), "404") {
		return false, err
	}
	if err != nil && strings.Contains(err.Error(), "404") {
		return false, nil
	}
	return true, err
}

//获取BUCKET中的文件列表
func GetBuckFileList(BucketName string) (*cos.ListBucketResult, error) {
	ctx := cos.GetTimeoutCtx(time.Second * 5)
	res, err := Client.ListBucketContents(ctx, BucketName, &cos.QueryCondition{})
	if err != nil {
		return nil, err
	}
	return res, nil
}

//任务上传列表  不知道怎么用。。
func ListUnLoading() {
	bucket := "ebook"
	ctx := cos.GetTimeoutCtx(time.Second * 5)
	lupr, err := Client.ListUploading(ctx, bucket, &cos.ListUploadParam{})
	if err != nil {

		return
	}

	for _, obj := range lupr.Upload {
		fmt.Println(obj.Key, " ", obj.UploadID)
		err = Client.Bucket(bucket).ObjectExists(ctx, obj.Key)
		fmt.Println(err)
	}
}

//删除Bucket
func DeleteBucket(BucketName string) error {
	ctx := cos.GetTimeoutCtx(time.Second * 20)
	return Client.DeleteBucket(ctx, BucketName)
}

//上传文件
func UnLoadFile(BucketName string, FileName string, FileData string) error {
	ctx := cos.GetTimeoutCtx(time.Second * 30)
	err := Client.Bucket(BucketName).UploadObject(ctx, FileName, strings.NewReader(FileData), &cos.AccessControl{})
	return err
}

//获取文件信息   获取不到conn
//func AcquireFileData(BucketName string, FileName string) {
//	ctx := cos.GetTimeoutCtx(time.Second * 10)
//	bucket := Client.Bucket(BucketName)

//	resq, err := bucket.conn.Do(ctx, "HEAD", bucket.Name, FileName, nil, nil, nil)
//	if err == nil {
//		defer resq.Body.Close()
//	} else {
//		for k, v := range resq.Header {
//			value := fmt.Sprintf("%s", v)
//			fmt.Printf("%-18s: %s\n", k, strings.Replace(strings.Replace(value, "[", "", -1), "]", "", -1))
//		}
//	}

//}

//删除文件
func DeleteFile(BucketName string, FileName string) error {
	ctx := cos.GetTimeoutCtx(time.Second * 30)
	err := Client.Bucket(BucketName).DeleteObject(ctx, FileName)
	return err
}

//批量看不懂
//func UnLoadFileBySlice() {
//	bu := "ebook"
//	ctx := cos.GetTimeoutCtx(time.Second * 30)
//	dst := "testfile_slice"
//	src := "C:\\Users\\dmx\\Desktop\\新建文件夹\\新建文本文档.txt"
//	Client.Bucket(bu).UploadObjectBySlice(ctx, dst, src, 3, nil)

//}

//下载文件
func DownLoadBucketFile(BucketName string, FileName string) {
	ctx := cos.GetTimeoutCtx(time.Second * 60)
	data := bytes.NewBuffer(make([]byte, 0))
	writer := bufio.NewWriter(data)
	Client.Bucket(BucketName).DownloadObject(ctx, FileName, writer)
	fmt.Println(data)
}
