package case_recorder

import "encoding/json"

func UnmarshalCaseStruct(data []byte) (CaseStruct, error) {
	var r CaseStruct
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *CaseStruct) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type CaseStruct struct {
	By       string           `json:"by"`       // 创建者
	ID       string           `json:"id"`       // 文件编号id
	Name     string           `json:"name"`     // case名称
	Coverage string           `json:"coverage"` // 默认full
	Tasks    []TaskItemStruct `json:"tasks"`    // 任务列表
}

type TaskItemStruct struct {
	ID            int64                       `json:"id"`
	Name          string                      `json:"name"`          // 方法名称
	Operation     string                      `json:"operation"`     // 绑定的方法路径
	Params        map[interface{}]interface{} `json:"params"`        // 参数
	Run           int64                       `json:"run"`           // 运行次数
	Fun           string                      `json:"fun"`           // 执行方法
	ExpectedError string                      `json:"expectedError"` //期待错误
	//Return        Return `json:"return"`
}

type Params struct {
	AppStuID          string `json:"app_stu_id"`
	AppStuMsgType     string `json:"app_stu_msg_type"`
	AppStuMsgPushTime string `json:"app_stu_msg_push_time"`
}
