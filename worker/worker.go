package worker

import (
	"fmt"
	"mymodule/model"
	"sync"
	"time"
)

func Worker(id int, taskChan <-chan model.Task, resultChan chan<- model.Task,wg *sync.WaitGroup) {
defer wg.Done()
	for task := range taskChan {
		if task.Status == "Done" {
			fmt.Println("Task already completed:", task.Name)
			continue
		}
		fmt.Println("Processing this task: ", id, task.Name)
		time.Sleep(2 * time.Second)

		task.Status = "Done"
		resultChan <- task
	}
}
