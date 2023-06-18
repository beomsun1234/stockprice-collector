package service

import (
	"context"
	"github/beomsun1234/stockprice-collector/domain"
	"github/beomsun1234/stockprice-collector/external/kis"
	"log"
	"sync"
	"time"
)

type StockPriceColletorServiceInterface interface {
	CollectStockPrices() []*domain.Stock
}

type StockPriceColletorService struct {
	KisClientSetvice kis.KisClientSetviceInterface
}

var (
	collected_stock_prices []*domain.Stock
	stockCodes             []string
)

func setStockCode() {
	stockCodes = []string{
		"329180", //중공업
		"035720", //카카오
		"005930", //삼전
		"373220", //lg에너지
		"207940", //삼바
		"051910", //lg화학
		"035420", //네이버
		"012330", //현대모비스
		"005380", //현대차
		"105560", //kb
		"086790", //신한
		"086790", //하나
		"323410", //카뱅
		"000270", //기아
		"005490", //포스코
		"032830", //삼성생명,
		"024110", //기업은행,
		"377300", //카카오페이
		"316140", //우리금융
		"352820", //하이브
	}

}

func GetStockCodes() []string {
	return stockCodes
}

func NewStockPriceColletorService(kisClientSetviceDi kis.KisClientSetviceInterface) StockPriceColletorServiceInterface {
	return &StockPriceColletorService{
		KisClientSetvice: kisClientSetviceDi,
	}
}

/*
주식가격 수집
*/
func (s *StockPriceColletorService) CollectStockPrices() []*domain.Stock {
	token, err := s.KisClientSetvice.GetAccesstoken()
	if err != nil {
		token = domain.NewToken()
	}
	collected_stock_prices = []*domain.Stock{}
	setStockCode()
	wg := &sync.WaitGroup{}
	for _, stock_code := range stockCodes {
		wg.Add(1)
		go s.getStockPrice(wg, stock_code, token.AccessToken)
	}
	//wait until stock price collection is done
	wg.Wait()
	return collected_stock_prices

}

func (s *StockPriceColletorService) getStockPrice(wg *sync.WaitGroup, stock_code string, access_token string) (*domain.Stock, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer func() {
		wg.Done()
		cancel()
	}()
	stock_info, err := s.KisClientSetvice.GetStockPrice(stock_code, access_token)
	if err := ctx.Err(); err != nil {
		log.Println("timeout")
		stock_info = &domain.Stock{
			Stock_Code: stock_code,
		}
	}
	if err != nil {
		log.Println("error")
		stock_info = &domain.Stock{
			Stock_Code: stock_code,
		}
	}
	collected_stock_prices = append(collected_stock_prices, stock_info)
	return stock_info, nil
}
