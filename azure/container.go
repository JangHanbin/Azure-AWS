package azure

import (
	"github.com/Azure/azure-storage-blob-go/azblob"
	"log"
)

// Create New container
func CreateContainer(client *Client, containerName string) {

	containerURL := client.ServiceURL.NewContainerURL(containerName) // Container names require lowercase
	_, err := containerURL.Create(client.Context, azblob.Metadata{}, azblob.PublicAccessBlob)
	if err != nil {
		println("CREATE ERROR")
		log.Fatal(err)

	}
}

func GetContainers(client *Client) {

}
