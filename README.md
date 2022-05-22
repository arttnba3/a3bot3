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

The a3bot3 satisfies the [OneBot API](https://onebot.dev/), with Camel-Case (Because of fucking art of golang)

## Plugins

Available plugins are as bellow:

- `PluginSystem` : The plugin to control all other plugins
- `ExamplePlugin` : Only a model of plugin
- `RepeaterPlugin` : Repeat the message that had been repeated for twice. The same message will only be repeated once.
- `AntiRecallPlugin` : Resend the message that had been recalled. _Fighting against all recalled message!_

## _I'd like to develop new plugin..._

### I.Basic

If you'd like to add a new plugin, you should create your plugin according to `ExamplePlugin` in `plugin/example_plugin.go`, make it inheritance the `PluginInfo` strutct and complete two method:

```go
type YourPlugin struct {
	PluginInfo
}

func (p *YourPlugin) SendPrivateMsg(bot api.BotAPI, privateEvent event.PrivateEvent, messages []string) int {
	// implement it
}

func (p *YourPlugin) SendGroupMsg(bot api.BotAPI, groupEvent event.GroupEvent, messages []string) int {
	// implement it
}
```

Then add a new instance under `plugin/config.go`:

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

- `Enable` : It decides whether your plugin is loaded when the bot starts.
- `Name` : Name of your plugin.
- `Command` : Command of your plugin. You can use `MatchCommand()` to check whether the user's input match the command.

### II.Advanced

I also provide you with other APIs to deal with other events. You can check for `plugin/api.go` to see the definition of available APIs and implement your own one, and it'll be call automatically by the bot when the event comes.

For example, if you'd like to deal with the message-recalling event in group, you just need to implement your own `GroupRecallHandler`:

```go
func (p *YourPlugin) GroupRecallHandler(bot api.BotAPI, groupEvent event.GroupEvent, messages []string) int {
	// implement it!
}
```

Then run the bot make the plugin enabled. When the group-recalling event comes, your own `GroupRecallHandler` will be called automatically by the bot.

## To-do list

- Add connection type:
  - reverse websocket
- Complete All OneBot APIs
- Add a more-complete plugin systems
- Add more plugins
  - move some plugins from [a3bot2](https://github.com/arttnba3/a3bot2)