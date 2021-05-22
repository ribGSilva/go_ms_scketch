package object

type responseTemplate struct {
	Status   int
	Response interface{}
}

type ResponseTemplate interface {
	GetStatus() int
	HasResponse() bool
	GetResponse() interface{}
}

func NewResponseTemplate(status int, response interface{}) ResponseTemplate {
	return &responseTemplate{Status: status, Response: response}
}

func (rt *responseTemplate) GetStatus() int {
	return rt.Status
}

func (rt *responseTemplate) HasResponse() bool {
	return rt.Response != nil
}

func (rt *responseTemplate) GetResponse() interface{} {
	return rt.Response
}
