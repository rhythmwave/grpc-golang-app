package service

import (
	"context"
	"encoding/json"
	"fmt"
	"go-bponline/m/config"
	"go-bponline/m/gen/model"
	"go-bponline/m/internal/database"
	"go-bponline/m/internal/service/dbgo"
	DataModel "go-bponline/m/internal/service/model"
	"go-bponline/m/internal/service/sekolah"
	"go-bponline/m/internal/service/worker"
	"go-bponline/m/internal/utils"
	"go-bponline/m/pb/kegiatan"
	"net/http"
	"strings"
	"sync"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Kegiatan struct {
	kegiatan.UnimplementedKegiatanServiceServer
}

func (m *Kegiatan) GetRencanaKegiatanList(ctx context.Context, req *kegiatan.GetRencanaKegiatanListRequest) (*kegiatan.RencanaKegiatanListRes, error) {
	var renKegiatanArr []model.TRencanaKegiatan

	// db_go := os.Getenv("DB_HOST")

	conditions := map[string]interface{}{
		"c_tahun_ajaran": req.TahunAjaran,
	}
	if req.Gedung > 0 {
		conditions["c_id_gedung"] = req.Gedung
	}

	fmt.Println(conditions)
	database.DB = database.DB.Debug()

	query := database.DB.Limit(int(req.ResultsPerPage)).Offset(int(req.PageNumber)).Where(conditions).Find(&renKegiatanArr, "c_bulan_kegiatan_awal >= ? AND c_bulan_kegiatan_awal <= ?", req.BulanAwal, req.BulanAkhir)

	if query.Error != nil {
		return nil, status.Errorf(
			codes.DataLoss,
			query.Error.Error(),
		)
	}

	listRencanaKegiatan := make([]*kegiatan.RencanaKegiatan, len(renKegiatanArr))

	url := fmt.Sprintf("%s/api/v1/gedung/%d", config.Cfg.SvcDBGo, renKegiatanArr[0].CIDGedung)
	responseBody, err := utils.CallAPI(http.MethodGet, url, nil)
	var gedungResponse dbgo.GetGedungResponse
	json.Unmarshal(responseBody, &gedungResponse)
	if err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
	}

	for i, r := range renKegiatanArr {
		listRencanaKegiatan[i] = &kegiatan.RencanaKegiatan{
			KodeKegiatan: r.CKodeKegiatan,
			KodeAcuan:    r.CKodeAcuan,
			// NamaPenanda           : r.CNamaKegiatan ,
			NamaGedung: gedungResponse.Data.NamaGedung,
			// NamaKegiatanAcuan     : r. ,
			NamaKegiatan:       r.CNamaKegiatan,
			BulanKegiatanAwal:  r.CBulanKegiatanAwal,
			BulanKegiatanAkhir: r.CBulanKegiatanAkhir,
			TahunAjaran:        r.CTahunAjaran,
			// NamaBidang            : r.
			IdTingkat:       r.CIDTingkat,
			IdTujuanTingkat: r.CIDTujuanTingkat,
			TargetKehadiran: r.CTargetKehadiran,
			TargetCapaian:   r.CTargetCapaian,
			Status:          r.CStatus,
			// TglRencana            : r.
			// TglUsulBiayaAnggaran  : r.
			// TglUsulBiayaBarang    : r.
			// TglVerifKacabAnggaran : r.
			// TglVerifKacabBarang   : r.
			// TglVerifLog           : r.
			// TglVerifBid           : r.
			// StatusItemAnggaran    : r.
			// StatusBarang          : r.
			StatusBiaya: "null",
			// TotalBiayaAnggaran    : r.
			// TotalBiayaBarang      : r.
			// TotalAnggaran         : r.
			// TglUsulBiaya          : r.
			// TglVerifKacab         : r.
		}
	}

	return &kegiatan.RencanaKegiatanListRes{
		RencanaKegiatan: listRencanaKegiatan,
	}, nil
}

