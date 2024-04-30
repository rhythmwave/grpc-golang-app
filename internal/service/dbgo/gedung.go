package dbgo

type GedungData struct {
	IDGedung   int    `json:"c_id_gedung"`
	NamaGedung string `json:"c_nama_gedung"`
	ZonaWaktu  string `json:"c_zona_waktu"`
}

type ListGedungData struct {
	IDGedung   int    `json:"c_id_gedung"`
	NamaGedung string `json:"c_nama_gedung"`
}

type GetGedungAllResponse struct {
	StatusCode int            `json:"status"`
	Message    string         `json:"message"`
	Data       ListGedungData `json:"data"`
}

type GetGedungResponse struct {
	Code    int        `json:"code"`
	Message string     `json:"message"`
	Data    GedungData `json:"data"`
}
