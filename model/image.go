package model

import "net/textproto"

type Image struct {
	Src      []byte               `json:"src"`
	FileName string               `json:"file_name"`
	Size     int64                `json:"size"`
	Header   textproto.MIMEHeader `json:"header"`
}