func (m *Kegiatan) CreateRencanaKegiatan(ctx context.Context, req *kegiatan.SetRencanaKegiatanRequest) (*kegiatan.SetRencanaKegiatanResponse, error) {

	var vStatus string
	if req.IsKirim {
		vStatus = "Kirim"
	} else {
		vStatus = "Create"
	}

	splitTahunAjaran := strings.Split(req.TahunAjaran, "/")
	splitKodeAcuan := strings.Split(req.KodeAcuan, ".")
	tempKodeKegiatan := splitTahunAjaran[0][len(splitTahunAjaran[0])-2:] + splitTahunAjaran[1][len(splitTahunAjaran[1])-2:]

	for i, part := range splitKodeAcuan {
		if part == "AK" {
			splitKodeAcuan[i] = "KG." + tempKodeKegiatan
		}
	}

	kodeAwal := strings.Join(splitKodeAcuan, ".")
	now := time.Now()
	bulanTanggal := now.Format("0102")
	preKodeKegiatan := fmt.Sprintf("%s.%d.%s", kodeAwal, req.IdGedung, bulanTanggal)
	KodeKegiatan, err := GetKodeKegiatan(preKodeKegiatan)

	if err != nil {
		return nil, status.Errorf(
			codes.DataLoss,
			"Gagal generate kode kegiatan",
		)
	}

	rencanaKegiatan := model.TRencanaKegiatan{
		CKodeKegiatan:       KodeKegiatan,
		CIDGedung:           req.IdGedung,
		CKodeAcuan:          req.KodeAcuan,
		CNamaKegiatan:       req.NamaKegiatan,
		CUpline:             KodeKegiatan,
		CBulanKegiatanAwal:  req.BulanAwal,
		CBulanKegiatanAkhir: req.BulanAkhir,
		CIDTingkat:          req.TingkatKelas,
		CIDTujuanTingkat:    req.TingkatKelasTarget,
		CTargetKehadiran:    req.TargetKehadiran,
		CTargetCapaian:      req.TargetCapaian,
		CTahunAjaran:        req.TahunAjaran,
		CStatus:             vStatus,
		CUpdater:            req.Updater,
	}

	query := database.DB.Create(&rencanaKegiatan)

	if query.Error != nil {
		fmt.Println("Error creating user:", query.Error)
	} else {
		fmt.Println("User created successfully:", rencanaKegiatan.CKodeKegiatan)
	}

	return &kegiatan.SetRencanaKegiatanResponse{
		Result:  true,
		Data:    rencanaKegiatan.CKodeKegiatan,
		Console: 20240420,
	}, nil
}

func GetKodeKegiatan(strSearch string) (string, error) {

	var newKode string

	database.DB = database.DB.Debug()

	query := database.DB.Raw(`
	SELECT
		COALESCE(
			(SELECT
				CONCAT($1, c_urut + 1) AS new
			FROM (
				SELECT
					CAST(SUBSTRING(c_kode_kegiatan, LENGTH($1) + 1, LENGTH(c_kode_kegiatan)) AS INT) AS c_urut
				FROM t_rencana_kegiatan
				WHERE c_kode_kegiatan LIKE $1
					AND c_kode_kegiatan = c_upline
				ORDER BY 1 DESC
				LIMIT 1
			) AS TB),
			$2
		) AS newkode
	`, strSearch+"%", strSearch+".1").Scan(&newKode)

	fmt.Println(newKode)
	if query.Error != nil {
		return "error", status.Errorf(
			codes.DataLoss,
			query.Error.Error(),
		)
	}

	return newKode, nil
}

func (m *Kegiatan) GetKegiatanPengikut(ctx context.Context, req *kegiatan.GetKegiatanPengikutRequest) (*kegiatan.GetKegiatanPengikutResponse, error) {

	return nil, nil
}

