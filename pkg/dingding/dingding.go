package dingding

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

//异常信息监控群
const DINGDING_FOR_EXCEPTION_URL = "https://oapi.dingtalk.com/robot/send?access_token=91b4af274b59a5f913265269f5bdc8eae2c31a4848814fd901d2e8ce13fe8250"


func BuildTextContext(context string) string {
	postContext := make(map[string]interface{})
	postContext["msgtype"] = "text"
	text := make(map[string]string)
	text["content"] = context
	at := make(map[string]interface{})
	at["atMobiles"] = make([]string, 0)
	at["isAtAll"] = false
	postContext["text"] = text
	postContext["at"] = at
	data, _ := json.Marshal(postContext)
	return string(data)
}

func SendDingDingReqest(url, method, requestBody string) (body []byte, statusCode int, err error) {
	client := &http.Client{
		Timeout: 30 * time.Second,
	}
	req, err := http.NewRequest(method, url, strings.NewReader(requestBody))
	if err != nil {
		fmt.Printf("http send request url %s fails -- %v ", url, err)
		return
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("http send request url %s fails -- %v ", url, err)
		return
	}

	defer resp.Body.Close()
	body, err = ioutil.ReadAll(resp.Body)

	statusCode = resp.StatusCode
	//status code not in [200, 300) fail
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		fmt.Printf("response status code %d, error messge: %s", resp.StatusCode, string(body))
		return
	}

	if err != nil {
		fmt.Printf("read the result of get url %s fails, response status code %d -- %v", url, resp.StatusCode, err)
	}

	return
}
