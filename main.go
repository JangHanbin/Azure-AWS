package main

import "mini-contents-hub/aws"

func main() {

	clientAWS := aws.GetClient() //AWS client uses ~/.aws/config file. maybe managed same config file

	buckets := aws.GetBuckets(clientAWS) // Get bucket list as a slice type
	//
	for _, bucket := range buckets {
		println(bucket) // print bucket line by line

	}
	//objects := aws.GetObjects(clientAWS, "unique-bucket-name")
	//for _, object := range objects { //read object from bucket
	//	println(object) //print object line by line
	//}

	//aws.CreateBucket(clientAWS, "unique-bucket-name", "ap-northeast-2")
	//aws.UploadObject(clientAWS, "unique-bucket-name", "file-name.txt", []byte("Hello Binary!"))
	//data := aws.DownloadObject(clientAWS, "mywavvebucket", "drm.js")

	//println(aws.GetPublicURL(clientAWS, "mywavvebucket", "drm.js")
	//aws.DeleteBucket(clientAWS, "testbucket4881235")
	//aws.DeleteObject(clientAWS, "mywavvebucket", "drm.js")

	//clientAZURE := azure.GetClient("./configs/config.json") //TODO : replace hard coded path to execution parameter
	//
	//containers := azure.GetContainers(clientAZURE)
	//
	//for _, container := range containers {
	//	println(container) // print container line by line
	//	blobs := azure.GetBlobs(clientAZURE, container)
	//
	//	for _, blob := range blobs {
	//		println(blob) // print blob line by line
	//	}
	//}

}
