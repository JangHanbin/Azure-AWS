package main

import (
	"mini-contents-hub/azure"
)

func main() {

	//clientAWS := aws.GetClient() //AWS client uses ~/.aws/config file. maybe managed same config file
	//
	//buckets := aws.GetBuckets(clientAWS) // Get bucket list as a slice type
	//
	//for _, bucket := range buckets {
	//	println(bucket)                                            // print bucket line by line
	//	for _, object := range aws.GetObjects(clientAWS, bucket) { //read object from bucket
	//		println(object) //print object line by line
	//	}
	//}
	//aws.CreateBucket(clientAWS, "codetest88839", "ap-northeast-2")
	//aws.UploadObject(clientAWS, "codetest88839", "stringbinary.txt", []byte("Hello Binary!"))
	clientAZURE := azure.GetClient("./configs/config.json") //TODO : replace hard coded path to execution parameter

	//
	//azure.CreateContainer(clientAZURE, "yiya")
	//azure.UploadBlob(clientAZURE, "yiya", "testdata.txt", []byte("Hello DATA Replaced!"))

	data := azure.DownloadBlob(clientAZURE, "yiya", "testdata.txt")
	println(string(data))

}
