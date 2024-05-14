package s3

import (
	"fmt"
	"testing"
	"time"
)

func Test(t *testing.T) {
	fileName := fmt.Sprintf("hello-%d.txt", time.Now().Unix())
	fileData := []byte(fmt.Sprintf("Hello World! %d\n", time.Now().Unix()))

	err := UploadToS3(
		"https://minio-sar.oxs.cz",
		"",
		"filedrop",
		"filedrop",
		"filedrop",
		fileName,
		fileData,
		"text/plain",
	)

	if err != nil {
		t.Errorf("error: %v", err)
	}
}
