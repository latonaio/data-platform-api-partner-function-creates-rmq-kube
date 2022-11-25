package dpfm_api_input_reader

import (
	"data-platform-api-partner-function-creates-rmq-kube/DPFM_API_Caller/requests"
)

func (sdc *SDC) ConvertToPartnerFunction() *requests.PartnerFunction {
	data := sdc.PartnerFunction
	return &requests.PartnerFunction{
		PartnerFunction: data.PartnerFunction,
	}
}

func (sdc *SDC) ConvertToPartnerFunctionText() *requests.PartnerFunctionText {
	dataPartnerFunction := sdc.PartnerFunction
	data := sdc.PartnerFunction.PartnerFunctionText
	return &requests.PartnerFunctionText{
		PartnerFunction:     dataPartnerFunction.PartnerFunction,
		Language:            data.Language,
		PartnerFunctionName: data.PartnerFunctionName,
	}
}
