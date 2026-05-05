package worker

import (
	"fmt"
	"mymodule/model"
	"mymodule/service"
	"sync"
)

func Worker(
	id int,
	taskChan <-chan model.Task,
	resultChan chan<- model.Task,
	wg *sync.WaitGroup,
	taskService *service.TaskService,
) {
	defer wg.Done()

	for task := range taskChan {

		fmt.Println("Worker", id, "processing:", task.Name)

		err := taskService.MarkDone(task.ID)
		if err != nil {
			fmt.Println("Error marking done:", err)
			continue
		}

		task.IsDone = true

		resultChan <- task
	}
}
