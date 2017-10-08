package models

type Command struct {
	Command string   `json:"command"`
	Args    []string `json:"Args"`
	Output  string   `json:"output"`
}
