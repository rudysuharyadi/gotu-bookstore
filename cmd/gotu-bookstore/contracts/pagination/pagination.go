package pagination

type Pagination struct {
	Limit   int    `form:"limit,default=100" mapstructure:"limit"`
	Page    int    `form:"page,default=0" mapstructure:"page"`
	SortBy  string `form:"sort_by" mapstructure:"sort_by"`
	Desc    bool   `form:"desc" mapstructure:"desc"`
	Keyword string `form:"keyword" mapstructure:"keyword"`
}

type Metadata struct {
	Pagination `mapstructure:",squash"`
	Size       int64 `mapstructure:"size"`
}
