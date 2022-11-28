package controllers

import (
	"CobraApp/pkg/elasticsearch"
	"fmt"
	"github.com/olivere/elastic/v7"
	"time"
)

type Elasticsearch struct {
}

const mapping = `{
  "mappings": {
    "properties": {
      "user": {
        "type": "keyword"
      },
      "message": {
        "type": "text"
      },
      "image": {
        "type": "keyword"
      },
      "created": {
        "type": "date"
      },
      "tags": {
        "type": "keyword"
      },
      "location": {
        "type": "geo_point"
      },
      "suggest_field": {
        "type": "completion"
      }
    }
  }
}`

type Weibo struct {
	User     string                `json:"user"`               // 用户
	Message  string                `json:"message"`            // 微博内容
	Retweets int                   `json:"retweets"`           // 转发数
	Image    string                `json:"image,omitempty"`    // 图片
	Created  time.Time             `json:"created,omitempty"`  // 创建时间
	Tags     []string              `json:"tags,omitempty"`     // 标签
	Location string                `json:"location,omitempty"` //位置
	Suggest  *elastic.SuggestField `json:"suggest_field,omitempty"`
}

var indexName = "weibo"

// Create 创建索引
func (es *Elasticsearch) Create() {
	var client = elasticsearch.Es
	isIndex, err := client.Client.IndexExists(indexName).Do(client.Context)
	if err != nil {
		fmt.Println(err)
		return
	}

	if isIndex == true {
		fmt.Println("索引已经存在了")
		return
	}

	_, err = client.Client.CreateIndex(indexName).BodyString(mapping).Do(client.Context)
	if err != nil {
		return
	}

	fmt.Printf("创建索引成功:%v \n", indexName)

}

func (es *Elasticsearch) Add() {
	// 创建创建一条微博
	var client = elasticsearch.Es
	msg1 := Weibo{User: "xiaoming", Message: "打酱油的一天", Retweets: 0}
	data, err := client.Client.Index().Index("indexName").Id("1").BodyJson(msg1).Do(client.Context)
	if err != nil {
		return
	}

	fmt.Printf("文档Id %s, 索引名 %s\n", data.Id, data.Index)
}
