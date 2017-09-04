package cos

import (
	"fmt"
	"testing"
	"time"
)

func TestClient_GetBucketList(t *testing.T) {
	setUp()
	ctx := GetTimeoutCtx(time.Second * 5)
	resp, err := client.GetBucketList(ctx)
	if err != nil {
		t.Error(err)
		return
	}

	for _, v := range resp.Buckets.Bucket {
		fmt.Println(v.Name)
	}
}

func TestClient_CreateBucket(t *testing.T) {
	setUp()
	bucket := "hellonewbuckettest"
	ctx := GetTimeoutCtx(time.Second * 5)
	err := client.CreateBucket(ctx, bucket, &AccessControl{})
	if err != nil {
		t.Error(err)
	}
}

func TestClient_BucketExists(t *testing.T) {
	setUp()
	bucket := "hellocos"
	ctx := GetTimeoutCtx(time.Second * 5)
	err := client.BucketExists(ctx, bucket)
	if err != nil {
		t.Error(err)
	}
}

func TestClient_ListBucketContents(t *testing.T) {
	setUp()
	bucket := "hellocos"
	ctx := GetTimeoutCtx(time.Second * 5)
	res, err := client.ListBucketContents(ctx, bucket, &QueryCondition{})
	if err != nil {
		t.Error(err)
		return
	}

	for _, obj := range res.Contents {
		fmt.Println(obj.Key)
	}
}

func TestClient_ListUploading(t *testing.T) {
	setUp()
	bucket := "hellocos"
	ctx := GetTimeoutCtx(time.Second * 5)
	lupr, err := client.ListUploading(ctx, bucket, &ListUploadParam{})
	if err != nil {
		t.Error(err)
		return
	}

	for _, obj := range lupr.Upload {
		fmt.Println(obj.Key, " ", obj.UploadID)
		err = client.Bucket(bucket).ObjectExists(ctx, obj.Key)
		fmt.Println(err)
	}
}
