package slave

type FileTransfer struct {
	Filename  string `json:"filename"`
	Data      []byte `json:"data"`
	MD5       string `json:"MD5"`
	Timestamp int64  `json:"timestamp"`
}
