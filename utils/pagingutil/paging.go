package pagingutil

import "math"

type Paging struct {
	Pages      []int `json:"pages"`
	TotalPages int   `json:"total_pages"`
	FirstPage  int   `json:"first_page"`
	LastPage   int   `json:"last_page"`
	CurrPage   int   `json:"curr_page"`
}

func Paginator(page, pageLimit int, nums int64) *Paging {
	var firstPage int                                                //前一页地址
	var lastPage int                                                 //后一页地址
	totalPages := int(math.Ceil(float64(nums) / float64(pageLimit))) //page总数
	if page > totalPages {
		page = totalPages
	}
	if page <= 0 {
		page = 1
	}
	pages := make([]int, totalPages)
	for i := range pages {
		pages[i] = i + 1
	}
	firstPage = int(math.Max(float64(1), float64(page-1)))
	lastPage = page + 1

	return &Paging{
		Pages:      pages,
		TotalPages: totalPages,
		FirstPage:  firstPage,
		LastPage:   lastPage,
		CurrPage:   page,
	}
}
