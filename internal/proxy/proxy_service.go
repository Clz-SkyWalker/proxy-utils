package proxy

import (
	"io"
	"log"
	"net/http"

	"unipro-proxy/internal/case_recorder"
	"unipro-proxy/internal/common"
	"unipro-proxy/internal/utils/myerr"
)

func NewTcpProxy() *TcpProxy {
	return &TcpProxy{RecorderSwitch: true,
		CaseRecorder: case_recorder.NewCaseRecorder("UniPro"),
		Filter: common.ResourceFilter{StaticFilterList: []string{
			"html", "css", "js", "jpg", "jpeg", "png",
		}}}
}

type TcpProxy struct {
	RecorderSwitch bool                        // case 监听者开关
	CaseRecorder   case_recorder.ICaseRecorder // case 记录器
	Filter         common.ResourceFilter       // 过滤器
}

//
//  recorderSwitch
//  @Description: 切换开关，创建case recorder
//  @receiver p
//  @param isOpen 是否开启记录
//  @param outputPath json 输出路径
//  @param recordName case recorder 名称
//  @return string 最终输出路径
//  @return error
//
func (p *TcpProxy) recorderSwitch(isOpen bool, outputPath, recordName string) (string, *myerr.MyErr) {
	if !isOpen {
		p.RecorderSwitch = false
		return p.CaseRecorder.OutputResult(outputPath)

	}
	p.CaseRecorder = case_recorder.NewCaseRecorder(recordName)
	p.RecorderSwitch = true
	return "", nil
}

func (p *TcpProxy) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	log.Printf("received request %s %s %s \n", req.Method, req.Host, req.RemoteAddr)
	if p.RecorderSwitch && p.CaseRecorder.IsTargetPath(req.RequestURI) {
		err := p.CaseRecorder.AddTaskItem(*req)
		if err != nil {
			log.Fatalf("add task err: %s", err)
		}
	}
	transport := http.DefaultTransport
	outReq := new(http.Request)
	*outReq = *req
	res, err := transport.RoundTrip(outReq)
	if err != nil {
		rw.WriteHeader(http.StatusBadGateway)
		return
	}
	for key, value := range res.Header {
		for _, v := range value {
			rw.Header().Add(key, v)
		}
	}
	rw.WriteHeader(res.StatusCode)
	_, err = io.Copy(rw, res.Body)
	if err != nil {
		log.Fatalf("io copy err: %s", err)
		return
	}
	err = res.Body.Close()
	if err != nil {
		log.Fatalf("io copy err: %s", err)
		return
	}
}
