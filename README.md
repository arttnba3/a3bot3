# a3bot3 - Documentation

## Introduction

_Yet another QQbot developed by arttnba3 with golang, based on go-cqhttp_

## Usage

### _*Requirement_

You need to have a [go-cqhttp](https://github.com/Mrs4s/go-cqhttp) firstly, config it, then run it with a3bot3

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
  "type":"http"
}


```

Supported connection type now:

- http

## API

The a3bot3 satisfies the [OneBot API](https://onebot.dev/), with Camel-Case

## I'd like to develop new plugin...

Add your plugin source file under `plugin` folder, write your plugin according to `ExamplePlugin`, then add it under `bot/handler.go`:

```go
var plugins = [...]plugin.Plugin{
    //...
	plugin.YourPLugin,
}
```

## To-do list

- Add connection type:
  - reverse websocket
- Complete All OneBot APIs
- Add a more-complete plugin systems
- Add more plugins