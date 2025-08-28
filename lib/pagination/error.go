package pagination

import "errors"

var (
	ErrorMaxPage = errors.New("maximum page limit exceeded")
	ErrorPage    = errors.New("invalid page number")
	ErrorPageEmpty = errors.New("page number is required")
	ErrorPageInvalid = errors.New("invalid page number")
)