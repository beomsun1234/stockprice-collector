package mocks

import (
	"bytes"
	"encoding/json"
	"github/beomsun1234/stockprice-collector/external/kis/dto"
	"io/ioutil"
	"net/http"
)

/*
MockStockPriceHttpClient
*/
type MockStockPriceHttpClient struct {
	MockKisStockPriceResponse *dto.KisStockPriceResponse
}

func (m *MockStockPriceHttpClient) Do(req *http.Request) (*http.Response, error) {
	byte_kisStockPriceResponse, _ := json.Marshal(m.MockKisStockPriceResponse)
	body := ioutil.NopCloser(bytes.NewReader(byte_kisStockPriceResponse))
	res := &http.Response{
		Body: body,
	}
	return res, nil
}

/*
	MockAccessTokenHttpClient
*/

type MockAccessTokenHttpClient struct {
	MockKisTokenResponse *dto.KisAccessTokenResponse
}

func (m *MockAccessTokenHttpClient) Do(req *http.Request) (*http.Response, error) {
	byte_kisAccessTokenResponse, _ := json.Marshal(m.MockKisTokenResponse)
	body := ioutil.NopCloser(bytes.NewReader(byte_kisAccessTokenResponse))
	res := &http.Response{
		Body: body,
	}
	return res, nil
}
