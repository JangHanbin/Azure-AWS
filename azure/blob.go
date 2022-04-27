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

func DeleteBlob(client *Client, containerName string, blobName string) {
	containerURL := client.ServiceURL.NewContainerURL(containerName)
	blobURL := containerURL.NewBlockBlobURL(blobName)
	_, err := blobURL.Delete(client.Context, azblob.DeleteSnapshotsOptionNone, azblob.BlobAccessConditions{})
	if err != nil {
		log.Fatal(err)
	}
}

func GetBlobs(client *Client, containerName string) (blobs []string) {
	containerURL := client.ServiceURL.NewContainerURL(containerName)
	for marker := (azblob.Marker{}); marker.NotDone(); { // The parens around Marker{} are required to avoid compiler error.
		// Get a result segment starting with the blob indicated by the current Marker.
		listBlob, err := containerURL.ListBlobsFlatSegment(client.Context, marker, azblob.ListBlobsSegmentOptions{})
		if err != nil {
			log.Fatal(err)
		}
		marker = listBlob.NextMarker
		// Process the blobs returned in this result segment (if the segment is empty, the loop body won't execute)
		for _, blobInfo := range listBlob.Segment.BlobItems {
			blobs = append(blobs, blobInfo.Name)
		}
	}

	return blobs
}
