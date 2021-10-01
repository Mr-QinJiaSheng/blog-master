package utils

import "math"

type Paginator struct {
	CurrentPage int64 `json:"currentPage"` //当前页
	PageSize    int64 `json:"pageSize"`    //每页数量
	TotalPage   int64 `json:"totalPage"`   //总页数
	TotalCount  int64 `json:"totalCount"`  //总数量
}

func GenPaginator(page, limit, count int64) Paginator {
	var paginator Paginator
	paginator.TotalCount = count
	paginator.TotalPage = int64(math.Ceil(float64(count) / float64(limit)))
	paginator.PageSize = limit
	paginator.CurrentPage = page
	return paginator
}
