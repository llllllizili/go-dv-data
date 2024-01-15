package main

import (
	"sre-dashboard/pkg/gitlab"
	"sre-dashboard/pkg/logging"
	"sre-dashboard/pkg/tapd"
)

func main() {
	logging.CLILog.Info("============tapd story=================================================")
	tapd.SaveTapdStoryToDB()
	logging.CLILog.Info("============gitlab merge================================================")
	gitlab.SaveMergeDataToDB()
	logging.CLILog.Info("============gitlab ppl=================================================")
	gitlab.SavePPLToDB()
	//ppdata, _, err := gitlab.GetPipelineData(735)
	//logging.CLILog.Info(ppdata)
	//logging.CLILog.Error(err)
	//gitlab.GitlabConfigTest()
}
