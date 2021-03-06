package haproxy

import (
	"github.com/samuel/go-zookeeper/zk"

	conf "github.com/seomoz/roger-bamboo/configuration"
	"github.com/seomoz/roger-bamboo/services/marathon"
	"github.com/seomoz/roger-bamboo/services/service"
)

type TemplateData struct {
	Apps     marathon.AppList
	Services map[string]service.Service
}

func GetTemplateData(config *conf.Configuration, conn *zk.Conn) TemplateData {

	apps, _ := marathon.FetchApps(config.Marathon)
	services, _ := service.All(conn, config.Bamboo.Zookeeper)

	return TemplateData{apps, services}
}
