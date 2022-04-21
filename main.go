package main

import (
	"mini-contents-hub/aws"
	"mini-contents-hub/azure"
)

func main() {

	clientAWS := aws.GetClient()
	buckets := aws.GetBuckets(clientAWS)

	for _, bucket := range buckets {
		println(bucket)
		for _, object := range aws.GetObjects(clientAWS, bucket) {
			println(object)
		}
	}

	clientAZURE := azure.GetClient("./configs/config.json") //TODO : replace hard coded path to execution parameter

	azure.CreateContainer(clientAZURE, "yiya")
	azure.UploadBlob(clientAZURE, "heya", "testdata.txt", "Hello DATA")

}
