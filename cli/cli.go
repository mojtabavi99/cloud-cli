package cli

import (
	"bufio"
	"cloud-cli/models"
	"cloud-cli/repository"
	"cloud-cli/services"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func RunCLI() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Welcome to Cloud CLI. Type 'exit' to quit.")
	fmt.Println("If you need help type 'help' too see commands")

	repo := repository.NewResourceRepository()

	// ساخت سرویس‌ها با استفاده از repository
	serverService := services.NewServerService(repo)
	dbService := services.NewDatabaseService(repo)
	storageService := services.NewStorageService(repo)

	for {
		fmt.Print("> ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input == "exit" {
			break
		}

		args := strings.Fields(input)
		if len(args) == 0 {
			continue
		}

		cmd := args[0]

		switch cmd {
		case "create":
			if len(args) < 2 {
				fmt.Println("Usage: create <service> <spec1=value1> ...")
				continue
			}
			service := args[1]
			specs, err := parseSpecs(args[2:])
			if err != nil {
				fmt.Println("Error:", err)
				continue
			}
			handleCreateCommand(service, repo, specs)

		case "list":
			handleListCommand(repo)

		case "filter":
			if len(args) < 2 {
				fmt.Println("Usage: filter [type=<vm|db|storage>] [status=<Running|Stopped|Terminated>]")
				continue
			}
			handleFilterCommand(repo, args[1:])

		case "help":
			handelHelpCommand()

		case "start", "stop", "terminate", "restart", "status":
			if len(args) != 3 {
				fmt.Println("Usage: <command> <id> <service>")
				continue
			}
			id, err := strconv.Atoi(args[1])
			if err != nil {
				fmt.Println("Invalid ID")
				continue
			}
			service := args[2]
			executeCommand(cmd, id, service, serverService, dbService, storageService)

		default:
			fmt.Println("Unknown command. Valid commands: create, list, filter, start, stop, terminate, restart, status")
		}
	}

}

func executeCommand(cmd string, id int, service string, server *services.ServerService, db *services.DatabaseService, storage *services.StorageService) {
	var err error
	switch service {
	case "server":
		err = handleServerCommand(cmd, id, server)
	case "db":
		err = handleDBCommand(cmd, id, db)
	case "storage":
		err = handleStorageCommand(cmd, id, storage)
	default:
		fmt.Println("Unknown service. Use 'server', 'db', or 'storage'")
		return
	}

	if err != nil {
		fmt.Println("Error:", err)
	}
}

// ====================== List ======================
func handleListCommand(repo *repository.ResourceRepository) {
	resources := repo.GetAllResources()
	if len(resources) == 0 {
		fmt.Println("No resources found.")
		return
	}

	for _, r := range resources {
		fmt.Printf("ID: %d | Type: %s | Status: %s | Specs: %v\n",
			r.ID, r.Type, r.Status, r.Specs)
	}
}

// ====================== Filter Services ======================
func handleFilterCommand(repo *repository.ResourceRepository, args []string) {
	var filterType *models.ResourceType
	var filterStatus *models.ResourceStatus

	// پردازش ورودی‌ها
	for _, arg := range args {
		parts := strings.SplitN(arg, "=", 2)
		if len(parts) != 2 {
			fmt.Printf("Invalid argument format: %s\n", arg)
			return
		}
		key := strings.ToLower(parts[0])
		value := parts[1]

		switch key {
		case "type":
			switch value {
			case string(models.Server):
				t := models.Server
				filterType = &t
			case string(models.Database):
				t := models.Database
				filterType = &t
			case string(models.Storage):
				t := models.Storage
				filterType = &t
			default:
				fmt.Println("Invalid type. Valid types: Server, Database, Storage")
				return
			}

		case "status":
			switch value {
			case string(models.Running):
				s := models.Running
				filterStatus = &s
			case string(models.Stopped):
				s := models.Stopped
				filterStatus = &s
			case string(models.Terminated):
				s := models.Terminated
				filterStatus = &s
			default:
				fmt.Println("Invalid status. Valid statuses: Running, Stopped, Terminated")
				return
			}

		default:
			fmt.Printf("Unknown filter key: %s\n", key)
			return
		}
	}

	// فیلتر کردن منابع
	resources := repo.GetAllResources()
	var filtered []*models.Resource
	for _, r := range resources {
		if filterType != nil && r.Type != *filterType {
			continue
		}
		if filterStatus != nil && r.Status != *filterStatus {
			continue
		}
		filtered = append(filtered, r)
	}

	// نمایش منابع فیلتر شده
	if len(filtered) == 0 {
		fmt.Println("No resources found with the given filters.")
		return
	}

	for _, r := range filtered {
		fmt.Printf("ID: %d | Type: %s | Status: %s | Specs: %v\n",
			r.ID, r.Type, r.Status, r.Specs)
	}
}

