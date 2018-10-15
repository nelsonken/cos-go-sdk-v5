package services

import (
	"context"
	"io"
	"log"
	"os"
	"time"

	"github.com/fpay/cos-go-sdk-v5/cos"
)

type UploadService struct {
	client *cos.Client
	bucket string
}

func NewUploadService(client *cos.Client, bucket string) (*UploadService, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := client.BucketExists(ctx, bucket); err != nil {

		err := client.CreateBucket(ctx, bucket, nil)
		if err != nil {
			return nil, err
		}
	}
	return &UploadService{
		client: client,
		bucket: bucket,
	}, nil
}

func (us *UploadService) Upload(src string, dst string) error {
	log.Println("uploading...")

	fi, err := os.Stat(src)
	if err != nil {
		return err
	}

	if fi.Size() < us.client.PartSize() {
		r, err := os.Open(src)
		if err != nil {
			return err
		}

		return us.upload(r, dst)
	}

	return us.multipartUpload(src, dst)
}

func (us *UploadService) upload(src io.Reader, dst string) error {
	return us.client.Bucket(us.bucket).UploadObject(context.Background(), dst, src, nil)
}

func (us *UploadService) multipartUpload(src string, dst string) error {
	return us.client.Bucket(us.bucket).UploadObjectBySlice(context.Background(), dst, src, 2, nil)
}
