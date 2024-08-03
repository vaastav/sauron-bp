package sauron

import (
	"github.com/blueprint-uservices/blueprint/blueprint/pkg/ir"
	"github.com/blueprint-uservices/blueprint/blueprint/pkg/wiring"
	"github.com/blueprint-uservices/blueprint/plugins/goproc"
)

// [Logger] can be used by wiring specs to install a process-level sauron logger for process `processName`. Replaces the existing logger installed for the process.
//
// # Wiring Spec Usage:
//
//	sauron.Logger(spec, "my_process")
func Logger(spec wiring.WiringSpec, processName string, buffer_size int64) string {
	logger := processName + "_sauron_logger"
	spec.Define(logger, &SauronLogger{}, func(ns wiring.Namespace) (ir.IRNode, error) {
		return newSauronLogger(logger, buffer_size)
	})
	goproc.SetLogger(spec, processName, logger)
	return logger
}
