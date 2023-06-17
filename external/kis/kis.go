package kis

import (
	"bytes"
	"encoding/json"
	"github/beomsun1234/stockprice-collector/config"
	"github/beomsun1234/stockprice-collector/domain"
	"github/beomsun1234/stockprice-collector/external"
	"github/beomsun1234/stockprice-collector/external/kis/dto"
	"github/beomsun1234/stockprice-collector/repository"
	"io"
	"log"
	"net/http"
	"net/url"
	"time"
)

const (
	kis_base = "https://openapi.koreainvestment.com:9443"
)

var (
	requestURL *url.URL
)

type KisClientSetviceInterface interface {
	GetStockPrice(stock_code string) (*domain.Stock, error)
	GetAccesstoken() (*domain.Token, error)
}

type KisClientSetvice struct {
	HttpClient               external.HttpClient
	KisConfig                *config.KisConfig
	KisAccessTokenRepository repository.KisAccessTokenRepositoryInterface
}

func NewKisClientSetvice(http_client_di external.HttpClient, kis_config *config.KisConfig, kisAccessTokenRepository repository.KisAccessTokenRepositoryInterface) KisClientSetviceInterface {
	return &KisClientSetvice{
		HttpClient:               http_client_di,
		KisConfig:                kis_config,
		KisAccessTokenRepository: kisAccessTokenRepository,
	}
}

/*
주식 현재가 요청
*/
func (k *KisClientSetvice) GetStockPrice(stock_code string) (*domain.Stock, error) {
	// token 체크
	token, err := k.KisAccessTokenRepository.GetKisAccessToken()
	if token == nil || token.IsTokenExpired() || err != nil {
		log.Println("token reissue")
		token, err = k.GetAccesstoken()
		if err != nil {
			return nil, err
		}
	}

	kis_stock_price_response, err := k.requestStockPriceToKis(stock_code, token.AccessToken)
	log.Println("msg1 : ", kis_stock_price_response.Msg1)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return kis_stock_price_response.KisStockPriceResDetails.ToStock(stock_code), nil
}

func (k *KisClientSetvice) requestStockPriceToKis(stock_code string, accessToken string) (*dto.KisStockPriceResponse, error) {
	k.setStockPriceUrl(stock_code)
	request, err := http.NewRequest("GET", requestURL.String(), nil)
	k.setStockPriceRequestHeader(request, accessToken)
	if err != nil {
		return nil, err
	}

	response, err := k.HttpClient.Do(request)

	if err != nil {
		log.Println("request err : ", err)
		return nil, err
	}

	defer response.Body.Close()

	responseBody, err := io.ReadAll(response.Body)

	if err != nil {
		log.Println("resonposebody read err : ", err)
		return nil, err
	}

	kis_stock_price_response := &dto.KisStockPriceResponse{}
	json.Unmarshal(responseBody, &kis_stock_price_response)

	return kis_stock_price_response, nil
}

func (k *KisClientSetvice) setStockPriceUrl(stock_code string) {
	requestURL, _ = url.Parse(kis_base + "/uapi/domestic-stock/v1/quotations/inquire-price")
	//set queryparm
	params := url.Values{}
	params.Add("FID_COND_MRKT_DIV_CODE", "J")
	params.Add("FID_INPUT_ISCD", stock_code)
	requestURL.RawQuery = params.Encode()
}

func (k *KisClientSetvice) setStockPriceRequestHeader(req *http.Request, accessToken string) {
	req.Header.Add("authorization", "Bearer "+accessToken)
	req.Header.Add("appkey", k.KisConfig.Key)
	req.Header.Add("appsecret", k.KisConfig.Secret)
	req.Header.Add("tr_id", "FHKST01010100")
}

/*
토큰 요청
*/
func (k *KisClientSetvice) GetAccesstoken() (*domain.Token, error) {
	tokenRes, err := k.requestToken()
	if err != nil {
		return nil, err
	}
	k.KisAccessTokenRepository.DeleteKisAccessToken()
	return k.saveToken(tokenRes), err
}

func (k *KisClientSetvice) saveToken(res_token *dto.KisAccessTokenResponse) *domain.Token {
	now := time.Now().Format("2006-01-02 15:04:05")
	token := res_token.ToToken(now)
	k.KisAccessTokenRepository.InsertKisAccessToken(token)
	return token
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
