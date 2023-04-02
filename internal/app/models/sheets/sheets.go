package sheets

type Response struct {
	Id         string      `json:"spreadsheetId"`
	Properties interface{} `json:"properties"`
	Sheets     []Sheets    `json:"sheets"`
}

type Sheets struct {
	Properties interface{}  `json:"properties"`
	Data       []SheetsData `json:"data"`
}

type SheetsData struct {
	StartRow    int       `json:"startRow"`
	StartColumn int       `json:"startColumn"`
	RowData     []RowData `json:"rowDATA"`
}

type RowData struct {
	Values []Values `json:"values"`
}

type Values struct {
	Value string `json:"formattedValue"`
}

type Storage struct {
	Data []SheetsData
}

type SheetStorage interface {
	GetData() []SheetsData
}

func NewStorage(d []SheetsData) *Storage {
	return &Storage{
		Data: d,
	}
}

func (s *Storage) GetData() []SheetsData {
	return s.Data
}
