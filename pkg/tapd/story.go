package tapd

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"sre-dashboard/pkg/db"
	"sre-dashboard/pkg/global"
	"sre-dashboard/pkg/logging"
	"sre-dashboard/pkg/tools"
	"strings"
)

func GetTapdStory(date string) ([]Story, int, error) {

	c := global.TapdConfig

	targetUrl := c.Endpoint + "/stories"
	username := c.Username
	password := c.Password
	created := date

	// 创建URL参数
	params := url.Values{}
	params.Add("limit", "200")
	params.Add("workspace_id", "61873854")
	params.Add("created", created)

	targetUrl += "?" + params.Encode()

	//get http client
	client := &http.Client{}

	resp, err := http.NewRequest("GET", targetUrl, nil)
	if err != nil {
		logging.CLILog.Error("发起GET请求时出错:", err)
		return nil, 0, nil
	}
	// set Header and Auth
	resp.Header.Set("Content-Type", "application/json")
	resp.SetBasicAuth(username, password)

	response, err := client.Do(resp)
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)

	//创建结构体实例来反序列化json数据
	var request Request
	err = json.Unmarshal(body, &request)
	if err != nil {
		return nil, 0, nil
	}
	//request 中data字段值是interface类型
	//需要类型断言成 map[string]interface 类型的数据，以便获取
	var data map[string]interface{}
	var tapd []Story = make([]Story, 0)
	var temp Story
	for i, _ := range request.Data {
		data = request.Data[i]["Story"].(map[string]interface{})
		temp.ID = data["id"].(string)
		temp.Name = data["name"].(string)
		temp.Owner = data["owner"].(string)
		temp.Status = data["status"].(string)
		temp.Creator = data["creator"].(string)
		temp.Created = data["created"].(string)
		temp.Modified = data["modified"].(string)
		temp.Pm = data["custom_field_five"].(string)
		tapd = append(tapd, temp)
	}
	return tapd, len(tapd), nil

}

func SaveTapdStoryToDB() {
	c := global.TapdConfig

	var storyData []Story
	var tapdNum int
	var tapdErr error

	q, firstDay, lastDay := tools.GetCurrentQuarterDates()

	logging.CLILog.Info("Q", q, ":", firstDay, lastDay)
	if c.Created != "" {
		storyData, tapdNum, tapdErr = GetTapdStory(c.Created)
	} else {
		storyData, tapdNum, tapdErr = GetTapdStory(firstDay + "~" + lastDay)
	}

	logging.CLILog.Info("tapd story num is : ", tapdNum)
	if tapdErr != nil {
		logging.CLILog.Error(tapdErr)

	}
	for _, v := range storyData {
		Owner := strings.TrimRight(v.Owner, ";")
		if strings.Contains(Owner, ";") {
			splitStr := strings.Split(Owner, ";")
			for _, s := range splitStr {
				story_db := db.Story{
					ID:       v.ID,
					Name:     v.Name,
					Owner:    s,
					Status:   v.Status,
					Creator:  v.Creator,
					Created:  v.Created,
					Modified: v.Modified,
					Pm:       v.Pm,
				}
				//story_db.Init()
				res := story_db.IDSearch()
				if !res {
					logging.CLILog.Info("Add story data:", story_db.ID, " | ", story_db.Name)
					story_db.Add()
				} else {
					logging.CLILog.Info("Already exists:", story_db.ID, " | ", story_db.Name)
				}
			}
		} else {
			story_db := db.Story{
				ID:       v.ID,
				Name:     v.Name,
				Owner:    Owner,
				Status:   v.Status,
				Creator:  v.Creator,
				Created:  v.Created,
				Modified: v.Modified,
				Pm:       v.Pm,
			}
			//story_db.Init()
			res := story_db.IDSearch()
			if !res {
				logging.CLILog.Info("Add story data:", story_db.ID, " | ", story_db.Name)
				story_db.Add()
			} else {
				logging.CLILog.Info("Already exists:", story_db.ID, " | ", story_db.Name)
			}
		}

	}
}
