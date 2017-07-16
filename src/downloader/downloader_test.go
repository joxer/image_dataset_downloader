package downloader

import (
	"os"
	"testing"
)

func TestDownloadWork(t *testing.T) {

	d := Downloader{api_key: os.Getenv("api_key"), secret_key: os.Getenv("secret_key"), prefix: os.Getenv("prefix")};
	d.getImage("Italia", 10)
}
