# a3bot3 - Documentation

## Introduction

_Yet another QQbot developed by arttnba3 with golang, based on go-cqhttp_

## Usage

### _*Requirement_

You need to have a [go-cqhttp](https://github.com/Mrs4s/go-cqhttp) firstly, config it, then run it with a3bot3

### Build

The project doesn't need extra packages temporarily, so you can build it simply with:

```shell
$ go build
```

Then an executable binary file `a3bot` will appear. JUST RUN TI WITH `go-cqhttp`.

## Config

You should config your settings in file `config.json` as following:

```json
{
    "post":
    {
        "host":"127.0.0.1",
        "port":5700
    },
    "listen":
    {
        "host":"127.0.0.1",
        "port":5701
    },
    "type":"http",
    "parse_command_only":true,
    "admin":1145141919
}
```

Supported connection type now:

- http

## API

The a3bot3 satisfies the [OneBot API](https://onebot.dev/), with Camel-Case

## _I'd like to develop new plugin..._

If you'd like to add a new plugin, you should create your plugin according to `ExamplePlugin` in `plugin/example_plugin.go`, make it inheritance the `PluginInfo` strutct and complete two method:

```go
type YourPlugin struct {
	PluginInfo
}

func (p *ExamplePlugin) SendPrivateMsg(bot api.BotAPI, privateEvent event.PrivateEvent, messages []string) int {
	// implement it
}

func (p *ExamplePlugin) SendGroupMsg(bot api.BotAPI, groupEvent event.GroupEvent, messages []string) int {
	// implement it
}
```

Then add a new instance under `plugin/plugin.go`:

```go
var Plugins = [...]plugin.Plugin{
    //...
	&plugin.YourPlugin{
		PluginInfo: plugin.PluginInfo{
			Enable:  true,
			Name:    "YourPlugin",
			Command: "/your_command",
		},
	},
}
```

## To-do list

- Add connection type:
  - reverse websocket
- Complete All OneBot APIs
- Add a more-complete plugin systems
- Add more plugins