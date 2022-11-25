package dpfm_api_output_formatter

type PartnerFunction struct {
	PartnerFunction     string              `json:"PartnerFunction"`
	PartnerFunctionText PartnerFunctionText `json:"PartnerFunctionText"`
}

type PartnerFunctionText struct {
	PartnerFunction     string `json:"PartnerFunction"`
	Language            string `json:"Language"`
	PartnerFunctionName string `json:"PartnerFunctionName"`
}
