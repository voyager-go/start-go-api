package schedule_worker

import "github.com/RichardKnop/machinery/v1"

// StartWorker 开启worker
func StartWorker(taskServer *machinery.Server) error {
	worker := taskServer.NewWorker("start_go_api_worker", 10)
	if err := worker.Launch(); err != nil {
		return err
	}
	return nil
}
