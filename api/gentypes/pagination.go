package gentypes

// PageInfo - used for pagination of results
type PageInfo struct {
	Total  int32 // Total number of datapoints
	Offset int32 // The offset of this page from the start
	Limit  int32 // The amount requested to be returned
	Given  int32 // The number of items returned in the query
}

// Page is the PaginatedInput type used for queries to specify how much data to return
type Page struct {
	Offset *int32
	Limit  *int32
}
