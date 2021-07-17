package pagingutil

import "math"

type Paging struct {
	Pages    []int `json:"pages"`
	Total    int   `json:"total"`
	PrevPage int   `json:"prev_page"`
	NextPage int   `json:"next_page"`
	CurrPage int   `json:"curr_page"`
}

func Paginator(page, pageLimit int, nums int64) *Paging {
	var PrevPage int                                            //前一页地址
	var NextPage int                                            //后一页地址
	Total := int(math.Ceil(float64(nums) / float64(pageLimit))) //page总数
	if page > Total {
		page = Total
	}
	if page <= 0 {
		page = 1
	}
	pages := make([]int, Total)
	for i := range pages {
		pages[i] = i + 1
	}
	PrevPage = int(math.Max(float64(1), float64(page-1)))
	NextPage = page + 1

	return &Paging{
		Pages:    pages,
		Total:    Total,
		PrevPage: PrevPage,
		NextPage: NextPage,
		CurrPage: page,
	}
}
