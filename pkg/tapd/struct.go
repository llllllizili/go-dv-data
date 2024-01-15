package tapd

// 定义Tapd返回结构体
type Issues struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Operator string `json:"operator"`
	Priority string `json:"priority"`
	Status   string `json:"status"`
	Url      string `json:"url"`
}

type Story struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Owner    string `json:"owner"`
	Status   string `json:"status"`
	Creator  string `json:"creator"`
	Created  string `json:"created"`
	Modified string `json:"modified"`
	Pm       string `json:"custom_field_five"` // 责任人、项目经理
}

// 定义反序列化结构体结构
type Request struct {
	Status int    `json:"status"`
	Data   []Bug  `json:"data"`
	Info   string `json:"info"`
}

// 定义序列化Bug字段
type Bug map[string]interface{}
