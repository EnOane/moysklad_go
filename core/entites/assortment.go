package entites

type Assortment struct {
	Id                  string  `json:"id"`
	AccountId           string  `json:"accountId"`
	Shared              bool    `json:"shared"`
	Updated             string  `json:"updated"`
	Name                string  `json:"name"`
	Code                string  `json:"code"`
	ExternalCode        string  `json:"externalCode"`
	Archived            bool    `json:"archived"`
	PathName            string  `json:"pathName"`
	UseParentVat        bool    `json:"useParentVat"`
	Vat                 int     `json:"vat"`
	VatEnabled          bool    `json:"vatEnabled"`
	EffectiveVat        int     `json:"effectiveVat"`
	EffectiveVatEnabled bool    `json:"effectiveVatEnabled"`
	Article             string  `json:"article"`
	Weight              float64 `json:"weight"`
	Volume              float64 `json:"volume"`
	VariantsCount       int     `json:"variantsCount"`
	IsSerialTrackable   bool    `json:"isSerialTrackable"`
	Stock               float64 `json:"stock"`
	Reserve             float64 `json:"reserve"`
	InTransit           float64 `json:"inTransit"`
	Quantity            float64 `json:"quantity"`
}
