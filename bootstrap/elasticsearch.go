package bootstrap

import (
	"CobraApp/pkg/config"
	"CobraApp/pkg/elasticsearch"
	"fmt"
)

func SetupEs() {
	address := fmt.Sprintf("http://%v:%v", config.GetString("elasticsearch.host"), config.GetString("elasticsearch.port"))
	elasticsearch.ConnectElasticsearch(address)
}