func (m *Kegiatan) GetKegiatanPengikutDetail(ctx context.Context, req *kegiatan.GetKegiatanPengikutDetailRequest) (*kegiatan.GetKegiatanPengikutDetailResponse, error) {

	if req.KodeKegiatan == "" {
		return nil, status.Errorf(
			codes.Canceled,
			"Mandatory Field Required",
		)
	}

	// var renKegiatanArr []model.TRencanaKegiatan
	var resultQuery []struct {
		KodeAcuan                         string `gorm:"column:c_kode_acuan"`
		KodeKegiatan                      string `gorm:"column:c_kode_kegiatan"`
		NamaKegiatan                      string `gorm:"column:c_nama_kegiatan"`
		BulanKegiatan                     string `gorm:"column:c_bulan_kegiatan_awal"`
		AcuanKegiatanCNama                string `gorm:"column:acuan_nama_kegiatan"`
		AcuanKegiatanTahunanCTanggalAwal  string `gorm:"column:c_tanggal_awal"`
		AcuanKegiatanTahunanCTanggalAkhir string `gorm:"column:c_tanggal_akhir"`
	}

	database.DB = database.DB.Debug()

	query := database.DB.Table(model.TableNameTRencanaKegiatan+" rk").
		Joins("Join "+model.TableNameTAcuanKegiatan+" ak ON ak.c_kode_acuan=rk.c_kode_acuan").
		Joins("Join "+model.TableNameTAcuanKegiatanTahunan+" akt ON akt.c_kode_acuan=rk.c_kode_acuan AND akt.c_tahun_ajaran = rk.c_tahun_ajaran").
		Where("rk.c_upline = ? AND rk.c_upline <> rk.c_kode_kegiatan", req.KodeKegiatan).
		Order("rk.c_bulan_kegiatan_awal ASC, rk.c_kode_kegiatan ASC, rk.c_nama_kegiatan ASC").
		Select("rk.*", "ak.c_nama_kegiatan as acuan_nama_kegiatan", "akt.c_tanggal_awal", "akt.c_tanggal_akhir").
		Scan(&resultQuery)

	if query.Error != nil {
		return nil, status.Errorf(
			codes.DataLoss,
			query.Error.Error(),
		)
	}

	resData := make([]*kegiatan.GetKegiatanPengikutDetailData, len(resultQuery))

	for i, r := range resultQuery {
		resData[i] = &kegiatan.GetKegiatanPengikutDetailData{
			KodeAcuan:         r.KodeAcuan,
			NamaKegiatanAcuan: r.KodeAcuan + " - " + r.AcuanKegiatanCNama,
			KodeKegiatan:      r.KodeKegiatan,
			NamaKegiatan:      r.NamaKegiatan,
			KodeNamaKegiatan:  r.KodeKegiatan + " - " + r.NamaKegiatan,
			BulanKegiatan:     r.BulanKegiatan,
			TanggalAwal:       r.AcuanKegiatanTahunanCTanggalAwal,
			TanggalAkhir:      r.AcuanKegiatanTahunanCTanggalAkhir,
		}
	}

	return &kegiatan.GetKegiatanPengikutDetailResponse{
		Result: true,
		Data:   resData,
	}, nil

}

