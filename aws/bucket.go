package aws

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

func CreateBucket(client *s3.Client, bucketName string, region types.BucketLocationConstraint) {
	// 버킷 생성하기
	output, err := client.CreateBucket(context.TODO(), &s3.CreateBucketInput{
		Bucket: &bucketName,
		CreateBucketConfiguration: &types.CreateBucketConfiguration{
			LocationConstraint: region,
		},
	})

	if err != nil {
		panic(err)
	}
	fmt.Println(output.Location)
}

func GetBuckets(client *s3.Client) (buckets []string) {
	output, err := client.ListBuckets(context.TODO(), &s3.ListBucketsInput{})
	if err != nil {
		panic(err)
	}
	for _, bucket := range output.Buckets {
		buckets = append(buckets, *bucket.Name)
	}
	return buckets //may can return output but it's type depend on s3 package
}
