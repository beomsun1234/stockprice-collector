package scheduler

import (
	"encoding/json"
	"github/beomsun1234/stockprice-collector/domain"
	"github/beomsun1234/stockprice-collector/infra/messagequeue"
	"github/beomsun1234/stockprice-collector/service"
	"log"
)

type StockPriceCollectionSchedulerInterface interface {
	CollectStockPricesEverySecond()
}

type StockPriceCollectionScheduler struct {
	StockPriceColletorService service.StockPriceColletorServiceInterface
	Scheduler                 Scheduler
	Kafka                     messagequeue.Messagequeue
}

func NewStockPriceCollectionScheduler(stockPriceColletorService service.StockPriceColletorServiceInterface, scheduler Scheduler, mq messagequeue.Messagequeue) StockPriceCollectionSchedulerInterface {
	return &StockPriceCollectionScheduler{
		StockPriceColletorService: stockPriceColletorService,
		Scheduler:                 scheduler,
		Kafka:                     mq,
	}
}

func (s *StockPriceCollectionScheduler) CollectStockPricesEverySecond() {
	log.Println("Start Scheduler")
	s.Scheduler.AddFunc("* * * * * *", s.collectStockPrices)
	s.Scheduler.Start()

}

func (s *StockPriceCollectionScheduler) collectStockPrices() {
	stock_prices := s.StockPriceColletorService.CollectStockPrices()
	msg := s.convertStocksToBytes(stock_prices)
	err := s.Kafka.SendMessage(msg)
	if err != nil {
		log.Fatalln(err)
	}
	s.printLogs(msg)
}

func (s *StockPriceCollectionScheduler) convertStocksToBytes(stock_prices []*domain.Stock) []byte {
	byte_stock_prices, _ := json.Marshal(stock_prices)
	return byte_stock_prices
}

func (s *StockPriceCollectionScheduler) printLogs(stock_prices []byte) {
	log.Println(string(stock_prices))
	log.Println("---------------------------------------------------------------------")
}
