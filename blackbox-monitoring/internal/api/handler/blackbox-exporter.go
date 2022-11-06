package handler

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/blackbox_exporter/config"
	"github.com/prometheus/common/promlog"
	"gopkg.in/yaml.v3"
)

type BlackboxExporter struct{}

type Module struct {
	Prober  string           `yaml:"prober,omitempty"`
	Timeout time.Duration    `yaml:"timeout,omitempty"`
	HTTP    config.HTTPProbe `yaml:"http,omitempty"`
	// TCP     TCPProbe      `yaml:"tcp,omitempty"`
	// ICMP    ICMPProbe     `yaml:"icmp,omitempty"`
	// DNS     DNSProbe      `yaml:"dns,omitempty"`
	// GRPC    GRPCProbe     `yaml:"grpc,omitempty"`
}

type Modules struct {
	Modules map[string]Module `yaml:"modules"`
}

type HTTPProbe struct {
	Name    string           `yaml:"name,omitempty"`
	URL     string           `yaml:"url,omitempty"`
	HTTP    config.HTTPProbe `yaml:"http,omitempty"`
	Contact string           `yaml:"contact,omitempty"`
}

func (p *BlackboxExporter) Get(c *gin.Context) {
	promlogConfig := &promlog.Config{}
	logger := promlog.New(promlogConfig)
	sc := &config.SafeConfig{
		C: &config.Config{},
	}

	err := sc.ReloadConfig("/Users/duxianghua/workspace/demoandpoc/blackbox-exporter/config.yml", logger)
	if err != nil {
		fmt.Println(err)
	}
	for _, expression := range sc.C.Modules["http_baidu"].HTTP.FailIfBodyMatchesRegexp {
		fmt.Println(expression.Regexp.String())
	}

	yamlData, _ := yaml.Marshal(&sc.C)
	fmt.Println(" --- YAML ---")
	fmt.Println(string(yamlData))
	c.JSON(200, sc)
}

func (p *BlackboxExporter) GetHttpProbe(c *gin.Context) {
	body := &HTTPProbe{
		Name:    "test1",
		URL:     "www.bbaidu.com",
		HTTP:    config.HTTPProbe{},
		Contact: "asdfadf@qq.com",
	}
	yamlData, _ := yaml.Marshal(body)
	fmt.Println(" --- YAML ---")
	fmt.Println(string(yamlData))
	c.JSON(200, body)
}
