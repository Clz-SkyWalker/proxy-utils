package case_recorder

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"unipro-proxy/internal/utils/myerr"
)

//
//  UniProRecorder
//  @Description: UniPro 项目的CASE记录者
//
type UniProRecorder struct {
	Case       CaseStruct // 真正的case
	TargetPath []string   // uniPro 针对的路径
	UrlMapping map[string]string
}

func (r *UniProRecorder) GetMapping(url string) string {
	return r.UrlMapping[url]
}

func (r *UniProRecorder) AddTaskItem(req http.Request) *myerr.MyErr {
	urlPath := req.RequestURI
	//taskItem:=TaskItemStruct{}
	// 1.替换掉源链接
	for _, value := range r.TargetPath {
		index := strings.Index(urlPath, value)
		if index == 0 {
			urlPath = strings.Replace(urlPath, value, "", 1)
		}
	}
	// 2.提取出 router 路径与 query 参数
	queryMap := make(map[string]interface{})
	paramList := strings.Split(req.URL.RawQuery, "&")
	for _, param := range paramList {
		paramSplit := strings.SplitN(param, "=", 2)
		if len(paramSplit) == 2 {
			queryMap[paramSplit[0]] = paramSplit[1]
		}
	}

	// 3.提取body参数
	bodyMap := make(map[string]interface{})
	bodyBytes, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Fatalf("body err:%s", err)
	}
	err = json.Unmarshal(bodyBytes, &bodyMap)
	if err != nil {
		log.Fatalf("body unmarshal:%s", err)
	}
	log.Println(bodyBytes)

	// 4.制作 task item

	// 5. 添加到taskList中

	return nil
}

func (r *UniProRecorder) OutputResult(path string) (string, *myerr.MyErr) {
	return "", nil
}

func (r *UniProRecorder) GetTargetPath() []string {
	return r.TargetPath
}

func (r *UniProRecorder) IsTargetPath(url string) bool {
	for _, value := range r.TargetPath {
		index := strings.Index(url, value)
		if index == 0 {
			return true
		}
	}
	return false
}

func NewUniProRecorder(caseValue CaseStruct) *UniProRecorder {
	basePath := "/api/v1"
	return &UniProRecorder{
		Case: caseValue,
		TargetPath: []string{"http://home.unipro.innovsharing.com",
			"https://home.unipro.innovsharing.com",
			"http://ucenter-test.innovsharing.com"},
		UrlMapping: map[string]string{
			basePath + "/login/mobile": "get_user_token",
			basePath + "/user/info":    "get_user_info",
			basePath + "/apps":         "get_apps",
		}}
}
