package cos

import (
	"context"
	"fmt"
	"io"
	"os"
)

// Bucket bucket
type Bucket struct {
	Name string
	conn *Conn
}

// ObjectSlice object slice
type ObjectSlice struct {
	UploadID string
	Size     int64
	Offset   int64
	Number   int
	MD5      string
	src      string
}

// UploadObject 上传文件
func (b *Bucket) UploadObject(ctx context.Context, object string, content io.Reader, acl *AccessControl) error {
	_, err := b.conn.Do(ctx, "PUT", b.Name, object, nil, acl.GenHead(), content)

	return err
}

// CopyObject 复制对象
func (b *Bucket) CopyObject(ctx context.Context, src, dst string, acl *AccessControl) error {
	srcURL := fmt.Sprintf("%s-%s.cos.%s.%s/%s", b.Name, b.conn.conf.AppID, b.conn.conf.Region, b.conn.conf.Domain, dst)
	header := map[string]string{
		"x-cos-source-url": srcURL,
	}

	_, err := b.conn.Do(ctx, "PUT", b.Name, dst, nil, header, nil)

	return err
}

// DeleteObject delete object
func (b *Bucket) DeleteObject(ctx context.Context, obj string) error {
	_, err := b.conn.Do(ctx, "DELETE", b.Name, obj, nil, nil, nil)

	return err
}

// DownloadObject 下载对象
func (b *Bucket) DownloadObject(ctx context.Context, object string, w io.Writer) error {
	res, err := b.conn.Do(ctx, "GET", b.Name, object, nil, nil, nil)
	if err != nil {
		return err
	}

	_, err = io.Copy(w, res.Body)

	return err
}

// UploadObjectBySlice upload by slice
func (b *Bucket) UploadObjectBySlice(ctx context.Context, dst, src string, taskNum uint) error {
	uploadID, err := b.InitSliceUpload(ctx, dst)
	if err != nil {
		return err
	}

	fd, err := os.Open(src)
	if err != nil {
		return err
	}

	slices, err := b.PerformSliceUplaod(ctx, dst, uploadID, fd)
	if err != nil {
		return err
	}

	err = b.CompleteSliceUpload(ctx, dst, uploadID, fd, slices)

	return err
}

// InitSliceUpload init upload by slice
func (b *Bucket) InitSliceUpload(ctx context.Context, obj string) (string, error) {
	param := map[string]interface{}{
		"uploads": "",
	}
	res, err := b.conn.Do(ctx, "PUT", b.Name, obj, param, nil, nil)
	if err != nil {
		return "", err
	}
	imur := &InitiateMultipartUploadResult{}
	err := XMLDecode(res.Body, imur)
	if err != nil {
		return "", err
	}

	return imur.UploadID, nil
}

// CompleteSliceUpload finish slice Upload
func (b *Bucket) CompleteSliceUpload(ctx context.Context, dst, uploadID string, fd *os.File, slice []*ObjectSlice) error {
}
