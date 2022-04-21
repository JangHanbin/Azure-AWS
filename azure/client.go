package azure

import (
	"context"
	"encoding/json"
	"github.com/Azure/azure-pipeline-go/pipeline"
	"github.com/Azure/azure-storage-blob-go/azblob"
	"log"
	"net/url"
	"os"
)

type Configuration struct {
	ConnectionString   string
	SASToken           string
	BlobServiceSASURL  string
	FileServiceSASURL  string
	QueueServiceSASURL string
	TableServiceSASURL string
}

func getBlobSASURLFromFile(path string) string { //may function name is too long
	// From the file, get your Storage account's name and account key.
	file, fileErr := os.Open(path)
	if fileErr != nil {
		log.Fatalf("Config file open failure: %+v", fileErr)
		panic(fileErr)
	}
	defer file.Close()
	decoder := json.NewDecoder(file)

	configuration := Configuration{}

	err := decoder.Decode(&configuration)
	if err != nil {
		log.Fatalf("Config parsing failure: %+v", err)
		panic(err)
	}

	return configuration.BlobServiceSASURL

}

type Client struct {
	//BlobSASURL string
	Credential azblob.Credential
	PipeLine   pipeline.Pipeline
	Context    context.Context
	ServiceURL azblob.ServiceURL
}

func newClient(path string) *Client {
	client := Client{}
	client.Credential = azblob.NewAnonymousCredential()
	client.PipeLine = azblob.NewPipeline(client.Credential, azblob.PipelineOptions{})
	client.Context = context.Background()

	blobURL := getBlobSASURLFromFile(path) //Set blob URL from config file path
	parsedURL, _ := url.Parse(blobURL)

	client.ServiceURL = azblob.NewServiceURL(*parsedURL, client.PipeLine)

	return &client
}

func GetClient(path string) (client *Client) {
	client = newClient(path)
	return client

}
