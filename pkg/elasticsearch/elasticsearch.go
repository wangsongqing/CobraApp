package elasticsearch

import (
	"context"
	"fmt"
	es "github.com/olivere/elastic/v7"
	"sync"
)

type ElasticsearchClient struct {
	Client  *es.Client
	Context context.Context
}

// once 确保全局的 Elasticsearch 对象只实例一次
var once sync.Once

var Es *ElasticsearchClient

// ConnectElasticsearch ConnectRedis ConnectElasticsearch 连接 elasticsearch 数据库，设置全局的 elasticsearch 对象
func ConnectElasticsearch(address string) {
	once.Do(func() {
		Es = NewClient(address)
	})
}

func NewClient(address string) *ElasticsearchClient {
	esc := &ElasticsearchClient{}

	// 创建ES client用于后续操作ES
	client, err := es.NewClient(
		// 设置ES服务地址，支持多个地址
		es.SetURL(address),
		// 设置基于http base auth验证的账号和密码
		//es.SetBasicAuth("user", "secret")
	)
	if err != nil {
		// Handle error
		fmt.Printf("连接失败: %v\n", err)
	}

	esc.Client = client
	// 使用默认的 context
	esc.Context = context.Background()

	return esc
}
