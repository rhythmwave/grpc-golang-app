package sekolah

import (
	"encoding/json"
	"fmt"
	"go-bponline/m/internal/utils"
	"net/http"
	"os"
	"sort"
)

type MetaResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type GetTahunAjaranData struct {
	ID                string `json:"id"`
	TahunAjaran       string `json:"tahun_ajaran"`
	IsCurrentlyActive bool   `json:"is_currently_active"`
	Status            string `json:"status"`
	Awal              string `json:"awal"`
	Akhir             string `json:"akhir"`
}

type GetTahunAjaranResponse struct {
	Data []GetTahunAjaranData `json:"data"`
	Meta MetaResponse         `json:"_meta"`
}

type lessFunc func(p1, p2 *GetTahunAjaranData) bool

// multiSorter implements the Sort interface, sorting the changes within.
type multiSorter struct {
	changes []GetTahunAjaranData
	less    []lessFunc
}

// Sort sorts the argument slice according to the less functions passed to OrderedBy.
func (ms *multiSorter) Sort(changes []GetTahunAjaranData) {
	ms.changes = changes
	sort.Sort(ms)
}

// OrderedBy returns a Sorter that sorts using the less functions, in order.
// Call its Sort method to sort the data.
func OrderedBy(less ...lessFunc) *multiSorter {
	return &multiSorter{
		less: less,
	}
}

// Len is part of sort.Interface.
func (ms *multiSorter) Len() int {
	return len(ms.changes)
}

// Swap is part of sort.Interface.
func (ms *multiSorter) Swap(i, j int) {
	ms.changes[i], ms.changes[j] = ms.changes[j], ms.changes[i]
}

func (ms *multiSorter) Less(i, j int) bool {
	p, q := &ms.changes[i], &ms.changes[j]
	// Try all but the last comparison.
	var k int
	for k = 0; k < len(ms.less)-1; k++ {
		less := ms.less[k]
		switch {
		case less(p, q):
			// p < q, so we have a decision.
			return true
		case less(q, p):
			// p > q, so we have a decision.
			return false
		}
		// p == q; try the next comparison.
	}
	// All comparisons to here said "equal", so just return whatever
	// the final comparison reports.
	return ms.less[k](p, q)
}

func FetchTahunAjaran(tahun_ajaran string) ([]GetTahunAjaranData, error) {

	descTahunAjaran := func(c1, c2 *GetTahunAjaranData) bool {
		return c1.TahunAjaran > c2.TahunAjaran // Note: > orders downwards.
	}

	endpoint := "/api/v1/tahun-ajaran"
	url := fmt.Sprintf("%s%s", os.Getenv("SERVICE_DB_SEKOLAH"), endpoint)

	responseBody, err := utils.CallAPI(http.MethodGet, url, nil)
	if err != nil {
		fmt.Println("Error calling API:", endpoint, err)
		return nil, err // Or send an error message on the channel
	}

	var tahunAjaran GetTahunAjaranResponse
	err = json.Unmarshal(responseBody, &tahunAjaran)
	if err != nil {
		fmt.Println("Error unmarshaling JSON:", err, " from tahun-ajaran")
	}
	fmt.Println(tahunAjaran.Data)
	if len(tahunAjaran.Data) < 1 {
		fmt.Println("API: fetch service DB GO /api/v1/tahun-ajaran - Empty Data")
		return nil, err
	}
	var filteredData []GetTahunAjaranData
	for _, element := range tahunAjaran.Data {
		if utils.IsYearRangeLessOrEqual(element.TahunAjaran, tahun_ajaran) { // Filter based on tahun_ajaran
			filteredData = append(filteredData, element)
		}
	}

	OrderedBy(descTahunAjaran).Sort(filteredData)

	return filteredData[:3], nil
}
