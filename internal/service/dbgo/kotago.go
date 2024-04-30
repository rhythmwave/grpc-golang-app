package dbgo

type KotaDataAll struct {
	IDKota    int    `json:"id_kota"`
	Kota      string `json:"kota"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type GetKotaGoResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    KotaDataAll `json:"data"`
}
