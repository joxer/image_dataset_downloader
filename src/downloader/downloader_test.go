package downloader

import (
	"os"
	"testing"
)

func TestDownloadWork(t *testing.T) {
	var d Downloader;
	d = instance(os.Getenv("api_key"), os.Getenv("secret_key"), os.Getenv("prefix"))
	d.getImage("Italia", 10)
}
