package azure

import (
	"bytes"
	"github.com/Azure/azure-storage-blob-go/azblob"
	"io/ioutil"
	"log"
)

func UploadBlob(client *Client, containerName string, blobName string, data []byte) {
	containerURL := client.ServiceURL.NewContainerURL(containerName)
	blobURL := containerURL.NewBlockBlobURL(blobName)

	_, err := blobURL.Upload(client.Context, bytes.NewReader(data), azblob.BlobHTTPHeaders{ContentType: "text/plain"}, azblob.Metadata{}, azblob.BlobAccessConditions{}, azblob.DefaultAccessTier, nil, azblob.ClientProvidedKeyOptions{})
	if err != nil {
		log.Fatal(err)
	}

}

func DownloadBlob(client *Client, containerName string, blobName string) (downloadedData []byte) {
	containerURL := client.ServiceURL.NewContainerURL(containerName)
	blobURL := containerURL.NewBlockBlobURL(blobName)

	get, err := blobURL.Download(client.Context, 0, 0, azblob.BlobAccessConditions{}, false, azblob.ClientProvidedKeyOptions{})
	if err != nil {
		log.Fatal(err)
	}

	//downloadedData = &bytes.Buffer{}
	reader := get.Body(azblob.RetryReaderOptions{})
	downloadedData, err = ioutil.ReadAll(reader)
	reader.Close() // The client must close
	if err != nil {
		log.Fatal(err)
	}

	return downloadedData

}
