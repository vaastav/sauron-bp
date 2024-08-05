package specs

import (
	"strings"

	"github.com/blueprint-uservices/blueprint/blueprint/pkg/wiring"
	"github.com/blueprint-uservices/blueprint/plugins/cmdbuilder"
	"github.com/blueprint-uservices/blueprint/plugins/goproc"
	"github.com/blueprint-uservices/blueprint/plugins/http"
	"github.com/blueprint-uservices/blueprint/plugins/linuxcontainer"
	"github.com/blueprint-uservices/blueprint/plugins/simple"
	"github.com/blueprint-uservices/blueprint/plugins/workflow"
	"github.com/vaastav/sauron-bp/examples/leaf/workflow/leaf"
	"github.com/vaastav/sauron-bp/plugin/sauron"
)

var Default = cmdbuilder.SpecOption{
	Name:        "default",
	Description: "Deploys each service in a separate container, communicating using HTTP. Wraps each service in XTrace tracing and sets the XTraceLogger for each process.",
	Build:       makeSauronLoggerSpec,
}

func makeSauronLoggerSpec(spec wiring.WiringSpec) ([]string, error) {
	applyLoggerDefaults := func(service_name string) string {

		procName := strings.ReplaceAll(service_name, "service", "process")
		cntrName := strings.ReplaceAll(service_name, "service", "container")
		http.Deploy(spec, service_name)
		goproc.CreateProcess(spec, procName, service_name)
		sauron.Logger(spec, procName, "1", procName+"_log.txt")
		return linuxcontainer.CreateContainer(spec, cntrName, procName)
	}
	leaf_db := simple.NoSQLDB(spec, "leaf_db")
	leaf_cache := simple.Cache(spec, "leaf_cache")
	leaf_service := workflow.Service[*leaf.LeafServiceImpl](spec, "leaf_service", leaf_cache, leaf_db)
	leaf_proc := applyLoggerDefaults(leaf_service)

	nonleaf_service := workflow.Service[leaf.NonLeafService](spec, "nonleaf_service", leaf_service)
	nonleaf_proc := applyLoggerDefaults(nonleaf_service)

	return []string{leaf_proc, nonleaf_proc}, nil
}
