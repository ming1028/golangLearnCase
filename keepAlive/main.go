package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/spf13/cast"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

var template_send_url = "https://api.weixin.qq.com/cgi-bin/message/template/send?access_token=%s"
var accToken = "66_tc4P5004CZTXC-Lqx1oguxmaBvCoKrE60XBzBRAoC8wpG3_Gzd9yMcbuJj9TICct5XAIsmduEiGZ-n124nHbIAxIMWe69drCoqiP0405wzYpf5-GOTkAkK_zGQQVHUgACALXK"
var str = `{"wx_send_req":{"template_id":"UZPRKusPeOoJ9oQ0LJC5_xqifKpSOkNKv2SAnkP1l8I","url":"http://m.qihuo18.com/smp/article/stock?appid=wxa6d9a782033a426f&type=MjUx&tid=MTk%3D&aid=MzA%3D","data":{"first":{"value":"2342","color":"#173177"},"remark":{"value":"234","color":"#173177"},"keyword1":{"value":"234","color":"#173177"},"keyword2":{"value":"234","color":"#173177"},"keyword3":{"value":"2","color":"#173177"},"keyword4":{"value":"2","color":"#173177"},"keyword5":{"value":"2","color":"#173177"}}},"extra_info":{"app_id":4,"template_id":19,"tag_id":[7],"corp_id":"wwc5c5dbbc5c2ccf25","agent_id":1000033}}`
var openId1 = "oOLpN59oaX9Wni72l_YDazmqy0tE"
var openId2 = ""
var url = "http://localhost:8080"
var chReq chan *http.Request
var httpRequest *http.Request
var hreqOnce sync.Once

func main() {
	client := http.Client{}
	// apiUrl := fmt.Sprintf(template_send_url, accToken)
	for i := 1; i < 13; i++ {
		j := i
		go func(i int) {
			s := time.Now()
			req := GetRequest()
			wxReq := new(ServiceSendReq)
			_ = json.Unmarshal([]byte(str), wxReq)
			wxReq.WxSendReq.Touser = openId1
			tmp := new(TempMsgData)
			tmp.Value = cast.ToString(i)
			tmp.Color = "#173177"
			wxReq.WxSendReq.Data["first"] = *tmp
			js, _ := json.Marshal(wxReq.WxSendReq)
			body := bytes.NewBuffer(js)
			req.Body = ioutil.NopCloser(body)
			req.ContentLength = int64(len(js))
			resp, err := client.Do(req)
			if err != nil {
				fmt.Println(err)
			}
			b, err := ioutil.ReadAll(resp.Body)
			resp.Body.Close()
			d := time.Since(s)
			fmt.Println(d, string(b), err)
			PushRequest(req)
		}(j)
	}

	time.Sleep(time.Hour)
}

type WxSendReq struct {
	TemplateId string                 `json:"template_id"`
	Touser     string                 `json:"touser"`
	Url        string                 `json:"url"`
	Data       map[string]TempMsgData `json:"data"`
}

type TempMsgData struct {
	Value string `json:"value"`
	Color string `json:"color"`
}

type ServiceSendReq struct {
	WxSendReq WxSendReq `json:"wx_send_req"`
}

func init() {
	chReq = make(chan *http.Request, 2)
	for i := 0; i < 2; i++ {
		req, err := http.NewRequest("POST", url, nil)
		req.Header.Set("Content-Type", "application/json")
		if err != nil {
			fmt.Println("newRequestError", err)
			return
		}
		chReq <- req
	}
}

func GetRequest() *http.Request {
	select {
	case req := <-chReq:
		return req
	}
}

func PushRequest(req *http.Request) {
	chReq <- req
}