func (m *Kegiatan) GetKepanitian(ctx context.Context, req *kegiatan.GetKepanitiaanRequest) (*kegiatan.GetKepanitiaanResponse, error) {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	var resultQuery []struct {
		KodeKegiatan string `gorm:"column:c_kode_kegiatan"`
		KodePanitia  string `gorm:"column:c_kode_panitia"`
		NamaPanitia  string `gorm:"column:c_nama_panitia"`
		NIK          string `gorm:"column:c_nik"`
	}

	database.DB = database.DB.Debug()

	query := database.DB.Table(model.TableNameTRencanaKegiatanPanitium+" rkp").
		Joins("Join "+model.TableNameTNamaPanitium+" np ON rkp.c_kode_panitia=np.c_kode_panitia").
		Where("rkp.c_kode_kegiatan", req.KodeKegiatan).
		Order("np.c_nama_panitia ASC").
		Select("rkp.*", "np.c_nama_panitia").
		Scan(&resultQuery)

	if query.Error != nil {
		return nil, status.Errorf(
			codes.DataLoss,
			query.Error.Error(),
		)
	}

	countResultRow := len(resultQuery)
	resData := make([]*kegiatan.GetKepanitiaanData, len(resultQuery))
	result := true

	if countResultRow < 1 {
		return &kegiatan.GetKepanitiaanResponse{
			Result: false,
			Data:   nil,
		}, nil
	}

	ch := make(chan dbgo.KaryawanDetailInfo, countResultRow)
	var wg sync.WaitGroup
	wg.Add(countResultRow)

	for i, r := range resultQuery {
		go worker.FetchKaryawan(ctx, i, r.NIK, ch, &wg)
	}

	wg.Wait()
	karyawanDetail := make([]*dbgo.KaryawanDetailInfo, countResultRow)

	for i := 0; i < countResultRow; i++ {
		KaryawanDetailInfo := <-ch
		karyawanDetail[KaryawanDetailInfo.IDX] = &KaryawanDetailInfo
	}
	close(ch)
	println(karyawanDetail)

	for i, r := range resultQuery {
		resData[i] = &kegiatan.GetKepanitiaanData{
			KodeKegiatan: r.KodeKegiatan,
			KodePanitia:  r.KodePanitia,
			NamaPanitia:  r.NamaPanitia,
			Nik:          r.NIK,
			NamaLengkap:  karyawanDetail[i].Data.CNamaLengkap,
			NamaKaryawan: r.NIK + " - " + karyawanDetail[i].Data.CNamaLengkap,
		}

	}

	return &kegiatan.GetKepanitiaanResponse{
		Result: result,
		Data:   resData,
	}, nil
}

func (m *Kegiatan) GetKegiatanSekolah(ctx context.Context, req *kegiatan.GetKegiatanSekolahRequest) (*kegiatan.GetKegiatanSekolahResponse, error) {

	var resultQuery []struct {
		KodeKegiatan    string `gorm:"column:c_kode_kegiatan"`
		IdSekolah       int32  `gorm:"column:c_id_sekolah"`
		TargetKehadiran int32  `gorm:"column:c_target_kehadiran"`
		TargetCapaian   int32  `gorm:"column:c_target_capaian"`
	}

	database.DB = database.DB.Debug()

	query := database.DB.Model(model.TRencanaKegiatanSekolah{}).
		Where("c_kode_kegiatan = ?", req.KodeKegiatan).
		Scan(&resultQuery)

	if query.Error != nil {
		return nil, status.Errorf(
			codes.DataLoss,
			query.Error.Error(),
		)
	}

	countResultRow := len(resultQuery)
	resData := make([]*kegiatan.GetKegiatanSekolahData, len(resultQuery))
	result := true

	if countResultRow < 1 {
		return &kegiatan.GetKegiatanSekolahResponse{
			Result: false,
			Data:   nil,
		}, nil
	}

	var ids string
	for i, sekolah := range resultQuery {
		ids += fmt.Sprintf("%d", sekolah.IdSekolah)
		if i != countResultRow-1 {
			ids += ","
		}
	}

	url := fmt.Sprintf("%s/api/v1/sekolah/data-sekolah?list_id_sekolah=%s", config.Cfg.SvcDBSekolah, ids)

	responseBody, err := utils.CallAPI(http.MethodGet, url, nil)
	if err != nil {
		fmt.Println("Error calling API get list Sekolah:", err)
	}

	var sekolahResponse []sekolah.GetListSekolah
	json.Unmarshal(responseBody, &sekolahResponse)
	if err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
	}

	sekolahMap := make(map[int]sekolah.GetListSekolah)
	for _, sekolah := range sekolahResponse {
		sekolahMap[sekolah.IdSekolah] = sekolah
	}

	for i, r := range resultQuery {
		if sekolah, ok := sekolahMap[int(r.IdSekolah)]; ok {
			resData[i] = &kegiatan.GetKegiatanSekolahData{
				KodeKegiatan:    r.KodeKegiatan,
				IdSekolah:       r.IdSekolah,
				NamaSekolah:     sekolah.NamaSekolah,
				TargetKehadiran: r.TargetKehadiran,
				TargetCapaian:   r.TargetCapaian,
			}
		}

	}

	return &kegiatan.GetKegiatanSekolahResponse{
		Result: result,
		Data:   resData,
	}, nil
}

