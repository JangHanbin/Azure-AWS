# Azure-AWS

Azure-AWS module for use storages. you can use both storage easiler!



## Getting Started

Azure와 AWS S3에 접근하기 위한 인증 정보 준비

### Prerequisites

AWS S3를 이용하기 위해서는 config 파일과 credentials 파일에 AWS의 [IAM](https://console.aws.amazon.com/iam/home)에 대한 정보를 저장해 놓아야 함 

- ~/.aws/config

  ```ini
  [default]
  region = ap-northeast-2
  output = json
  ```

* ~/.aws/credentials

  ```ini
  [default]
  aws_access_key_id = BASE64
  aws_secret_access_key = BASE64
  ```




Azure를 이용하기 위해서는 config 파일에 [SAS token](https://portal.azure.com/)에 대한 정보를 json 형식으로 저장해 놓아야 함 

- $PROJECTPATH/config.json

  ```json
  {
    "ConnectionString": "URL",
    "SASToken": "URL",
    "BlobServiceSASURL": "URL",
    "FileServiceSASURL": "URL",
    "QueueServiceSASURL": "URL",
    "TableServiceSASURL": "URL"
  }
  ```

  

## Usage

Azure와 AWS S3의 사용 예제

### Azure

##### 클라이언트 생성

```go
package main

import (
	"mini-contents-hub/azure"
)

clientAZURE := azure.GetClient("./configs/config.json") //Azure의 SAS Token을 저장한 json 파일의 경로
```



##### Container 탐색

```go
containers := azure.GetContainers(clientAZURE)

for _, container := range containers {
		println(container) // print container line by line
}
```



##### Container 생성

```go
azure.CreateContainer(clientAZURE, "container-name")
```



##### Container 삭제

```go
azure.DeleteContainer(clientAZURE, "container-name")
```





##### Blob 탐색

```go
blobs := azure.GetBlobs(clientAZURE, "container-name")

for _, blob := range blobs {
		println(blob) // print blob line by line
}
```



##### Blob 업로드 

```go
azure.UploadBlob(clientAZURE, "container-name", "file-name.txt", []byte("Hello DATA Replaced!"))
```



##### Blob 다운로드

```go
data := azure.DownloadBlob(clientAZURE, "container-name", "file-name.txt")
println(string(data))

//실행결과
//Hello DATA Replaced!
```



##### Blob 삭제

```go
azure.DeleteBlob(clientAZURE, "container-name", "file-name.txt")
```



##### 모든 Container와 Blob 출력

```go
package main

import (
	"mini-contents-hub/azure"
)

func main() {
	clientAZURE := azure.GetClient("./configs/config.json")
	for _, container := range containers {
		println(container) // print container line by line
		blobs := azure.GetBlobs(clientAZURE, container)

		for _, blob := range blobs {
			println(blob) // print blob line by line
		}
	}
}
```





### AWS S3

##### 클라이언트 생성

```go
package main

import (
	"mini-contents-hub/aws"
)

clientAWS := aws.GetClient() //AWS client uses ~/.aws/config file. 
```



##### Bucket 탐색

```go
buckets := aws.GetBuckets(clientAWS) // Get bucket list as a slice type

for _, bucket := range buckets {
  println(bucket)// print bucket line by line
}
```



##### Bucket 생성

```go
aws.CreateBucket(clientAWS, "unique-bucket-name", "ap-northeast-2")
```



###### Avaliable regions

```go
const (
	BucketLocationConstraintAfSouth1     BucketLocationConstraint = "af-south-1"
	BucketLocationConstraintApEast1      BucketLocationConstraint = "ap-east-1"
	BucketLocationConstraintApNortheast1 BucketLocationConstraint = "ap-northeast-1"
	BucketLocationConstraintApNortheast2 BucketLocationConstraint = "ap-northeast-2"
	BucketLocationConstraintApNortheast3 BucketLocationConstraint = "ap-northeast-3"
	BucketLocationConstraintApSouth1     BucketLocationConstraint = "ap-south-1"
	BucketLocationConstraintApSoutheast1 BucketLocationConstraint = "ap-southeast-1"
	BucketLocationConstraintApSoutheast2 BucketLocationConstraint = "ap-southeast-2"
	BucketLocationConstraintCaCentral1   BucketLocationConstraint = "ca-central-1"
	BucketLocationConstraintCnNorth1     BucketLocationConstraint = "cn-north-1"
	BucketLocationConstraintCnNorthwest1 BucketLocationConstraint = "cn-northwest-1"
	BucketLocationConstraintEu           BucketLocationConstraint = "EU"
	BucketLocationConstraintEuCentral1   BucketLocationConstraint = "eu-central-1"
	BucketLocationConstraintEuNorth1     BucketLocationConstraint = "eu-north-1"
	BucketLocationConstraintEuSouth1     BucketLocationConstraint = "eu-south-1"
	BucketLocationConstraintEuWest1      BucketLocationConstraint = "eu-west-1"
	BucketLocationConstraintEuWest2      BucketLocationConstraint = "eu-west-2"
	BucketLocationConstraintEuWest3      BucketLocationConstraint = "eu-west-3"
	BucketLocationConstraintMeSouth1     BucketLocationConstraint = "me-south-1"
	BucketLocationConstraintSaEast1      BucketLocationConstraint = "sa-east-1"
	BucketLocationConstraintUsEast2      BucketLocationConstraint = "us-east-2"
	BucketLocationConstraintUsGovEast1   BucketLocationConstraint = "us-gov-east-1"
	BucketLocationConstraintUsGovWest1   BucketLocationConstraint = "us-gov-west-1"
	BucketLocationConstraintUsWest1      BucketLocationConstraint = "us-west-1"
	BucketLocationConstraintUsWest2      BucketLocationConstraint = "us-west-2"
)
```



##### Bucket 삭제

```go
TODO
```





##### Object 탐색

```go
objects := aws.GetObjects(clientAWS, "unique-bucket-name")
for _, object := range objects { //read object from bucket
  println(object) //print object line by line
}
```



##### Object 업로드 

```go
aws.UploadObject(clientAWS, "unique-bucket-name", "file-name.txt", []byte("Hello Binary!"))
```



##### Object 다운로드

```go
data := aws.DownloadObject(clientAWS, "unique-bucket-name", "file-name.txt")

println(string(data))

//실행결과
//Hello Binary!
```



##### Object 삭제

```go
TODO
```



