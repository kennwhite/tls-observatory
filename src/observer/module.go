package observer

import (
	"fmt"
	"os"
)

// ModuleResult implement the base type for results returned by modules.
// All modules must return this type of result. The fields are:
//
// - FoundAnything: a boolean that must be set to true if the module ran
//                  a search that returned at least one positive result
//
// - Success: a boolean that must be set to true if the module ran without
//            fatal errors. soft errors are reported in Errors
//
// - Elements: an undefined type that can be customized by the module to
//             contain the detailled results
//
// - Statistics: an undefined type that can be customized by the module to
//               contain some information about how it ran
//
// - Errors: an array of strings that contain non-fatal errors encountered
//           by the module
type ModuleResult struct {
	FoundAnything bool        `json:"foundanything"`
	Success       bool        `json:"success"`
	Elements      interface{} `json:"elements"`
	Statistics    interface{} `json:"statistics"`
	Errors        []string    `json:"errors"`
}

type ModuleInfo struct {
	InputQueue string
	RunnerFunc func() interface{}
}

// AvailableModules stores a list of activated module with their runner
var AvailableModules = make(map[string]ModuleInfo)

// RegisterModule adds a module to the list of available modules
func RegisterModule(name string, info ModuleInfo) {
	if _, exist := AvailableModules[name]; exist {
		fmt.Fprintf(os.Stderr, "RegisterModule: a module named '%s' has already been registered.\nAre you trying to import the same module twice?\n", name)
		os.Exit(1)
	}
	AvailableModules[name] = info
}

// Moduler provides the interface to a Module
type Moduler interface {
	Run([]byte) string
}

// HasResultsPrinter implements functions used by module to print information
type HasResultsPrinter interface {
	PrintResults([]byte, bool) ([]string, error)
}

// HasParamsCreator implements a function that creates module parameters
type HasParamsCreator interface {
	ParamsCreator() (interface{}, error)
}
