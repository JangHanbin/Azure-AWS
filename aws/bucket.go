package aws

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

func CreateBucket(client *s3.Client, bucketName string, region types.BucketLocationConstraint) *string {

	// create bucket
	output, err := client.CreateBucket(context.TODO(), &s3.CreateBucketInput{
		Bucket: &bucketName,
		CreateBucketConfiguration: &types.CreateBucketConfiguration{
			LocationConstraint: region,
		},
	})

	if err != nil {
		panic(err)
	}

	return output.Location
}

func DeleteBucket(client *s3.Client, bucketName string) {

	// delete bucket
	_, err := client.DeleteBucket(context.TODO(), &s3.DeleteBucketInput{
		Bucket: &bucketName,
	})

	if err != nil {
		panic(err)
	}

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
