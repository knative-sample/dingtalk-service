package dingding

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

//sigma agility群机器人url
const DINGDING_OPENAPI_URL = "https://oapi.dingtalk.com/robot/send?access_token=189a15a27d538787af15575fca02d9a4108597ec8a505dcd92d3c3696501134a"

//异常信息监控群
const DINGDING_FOR_EXCEPTION_URL = "https://oapi.dingtalk.com/robot/send?access_token=38b52342dee9149d64a1b036d7ef3b8b9dcb9d518a1d33dc8741d514cc530bb4"

//测试群url
//const DINGDING_OPENAPI_URL = "https://oapi.dingtalk.com/robot/send?access_token=08b69305f92c298e77c7e0b727bf4b697a8a9fcf72f92aefff2658cc4a96cbf5"

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
	fmt.Println("11:" + string(data))
	return string(data)
}

func SendDingDingReqest(url, method, requestBody string) (body []byte, statusCode int, err error) {
	fmt.Println(requestBody)
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
