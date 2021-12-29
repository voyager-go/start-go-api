package task_server

import (
	"fmt"
	"github.com/RichardKnop/machinery/v1"
	"github.com/RichardKnop/machinery/v1/config"
	"github.com/RichardKnop/machinery/v1/log"
	"github.com/RichardKnop/machinery/v1/tasks"
	task_lists "github.com/voyager-go/start-go-api/schedule/task"
	task_util "github.com/voyager-go/start-go-api/schedule/util"
	"time"
)

// StartServer 任务启动
func StartServer(taskServer *machinery.Server) {
	var (
		sendEmail = tasks.Signature{
			UUID: task_util.JoinUUIDPrefix("send_email"),
			Name: "send_email",
		}
		addInt = tasks.Signature{
			UUID: task_util.JoinUUIDPrefix("add_int"),
			Name: "add",
			Args: []tasks.Arg{
				{
					Type:  "int64",
					Value: 1,
				},
				{
					Type:  "int64",
					Value: 2,
				},
			},
		}
		concatStr = tasks.Signature{
			UUID: task_util.JoinUUIDPrefix("concat"),
			Name: "concat",
			Args: []tasks.Arg{
				{
					Type:  "[]string",
					Value: []string{"Hello ", "world"},
				},
			},
		}
	)

	group, _ := tasks.NewGroup(&sendEmail, &addInt, &concatStr)
	asyncResults, err := taskServer.SendGroup(group, 0)
	log.INFO.Print("send success")
	if err != nil {
		log.INFO.Print("taskServer.SendGroup error: ", err.Error())
		return
	}
	for _, asyncResult := range asyncResults {
		results, err := asyncResult.Get(time.Millisecond * 5)
		if err != nil {
			log.INFO.Print("asyncResult.Get error: ", err.Error())
			return
		}
		for _, result := range results {
			fmt.Println(result.Interface())
		}
	}
}

// InitMachineryServer 初始化Machinery的服务 并注册任务
func InitMachineryServer() (*machinery.Server, error) {
	customerCnf, err := config.NewFromYaml("config_task.yaml", true)
	if err != nil {
		return nil, err
	}
	var cnf = &config.Config{
		Broker:          customerCnf.Broker,
		DefaultQueue:    customerCnf.DefaultQueue,
		ResultBackend:   customerCnf.ResultBackend,
		ResultsExpireIn: customerCnf.ResultsExpireIn,
		Redis:           customerCnf.Redis,
	}
	server, err := machinery.NewServer(cnf)
	if err != nil {
		return nil, err
	}
	err = server.RegisterTasks(task_lists.GetTaskLists())
	return server, err
}
