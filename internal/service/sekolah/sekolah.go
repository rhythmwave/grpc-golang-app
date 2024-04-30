package sekolah

type GetListSekolah struct {
	IdSekolah   int    `json:"id_sekolah"`
	NamaSekolah string `json:"nama_sekolah"`
}

type SekolahKelasList struct {
	ID           int    `json:"id"`
	TingkatKelas string `json:"tingkat_kelas"`
	Jenjang      string `json:"jenjang"`
}

type GetSekolahListResponse struct {
	Data []SekolahKelasList `json:"data"`
	Meta MetaResponse       `json:"_meta"`
}

type GetTargetSiswaData struct {
	IdGedung     int32
	TahunAjaran  string
	IdSekolah    int32
	NamaSekolah  string
	IdTingkat    int32
	TingkatKelas string
	JmlTarget    int32
}

type ChannelTargetSiswa struct {
	IDX  uint16
	Data []GetTargetSiswaData
}
