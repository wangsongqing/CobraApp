package controllers

import (
	"CobraApp/pkg/elasticsearch"
	"fmt"
	"github.com/olivere/elastic/v7"
	"reflect"
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
	msg1 := Weibo{User: "mengmeng", Message: "打酱油的第二天", Retweets: 0}
	data, err := client.Client.Index().Index(indexName).Id("2").BodyJson(msg1).Do(client.Context)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("文档Id %s, 索引名 %s\n", data.Id, data.Index)
}

func (es *Elasticsearch) Search() {
	var client = elasticsearch.Es
	//termQuery := elastic.NewTermQuery("user", "xiaoming") // 精确匹配单个字段
	//termQuery := elastic.NewTermsQuery("user", "mengmeng", "xiaohong") // 通过terms实现SQL的in查询
	termQuery := elastic.NewMatchQuery("message", "大家") // 匹配单个字段, 某个字段使用全文搜索，也就是ES的match语法
	//termQuery := elastic.NewRangeQuery("id").Gte(1).Lte(10) // 等价表达式： id >= 1 and id < 10

	// must查询，类型and
	// 设置bool查询的must条件, 组合了两个子查询
	// 表示搜索匹配user=mengmeng且message匹配"二"的文档
	//termQuery := elastic.NewBoolQuery().Must()
	//termQuery1 := elastic.NewTermQuery("user", "mengmeng")
	//matchQuery := elastic.NewMatchQuery("message", "二")
	//termQuery.Must(termQuery1, matchQuery)

	searchResult, err := client.Client.Search().
		Index(indexName).      // 设置索引名
		Query(termQuery).      // 设置查询条件
		Sort("created", true). // 设置排序字段，根据Created字段升序排序，第二个参数false表示逆序
		From(0).               // 设置分页参数 - 起始偏移量，从第0行记录开始
		Size(10).              // 设置分页参数 - 每页大小
		Pretty(true).          // 查询结果返回可读性较好的JSON格式
		Do(client.Context)     // 执行请求

	if err != nil {
		panic(err)
	}

	fmt.Printf("查询消耗时间 %d ms, 结果总数: %d\n", searchResult.TookInMillis, searchResult.TotalHits())

	if searchResult.TotalHits() == 0 {
		return
	}

	// 查询结果不为空，则遍历结果
	var b1 Weibo
	// 通过Each方法，将es结果的json结构转换成struct对象
	for _, item := range searchResult.Each(reflect.TypeOf(b1)) {
		// 转换成Article对象
		if t, ok := item.(Weibo); ok {
			fmt.Println(t.Message)
		}
	}
}

// Update 更新数据
func (es *Elasticsearch) Update() {
	var client = elasticsearch.Es
	data := map[string]interface{}{"retweets": 521, "message": "我笑了一下下大家"}
	result, err := client.Client.Update().
		Index(indexName).  // 设置索引名
		Id("2").           // 文档id
		Doc(data).         // 更新retweets=0，支持传入键值结构
		Do(client.Context) // 执行ES查询
	if err != nil {
		// Handle error
		panic(err)
	}

	fmt.Println(result)
}

// Delete 删除数据
func (es *Elasticsearch) Delete() {
	// 根据id删除一条数据
	var client = elasticsearch.Es
	_, err := client.Client.Delete().
		Index("weibo").
		Id("1").
		Do(client.Context)
	if err != nil {
		// Handle error
		panic(err)
	}
}
