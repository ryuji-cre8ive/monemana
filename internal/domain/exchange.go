package domain

type Exchange struct {
	Success   bool
	Timestamp uint64
	Base      string
	Date      string
	Rates     map[string]float64
}
