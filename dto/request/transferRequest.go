package request

type TransferRequest struct {
	From   string
	To     string
	Amount int64
	Notes  string
}
