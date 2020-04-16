package result

type Page struct {
	Total int32       `json:"total"`
	Page  int32       `json:"page"`
	Limit int32       `json:"limit"`
	Data  interface{} `json:"data"`
}

func NewPage(total int32, page int32, limit int32, data interface{}) *Page {
	return &Page{Total: total, Page: page, Limit: limit, Data: data}
}
