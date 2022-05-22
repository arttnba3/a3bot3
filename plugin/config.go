package plugin

type PluginInfo struct {
	Enable  bool
	Name    string
	Command string
	Plugin
}

// Plugins :
// This is the array of instances of all available plugins.
var Plugins = [...]Plugin{
	&PluginManager{
		PluginInfo: PluginInfo{
			Enable:  true,
			Name:    "PluginManager",
			Command: "/plugin",
		},
	},
	&ExamplePlugin{
		PluginInfo: PluginInfo{
			Enable:  false,
			Name:    "ExamplePlugin",
			Command: "/hello",
		},
	},
	&RepeaterPlugin{
		PluginInfo: PluginInfo{
			Enable:  true,
			Name:    "RepeaterPlugin",
			Command: "",
		},
	},
	&AntiRecallPlugin{
		PluginInfo: PluginInfo{
			Enable: false,
			Name:   "AntiRecallPlugin",
		},
	},
}
