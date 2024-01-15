package gitlab

import (
	"sre-dashboard/pkg/db"
	"sre-dashboard/pkg/global"
	"sre-dashboard/pkg/logging"
	"sync"
)

func SaveMergeDataToDB() {
	_, conf := GitlabConn()

	//logging.CLILog.Info(conf.MRProjectId)
	for _, projectId := range conf.MRProjectId {
		mrData, _, mrErr := GetMergeReqData(projectId)
		//logging.CLILog.Info(mrData)
		if mrErr != nil {
			logging.CLILog.Error(mrErr)
		}

		for _, mr := range mrData {
			merge_db := db.MergeRequest{
				ID:        mr.ID,
				Title:     mr.Title,
				MergeAt:   mr.MergedAt,
				ProjectID: projectId,
			}
			res := merge_db.IDSearch()

			if !res {
				logging.CLILog.Info("Add merge data:", mr.ID, " | ", mr.Title)
				merge_db.Add()
			} else {
				logging.CLILog.Info("Already exists:", mr.ID, " | ", mr.Title)
			}
		}

	}
}

func SavePPLToDB() {
	conf := global.GitlabConfig

	var wg sync.WaitGroup
	wg.Add(len(conf.PPL))

	for _, v := range conf.PPL {

		go func(id int, name string) {
			defer wg.Done()

			pplData, _, err := GetPipelineData(id)
			if err != nil {
				logging.CLILog.Error(err)
			}
			//logging.CLILog.Info(pplData)

			for _, pp := range pplData {
				//logging.CLILog.Info(pp)
				ppl_data := db.PipelineData{
					ID:       pp.ID,
					Status:   pp.Status,
					Project:  name,
					Duration: pp.Duration,
					UpdataAt: pp.UpdataAt,
				}
				res := ppl_data.IDSearch()
				if !res {
					logging.CLILog.Info("Add ppl data:", ppl_data.ID, " | ", ppl_data.Project, " | ", ppl_data.Status)
					ppl_data.Add()
				} else {
					logging.CLILog.Info("Already exists:", ppl_data.ID, " | ", ppl_data.Project, " | ", ppl_data.Status)
				}
			}
		}(v.ID, v.Name)
	}
	wg.Wait()
	logging.CLILog.Info("Done")
	return
}
