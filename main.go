package main

import (
	"github.com/turbot/steampipe-plugin-azure/azure"

	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		PluginFunc: azure.Plugin})
}
