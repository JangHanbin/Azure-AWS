package routes

import (
	"github.com/gofiber/fiber/v2"
	"io/ioutil"
	"mini-contents-hub/aws"
	"mini-contents-hub/azure"
)

var (
	clientAZURE = azure.GetClient("./configs/config.json")
	clientAWS   = aws.GetClient()
)

type ResponseHTTP struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

func StorageController(app *fiber.App) {
	route := app.Group("/api/v1")

	route.Get("/:container/:blob", downloadFile)
	route.Post("/:container/:blob", uploadFile)
}

func downloadFile(c *fiber.Ctx) error {
	containerName := c.Params("container")
	blobName := c.Params("blob")

	blob := azure.DownloadBlob(clientAZURE, containerName, blobName)
	return c.Send(blob)
}

func uploadFile(c *fiber.Ctx) error {
	containerName := c.Params("container")
	blob := c.Params("blob")
	formFile, _ := c.FormFile("file")
	file, _ := formFile.Open()
	buf, _ := ioutil.ReadAll(file)

	azure.UploadBlob(clientAZURE, containerName, blob, buf)
	aws.UploadObject(clientAWS, containerName, blob, buf)
	//synchronize.Sync(clientAZURE, clientAWS, containers, buckets)

	return c.SendStatus(200)
}
