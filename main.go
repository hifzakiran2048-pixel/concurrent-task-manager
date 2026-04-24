package main

import (
	"bufio"
	"fmt"
	"mymodule/model"
	"mymodule/service"
	"mymodule/worker"
	"os"
	"strconv"
	"strings"
	"sync"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	taskChan := make(chan model.Task, 10)
	resultChan := make(chan model.Task, 10)

	var wg sync.WaitGroup

	for i := 1; i < 2; i++ {
		wg.Add(1)
		go worker.Worker(i, taskChan, resultChan, &wg)
	}

	go func() {
		for result := range resultChan {
			service.MarkDone(result.ID)
			fmt.Println("Task Completed:", result.Name)

		}
	}()

	for {
		fmt.Println("\n---Task Manager---")
		fmt.Println("1. Add Task")
		fmt.Println("2. Show Task")
		fmt.Println("3. Mark Task as Done")
		fmt.Println("4. Update Task")
		fmt.Println("5. Delete Task")
		fmt.Println("6. Exit")
		fmt.Println("Choose Option: ")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		choose, err := strconv.Atoi(input)

		if err != nil {
			fmt.Println("Please enter a valid number!")
			continue
		}

		switch choose {

		case 1:
			fmt.Println("Enter Task: ")
			name, _ := reader.ReadString('\n')
			name = strings.TrimSpace(name)

			task := service.AddTask(name)
			fmt.Println("Task Added: ", task.Name)

		case 2:
			tasks := service.GetTask()

			if len(tasks) == 0 {
				fmt.Println("No tasks yet!")
				continue
			}

			for _, t := range tasks {
				fmt.Println("ID:", t.ID, "|", t.Name, "|", t.Status)
			}

		case 3:
			fmt.Println("Enter Task ID: ")
			idStr, _ := reader.ReadString('\n')
			idStr = strings.TrimSpace(idStr)

			id, _ := strconv.Atoi(idStr)

			task, err := service.GetTaskByID(id)
			if err != nil {
				fmt.Println(err)
				continue
			}
			taskChan <- *task
			fmt.Println("Send to Worker:", task.Name)
		case 4:
			fmt.Print("Enter Task ID to update: ")
			idStr, _ := reader.ReadString('\n')
			idStr = strings.TrimSpace(idStr)

			id, _ := strconv.Atoi(idStr)

			fmt.Print("Enter new task name: ")
			newName, _ := reader.ReadString('\n')
			newName = strings.TrimSpace(newName)

			err := service.UpdateTask(id, newName)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Task updated successfully!")
			}
		case 5:
			fmt.Print("Enter Task ID to delete: ")
			idStr, _ := reader.ReadString('\n')
			idStr = strings.TrimSpace(idStr)

			id, _ := strconv.Atoi(idStr)

			err := service.DeleteTask(id)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Task deleted successfully!")
			}

		case 6:
			fmt.Println("Exiting")
			close(taskChan)
			wg.Wait()
			close(resultChan)
			return
		}
	}

}
