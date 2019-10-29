package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"

	"os"
	cloudevents "github.com/cloudevents/sdk-go"
	"github.com/knative-sample/dingtalk-service/pkg/dingding"
	"github.com/knative-sample/dingtalk-service/pkg/kncloudevents"
	"encoding/json"
	"strings"
)

/*
Example Output:

☁  cloudevents.Event:
Validation: valid
Context Attributes,
  SpecVersion: 0.2
  Type: dev.knative.eventing.samples.heartbeat
  Source: https://github.com/knative/eventing-sources/cmd/heartbeats/#local/demo
  ID: 3d2b5a1f-10ca-437b-a374-9c49e43c02fb
  Time: 2019-03-14T21:21:29.366002Z
  ContentType: application/json
  Extensions:
    the: 42
    beats: true
    heart: yes
Transport Context,
  URI: /
  Host: localhost:8080
  Method: POST
Data,
  {
    "id":162,
    "label":""
  }
*/
type Info struct {
	Title string `json:"title"`
	Content string `json:"content"`
}
func dispatch(ctx context.Context, event cloudevents.Event) {
	//tctx := cloudevents.HTTPTransportContextFrom(ctx)
	//h, _ := json.Marshal(tctx)
	fmt.Printf(event.String())
	payload := &Info{}
	if event.Data == nil {
		fmt.Printf("cloudevents.Event\n  Type:%s\n  Data is empty", event.Context.GetType())
	}

	data, ok := event.Data.([]byte)
	if !ok {
		var err error
		data, err = json.Marshal(event.Data)
		if err != nil {
			data = []byte(err.Error())
		}
	}
	json.Unmarshal(data, payload)
	ms := make([]string, 0)
	if mobiles != ""{
		ms = strings.Split(mobiles, ",")
	}
	dingding.SendDingDingReqest(url, http.MethodPost, dingding.BuildTextContext(payload.Title, payload.Content, ms));


}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "ok")
}
var (
	url string
	mobiles string
)
func init() {
	flag.StringVar(&url, "dingtalkurl", "", "dingtalk url.")
	flag.StringVar(&url, "mobiles", "", "mobiles.")
}
func main() {
	flag.Parse()


	go func() {
		http.HandleFunc("/health", handler)
		port := os.Getenv("PORT")
		if port == "" {
			port = "8022"
		}
		http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
	}()

	c, err := kncloudevents.NewDefaultClient()
	if err != nil {
		log.Fatal("Failed to create client, ", err)
	}
	log.Fatal(c.StartReceiver(context.Background(), dispatch))
}

