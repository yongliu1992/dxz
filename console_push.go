package dxz
type MailMessage struct {
	Type int `json:"type"` //1代表控制台线索信息
	Content interface{} `json:"content"`
}
type ConsolePush struct {
	Tel     string `json:"tel"`
	Name    string `json:"name"`
	Country string `json:"country"`
	Time    string `json:"time"`
	Id      int    `json:"id"` //resource_allocation id
	Remark  string `json:"remark"`
	Wx      string `json:"wx"`
	Qq      string `json:"qq"`
}
