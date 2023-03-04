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

package controllers

import (
	"github.com/nvwa-io/nvwa-io/nvwa-server/libs/logger"
	. "github.com/nvwa-io/nvwa-io/nvwa-server/svrs"
)

type SystemController struct {
	BaseAuthController
}

// @Title Get system config
// @router / [get]
func (t *SystemController) Get() {
	sys := DefaultSystemSvr.Get()

	logger.Debugf("查询系统配置", sys)
	t.SuccJson(RespData{
		"system": sys,
	})
}
