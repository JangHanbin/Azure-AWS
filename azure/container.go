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

func DeleteContainer(client *Client, containerName string) {

	containerURL := client.ServiceURL.NewContainerURL(containerName) // Container names require lowercase
	_, err := containerURL.Delete(client.Context, azblob.ContainerAccessConditions{})
	if err != nil {
		log.Fatal(err)
	}

}

func GetContainers(client *Client) (containers []string) {

	for marker := (azblob.Marker{}); marker.NotDone(); { // The parens around Marker{} are required to avoid compiler error.
		// Get a result segment starting with the blob indicated by the current Marker.
		listContainer, err := client.ServiceURL.ListContainersSegment(client.Context, marker, azblob.ListContainersSegmentOptions{})
		if err != nil {
			log.Fatal(err)
		}
		marker = listContainer.NextMarker

		// Process the blobs returned in this result segment (if the segment is empty, the loop body won't execute)
		for _, containerInfo := range listContainer.ContainerItems {
			containers = append(containers, containerInfo.Name)
		}
	}

	return containers
}
