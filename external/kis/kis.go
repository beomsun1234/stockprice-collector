package kis

import (
	"bytes"
	"encoding/json"
	"github/beomsun1234/stockprice-collector/config"
	"github/beomsun1234/stockprice-collector/domain"
	"github/beomsun1234/stockprice-collector/external"
	"github/beomsun1234/stockprice-collector/external/kis/dto"
	"io"
	"net/http"
	"net/url"
)

const (
	kis_base = "https://openapi.koreainvestment.com:9443"
)

var requestURL *url.URL

type KisClientSetviceInterface interface {
	GetStockPrice(stock_code string) (*domain.Stock, error)
	GetAccesstoken() (*dto.KisAccessTokenResponse, error)
}

type KisClientSetvice struct {
	HttpClient external.HttpClient
	KisConfig  *config.KisConfig
}

func NewKisClientSetvice(http_client_di external.HttpClient, kis_config *config.KisConfig) KisClientSetviceInterface {
	return &KisClientSetvice{
		HttpClient: http_client_di,
		KisConfig:  kis_config,
	}
}

/*
주식 현재가 요청
*/
func (k *KisClientSetvice) GetStockPrice(stock_code string) (*domain.Stock, error) {
	kis_stock_price_response, err := k.requestStockPriceToKis(stock_code)
	if err != nil {
		return nil, err
	}
	return kis_stock_price_response.KisStockPriceResDetails.ToStock(stock_code), nil
}

func (k *KisClientSetvice) setStockPriceUrl(stock_code string) {
	requestURL, _ = url.Parse(kis_base + "/uapi/domestic-stock/v1/quotations/inquire-price")
	//set queryparm
	params := url.Values{}
	params.Add("FID_COND_MRKT_DIV_CODE", "J")
	params.Add("FID_INPUT_ISCD", stock_code)
	requestURL.RawQuery = params.Encode()
}

func (k *KisClientSetvice) setStockPriceRequestHeader(req *http.Request) {
	req.Header.Add("authorization", "Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzUxMiJ9.eyJzdWIiOiJ0b2tlbiIsImF1ZCI6ImYyZGVkZjg1LWZlNmQtNGI2OC05ODg5LTI3YTYxNjAyZGYyOCIsImlzcyI6InVub2d3IiwiZXhwIjoxNjg2MzkzNTkzLCJpYXQiOjE2ODYzMDcxOTMsImp0aSI6IlBTb0VZcEk3TW5zQUpJSGxwQUc0ZXFySEVOWDVnR2JJQ3NkbSJ9.-LonwsKY9wZR3PvkZ5gnGCiPN0QEcdsfAQ0KOMJeBJByhIcLsJR6yHeGWM4_CFn-LzcHlmEHp2AUAqe3-7LvgA")
	req.Header.Add("appkey", k.KisConfig.Key)
	req.Header.Add("appsecret", k.KisConfig.Secret)
	req.Header.Add("tr_id", "FHKST01010100")
}

func (k *KisClientSetvice) requestStockPriceToKis(stock_code string) (*dto.KisStockPriceResponse, error) {
	k.setStockPriceUrl(stock_code)
	request, err := http.NewRequest("GET", requestURL.String(), nil)
	k.setStockPriceRequestHeader(request)
	if err != nil {
		return nil, err
	}

	response, err := k.HttpClient.Do(request)

	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	responseBody, err := io.ReadAll(response.Body)

	if err != nil {
		return nil, err
	}

	kis_stock_price_response := &dto.KisStockPriceResponse{}
	err = json.Unmarshal(responseBody, &kis_stock_price_response)

	if err != nil {
		return nil, err
	}

	return kis_stock_price_response, nil

}

/*
토큰 요청
*/
func (k *KisClientSetvice) GetAccesstoken() (*dto.KisAccessTokenResponse, error) {

	return k.requestToken()
}

func (k *KisClientSetvice) requestToken() (*dto.KisAccessTokenResponse, error) {
	kisAccessTokenRequest := dto.KisAccessTokenRequest{
		GrantType: "client_credentials",
		AppSecret: k.KisConfig.Secret,
		AppKey:    k.KisConfig.Key,
	}
	bytes_kisAccessTokenRequest, _ := json.Marshal(kisAccessTokenRequest)
	request_body := bytes.NewBuffer(bytes_kisAccessTokenRequest)
	request, err := http.NewRequest("POST", kis_base+"/oauth2/tokenP", request_body)

	if err != nil {
		return nil, err
	}

	request.Header.Add("Content-Type", "application/json")
	response, err := k.HttpClient.Do(request)

	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	kisAccessTokenResponse := &dto.KisAccessTokenResponse{}
	respBody, _ := io.ReadAll(response.Body)
	json.Unmarshal(respBody, &kisAccessTokenResponse)

	return kisAccessTokenResponse, nil
}
