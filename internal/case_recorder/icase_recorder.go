package case_recorder

import (
	"net/http"

	"unipro-proxy/internal/utils/myerr"
)

//
//  ICaseRecorder
//  @Description: case recorder
//
type ICaseRecorder interface {
	GetTargetPath() []string                         // 获取目标路径
	IsTargetPath(url string) bool                    // 是否为需要拦截的路径
	GetMapping(url string) string                    // 获取url的方法映射路径
	AddTaskItem(resp http.Request) *myerr.MyErr      // 添加 task item
	OutputResult(path string) (string, *myerr.MyErr) // 输出 case 到 json
}

//
//  NewCaseRecorder
//  @Description: case recorder 生成器
//  @param name
//  @return ICaseRecorder
//
func NewCaseRecorder(name string) ICaseRecorder {
	switch name {
	case "UniPro":
		return NewUniProRecorder(CaseStruct{
			By:       "Angular",
			ID:       "1000",
			Name:     "P_login",
			Coverage: "full",
			Tasks:    make([]TaskItemStruct, 0),
		})
	default:
		return NewUniProRecorder(CaseStruct{
			By:       "Angular",
			ID:       "1000",
			Name:     "P_login",
			Coverage: "full",
			Tasks:    make([]TaskItemStruct, 0),
		})
	}
}
