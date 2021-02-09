package controller

import (
	"k8s.io/ingress-nginx/internal/file"
	"k8s.io/ingress-nginx/internal/ingress/metric"
)

type CustomNGINEXController struct {
	*NGINXController
}

func NewCustomNGINXController(conf *Configuration, mc metric.Collector, fs file.Filesystem, ng NginxExecTester) *CustomNGINEXController {
	ngCon := NewNGINXController(conf, mc)
	customNgCon := &CustomNGINEXController{ngCon}
	customNgCon.command = ng
	return customNgCon
}
