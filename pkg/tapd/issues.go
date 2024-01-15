package tapd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sre-dashboard/pkg/config"
)

func GetTAPDIssues(c config.TapdConfig) ([]Issues, int, error) {
	url := c.Endpoint + "/bugs?workspace_id=61873854"
	username := c.Username
	password := c.Password

	//get http client
	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println("request tapd api failed", err)
	}
	// set Header and Auth
	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(username, password)

	response, err := client.Do(req)
	defer response.Body.Close()
	//获取response.Body中的数据，返回是[]byte类型的数据
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
	var tapd []Issues = make([]Issues, 0)
	var temp Issues
	for i, _ := range request.Data {
		data = request.Data[i]["Bug"].(map[string]interface{})
		if fmt.Sprintf("%v", data["status"]) != "closed" && data["priority"].(string) == "high" || data["priority"].(string) == "urgent" {
			temp.ID = data["id"].(string)
			temp.Title = data["title"].(string)
			temp.Operator = data["current_owner"].(string)
			temp.Priority = data["priority"].(string)
			temp.Status = data["status"].(string)
			tapd = append(tapd, temp)
		}
	}
	return tapd, len(tapd), nil
}
