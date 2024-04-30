package service

import (
	"context"
	"errors"
	"go-bponline/m/gen/model"
	"go-bponline/m/internal/database"
	"go-bponline/m/pb/anggaran"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AnggaranService struct {
	anggaran.UnimplementedAnggaranServiceServer
}

func (s *AnggaranService) GetKegiatan(ctx context.Context, req *anggaran.KegiatanRequest) (*anggaran.KegiatanResponse, error) {
	if req.Pagination.PageNumber == 0 {
		req.Pagination.PageNumber = 1
	}
	
	if req.Pagination.ResultsPerPage == 0 {
		req.Pagination.ResultsPerPage = 10
	}
	
	offset := (req.Pagination.PageNumber - 1) * req.Pagination.ResultsPerPage
	
	var results []struct {
		CKodeKegiatan          string `json:"c_kode_kegiatan"`
		CKodeAcuan             string `json:"c_kode_acuan"`
		CNamaKegiatan          string `json:"c_nama_kegiatan"`
		CTargetKehadiran       int32  `json:"c_target_kehadiran"`
		CTargetCapaian         int32  `json:"c_target_capaian"`
		CKodeKepanitiaan       string `json:"c_kode_kepanitiaan"`
		CKodeKelAnggaran       string `json:"c_kode_kel_anggaran"`
		CKodeKelBarang         string `json:"c_kode_kel_barang"`
		IsKegiatan             int32  `json:"is_kegiatan"`
		CBidangPenanggungJawab string `json:"c_bidang_penanggung_jawab"`
		CTahunAjaran           string `json:"c_tahun_ajaran"`
	}
	
	query := database.DB.WithContext(ctx).Model(model.TRencanaKegiatan{}).
		Select("*").
		Where("t_rencana_kegiatan.c_id_gedung = ?", req.Gedung).
		Where("ac.c_bidang_penanggung_jawab = ?", req.Bidang).
		Where("t_rencana_kegiatan.c_tahun_ajaran = ?", req.TahunAjaran).
		Joins("left join t_acuan_kegiatan as ac on t_rencana_kegiatan.c_kode_acuan = ac.c_kode_acuan").
		Limit(int(req.Pagination.ResultsPerPage)).
		Offset(int(offset)).
		Scan(&results)
	
	if query.Error != nil {
		return nil, status.Errorf(
			codes.DataLoss,
			query.Error.Error(),
		)
	}
	
	if query.RowsAffected == 0 {
		return &anggaran.KegiatanResponse{Data: nil}, nil
	}
	
	var kegiatanRes []*anggaran.Kegiatan
	
	for _, value := range results {
		kegiatan := &anggaran.Kegiatan{
			KodeKegiatan:    value.CKodeKegiatan,
			KodeAcuan:       value.CKodeAcuan,
			NamaKegiatan:    value.CNamaKegiatan,
			TargetKehadiran: value.CTargetKehadiran,
			TargetCapaian:   value.CTargetCapaian,
			KodeKepanitiaan: value.CKodeKepanitiaan,
			KodeKelAnggaran: value.CKodeKelAnggaran,
			KodeKelBarang:   value.CKodeKelBarang,
			TahunAjaran:     value.CTahunAjaran,
			IsKegiatan:      true,
		}
		
		kegiatanRes = append(kegiatanRes, kegiatan)
	}
	
	var anggaranResponse = &anggaran.KegiatanResponse{Data: kegiatanRes}
	
	return anggaranResponse, nil
}

func (s *AnggaranService) GetAcuanHargadanStatus(ctx context.Context, req *anggaran.AcuanHargaDanStatusRequest) (*anggaran.AcuanHargaDanStatusResponse, error) {
	
	var acuanHargaAnggaranChannel = make(chan int32)
	defer close(acuanHargaAnggaranChannel)
	
	var acuanHargaBarangChannel = make(chan int32)
	defer close(acuanHargaBarangChannel)
	
	var statusAnggaranChannel = make(chan string)
	defer close(statusAnggaranChannel)
	
	var statusBarangChannel = make(chan string)
	defer close(statusBarangChannel)
	
	go func() {
		var acuanHargaAnggaranArr []model.TAcuanHargaAnggaran
		
		var totalAcuanAnggaran struct {
			TotalAcuanAnggaran int32 `json:"total_acuan_anggaran"`
		}
		
		database.DB.WithContext(ctx).Model(&acuanHargaAnggaranArr).
			Select("sum(c_acuan_harga) as total_acuan_anggaran").
			Where("c_id_gedung = ?", req.Gedung).
			Where("c_kode_acuan = ?", req.KodeAcuan).
			Scan(&totalAcuanAnggaran)
		
		acuanHargaAnggaranChannel <- totalAcuanAnggaran.TotalAcuanAnggaran
	}()
	
	go func() {
		var acuanHargaBarangArr []model.TAcuanHargaBarang
		
		var totalAcuanHargaBarang struct {
			TotalAcuanBarang int32 `json:"total_acuan_barang"`
		}
		
		database.DB.WithContext(ctx).
			Select("sum(t_acuan_harga_barang.c_acuan_harga) as total_acuan_barang").
			Model(&acuanHargaBarangArr).
			Where("brg.c_kode_kel_barang = ?", req.KodeKelBarang).
			Where("t_acuan_harga_barang.c_id_kota = ?", req.Kota).
			Joins("left join t_isi_i_kel_barang as brg  on t_acuan_harga_barang.c_kode_barang = brg.c_kode_barang").
			Scan(&totalAcuanHargaBarang)
		
		acuanHargaBarangChannel <- totalAcuanHargaBarang.TotalAcuanBarang
	}()
	
	go func() {
		var kebutuhanAnggaran model.TKebutuhanAnggaran
		var statusKebutuhanAnggaran struct {
			CStatus string `json:"c_status"`
		}
		
		database.DB.WithContext(ctx).Select("c_status").
			Model(&kebutuhanAnggaran).
			Where("c_kode_kegiatan = ?", req.KodeKegiatan).
			Scan(&statusKebutuhanAnggaran)
		
		switch statusKebutuhanAnggaran.CStatus {
		case "":
			statusAnggaranChannel <- "-"
		default:
			statusAnggaranChannel <- statusKebutuhanAnggaran.CStatus
		}
		
	}()
	
	go func() {
		var kebutuhanBarang model.TKebutuhanBarang
		var statusKebutuhanBarang struct {
			CStatus string `json:"c_status"`
		}
		
		database.DB.WithContext(ctx).Select("c_status").
			Model(&kebutuhanBarang).
			Where("c_kode_kegiatan = ?", req.KodeKegiatan).
			Scan(&statusKebutuhanBarang)
		
		switch statusKebutuhanBarang.CStatus {
		case "":
			statusBarangChannel <- "-"
		default:
			statusBarangChannel <- statusKebutuhanBarang.CStatus
		}
		
	}()
	
	return &anggaran.AcuanHargaDanStatusResponse{
		Data: &anggaran.AcuanHargaDanStatus{
			TotalAcuanAnggaran: <-acuanHargaAnggaranChannel,
			TotoalAcuanBarang:  <-acuanHargaBarangChannel,
			CStatusAnggaran:    <-statusAnggaranChannel,
			CStatusBarang:      <-statusBarangChannel,
		},
	}, nil
}

func (s *AnggaranService) Getkegiatanikut(ctx context.Context, req *anggaran.KegiatanIkutRequest) (*anggaran.KegiatanIkutResponse, error) {
	var kegiatanIkutArr []*anggaran.KegiatanIkut
	
	var rencanaKegiatanArr []model.TRencanaKegiatan
	
	var databaseScan []struct {
		CKodeKegiatan      string `json:"c_kode_kegiatan"`
		CNamaKegiatan      string `json:"c_nama_kegiatan"`
		CBulanKegiatanAwal string `json:"c_bulan_kegiatan_awal"`
		CKodeKepanitiaan   string `json:"c_kode_kepanitiaan"`
		JumlahPanitia      int32  `json:"jumlah_panitia"`
	}
	
	query := database.DB.WithContext(ctx).Model(&rencanaKegiatanArr).
		Select("t_rencana_kegiatan.c_kode_kegiatan, t_rencana_kegiatan.c_nama_kegiatan, t_rencana_kegiatan.c_bulan_kegiatan_awal, COALESCE(b.c_kode_kepanitiaan, '-') as c_kode_kepanitiaan, (select count(*) from t_rencana_kegiatan_panitia c where c.c_kode_kegiatan = t_rencana_kegiatan.c_kode_kegiatan) as jumlah_panitia").
		Where("t_rencana_kegiatan.c_upline = ?", req.KodeKegiatan).
		Where("t_rencana_kegiatan.c_kode_kegiatan != ?", req.KodeKegiatan).
		Joins("left join t_acuan_kegiatan as b on b.c_kode_acuan = t_rencana_kegiatan.c_kode_acuan ").
		Scan(&databaseScan)
	
	if query.Error != nil {
		return nil, errors.New("error database : " + query.Error.Error())
	}
	
	if query.RowsAffected == 0 {
		return &anggaran.KegiatanIkutResponse{Data: nil}, nil
	}
	
	for _, v := range databaseScan {
		kegiatanIkutArr = append(kegiatanIkutArr, &anggaran.KegiatanIkut{
			KodeKegiatan:    v.CKodeKegiatan,
			NamaKegiatan:    v.CNamaKegiatan,
			BulanKegiatan:   v.CBulanKegiatanAwal,
			KodeKepanitiaan: v.CKodeKepanitiaan,
			JumlahPanitia:   v.JumlahPanitia,
		})
	}
	
	var kegiatanIkutResponse = &anggaran.KegiatanIkutResponse{Data: kegiatanIkutArr}
	return kegiatanIkutResponse, nil
}

func (s *AnggaranService) Getitemanggaran(context.Context, *anggaran.ItemAnggaranRequest) (*anggaran.ItemAnggaranResponse, error) {
	var itemAnggaranArr []*anggaran.ItemAnggaran
	
	itemAnggaran1 := &anggaran.ItemAnggaran{
		KodeKelAnggaran:  "2",
		Kode:             "32",
		KodeItemAnggaran: "21",
		NamaItemAnggaran: "OKE",
		AcuanHarga:       3,
		AcuanHargaMax:    3,
		Jumlah:           3,
		NilaiSatuan:      3,
		IdSatuan:         3,
		Satuan:           "223",
		Total:            2323,
		IdVendor:         23,
		NamaVendor:       "eff",
		Dokumen:          "fff",
		Pajak:            2,
		PajakOpsi:        23,
		ListPajak:        "ff",
		Split:            "wqwe",
	}
	
	itemAnggaranArr = append(itemAnggaranArr, itemAnggaran1)
	
	var itemAnggaranResponse = &anggaran.ItemAnggaranResponse{Data: itemAnggaranArr}
	return itemAnggaranResponse, nil
}
