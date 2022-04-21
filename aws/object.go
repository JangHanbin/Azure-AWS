package aws

import (
	"bytes"
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	v4 "github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type S3PresignGetObjectAPI interface {
	PresignGetObject(
		ctx context.Context,
		params *s3.GetObjectInput,
		optFns ...func(*s3.PresignOptions)) (*v4.PresignedHTTPRequest, error)
}

func GetObjects(client *s3.Client, bucketName string) (objects []string) {
	// Get the first page of results for ListObjectsV2 for a bucket
	output, err := client.ListObjectsV2(context.TODO(), &s3.ListObjectsV2Input{
		Bucket: aws.String(bucketName),
	})
	if err != nil {
		log.Fatal(err)
	}

	for _, object := range output.Contents {
		objects = append(objects, aws.ToString(object.Key))
	}

	return objects
}

func DownloadObject(client *s3.Client, bucketName string, path string, key string) error {
	// Create the directories in the path
	splitKeyArr := strings.Split(key, "/")
	file := filepath.Join(path, splitKeyArr[len(splitKeyArr)-1])
	if err := os.MkdirAll(filepath.Dir(file), 0775); err != nil {
		return err
	}

	// Set up the local file
	fd, err := os.Create(file)
	if err != nil {
		return err
	}

	defer fd.Close()

	downloader := manager.NewDownloader(client)
	_, err = downloader.Download(context.TODO(), fd,
		&s3.GetObjectInput{
			Bucket: &bucketName,
			Key:    &key,
		})
	return err

}
func UploadObject(client *s3.Client, bucketName string, fileName string, data []byte) *manager.UploadOutput {

	//file, err := ioutil.ReadFile(fileName)
	uploader := manager.NewUploader(client)
	result, err := uploader.Upload(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(fileName),
		Body:   bytes.NewReader(data),
	})
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	return result

}

func GetPresignedURL(c context.Context, api S3PresignGetObjectAPI, input *s3.GetObjectInput) (*v4.PresignedHTTPRequest, error) {
	return api.PresignGetObject(c, input)
}

func GetPublicURL(client *s3.Client, bucketName *string, key *string) (publicURL string) {
	input := &s3.GetObjectInput{
		Bucket: bucketName,
		Key:    key,
	}

	psClient := s3.NewPresignClient(client)

	resp, err := GetPresignedURL(context.TODO(), psClient, input)
	if err != nil {
		fmt.Println("Got an error retrieving pre-signed object:")
		fmt.Println(err)
		return
	}

	return resp.URL
}
