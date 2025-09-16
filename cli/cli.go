package cli

import (
    "cloud-cli/models"
    "cloud-cli/services"
    "fmt"
)

func Run() {
    services.AddResource(1, models.Server, models.Stopped, map[models.ResourceSpecs]string{models.CPU: "4", models.Memory: "16GB"})
    services.AddResource(2, models.Database, models.Running, map[models.ResourceSpecs]string{models.Stotage: "100GB"})

    fmt.Println("Resources:")
    for _, r := range services.Resources {
        fmt.Println(r.Display())
    }
}
