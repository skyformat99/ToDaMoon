package exchanges

//Ticker 是ticker的数据结构。
type Ticker struct {
	Last float64
	Buy  float64
	Sell float64
	High float64
	Low  float64
	Vol  float64
}

//Depth 记录深度信息
type Depth struct {
	Asks [][2]float64
	Bids [][2]float64
}

//Trade 记录一个成交记录的细节
type Trade struct {
	Tid    int64
	Date   int64
	Price  float64
	Amount float64
	Type   string
}

func (t Trade) Attributes() (int64, int64, float64, float64, string) {
	return t.Tid, t.Date, t.Price, t.Amount, t.Type
}

type Trades []Trade

type Balance struct {
}
