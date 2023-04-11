package pagination

import (
	"math"

	"go.mongodb.org/mongo-driver/mongo/options"
)

// T is a generic to your Items
// We already convert to slice internally, provide a raw struct
// `pagination.Paginate[YourStruct]`
type Paginate[T interface{}] struct {
	Items []T `json:"items"`
	*MongoPaginate
}

type MongoPaginate struct {
	Limit       int64  `json:"limit"`
	Page        int64  `json:"page"`
	TotalItems  int64  `json:"totalItems"`
	TotalPages  int64  `json:"totalPages"`
	NextPage    *int64 `json:"nextPage"`
	PrevPage    *int64 `json:"prevPage"`
	HasNextPage bool   `json:"hasNextPage"`
	HasPrevPage bool   `json:"hasPrevPage"`
}

// This will provide all struct to return base structure to pagination
/*
Limit       int64
Page        int64
TotalItems  int64
TotalPages  int64
NextPage    *int64
PrevPage    *int64
HasNextPage bool
HasPrevPage bool

To have your items into this struct use `pagination.Paginate[T]` struct
Provide then into your dto response
*/
func NewMongoPaginate(limit, page, total int64) *MongoPaginate {
	var (
		// access pointer only to can set nil value when not satisfies logic
		nextPage, prevPage *int64
	)

	// Total of pages. Calc ceil to provide max os pages
	totalPages := int64(math.Ceil(float64(total) / float64(limit)))

	hasNextPage := page < totalPages
	hasPrevPage := page > 1

	// If have more pages the next page will sum one page
	if hasNextPage {
		value := (page + 1)
		nextPage = &value
	} else {
		// If dont have more pages will set null value to json
		nextPage = nil
	}

	// If have previous pages the next page will subtract one page
	if hasPrevPage {
		value := page - 1
		prevPage = &value
	} else {
		// If dont have previous pages will set null value to json
		prevPage = nil
	}

	return &MongoPaginate{
		Limit:       limit,
		Page:        page,
		TotalItems:  total,
		TotalPages:  totalPages,
		NextPage:    nextPage,
		PrevPage:    prevPage,
		HasNextPage: hasNextPage,
		HasPrevPage: hasPrevPage,
	}
}

// This func will generate necessary options to mongo paginate your data.
// paginate := pagination.NewMongoPaginate(ctx, count)
// cur, err := db.Col.Find(ctx.Context(), bson.D{}, paginate.Options())
func (p MongoPaginate) Options() *options.FindOptions {
	skip := (p.Page * p.Limit) - p.Limit
	limit := p.Limit

	return &options.FindOptions{
		Limit: &limit,
		Skip:  &skip,
	}
}
