package sauron

import (
	"fmt"

	"github.com/blueprint-uservices/blueprint/blueprint/pkg/coreplugins/service"
	"github.com/blueprint-uservices/blueprint/blueprint/pkg/ir"
	"github.com/blueprint-uservices/blueprint/plugins/golang"
	"github.com/blueprint-uservices/blueprint/plugins/workflow/workflowspec"
	runtime "github.com/vaastav/sauron-bp/sauron_runtime/sauron"
	"golang.org/x/exp/slog"
)

// Blueprint IR Node that represents a process-level OT trace logger
type SauronLogger struct {
	golang.Node
	golang.Instantiable

	LoggerName string
	BufferSize string
	Outfile    string
	Spec       *workflowspec.Service
}

func newSauronLogger(name string, buffer_size string, outfile string) (*SauronLogger, error) {
	spec, err := workflowspec.GetService[runtime.SauronLogger]()
	node := &SauronLogger{
		LoggerName: name,
		Spec:       spec,
		BufferSize: buffer_size,
		Outfile:    outfile,
	}
	return node, err
}

// Implements ir.IRNode
func (node *SauronLogger) Name() string {
	return node.LoggerName
}

// Implements ir.IRNode
func (node *SauronLogger) String() string {
	return node.Name() + " = SauronLogger()"
}

// Implements golang.ProvidesModule
func (node *SauronLogger) AddToWorkspace(builder golang.WorkspaceBuilder) error {
	return node.Spec.AddToWorkspace(builder)
}

// Implements golang.ProvidesInterface
func (node *SauronLogger) AddInterfaces(builder golang.ModuleBuilder) error {
	return node.Spec.AddToModule(builder)
}

// Implements golang.ProvidesInterface
func (node *SauronLogger) GetInterface(ctx ir.BuildContext) (service.ServiceInterface, error) {
	return node.Spec.Iface.ServiceInterface(ctx), nil
}

// Implements golang.Instantiable
func (node *SauronLogger) AddInstantiation(builder golang.NamespaceBuilder) error {
	if builder.Visited(node.LoggerName) {
		return nil
	}

	slog.Info(fmt.Sprintf("Instantiating SauronLogger %v in %v/%v", node.LoggerName, builder.Info().Package.PackageName, builder.Info().FileName))

	return builder.DeclareConstructor(node.LoggerName, node.Spec.Constructor.AsConstructor(), []ir.IRNode{&ir.IRValue{Value: node.BufferSize}, &ir.IRValue{Value: node.Outfile}})
}

// Implements ir.IRNode
func (node *SauronLogger) ImplementsGolangNode() {}
