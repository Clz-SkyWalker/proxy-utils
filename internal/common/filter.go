package common

import "strings"

type ResourceFilter struct {
	StaticFilterList []string // 资静态源文件过滤
}

//
//  IsResource
//  @Description: 判断该请求是否为请求资源文件
//  @receiver r
//  @param path url 路径
//  @return bool
//
func (r ResourceFilter) IsResource(path string) bool {
	pathSplit := strings.Split(path, ".")
	if len(pathSplit) == 0 {
		return true
	}
	for _, value := range r.StaticFilterList {
		if value == pathSplit[len(pathSplit)-1] {
			return true
		}
	}
	return false
}
