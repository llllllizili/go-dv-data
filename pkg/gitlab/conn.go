package gitlab

import (
	"github.com/xanzy/go-gitlab"
	"sre-dashboard/pkg/config"
	"sre-dashboard/pkg/global"
	"sre-dashboard/pkg/logging"
)

func GitlabConn() (*gitlab.Client, *config.GitlabConfig) {
	c := global.GitlabConfig
	url := c.Endpoint + "/api/v4"

	//logging.RuntimeLog.Info(url)

	git, err := gitlab.NewClient(c.Token, gitlab.WithBaseURL(url))

	if err != nil {
		logging.RuntimeLog.Error("Failed to create client: %v", err)
	}
	return git, c
}
