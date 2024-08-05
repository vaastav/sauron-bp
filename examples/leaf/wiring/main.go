package main

import (
	"github.com/blueprint-uservices/blueprint/plugins/cmdbuilder"
	"github.com/blueprint-uservices/blueprint/plugins/workflow/workflowspec"
	"github.com/vaastav/sauron-bp/examples/leaf/wiring/specs"
)

func main() {
	// Configure the location of our workflow spec
	workflowspec.AddModule("github.com/blueprint-uservices/blueprint/examples/leaf/workflow")

	// Build a supported wiring spec
	name := "LeafApp"
	cmdbuilder.MakeAndExecute(
		name,
		specs.Default,
	)
}
