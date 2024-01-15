package gitlab

import "time"

type MergeReq struct {
	ID       int        `json:"id"`
	Title    string     `json:"title"`
	MergedAt *time.Time `json:"merged_at"`
}

type Pipeline struct {
	ID       int        `json:"id"`
	Status   string     `json:"status"`
	Project  string     `json:"project"`
	Duration int        `json:"duration"`
	UpdataAt *time.Time `json:"updataAt"`
}
