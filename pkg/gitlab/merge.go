package gitlab

import (
	"fmt"
	"github.com/xanzy/go-gitlab"
)

func GetUserList() {
	git, _ := GitlabConn()

	var Active bool = true

	users, _, err := git.Users.ListUsers(&gitlab.ListUsersOptions{Active: &Active})

	if err != nil {
		panic(err)
	}
	//logging.CLILog.Info(users)
	for _, user := range users {
		fmt.Println(user.Name, user.State)
	}
}

func GetMergeReqData(id int) ([]MergeReq, int, error) {
	git, _ := GitlabConn()

	var (
		state = "all"
	)
	//var data map[string]interface{}
	var gitlabList = make([]MergeReq, 0)
	var temp MergeReq
	var page = 1

	// 循环遍历
	for {
		project, resp, err := git.MergeRequests.ListProjectMergeRequests(id, &gitlab.ListProjectMergeRequestsOptions{
			ListOptions: gitlab.ListOptions{
				PerPage: 100,
				//PerPage: 10,
				Page: page,
			},
			State: &state})

		for _, p := range project {
			temp.ID = p.ID
			temp.Title = p.Title
			temp.MergedAt = p.MergedAt

			gitlabList = append(gitlabList, temp)
		}
		if err != nil {
			return nil, 0, err
		}
		if page >= resp.TotalPages {
			//if page <= resp.TotalPages {
			break
		}
		page = page + 1
	}
	return gitlabList, len(gitlabList), nil
}
