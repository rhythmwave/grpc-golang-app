package dbgo

type KaryawanMetaStruct struct {
	Code    int    `json:"code"`
	Status  string `json:"status"`
	Message string `json:"message"`
}

type KaryawanKontakInfo struct {
	HP   string `json:"HP"`
	HP2  string `json:"HP2"`
	TLP  string `json:"TLP"`
	TLP2 string `json:"TLP2"`
}

type KaryawanDetailInfoData struct {
	CID               int                `json:"c_id"`
	CNik              string             `json:"c_nik"`
	CIDBiodata        int                `json:"c_id_biodata"`
	CNamaLengkap      string             `json:"c_nama_lengkap"`
	CIDGedung         int                `json:"c_id_gedung"`
	CIDKomar          int                `json:"c_id_komar"`
	CIDKota           int                `json:"c_id_kota"`
	CIDBidang         int                `json:"c_id_bidang"`
	CIDJabatan        int                `json:"c_id_jabatan"`
	CIDKomarCakupan   int                `json:"c_id_komar_cakupan"`
	CInfoKontak       KaryawanKontakInfo `json:"c_info_kontak"`
	CStatus           string             `json:"c_status"`
	CTanggalAkhirPKWT string             `json:"c_tanggal_akhir_pkwt"`
	KotaIDs           string             `json:"kota_ids"`
	GedungIDs         string             `json:"gedung_ids"`
}

type KaryawanDetailInfo struct {
	Data KaryawanDetailInfoData `json:"data"`
	Meta KaryawanMetaStruct     `json:"_meta"`
	IDX  int                    ``
}
