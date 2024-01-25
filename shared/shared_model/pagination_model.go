package shared_model

type Paging struct {
	Page        int `json:"Page"`
	RowsPerPage int `json:"rowsPerPage"`
	TotalRows   int `json:"totalRows"`
	TotalPages  int `json:"totalPages"`
}

func (p Paging) Error() string {
	//TODO implement me
	panic("implement me")
}
