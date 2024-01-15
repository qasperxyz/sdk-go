package types

type Element struct {
	Elements []Element `json:"elements"`
	Value    Fields    `json:"value"`
}

type Fields map[string]any

type ElementBuilder struct {
	element Element
}

func NewElementBuilder() *ElementBuilder {
	return &ElementBuilder{
		element: Element{
			Elements: []Element{},
			Value:    make(Fields),
		},
	}
}

func (h *ElementBuilder) AddElement(el Element) *ElementBuilder {
	h.element.Elements = append(h.element.Elements, el)
	return h
}

func (h *ElementBuilder) AddField(columnId string, value any) *ElementBuilder {
	h.element.Value[columnId] = value
	return h
}

func (h *ElementBuilder) Build() *Element {
	return &h.element
}
