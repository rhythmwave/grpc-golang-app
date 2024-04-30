package worker

import (
	"context"
	"encoding/json"
	"fmt"
	"go-bponline/m/config"
	"go-bponline/m/gen/model"
	"go-bponline/m/internal/database"
	"go-bponline/m/internal/service/dbgo"
	"go-bponline/m/internal/service/sekolah"
	"go-bponline/m/internal/utils"
	"net/http"
	"sync"
)

func FetchKaryawan(ctx context.Context, index int, nik string, ch chan dbgo.KaryawanDetailInfo, wg *sync.WaitGroup) {
	defer wg.Done()
	var karyawanInfo dbgo.KaryawanDetailInfo
	karyawanInfo.IDX = index
	url := fmt.Sprintf("%s/api/v1/karyawan/detail/%s", config.Cfg.SvcDBGo, nik)

	responseBody, err := utils.CallAPI(http.MethodGet, url, nil)
	if err != nil {
		fmt.Println("Error calling API for nik:", nik, err)

		ch <- karyawanInfo
		return
	}
	err = json.Unmarshal(responseBody, &karyawanInfo)

	if err != nil {
		fmt.Println("Error unmarshaling JSON:", err, " from nik:", nik)
	}

	ch <- karyawanInfo
}

// func FetchTargetSiswa(index int, id_gedung int32, tahun_ajaran string) {
func FetchTargetSiswa(ctx context.Context, index int, id_gedung int32, tahun_ajaran string, ch chan sekolah.ChannelTargetSiswa, wg *sync.WaitGroup) {

	defer wg.Done()
	var targetSiswa sekolah.ChannelTargetSiswa

	targetSiswa.IDX = uint16(index)

	//query target sekolah
	var resultQuery []struct {
		IdGedung    int32  `gorm:"column:c_id_gedung"`
		TahunAjaran string `gorm:"column:c_tahun_ajaran"`
		IdSekolah   int32  `gorm:"column:c_id_sekolah"`
		IdTingkat   int32  `gorm:"column:c_id_tingkat"`
		JmlTarget   int32  `gorm:"column:c_jml_target"`
	}

	query := database.DB.Model(model.TTargetSekolah{}).
		Where("c_id_gedung = ? AND c_tahun_ajaran = ?", id_gedung, tahun_ajaran).
		Select("SUM( c_target_siswa ) c_jml_target, c_id_gedung, c_tahun_ajaran, c_id_sekolah, c_id_tingkat").
		Group("c_id_gedung, c_tahun_ajaran, c_id_sekolah, c_id_tingkat").
		Scan(&resultQuery)

	if query.Error != nil {

		ch <- targetSiswa
		return
	}

	url := fmt.Sprintf("%s/api/v1/sekolah-kelas", config.Cfg.SvcDBSekolah)
	responseBody, err := utils.CallAPI(http.MethodGet, url, nil)

	if err != nil {
		fmt.Println("Error calling API sekolah-kelas:")

		ch <- targetSiswa
		return
	}

	var sekolahKelasResponse sekolah.GetSekolahListResponse
	err = json.Unmarshal(responseBody, &sekolahKelasResponse)
	if err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
	}

	sekolahKelasMap := make(map[int]sekolah.SekolahKelasList)
	for _, sekolah := range sekolahKelasResponse.Data {
		sekolahKelasMap[sekolah.ID] = sekolah
	}

	var ids string
	for i, sekolah := range resultQuery {
		ids += fmt.Sprintf("%d", sekolah.IdSekolah)
		if i != len(resultQuery)-1 {
			ids += ","
		}
	}

	urlSekolah := fmt.Sprintf("%s/api/v1/sekolah/data-sekolah?list_id_sekolah=%s", config.Cfg.SvcDBSekolah, ids)
	// println(url)
	sekolahBody, err := utils.CallAPI(http.MethodGet, urlSekolah, nil)
	if err != nil {
		fmt.Println("Error calling API get list Sekolah:", err)
		ch <- targetSiswa
		return
	}

	var sekolahResponse []sekolah.GetListSekolah
	json.Unmarshal(sekolahBody, &sekolahResponse)
	if err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
	}

	sekolahMap := make(map[int]sekolah.GetListSekolah)
	for _, sekolah := range sekolahResponse {
		sekolahMap[sekolah.IdSekolah] = sekolah
	}

	resData := make([]*sekolah.GetTargetSiswaData, len(resultQuery))

	for i, r := range resultQuery {

		namaSekolah := ""
		tingkatKelas := ""

		if sekolah, ok := sekolahMap[int(r.IdSekolah)]; ok {
			namaSekolah = sekolah.NamaSekolah
		}
		if sekolahKelas, ok := sekolahKelasMap[int(r.IdTingkat)]; ok {
			tingkatKelas = sekolahKelas.TingkatKelas
		}
		resData[i] = &sekolah.GetTargetSiswaData{
			IdGedung:     r.IdGedung,
			TahunAjaran:  r.TahunAjaran,
			IdSekolah:    r.IdSekolah,
			NamaSekolah:  namaSekolah,
			TingkatKelas: tingkatKelas,
			IdTingkat:    r.IdTingkat,
			JmlTarget:    r.JmlTarget,
		}

		targetSiswa.Data = append(targetSiswa.Data, *resData[i])
	}

	ch <- targetSiswa

	// return

}
