package downloader

import (
	"github.com/toomore/lazyflickrgo/flickr"
	"log"
	"os"
	"io"
	"net/http"
	"strings"
)

type Downloader interface {
	getImage(string, int)
	downloadImage()
}

type FlickrDownloader struct {
	api_key string;
	secret_key string;
	messages chan string;
	prefix string;
}

func instance(api_key string, secret_key string, prefix string) *FlickrDownloader {

	fd := &FlickrDownloader{api_key,secret_key,nil,prefix}
	return fd
}

func (d *FlickrDownloader) getImage(image_type string, limit int) {
	var flickr = flickr.NewFlickr(d.api_key,d.secret_key);

	go d.downloadImage()
	args := make(map[string]string);
	args["text"] = image_type
	args["per_page"] = "100"
	args["safe_search"] = "true"

	pages := flickr.PhotosSearch(args)
	for idx, data := range pages {
		if idx  < limit {
			for _, val := range data.Photos.Photo {

				size := flickr.PhotosGetSizes(val.ID)

				if ( size.Sizes.Candownload == 1) {
					for _, size_real := range size.Sizes.Size {
						if size_real.Label == "Medium" {
							d.getChannel() <- size_real.Source
						}
					}
				}
			}
		}
	}
}

func (d *FlickrDownloader) downloadImage() {
	for  url := range d.getChannel() {
		response, e := http.Get(url)
		if e != nil {
			log.Fatal(e)
		}

		defer response.Body.Close()
		name_parts := strings.Split(url,"/")
		last := name_parts[len(name_parts)-1]

		//open a file for writing
		file, err := os.Create(d.prefix+last)
		if err != nil {
			log.Print("ERROR DOWNLOADING: "+err.Error())
		}

		_, err = io.Copy(file, response.Body)
		if err != nil {
			log.Print("ERROR DOWNLOADING: "+err.Error())
		}
		file.Close()

	}
}

func (d *FlickrDownloader) getChannel() chan string {
	if d.messages == nil {
		d.messages = make(chan string, 100)
	}
	return d.messages
}