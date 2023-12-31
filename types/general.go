package types

type Response struct {
	Messages []Message `json:"messages"`
	Columns  []Column  `json:"columns"`
	Elements []Element `json:"elements"`
}

type ResponseBuilder struct {
	response Response
}

func NewResponseBuilder() *ResponseBuilder {
	return &ResponseBuilder{
		response: Response{
			Messages: []Message{},
			Columns:  []Column{},
			Elements: []Element{},
		},
	}
}

func (h *ResponseBuilder) AddMessage(message Message) *ResponseBuilder {
	h.response.Messages = append(h.response.Messages, message)
	return h
}

func (h *ResponseBuilder) AddColumn(column Column) *ResponseBuilder {
	h.response.Columns = append(h.response.Columns, column)
	return h
}

func (h *ResponseBuilder) AddElement(element Element) *ResponseBuilder {
	h.response.Elements = append(h.response.Elements, element)
	return h
}

func (h *ResponseBuilder) Build() *Response {
	return &h.response
}
