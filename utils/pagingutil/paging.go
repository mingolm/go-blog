package pagingutil

import "math"

type Paging struct {
	Pages    []int `json:"pages"`
	Total    int   `json:"total"`
	PrevPage int   `json:"prev_page"`
	NextPage int   `json:"next_page"`
	CurrPage int   `json:"curr_page"`
}

func Paginator(page, pageLimit, total int) *Paging {
	var (
		prev      int // 前一页地址
		next      int // 后一页地址
		pageTotal int // 总页数
	)
	pageTotal = int(math.Ceil(float64(total) / float64(pageLimit)))
	page = int(math.Max(float64(1), float64(pageTotal)))
	prev = int(math.Max(float64(1), float64(page-1)))
	next = page + 1
	if pageTotal == 0 {
		next = 1
	}
	pages := make([]int, total)
	for i := range pages {
		pages[i] = i + 1
	}

	return &Paging{
		Pages:    pages,
		Total:    total,
		PrevPage: prev,
		NextPage: next,
		CurrPage: page,
	}
}
