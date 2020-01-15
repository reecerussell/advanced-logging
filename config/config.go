package config

type Configuration struct {
	PrintPrefix string    `json:"printPrefix"`
	FatalPrefix string    `json:"fatalPrefix"`
	PanicPrefix string    `json:"panicPrefix"`
	Plugins     []*Plugin `json:"plugins"`
}

type Plugin struct {
	Name    string                 `json:"name"`
	Enabled bool                   `json:"enabled"`
	Config  map[string]interface{} `json:"config"`
}
