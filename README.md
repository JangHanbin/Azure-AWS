# Azure-AWS





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



##### 컨테이너 탐색

```go
containers := azure.GetContainers(clientAZURE)

for _, container := range containers {
		println(container) // print container line by line
}
```



##### 컨테이너 생성

```go
azure.CreateContainer(clientAZURE, "container-name")
```



##### 컨테이너 삭제

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





##### 
