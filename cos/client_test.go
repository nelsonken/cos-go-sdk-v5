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
