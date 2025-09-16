package cli

import (
	"bufio"
	"cloud-cli/models"
	"cloud-cli/resources"
	"cloud-cli/services"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func RunCLI() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Welcome to Cloud CLI. Type 'exit' to quit.")

	for {
		fmt.Print("> ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input == "exit" {
			break
		}

		args := strings.Split(input, " ")
		if len(args) != 3 {
			fmt.Println("Invalid command. Format: <command> <id> <service>")
			continue
		}

		cmd := args[0]
		id, err := strconv.Atoi(args[1])
		if err != nil {
			fmt.Println("Invalid ID")
			continue
		}
		service := args[2]

		executeCommand(cmd, id, service)
	}
}

func executeCommand(cmd string, id int, service string) {
	switch service {
	case "vm":
		resource, ok := resources.VMMap[id]
		if !ok {
			fmt.Println("VM with this ID not found")
			return
		}
		handleVMCommand(cmd, resource)
	case "db":
		resource, ok := resources.DBMap[id]
		if !ok {
			fmt.Println("Database with this ID not found")
			return
		}
		handleDBCommand(cmd, resource)
	case "storage":
		resource, ok := resources.StorageMap[id]
		if !ok {
			fmt.Println("Storage with this ID not found")
			return
		}
		handleStorageCommand(cmd, resource)
	default:
		fmt.Println("Unknown service. Use 'vm', 'db', or 'storage'")
	}
}

func handleVMCommand(cmd string, r *models.Resource) {
	vmService := &services.VMService{}
	switch cmd {
	case "start":
		vmService.Start(r)
	case "stop":
		vmService.Stop(r)
	case "terminate":
		vmService.Terminate(r)
	case "restart":
		vmService.Restart(r)
	case "status":
		vmService.StatusCheck(r)
	default:
		fmt.Println("Unknown command")
	}
}

func handleDBCommand(cmd string, r *models.Resource) {
	dbService := &services.DatabaseService{}
	switch cmd {
	case "start":
		dbService.Start(r)
	case "stop":
		dbService.Stop(r)
	case "terminate":
		dbService.Terminate(r)
	case "restart":
		dbService.Restart(r)
	case "status":
		dbService.StatusCheck(r)
	default:
		fmt.Println("Unknown command")
	}
}

func handleStorageCommand(cmd string, r *models.Resource) {
	storageService := &services.StorageService{}
	switch cmd {
	case "start":
		storageService.Start(r)
	case "stop":
		storageService.Stop(r)
	case "terminate":
		storageService.Terminate(r)
	case "restart":
		storageService.Restart(r)
	case "status":
		storageService.StatusCheck(r)
	default:
		fmt.Println("Unknown command")
	}
}
