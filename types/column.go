package types

type Column struct {
	Id          string `json:"id"`
	Description string `json:"description"`
	DataType    string `json:"dataType"`
}

const (
	String    string = "String"
	Number    string = "Number"
	Boolean   string = "Boolean"
	Date      string = "Date"
	TimeStamp string = "TimeStamp"
	List      string = "List"
)

func NewColumn(id string, description string, dataType string) Column {
	return Column{
		Id:          id,
		Description: description,
		DataType:    dataType,
	}
}
