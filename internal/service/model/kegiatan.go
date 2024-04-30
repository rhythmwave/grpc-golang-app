package model

type CPengajuan struct {
	CPajak      int32  `json:"c_Pajak"`
	CUpdater    string `json:"c_Updater"`
	CPajakOpsi  string `json:"c_PajakOpsi"`
	CLastUpdate string `json:"c_LastUpdate"`
}

type CSetujuKacab struct {
	CPajak       int32  `json:"c_Pajak"`
	CTotal       int32  `json:"c_Total"`
	CJumlah      int32  `json:"c_Jumlah"`
	CUpdater     string `json:"c_Updater"`
	CPajakOpsi   string `json:"c_PajakOpsi"`
	CLastUpdate  string `json:"c_LastUpdate"`
	CNilaiSatuan int32  `json:"c_NilaiSatuan"`
}

type CSetujuBidang struct {
	CPajak       int32  `json:"c_Pajak"`
	CTotal       int32  `json:"c_Total"`
	CJumlah      int32  `json:"c_Jumlah"`
	CUpdater     string `json:"c_Updater"`
	CPajakOpsi   string `json:"c_PajakOpsi"`
	CLastUpdate  string `json:"c_LastUpdate"`
	CNilaiSatuan int32  `json:"c_NilaiSatuan"`
}

type CSetujuAkuntansi struct {
	CPajak       int32  `json:"c_Pajak"`
	CTotal       int32  `json:"c_Total"`
	CJumlah      int32  `json:"c_Jumlah"`
	CUpdater     string `json:"c_Updater"`
	CPajakOpsi   string `json:"c_PajakOpsi"`
	CLastUpdate  string `json:"c_LastUpdate"`
	CNilaiSatuan int32  `json:"c_NilaiSatuan"`
}

type TKebutuhanBarangCAcc struct {
	CPengajuan       CPengajuan       `json:"c_Pengajuan"`
	CSetujuKacab     CSetujuKacab     `json:"c_SetujuKacab"`
	CSetujuBidang    CSetujuBidang    `json:"c_SetujuBidang"`
	CSetujuAkuntansi CSetujuAkuntansi `json:"c_SetujuAkuntansi"`
}
