package cos

import (
	"strings"
	"testing"
	"time"
)

var client *Client

func setUp() {
	client = New(&Option{
		AppID:     "",
		SecretID:  "",
		SecretKey: "",
		Region:    "",
		Domain:    "", // default is myqcloud.com
		Bucket:    "",
	})
}

func TestBucket_UploadObject(t *testing.T) {
	setUp()
	bu := "hellocos"
	ctx := GetTimeoutCtx(time.Second * 30)
	objName := "testfile"
	objContent := strings.Repeat("t", 1024*1024)
	err := client.Bucket(bu).UploadObject(ctx, objName, strings.NewReader(objContent), &AccessControl{})

	if err != nil {
		t.Error(err)
	}
}

func TestBucket_DeleteObject(t *testing.T) {
	setUp()
	bu := "hellocos"
	ctx := GetTimeoutCtx(time.Second * 30)
	obj := "testfile"
	err := client.Bucket(bu).DeleteObject(ctx, obj)

	if err != nil {
		t.Error(err)
	}
}

func TestBucket_UploadObjectBySlice(t *testing.T) {
	setUp()
	bu := "hellocos"
	ctx := GetTimeoutCtx(time.Second * 30)
	dst := "testfile_slice"
	src := "testfiles/test.zip"
	err := client.Bucket(bu).UploadObjectBySlice(ctx, dst, src, 3, nil)
	if err != nil {
		t.Error(err)
	}
}
