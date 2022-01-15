package models

type LeaseOffer struct {
	ID                    int     `json:"id"`
	UserID                int     `json:"userID"`
	Description           string  `json:"description"`
	FirstInstallmentPtMin string  `json:"firstInstallmentPtMin"`
	FirstInstallmentPtMax string  `json:"firstInstallmentPtMax"`
	InterestRate          float32 `json:"interestRate"`
	PeriodInMonths        uint    `json:"periodInMonths"`
	MinValue              uint    `json:"minValue"`
	MaxValue              uint    `json:"maxValue"`
	IsInsuranceRequired   bool    `json:"isInsuranceRequired"` //required_insurance
}
