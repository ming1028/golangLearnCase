package main

import (
	"encoding/csv"
	"encoding/json"
	"github.com/go-resty/resty/v2"
	"net/http"
	"os"
)

var (
	Token = "eyJhbGciOiJIUzUxMiJ9.eyJzdWIiOiJZUjE3NDUxNzQzNjQxMDAwMSIsImV4cCI6MTY5ODk3OTk0NCwiaWF0IjoxNjk4Mzc1MTQ0fQ.2rYiIo7myAL8Zx7WMx0lexlKodQzKgl3u_F-5aHTF6FWFzzXObmImbSdZOKjIUFvteGTtZ_uijkc87LuXtNX2g"
	url   = "http://quote.youruitech.com/trade/v1/constituent_list"
)

func main() {
	client := resty.New()
	client.SetTransport(&http.Transport{
		DisableKeepAlives: false,
		MaxConnsPerHost:   5,
	})
	resp, _ := client.R().SetHeaders(map[string]string{
		"Origin":        "http://quote.youruitech.com",
		"Content-Type":  "application/json",
		"Authorization": Token,
	}).SetBody(map[string]string{}).Post(url)

	file, _ := os.Create("constituent.csv")
	defer file.Close()
	file.WriteString("\xEF\xBB\xBF")
	w := csv.NewWriter(file)
	respData := struct {
		Code int `json:"code"`
		Data struct {
			Records [][]interface{} `json:"records"`
			Fields  interface{}     `json:"fileds"`
		}
		Message string `json:"message"`
		Success bool   `json:"success"`
	}{}
	json.Unmarshal(resp.Body(), &respData)
	fileds := respData.Data.Fields.([]interface{})
	w.Write([]string{fileds[0].(string), fileds[1].(string)})

	for _, record := range respData.Data.Records {
		w.Write([]string{record[0].(string), record[1].(string)})
	}
	w.Flush()
}
