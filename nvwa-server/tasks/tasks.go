// Copyright 2019 - now The https://github.com/nvwa-io/nvwa-io Authors
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package tasks

import (
	"github.com/astaxie/beego/toolbox"
	"github.com/nvwa-io/nvwa-io/nvwa-server/tasks/workers"
)

func init() {

	// notify consumer
	builderWorker := workers.BuildWorker{}
	go builderWorker.Deal()

	deployWorker := workers.DeployWorker{}
	go deployWorker.Deal()

	//长连接监听开始执行和执行结束的任务列表
	//consoles.DealRunningTaskList()
	//consoles.DealTaskReturnList()
	/**
	 * 每秒执行
	 */
	/**
	 * 每n秒执行
	 */
	//	AddTask("task_1", "0/5 * * * * *", consoles.AddTaskList)

	// Deal build task every 5 seconds
	//AddTask("build-task", "*/3 * * * * *", workers.DefaultBuildWorker.DealTask)
	toolbox.StartTask()
	//defer toolbox.StopTask()
}

func AddTask(name string, spec string, f toolbox.TaskFunc) {
	toolbox.AddTask(name, toolbox.NewTask(name, spec, f))
}
