package azure

import (
	"github.com/Azure/azure-storage-blob-go/azblob"
	"log"
	"strings"
)

func UploadBlob(client *Client, containerName string, blobName string, data string) { //TODO: DATA TYPE MUST BE BINARY TYPE
	containerURL := client.ServiceURL.NewContainerURL(containerName)
	blobURL := containerURL.NewBlockBlobURL(blobName)

	_, err := blobURL.Upload(client.Context, strings.NewReader(data), azblob.BlobHTTPHeaders{ContentType: "text/plain"}, azblob.Metadata{}, azblob.BlobAccessConditions{}, azblob.DefaultAccessTier, nil, azblob.ClientProvidedKeyOptions{})
	if err != nil {
		log.Fatal(err)
	}

}
