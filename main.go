package main

import (
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"mini-contents-hub/aws"
	"mini-contents-hub/azure"
)

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func getDiff(containers []string, buckets []string) (map[string]bool, map[string]bool) {

	var az map[string]bool
	var s3 map[string]bool

	az = make(map[string]bool)
	s3 = make(map[string]bool)

	for _, container := range containers {
		az[container] = contains(buckets, container)
	}

	for _, bucket := range buckets { // TODO : avoid searching that already checked in previous for loop
		s3[bucket] = contains(containers, bucket)
	}

	return az, s3
}
func sync(clientAzure *azure.Client, clientAWS *s3.Client, containers []string, buckets []string) {

	// Sync Azure Containers and AWS S3 Buckets
	az, s3 := getDiff(containers, buckets)

	for container, isContain := range az {
		if !isContain { //if not in AWS s3
			println("Create " + container + " bucket in AWS S3")
			aws.CreateBucket(clientAWS, container, "ap-northeast-2")
		}
	}
	for bucket, isContain := range s3 {
		if !isContain { //if not in Azure
			azure.CreateContainer(clientAzure, bucket)
			println("Create " + bucket + " container in Azure")
		}
	}

	//Must be synchronized Containers and Buckets before the Blob Object syncing
	//Sync Azure blob and AWS Object
	for _, name := range buckets {
		objects := aws.GetObjects(clientAWS, name)
		blobs := azure.GetBlobs(clientAzure, name)
		az, s3 = getDiff(blobs, objects)

		for blob, isContain := range az {
			if !isContain { //if not in AWS s3
				println("Create " + blob + " object in AWS S3 " + name)
				aws.UploadObject(clientAWS, name, blob, azure.DownloadBlob(clientAzure, name, blob))
			}
		}
		for object, isContain := range s3 {
			if !isContain { //if not in Azure
				println("Create " + object + " blob in Azure " + name)
				azure.UploadBlob(clientAzure, name, object, aws.DownloadObject(clientAWS, name, object))

			}
		}
	}
}

func main() {

	clientAWS := aws.GetClient() //AWS client uses ~/.aws/config file. maybe managed same config file

	buckets := aws.GetBuckets(clientAWS) // Get bucket list as a slice type

	clientAZURE := azure.GetClient("./configs/config.json") //TODO : replace hard coded path to execution parameter
	containers := azure.GetContainers(clientAZURE)

	sync(clientAZURE, clientAWS, containers, buckets)

}
