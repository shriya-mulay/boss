package main

type Config struct {
	ID        string     `toml:"id"`
	Image     string     `toml:"image"`
	Resources *Resources `toml:"resources"`
	Network   Network    `toml:"network"`
	Mounts    []Mount    `toml:"mounts"`
	Env       []string   `toml:"env"`
}

type Resources struct {
	CPU    float64 `toml:"cpu"`
	Memory int64   `toml:"memory"`
	Score  int     `toml:"score"`
}

type Network struct {
	Host bool `toml:"host"`
	CNI  bool `toml:"cni"`
}

type Mount struct {
	Type        string   `toml:"type"`
	Source      string   `toml:"source"`
	Destination string   `toml:"destination"`
	Options     []string `toml:"options"`
}