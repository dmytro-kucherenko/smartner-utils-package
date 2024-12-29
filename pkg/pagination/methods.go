package pagination

import (
	"net/http"

	"github.com/dmytro-kucherenko/smartner-utils-package/pkg/pagination/dtos"
	"github.com/dmytro-kucherenko/smartner-utils-package/pkg/server/errors"
)

func GetPageMeasures(total uint64, query dtos.PageQueryRequest) (measures PageMeasures, err error) {
	var page uint64 = 1
	if query.Page != nil {
		page = *query.Page
	}

	var limit uint64 = 20
	if query.Size != nil {
		limit = *query.Size
	}

	offset := (page - 1) * limit
	if offset >= total && (total != 0 || page != 1) {
		err = errors.NewHttpError(http.StatusBadRequest, "Page is out of bounds.")

		return
	}

	measures = PageMeasures{Offset: offset, Limit: limit, Page: page, Size: limit}

	return
}

func GetPageMeta(total uint64, measures PageMeasures) dtos.PageMetaResponse {
	meta := dtos.PageMetaResponse{
		Total: total,
		Page:  measures.Page,
		Size:  measures.Size,
	}

	if measures.Offset+measures.Limit < total {
		next := meta.Page + 1
		meta.Next = &next
	}

	if meta.Page > 1 {
		prev := meta.Page - 1
		meta.Prev = &prev
	}

	return meta
}