func (m *Kegiatan) GetAcuanKegiatan(ctx context.Context, req *kegiatan.GetAcuanKegiatanRequest) (*kegiatan.GetAcuanKegiatanResponse, error) {

	var resultQuery []struct {
		KodeAcuan    string `gorm:"column:c_kode_acuan"`
		NamaKegiatan string `gorm:"column:c_nama_kegiatan"`
		TanggalAwal  string `gorm:"column:c_tanggal_awal"`
		TanggalAkhir string `gorm:"column:c_tanggal_akhir"`
		IsPromosi    int32  `gorm:"column:c_is_promosi"`
		IsPusat      int32  `gorm:"column:c_is_pusat"`
	}

	ispusat := "0"
	if req.IdKomarCakupan == 1 {
		ispusat = "1"
		if req.IdBidang == 12 || req.IdBidang == 16 {
			ispusat = "0,1"
		}
	}

	query := database.DB.Table(model.TableNameTAcuanKegiatanTahunan+" a").
		Joins("Join "+model.TableNameTAcuanKegiatan+" b ON b.c_kode_acuan=a.c_kode_acuan").
		Where("a.c_tahun_ajaran = ?", req.TahunAjaran).
		Where("b.c_bidang_penanggung_jawab = ?", req.Bidang).
		Where("b.c_is_pusat in (?)", ispusat).
		Where("b.c_upline = b.kode_acuan").
		Where("b.c_Status = 'Aktif'").
		Order("a.c_kode_acuan ASC, b.c_nama_kegiatan").
		Select("a.*", "b.c_nama_kegiatan").Scan(&resultQuery)

	if query.Error != nil {
		return nil, status.Errorf(
			codes.DataLoss,
			query.Error.Error(),
		)
	}

	countResultRow := len(resultQuery)
	resData := make([]*kegiatan.GetAcuanKegiatanData, len(resultQuery))
	result := true

	if countResultRow < 1 {
		return &kegiatan.GetAcuanKegiatanResponse{
			Result: false,
			Data:   nil,
		}, nil
	}

	for i, r := range resultQuery {
		resData[i] = &kegiatan.GetAcuanKegiatanData{
			KodeAcuan:         r.KodeAcuan,
			NamaKegiatan:      r.NamaKegiatan,
			NamaKegiatanAcuan: r.KodeAcuan + " - " + r.NamaKegiatan,
			TanggalAwal:       r.TanggalAwal,
			TanggalAkhir:      r.TanggalAkhir,
			IsPromosi:         r.IsPromosi,
			IsPusat:           r.IsPusat,
			KomarLogin:        req.IdKomarCakupan,
			BidLogin:          req.IdBidang,
		}
	}

	return &kegiatan.GetAcuanKegiatanResponse{
		Result: result,
		Data:   resData,
	}, nil
}

