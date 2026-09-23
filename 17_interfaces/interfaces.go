package main

import "fmt"

type Uploader interface {
	upload(path string)
}

type File struct {
	uploader Uploader
}

type S3Uploader struct{}

func (s3 S3Uploader) uploadFile(path string) {
	fmt.Println("Uploading to S3.", path)
}

type CloudinaryUploader struct{}

func (cloudinary CloudinaryUploader) upload(path string) {
	fmt.Println("Uploading to Cloudinary.", path)
}

type GCPUploader struct{}

func (gcp GCPUploader) upload(path string) {
	fmt.Println("Uploading to GCP.", path)
}

func main() {
	// S3Bucket := S3Uploader{}
	// cloudinary := CloudinaryUploader{}
	gcp := GCPUploader{}
	myFile := File{uploader: gcp}

	myFile.uploader.upload("sandeep.png")
}
