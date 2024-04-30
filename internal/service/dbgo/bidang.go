package dbgo

type BidangDetail struct {
	IDKota    int    `json:"id_kota"`
	Kota      string `json:"kota"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type GetBidangFindOneRespose struct {
	Status    int         `json:"status"`
	Message string      `json:"message"`
	Data    BidangDetail `json:"data"`
}
