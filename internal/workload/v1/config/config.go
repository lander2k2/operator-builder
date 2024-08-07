// Copyright 2024 Nukleros
// SPDX-License-Identifier: Apache-2.0

package config

const PluginKey = "operatorBuilder"

// Plugin contains the project config values which are stored in the
// PROJECT file under plugins.operatorBuilder.
type Plugin struct {
	WorkloadConfigPath string `json:"workloadConfigPath" yaml:"workloadConfigPath"`
	CliRootCommandName string `json:"cliRootCommandName" yaml:"cliRootCommandName"`
	ControllerImg      string `json:"controllerImg" yaml:"controllerImg"`
	EnableOLM          bool   `json:"enableOlm" yaml:"enableOlm"`
}