func (m *Kegiatan) GetRencanaKegiatan(ctx context.Context, req *kegiatan.GetRencanaKegiatanRequest) (*kegiatan.GetRencanaKegiatanResponse, error) {

	var resultQuery []struct {
		KodeKegiatan       string `gorm:"column:c_kode_kegiatan"`
		KodeAcuan          string `gorm:"column:c_kode_acuan"`
		IdGedung           int32  `gorm:"column:c_id_gedung"`
		NamaKegiatan       string `gorm:"column:c_nama_kegiatan"`
		BulanKegiatanAwal  string `gorm:"column:c_bulan_kegiatan_awal"`
		BulanKegiatanAkhir string `gorm:"column:c_bulan_kegiatan_akhir"`
		IdTingkat          int32  `gorm:"column:c_id_tingkat"`
		IdTujuanTingkat    int32  `gorm:"column:c_id_tujuan_tingkat"`
		TargetKehadiran    int32  `gorm:"column:c_target_kehadiran"`
		TargetCapaian      int32  `gorm:"column:c_target_capaian"`
		Status             string `gorm:"column:c_status"`
	}

	query := database.DB.Model(model.TRencanaKegiatan{}).
		Where("c_id_gedung = ? AND c_tahun_ajaran = ? AND c_kode_acuan = ? AND c_upline = c_kode_kegiatan", req.Gedung, req.TahunAjaran, req.KodeAcuan).
		Find(&resultQuery)

	if query.Error != nil {
		return nil, status.Errorf(
			codes.DataLoss,
			query.Error.Error(),
		)
	}

	countResultRow := len(resultQuery)
	resData := make([]*kegiatan.GetRencanaKegiatanData, len(resultQuery))
	result := true

	if countResultRow < 1 {
		return &kegiatan.GetRencanaKegiatanResponse{
			Result: false,
			Data:   nil,
		}, nil
	}

	for i, r := range resultQuery {
		resData[i] = &kegiatan.GetRencanaKegiatanData{
			KodeAcuan:          r.KodeAcuan,
			NamaKegiatan:       r.NamaKegiatan,
			KodeNamaKegiatan:   r.KodeKegiatan + " - " + r.NamaKegiatan,
			KodeKegiatan:       r.KodeKegiatan,
			IdGedung:           r.IdGedung,
			BulanKegiatanAwal:  r.BulanKegiatanAwal,
			BulanKegiatanAkhir: r.BulanKegiatanAkhir,
			IdTingkat:          r.IdTingkat,
			IdTujuanTingkat:    r.IdTujuanTingkat,
			TargetKehadiran:    r.TargetKehadiran,
			TargetCapaian:      r.TargetCapaian,
			Status:             r.Status,
		}
	}

	return &kegiatan.GetRencanaKegiatanResponse{
		Result: result,
		Data:   resData,
	}, nil
}

func (m *Kegiatan) GetTargetSiswa(ctx context.Context, req *kegiatan.GetTargetSiswaRequest) (*kegiatan.GetTargetSiswaResponse, error) {

	if req.TahunAjaran == "" {
		return nil, status.Errorf(
			codes.Canceled,
			"Mandatory Field Required : tahun_ajaran with format '{YYYY}/{YYYY+1year}'",
		)
	}

	data, err := sekolah.FetchTahunAjaran(req.TahunAjaran)
	if err != nil {
		fmt.Println(err)
	}

	result := true
	ch := make(chan sekolah.ChannelTargetSiswa, len(data))
	var wg sync.WaitGroup
	wg.Add(len(data))

	for i, r := range data {
		go worker.FetchTargetSiswa(ctx, i, req.Gedung, r.TahunAjaran, ch, &wg)
	}
	wg.Wait()

	targetSiswa := make([]sekolah.GetTargetSiswaData, 0)
	for i := 0; i < len(data); i++ {
		fmt.Println(i)
		info := <-ch
		targetSiswa = append(targetSiswa, info.Data...)
	}

	close(ch)

	resData := make([]*kegiatan.GetTargetSiswaData, len(targetSiswa))
	for i, r := range targetSiswa {
		resData[i] = &kegiatan.GetTargetSiswaData{
			IdGedung:     r.IdGedung,
			TahunAjaran:  r.TahunAjaran,
			IdSekolah:    r.IdSekolah,
			NamaSekolah:  r.NamaSekolah,
			IdTingkat:    r.IdTingkat,
			TingkatKelas: r.TingkatKelas,
			JmlTarget:    r.JmlTarget,
		}
	}

	return &kegiatan.GetTargetSiswaResponse{
		Result: result,
		Data:   resData,
	}, nil
}

