package scheduler

import (
	"encoding/json"
	"github/beomsun1234/stockprice-collector/domain"
	"github/beomsun1234/stockprice-collector/service"
	"log"
)

type StockPriceCollectionSchedulerInterface interface {
	CollectStockPricesEverySecond()
}

type StockPriceCollectionScheduler struct {
	StockPriceColletorService service.StockPriceColletorServiceInterface
	Scheduler                 Scheduler
}

func NewStockPriceCollectionScheduler(stockPriceColletorService service.StockPriceColletorServiceInterface, scheduler Scheduler) StockPriceCollectionSchedulerInterface {
	return &StockPriceCollectionScheduler{
		StockPriceColletorService: stockPriceColletorService,
		Scheduler:                 scheduler,
	}
}

func (s *StockPriceCollectionScheduler) CollectStockPricesEverySecond() {
	log.Println("Start Scheduler")
	s.Scheduler.AddFunc("* * * * * *", s.collectStockPrices)
	s.Scheduler.Start()

}

func (s *StockPriceCollectionScheduler) collectStockPrices() {
	stock_prices := s.StockPriceColletorService.CollectStockPrices()
	s.printLogs(stock_prices)
}

func (s *StockPriceCollectionScheduler) printLogs(stock_prices []*domain.Stock) {
	byte_stock_prices, _ := json.Marshal(stock_prices)
	log.Println(string(byte_stock_prices))
}
