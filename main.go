package main

import (
	"cloud-cli/cli"
	"cloud-cli/models"
	"cloud-cli/resources"
)

func main() {
	// ایجاد چند منبع نمونه
	resources.AddVMResource(1, map[models.ResourceSpecs]string{"CPU": "4", "Memory": "16GB"}, models.Stopped)
	resources.AddDBResource(101, map[models.ResourceSpecs]string{"Storage": "100GB"}, models.Stopped)
	resources.AddStorageResource(201, map[models.ResourceSpecs]string{"Capacity": "500GB"}, models.Stopped)

	// اجرای CLI
	cli.RunCLI()
}
