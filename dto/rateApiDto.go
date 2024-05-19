package dto

type RateApiDto struct {
	CCY     string `json:"ccy"`
	BaseCCY string `json:"base_ccy"`
	Buy     string `json:"buy"`
	Sale    string `json:"sale"`
}
