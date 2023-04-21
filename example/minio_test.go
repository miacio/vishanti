package example_test

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/miacio/vishanti/lib"
)

func TestPutObject(t *testing.T) {
	Runner()
	file, err := os.Open("C://Users/10428/Downloads/肖博夷 (1).pdf")
	if err != nil {
		t.Fatal(err)
	}
	fileStat, _ := file.Stat()

	objTmpl := "files/%s/%d/%s"
	objectName := fmt.Sprintf(objTmpl, "xxxxx", time.Now().UnixMicro(), fileStat.Name())

	err = lib.Minio.PutObject("miajiodb", "miajio", objectName, file, fileStat.Size())
	if err != nil {
		t.Fatal(err)
	}
}

func TestFPutObject(t *testing.T) {
	Runner()
	err := lib.Minio.FPutObject("miajiodb", "miajio", "/miajio/config2.toml", "./config.toml")
	if err != nil {
		t.Fatal(err)
	}
}
