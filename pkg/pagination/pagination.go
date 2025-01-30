package pagination

import (
	"github.com/morkid/paginate"
)

var pg *paginate.Pagination

func GetInstance() *paginate.Pagination {
	if pg == nil {
		pg = paginate.New(&paginate.Config{
			DefaultSize:        10,
			PageStart:          1,
			CustomParamEnabled: true,
			SizeParams:         []string{"size", "limit", "per_page"},
			OrderParams:        []string{"order", "sort"},
		})
	}

	return pg
}
