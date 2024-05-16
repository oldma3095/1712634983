package slave

type FileTransfer struct {
	Base
	Filename string `json:"filename"`
	Data     []byte `json:"data"`
	MD5      string `json:"MD5"`
}
