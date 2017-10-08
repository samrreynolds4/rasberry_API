package models

type fileTransfer struct {
	FileSource string      `json:"fileSource"`
	FileDest   string      `json:"fileDest"`
	Cred       Credentials `json:"credentials"`
}
