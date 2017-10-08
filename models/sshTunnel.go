package models

type SshTunnel struct {
	PiPort   string      `json:"piport"`
	DestPort string      `json:"destPort"`
	DestIP   string      `json:"destIP"`
	Cred     Credentials `json:"credentials"`
}
