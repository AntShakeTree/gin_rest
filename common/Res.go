package common

type Header struct {
	Error   int `json:"error"`
	Message string `json:"message,omitempty"`
}

//

func (header *Header) Success() {
	header.Error = GetCfg().Err["success"]
	header.Message = ""
}


type Res struct {
	Header Header `json:"header"`
	Body   interface{} `json:"body,omitempty"`
}

func (res *Res) Instances(body interface{}) Res {
	res.Body = body
	res.Header = Head
	res.Header.Success()
	return *res
}
func (res *Res) InstancesNoBody() Res {
	res.Header = Head
	res.Header.Success()

	return *res
}

var Response = Res{}
var Head = Header{}
