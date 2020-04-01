package gentypes

// PageInfo - used for pagination of results
type PageInfo struct {
	PagesAfter int32 //Number of pages after this one
	Offset     int32 //The offset of this page from the start
	Limit      int32 //The max shown per page
	Given      int32 // The number of items returned in the query
}

// Page is the PaginatedInput type used for queries to specify how much data to return
type Page struct {
	Offset *int32
	Limit  *int32
}
