package blackboxexporter

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/prometheus/blackbox_exporter/config"
)

func TestXxx(t *testing.T) {
	a := config.HTTPProbe{
		ValidStatusCodes: []int{200, 300},
	}
	fmt.Println(json.Marshal(a))
}
