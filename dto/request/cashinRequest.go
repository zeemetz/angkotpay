package request

type CashinRequest struct {
	From   string
	To     string
	Amount int64
	Notes  string
}
