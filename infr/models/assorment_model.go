package models

type AssortmentModel struct {
	Id                  string `gorm:"primaryKey"`
	AccountId           string
	Shared              bool
	Updated             string
	Name                string
	Code                string
	ExternalCode        string
	Archived            bool
	PathName            string
	UseParentVat        bool
	Vat                 int
	VatEnabled          bool
	EffectiveVat        int
	EffectiveVatEnabled bool
	Article             string
	Weight              float64
	Volume              float64
	VariantsCount       int
	IsSerialTrackable   bool
	Stock               float64
	Reserve             float64
	InTransit           float64
	Quantity            int
}
