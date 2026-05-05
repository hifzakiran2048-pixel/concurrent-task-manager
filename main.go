package main

import (
	"bufio"
	"fmt"
	"mymodule/database"
	"mymodule/model"
	"mymodule/repository"
	"mymodule/service"
	"mymodule/worker"
	"os"
	"strconv"
	"strings"
	"sync"
)

func main() {

	// ✅ STEP 1: Connect Database
	db := database.ConnectDB()
	defer db.Close()

	// ✅ STEP 2: Setup layers
	repo := repository.NewTaskRepo(db)
	taskService := service.NewTaskService(repo)

	// ✅ STEP 3: Channels + Worker Pool
	taskChan := make(chan model.Task, 10)
	resultChan := make(chan model.Task, 10)

	var wg sync.WaitGroup

	// Start 2 workers
	for i := 1; i <= 2; i++ {
		wg.Add(1)
		go worker.Worker(i, taskChan, resultChan, &wg, taskService)
	}

	// Result listener
	go func() {
		for res := range resultChan {
			fmt.Println("✅ Completed:", res.Name)
		}
	}()

	// ✅ STEP 4: CLI Menu
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\n--- TASK MANAGER ---")
		fmt.Println("1. Add Task")
		fmt.Println("2. List Tasks")
		fmt.Println("3. Mark Done (Worker)")
		fmt.Println("4. Delete Task")
		fmt.Println("5. Exit")
		fmt.Print("Enter choice: ")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		switch input {

		// ➕ Add Task
		case "1":
			fmt.Print("Enter task name: ")
			name, _ := reader.ReadString('\n')
			name = strings.TrimSpace(name)

			err := taskService.AddTask(name)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Task added!")
			}

		// 📋 List Tasks
		case "2":
			tasks, err := taskService.ListTasks()
			if err != nil {
				fmt.Println("Error:", err)
				continue
			}

			for _, t := range tasks {
				fmt.Println(t.ID, "-", t.Name, "-", t.IsDone)
			}

		// ⚙️ Mark Done via Worker
		case "3":
			fmt.Print("Enter task ID: ")
			idStr, _ := reader.ReadString('\n')
			idStr = strings.TrimSpace(idStr)

			id, _ := strconv.Atoi(idStr)

			taskChan <- model.Task{
				ID: id,
			}

		// ❌ Delete
		case "4":
			fmt.Print("Enter task ID: ")
			idStr, _ := reader.ReadString('\n')
			idStr = strings.TrimSpace(idStr)

			id, _ := strconv.Atoi(idStr)

			err := taskService.DeleteTask(id)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Deleted!")
			}

		// 🚪 Exit
		case "5":
			close(taskChan)
			wg.Wait()
			close(resultChan)

			fmt.Println("Goodbye!")
			return

		default:
			fmt.Println("Invalid choice")
		}
	}
}