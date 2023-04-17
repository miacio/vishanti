package example_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/miacio/vishanti/lib"
	"github.com/minio/minio-go/v7"
)

func TestAbc(t *testing.T) {
	Runner()
	found, err := lib.MinioClient.BucketExists(context.Background(), "miajio")
	if err != nil {
		t.Fatal(err)
	}
	if found {
		fmt.Println("buckent found")
	} else {
		if err := lib.MinioClient.MakeBucket(context.Background(), "miajio", minio.MakeBucketOptions{
			Region:        "us-east-1",
			ObjectLocking: true,
		}); err != nil {
			t.Fatal(err)
		}
		fmt.Println("successfully create miajio buckent")
	}
}
