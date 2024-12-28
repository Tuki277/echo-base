package contract

type QueryRequest interface {
	GetPage() uint
	GetLimit() uint
}

type ListRequest struct {
	Page  uint `json:"page"`
	Limit uint `json:"limit"`
}

func (r *ListRequest) GetPage() uint {
	return r.Page
}

func (r *ListRequest) GetLimit() uint {
	return r.Limit
}

type ListResponse struct {
	TotalCount uint `json:"total_count"`
	MaxPage    uint `json:"max_page"`
	Page       uint `json:"page"`
	PerPage    uint `json:"per_page"`
}
