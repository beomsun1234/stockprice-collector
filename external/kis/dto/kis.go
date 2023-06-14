package dto

import "github/beomsun1234/stockprice-collector/domain"

func NewKisAccessTokenRequest(grant_type string, appsecret string, appkey string) *KisAccessTokenRequest {
	return &KisAccessTokenRequest{
		GrantType: grant_type,
		AppSecret: appsecret,
		AppKey:    appkey,
	}
}

type KisAccessTokenRequest struct {
	GrantType string `json:"grant_type"`
	AppSecret string `json:"appsecret"`
	AppKey    string `json:"appkey"`
}

type KisAccessTokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
}

type KisStockPriceResponse struct {
	KisStockPriceResDetails KisStockPriceResponseDetails `json:"output"`
	Rt_Cd                   string                       `json:"rt_cd"`
	Msg_Cd                  string                       `json:"msg_cd"`
	Msg1                    string                       `json:"msg1"`
}

type KisStockPriceResponseDetails struct {
	Iscd_Stat_Cls_Code       string `json:"iscd_stat_cls_code"`
	Marg_Rate                string `json:"marg_rate"`
	Rprs_Mrkt_Kor_Name       string `json:"rprs_mrkt_kor_name"`
	Bstp_Kor_Isnm            string `json:"bstp_kor_isnm"`
	Temp_stop_yn             string `json:"temp_stop_yn"`
	Oprc_Rang_Cont_Yn        string `json:"oprc_rang_cont_yn"`
	Clpr_Rang_Cont_Yn        string `json:"clpr_rang_cont_yn"`
	Crdt_able_Yn             string `json:"crdt_able_yn"`
	Grmn_Rate_Cls_Code       string `json:"grmn_rate_cls_code"`
	Elw_Pblc_Yn              string `json:"elw_pblc_yn"`
	Stck_Prpr                string `json:"stck_prpr"` //주식현재가
	Prdy_Vrss                string `json:"prdy_vrss"`
	Prdy_Vrss_Sign           string `json:"prdy_vrss_sign"` //전일대비부호
	Prdy_Ctrt                string `json:"prdy_ctrt"`
	Acml_Tr_Pbmn             string `json:"acml_tr_pbmn"`
	Acml_Vol                 string `json:"acml_vol"` //거래량
	Prdy_Vrss_Vol_Rate       string `json:"prdy_vrss_vol_rate"`
	Stck_Oprc                string `json:"stck_oprc"`
	Stck_Hgpr                string `json:"stck_hgpr"` //최고가
	Stck_Lwpr                string `json:"stck_lwpr"` //최저가
	Stck_Mxpr                string `json:"stck_mxpr"`
	Stck_Llam                string `json:"stck_llam"`
	Stck_Sdpr                string `json:"stck_sdpr"`
	Wghn_Avrg_Stck_Prc       string `json:"wghn_avrg_stck_prc"`
	Hts_Frgn_Ehrt            string `json:"hts_frgn_ehrt"`
	Frgn_Ntby_Qty            string `json:"frgn_ntby_qty"`
	Pgtr_Ntby_Qty            string `json:"pgtr_ntby_qty"`
	Pvt_Scnd_Dmrs_Prc        string `json:"pvt_scnd_dmrs_prc"`
	Pvt_Frst_Dmrs_Prc        string `json:"pvt_frst_dmrs_prc"`
	Pvt_Pont_Val             string `json:"pvt_pont_val"`
	Pvt_Frst_Dmsp_Prc        string `json:"pvt_frst_dmsp_prc"`
	Pvt_Scnd_Dmsp_Prc        string `json:"pvt_scnd_dmsp_prc"`
	Dmrs_Val                 string `json:"dmrs_val"`
	Dmsp_Val                 string `json:"dmsp_val"`
	Cpfn                     string `json:"cpfn"`
	Rstc_Wdth_Prc            string `json:"rstc_wdth_prc"`
	Stck_Fcam                string `json:"stck_fcam"`
	Stck_Sspr                string `json:"stck_sspr"`
	Aspr_Unit                string `json:"aspr_unit"`
	Hts_Deal_Qty_Unit_Val    string `json:"hts_deal_qty_unit_val"`
	Lstn_Stcn                string `json:"lstn_stcn"`
	Hts_Avls                 string `json:"hts_avls"`
	Per                      string `json:"per"`
	Pbr                      string `json:"pbr"`
	Stac_Month               string `json:"stac_month"`
	Vol_Tnrt                 string `json:"vol_tnrt"`
	Eps                      string `json:"eps"`
	Bps                      string `json:"bps"`
	D250_Hgpr                string `json:"d250_hgpr"`
	D250_Hgpr_Date           string `json:"d250_hgpr_date"`
	D250_hgpr_Vrss_Prpr_Rate string `json:"d250_hgpr_vrss_prpr_rate"`
	D250_Lwpr                string `json:"d250_lwpr"`
	D250_Lwpr_Date           string `json:"d250_lwpr_date"`
	D250_Lwpr_Vrss_Prpr_Rate string `json:"d250_lwpr_vrss_prpr_rate"`
	Stck_Dryy_Hgpr           string `json:"stck_dryy_hgpr"`
	Dryy_Hgpr_Vrss_Prpr_Rate string `json:"dryy_hgpr_vrss_prpr_rate"`
	Dryy_Hgpr_Date           string `json:"dryy_hgpr_date"`
	Stck_Dryy_Lwpr           string `json:"stck_dryy_lwpr"`
	Dryy_Lwpr_Vrss_Prpr_Rate string `json:"dryy_lwpr_vrss_prpr_rate"`
	Dryy_Lwpr_Date           string `json:"dryy_lwpr_date"`
	W52_Hgpr                 string `json:"w52_hgpr"`
	W52_Hgpr_Vrss_Prpr_Ctrt  string `json:"w52_hgpr_vrss_prpr_ctrt"`
	W52_Hgpr_Date            string `json:"w52_hgpr_date"`
	W52_Lwpr                 string `json:"w52_lwpr"`
	W52_Lwpr_Vrss_Prpr_Ctrt  string `json:"w52_lwpr_vrss_prpr_ctrt"`
	W52_Lwpr_Date            string `json:"w52_lwpr_date"`
	Whol_Loan_Rmnd_Rate      string `json:"whol_loan_rmnd_rate"`
	Ssts_Yn                  string `json:"ssts_yn"`
	Stck_Shrn_Iscd           string `json:"stck_shrn_iscd"`
	Fcam_Cnnm                string `json:"fcam_cnnm"`
	Cpfn_Cnnm                string `json:"cpfn_cnnm"`
	Frgn_Hldn_Qty            string `json:"frgn_hldn_qty"`
	Vi_Cls_Code              string `json:"vi_cls_code"`
	Ovtm_Vi_Cls_Code         string `json:"ovtm_vi_cls_code"`
	Last_Ssts_Cntg_Qty       string `json:"last_ssts_cntg_qty"`
	Invt_Caful_Yn            string `json:"invt_caful_yn"`
	Mrkt_Warn_Cls_Code       string `json:"mrkt_warn_cls_code"`
	Short_Over_Yn            string `json:"short_over_yn"`
	Sltr_Yn                  string `json:"sltr_yn"`
}

func (s *KisStockPriceResponseDetails) ToStock(stock_code string) *domain.Stock {
	stock := domain.NewStock().BuildStockCode(stock_code).BuildStockPrice(s.Stck_Prpr).BuildStockPrdyVrssSign(s.Prdy_Vrss_Sign).BuildStockHighestPrice(s.Stck_Hgpr).BuildStockLowestPrice(s.Stck_Lwpr).BuildStockVolume(s.Acml_Vol)

	return stock
}