func (m *Kegiatan) GetBarang(ctx context.Context, req *kegiatan.GetBarangRequest) (*kegiatan.GetBarangResponse, error) {
	if req.KodeKegiatan == "" {
		return nil, status.Errorf(
			codes.Canceled,
			"Mandatory Field Required",
		)
	}

	var resultQuery []struct {
		KodeBarang  string `gorm:"column:c_kode_barang"`
		NamaBarang  string `gorm:"column:c_nama_barang"`
		Acc         string `gorm:"column:c_acc"`
		IDVendor    int32  `gorm:"column:c_id_vendor"`
		NamaVendor  string `gorm:"column:c_nama_vendor"`
		Dokumen     string `gorm:"column:c_dokumen"`
		IsPusat     int32  `gorm:"column:c_is_pusat"`
		Status      string `gorm:"column:c_status"`
		Jumlah      int32  `gorm:"column:c_jumlah"`
		NilaiSatuan int32  `gorm:"column:c_nilai_satuan"`
		Total       int32  `gorm:"column:c_total"`
		IdSatuan    int32  `gorm:"column:c_id_satuan"`
		Satuan      string `gorm:"column:c_satuan"`
	}

	query := database.DB.Table(model.TableNameTKebutuhanBarang+" a").
		Select("a.c_kode_barang", "b.c_nama_barang", "a.c_acc", "coalesce(concat(c.c_nama_vendor, ' (', c.c_no_telepon, ')'), '-') AS c_nama_vendor",
			"coalesce(a.c_dokumen, '-') AS c_dokumen", "a.c_is_pusat", "a.c_status", "a.c_id_satuan", "d.c_satuan").
		Joins("JOIN "+model.TableNameTNamaBarang+" b ON a.c_kode_barang = b.c_kode_barang").
		Joins("LEFT JOIN "+model.TableNameTVendor+" c ON a.c_id_vendor = c.c_id_vendor").
		Joins("LEFT JOIN "+model.TableNameTSatuan+" d ON a.c_id_satuan = d.c_id_satuan").
		Where("a.c_kode_kegiatan", req.KodeKegiatan).Order("b.c_nama_barang").Scan(&resultQuery)

	if query.Error != nil {
		return nil, status.Errorf(
			codes.DataLoss,
			query.Error.Error(),
		)
	}

	// var CAcc = make(map[string]interface{}, 3)
	var CAcc DataModel.TKebutuhanBarangCAcc
	resData := make([]*kegiatan.GetBarangData, len(resultQuery))
	for i, r := range resultQuery {
		json.Unmarshal([]byte(r.Acc), &CAcc)

		var jumlah int32
		var total int32
		var nilai_satuan int32
		fmt.Println(r.Status)
		if r.Status == "Create" || r.Status == "Kirim" {
			jumlah = r.Jumlah
			total = r.Total
			nilai_satuan = r.NilaiSatuan
		} else if r.Status == "Setuju" {
			jumlah = CAcc.CSetujuKacab.CJumlah
			total = CAcc.CSetujuKacab.CTotal
			nilai_satuan = CAcc.CSetujuKacab.CNilaiSatuan
		} else if r.Status == "Bidang" {
			jumlah = CAcc.CSetujuBidang.CJumlah
			total = CAcc.CSetujuBidang.CTotal
			nilai_satuan = CAcc.CSetujuBidang.CNilaiSatuan
		} else if r.Status == "Terima" {
			jumlah = CAcc.CSetujuAkuntansi.CJumlah
			total = CAcc.CSetujuAkuntansi.CTotal
			nilai_satuan = CAcc.CSetujuAkuntansi.CNilaiSatuan
		}
		fmt.Println(r.Acc)
		resData[i] = &kegiatan.GetBarangData{
			Kode:        r.KodeBarang,
			KodeBarang:  r.KodeBarang,
			NamaBarang:  r.NamaBarang,
			IdVendor:    r.IDVendor,
			NamaVendor:  r.NamaVendor,
			IsPusat:     r.IsPusat,
			Dokumen:     r.Dokumen,
			Status:      r.Status,
			Jumlah:      jumlah,
			NilaiSatuan: nilai_satuan,
			IdSatuan:    r.IdSatuan,
			Satuan:      r.Satuan,
			Total:       total,
			TotalAcc:    total,
			Split:       "-",
		}
	}

	return &kegiatan.GetBarangResponse{
		Result: true,
		Data:   resData,
	}, nil
}
