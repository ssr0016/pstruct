package pagination

const (
	DefaultPage      = 1
	DefaultPageLimit = 20
)

type PaginationConfig struct {
	Page      int
	PageLimit int
}
