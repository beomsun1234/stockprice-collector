package domain

type Stock struct {
	Stock_Code           string `json:"stockCode"`
	Stock_Price          string `json:"stockPrice"`
	Stock_Volume         string `json:"stockVolume"`
	Stock_Highest_Price  string `json:"stockHighestPrice"`
	Stock_Lowest_Price   string `json:"stockLowestPrice"`
	Stock_Prdy_Vrss_Sign string `json:"stockPrdyVrssSign"` //전일대비부호
}

func NewStock() *Stock {
	return &Stock{}
}

func (s *Stock) BuildStockCode(stock_code string) *Stock {
	s.Stock_Code = stock_code
	return s
}
func (s *Stock) BuildStockPrice(stock_price string) *Stock {
	s.Stock_Price = stock_price
	return s
}
func (s *Stock) BuildStockVolume(stock_volume string) *Stock {
	s.Stock_Volume = stock_volume
	return s
}

func (s *Stock) BuildStockHighestPrice(stock_h_price string) *Stock {
	s.Stock_Highest_Price = stock_h_price
	return s
}
func (s *Stock) BuildStockLowestPrice(stock_l_price string) *Stock {
	s.Stock_Lowest_Price = stock_l_price
	return s
}
func (s *Stock) BuildStockPrdyVrssSign(stock_prdy_vrss_sign string) *Stock {
	s.Stock_Prdy_Vrss_Sign = stock_prdy_vrss_sign
	return s
}
