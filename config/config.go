package config

type URLForm struct {
	Host string `json:"host"`
	Port int    `json:"port"`
}

type BotSettings struct {
	Listen           URLForm `json:"listen"`
	Post             URLForm `json:"post"`
	Type             string  `json:"type"`
	ParseCommandOnly bool    `json:"parse_command_only"`
	Admin            int64   `json:"admin"`
}

var Settings BotSettings
