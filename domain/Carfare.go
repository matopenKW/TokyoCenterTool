package domain

type Carfare struct {
	SeqNo      int    "json:seqNo"
	Date       string "json:date"
	StartPoint string "json:start_point"
	EndPoint   string "json:end_point"
	Price      int    "json:price"
}