// ====================== Create ======================
func handleCreateCommand(service string, repo *repository.ResourceRepository, specs map[models.ResourceSpecs]string) {
	var rType models.ResourceType
	switch service {
	case "server":
		rType = models.Server
	case "db":
		rType = models.Database
	case "storage":
		rType = models.Storage
	default:
		fmt.Println("Unknown service type. Valid types: server, db, storage")
		return
	}

	resource := &models.Resource{
		ID:     repo.NextID(),
		Type:   rType,
		Status: models.Stopped,
		Specs:  specs,
	}

	if err := repo.AddResource(resource); err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("%s created successfully with ID %d ✅\n", rType, resource.ID)
}

// ====================== Server ======================
func handleServerCommand(cmd string, id int, svc *services.ServerService) error {
	switch cmd {
	case "start":
		return svc.Start(id)
	case "stop":
		return svc.Stop(id)
	case "terminate":
		return svc.Terminate(id)
	case "restart":
		return svc.Restart(id)
	case "status":
		return svc.StatusCheck(id)
	default:
		return fmt.Errorf("unknown command")
	}
}

// ====================== DB ======================
func handleDBCommand(cmd string, id int, svc *services.DatabaseService) error {
	switch cmd {
	case "start":
		return svc.Start(id)
	case "stop":
		return svc.Stop(id)
	case "terminate":
		return svc.Terminate(id)
	case "restart":
		return svc.Restart(id)
	case "status":
		return svc.StatusCheck(id)
	default:
		return fmt.Errorf("unknown command")
	}
}

// ====================== Storage ======================
func handleStorageCommand(cmd string, id int, svc *services.StorageService) error {
	switch cmd {
	case "start":
		return svc.Start(id)
	case "stop":
		return svc.Stop(id)
	case "terminate":
		return svc.Terminate(id)
	case "restart":
		return svc.Restart(id)
	case "status":
		return svc.StatusCheck(id)
	default:
		return fmt.Errorf("unknown command")
	}
}

// ====================== Storage ======================
func handelHelpCommand() {
	fmt.Println("Available commands:")
	fmt.Println("create <service> <spec1=value1> ...                   - Create a new resource (service: server/db/storage)")
	fmt.Println("list                                                  - List all resources")
	fmt.Println("filter [type=<Server|Database|Storage>] [status=<Running|Stopped|Terminated>]  - Filter resources by type and/or status")
	fmt.Println("start <id> <service>                                 - Start a resource")
	fmt.Println("stop <id> <service>                                  - Stop a resource")
	fmt.Println("terminate <id> <service>                             - Terminate a resource")
	fmt.Println("restart <id> <service>                               - Restart a resource")
	fmt.Println("status <id> <service>                                - Show status of a resource")
	fmt.Println("help                                                  - Show this help message")
	fmt.Println("exit                                                  - Exit the CLI")
}

// ====================== Parse Specs ======================
func parseSpecs(specArgs []string) (map[models.ResourceSpecs]string, error) {
	specs := make(map[models.ResourceSpecs]string)
	for _, s := range specArgs {
		parts := strings.SplitN(s, "=", 2)
		if len(parts) != 2 {
			return nil, fmt.Errorf("invalid spec format: %s", s)
		}
		key := parts[0]
		value := parts[1]

		if value == "" {
			return nil, fmt.Errorf("spec value cannot be empty for key: %s", key)
		}

		var validKey models.ResourceSpecs
		switch key {
		case string(models.CPU):
			validKey = models.CPU
		case string(models.Memory):
			validKey = models.Memory
		case string(models.Stotage):
			validKey = models.Stotage
		default:
			return nil, fmt.Errorf("invalid spec key: %s", key)
		}

		specs[validKey] = value
	}
	return specs, nil
}
