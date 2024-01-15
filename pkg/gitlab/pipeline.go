package gitlab

import (
	"github.com/xanzy/go-gitlab"
	"sre-dashboard/pkg/logging"
	"sre-dashboard/pkg/tools"
	"time"
)

func getPPLData(ppid int) ([]int, int, error) {
	git, _ := GitlabConn()

	var pplList = make([]int, 0)
	var page = 1
	var err error
	_, firstDay, _ := tools.GetCurrentQuarterDates()

	t, err := time.Parse("2006-01-02", firstDay)

	for {
		ppl, ppresp, err := git.Pipelines.ListProjectPipelines(ppid, &gitlab.ListProjectPipelinesOptions{
			ListOptions: gitlab.ListOptions{
				PerPage: 100,
				Page:    page,
			},
			UpdatedAfter: &t,
		})
		//
		for _, p := range ppl {
			pplList = append(pplList, p.ID)
		}
		if err != nil {
			return nil, 0, err
		}
		if page >= ppresp.TotalPages {
			break
		}
		page = page + 1
	}
	return pplList, len(pplList), err
}

func GetPipelineData(ppid int) ([]Pipeline, int, error) {
	git, _ := GitlabConn()

	ppl, _, err := getPPLData(ppid)

	if err != nil {
		logging.CLILog.Error(err)
	}
	var pplist = make([]Pipeline, 0)
	for _, pid := range ppl {
		var temp Pipeline

		pl, _, err := git.Pipelines.GetPipeline(ppid, pid)

		if err != nil {
			return nil, 0, err
		}

		temp.ID = pl.ID
		temp.Duration = pl.Duration
		temp.Status = pl.Status
		temp.UpdataAt = pl.UpdatedAt

		pplist = append(pplist, temp)

	}
	return pplist, len(pplist), err

}
