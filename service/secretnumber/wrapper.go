package secretnumber

import (
	"encoding/json"
	"net/http"

	"github.com/volcengine/volc-sdk-golang/base"
)

func (p *SecretNumber) BindAXB(req *BindAXBRequest) (*SecretBindResponse, int, error) {
	resp := new(SecretBindResponse)
	if statusCode, err := p.handler("BindAXB", req, resp); err != nil {
		return nil, statusCode, err
	} else {
		return resp, statusCode, nil
	}
}

func (p *SecretNumber) SelectNumberAndBindAXB(req *SelectNumberAndBindAXBRequest) (*SecretBindResponse, int, error) {
	resp := new(SecretBindResponse)
	if statusCode, err := p.handler("SelectNumberAndBindAXB", req, resp); err != nil {
		return nil, statusCode, err
	} else {
		return resp, statusCode, nil
	}
}

func (p *SecretNumber) UnbindAXB(req *SpecificSubIdRequest) (*OperationResponse, int, error) {
	resp := new(OperationResponse)
	if statusCode, err := p.handler("UnbindAXB", req, resp); err != nil {
		return nil, statusCode, err
	} else {
		return resp, statusCode, nil
	}
}

func (p *SecretNumber) QuerySubscription(req *SpecificSubIdRequest) (*QuerySubscriptionResponse, int, error) {
	resp := new(QuerySubscriptionResponse)
	if statusCode, err := p.handler("QuerySubscription", req, resp); err != nil {
		return nil, statusCode, err
	} else {
		return resp, statusCode, nil
	}
}

func (p *SecretNumber) QuerySubscriptionForList(req *QuerySubscriptionForListRequest) (*QuerySubscriptionForListResponse, int, error) {
	resp := new(QuerySubscriptionForListResponse)
	if statusCode, err := p.handler("QuerySubscriptionForList", req, resp); err != nil {
		return nil, statusCode, err
	} else {
		return resp, statusCode, nil
	}
}

func (p *SecretNumber) UpgradeAXToAXB(req *UpgradeAXToAXBRequest) (*SecretBindResponse, int, error) {
	resp := new(SecretBindResponse)
	if statusCode, err := p.handler("UpgradeAXToAXB", req, resp); err != nil {
		return nil, statusCode, err
	} else {
		return resp, statusCode, nil
	}
}

func (p *SecretNumber) UpdateAXB(req *UpdateAXBRequest) (*OperationResponse, int, error) {
	resp := new(OperationResponse)
	if statusCode, err := p.handler("UpdateAXB", req, resp); err != nil {
		return nil, statusCode, err
	} else {
		return resp, statusCode, nil
	}
}

func (p *SecretNumber) BindAXN(req *BindAXNRequest) (*SecretBindResponse, int, error) {
	resp := new(SecretBindResponse)
	if statusCode, err := p.handler("BindAXN", req, resp); err != nil {
		return nil, statusCode, err
	} else {
		return resp, statusCode, nil
	}
}

func (p *SecretNumber) SelectNumberAndBindAXN(req *SelectNumberAndBindAXNRequest) (*SecretBindResponse, int, error) {
	resp := new(SecretBindResponse)
	if statusCode, err := p.handler("SelectNumberAndBindAXN", req, resp); err != nil {
		return nil, statusCode, err
	} else {
		return resp, statusCode, nil
	}
}

func (p *SecretNumber) UpdateAXN(req *UpdateAXNRequest) (*OperationResponse, int, error) {
	resp := new(OperationResponse)
	if statusCode, err := p.handler("UpdateAXN", req, resp); err != nil {
		return nil, statusCode, err
	} else {
		return resp, statusCode, nil
	}
}

func (p *SecretNumber) UnbindAXN(req *SpecificSubIdRequest) (*OperationResponse, int, error) {
	resp := new(OperationResponse)
	if statusCode, err := p.handler("UnbindAXN", req, resp); err != nil {
		return nil, statusCode, err
	} else {
		return resp, statusCode, nil
	}
}

func (p *SecretNumber) BindAXNE(req *BindAXNERequest) (*SecretBindResponse, int, error) {
	resp := new(SecretBindResponse)
	if statusCode, err := p.handler("BindAXNE", req, resp); err != nil {
		return nil, statusCode, err
	} else {
		return resp, statusCode, nil
	}
}

func (p *SecretNumber) UnbindAXNE(req *SpecificSubIdRequest) (*OperationResponse, int, error) {
	resp := new(OperationResponse)
	if statusCode, err := p.handler("UnbindAXNE", req, resp); err != nil {
		return nil, statusCode, err
	} else {
		return resp, statusCode, nil
	}
}

func (p *SecretNumber) UpdateAXNE(req *UpdateAXNERequest) (*OperationResponse, int, error) {
	resp := new(OperationResponse)
	if statusCode, err := p.handler("UpdateAXNE", req, resp); err != nil {
		return nil, statusCode, err
	} else {
		return resp, statusCode, nil
	}
}

func (p *SecretNumber) Click2Call(req *Click2CallRequest) (*Click2CallResponse, int, error) {
	resp := new(Click2CallResponse)
	if statusCode, err := p.handler("Click2Call", req, resp); err != nil {
		return nil, statusCode, err
	} else {
		return resp, statusCode, nil
	}
}

func (p *SecretNumber) Click2CallLite(req *Click2CallLiteRequest) (*Click2CallLiteResponse, int, error) {
	resp := new(Click2CallLiteResponse)
	if statusCode, err := p.handler("Click2CallLite", req, resp); err != nil {
		return nil, statusCode, err
	} else {
		return resp, statusCode, nil
	}
}

func (p *DataCenter) QueryAudioRecordFileUrl(req *QueryAudioRecordFileUrlRequest) (*QueryAudioRecordFileUrlResponse, int, error) {
	resp := new(QueryAudioRecordFileUrlResponse)
	if statusCode, err := p.handler("QueryAudioRecordFileUrl", req, resp); err != nil {
		return nil, statusCode, err
	} else {
		return resp, statusCode, nil
	}
}

func (p *DataCenter) QueryAudioRecordToTextFileUrl(req *QueryAudioRecordToTextFileRequest) (*QueryAudioRecordToTextFileResponse, int, error) {
	resp := new(QueryAudioRecordToTextFileResponse)
	if statusCode, err := p.handler("QueryAudioRecordToTextFileUrl", req, resp); err != nil {
		return nil, statusCode, err
	} else {
		return resp, statusCode, nil
	}
}

func (p *DataCenter) QueryCallRecordMsg(req *QueryCallRecordMsgRequest) (*QueryCallRecordMsgResponse, int, error) {
	resp := new(QueryCallRecordMsgResponse)
	if statusCode, err := p.handler("QueryCallRecordMsg", req, resp); err != nil {
		return nil, statusCode, err
	} else {
		return resp, statusCode, nil
	}
}

func (p *NumberPool) CreateNumberPool(req *CreateNumberPoolRequest) (*CreateNumberPoolResponse, int, error) {
	resp := new(CreateNumberPoolResponse)
	if statusCode, err := handler("CreateNumberPool", req, resp, *p.Client); err != nil {
		return nil, statusCode, err
	} else {
		return resp, statusCode, nil
	}
}

func (p *NumberPool) UpdateNumberPool(req *UpdateNumberPoolRequest) (*UpdateNumberPoolResponse, int, error) {
	resp := new(UpdateNumberPoolResponse)
	if statusCode, err := handler("UpdateNumberPool", req, resp, *p.Client); err != nil {
		return nil, statusCode, err
	} else {
		return resp, statusCode, nil
	}
}

func (p *NumberPool) NumberPoolList(req *NumberPoolListRequest) (*NumberPoolListResponse, int, error) {
	resp := new(NumberPoolListResponse)
	if statusCode, err := handler("NumberPoolList", req, resp, *p.Client); err != nil {
		return nil, statusCode, err
	} else {
		return resp, statusCode, nil
	}
}

func (p *NumberPool) NumberList(req *NumberListRequest) (*NumberListResponse, int, error) {
	resp := new(NumberListResponse)
	if statusCode, err := handler("NumberList", req, resp, *p.Client); err != nil {
		return nil, statusCode, err
	} else {
		return resp, statusCode, nil
	}
}

func (p *NumberPool) EnableOrDisableNumber(req *EnableOrDisableNumberRequest) (*EnableOrDisableNumberResponse, int, error) {
	resp := new(EnableOrDisableNumberResponse)
	if statusCode, err := handler("EnableOrDisableNumber", req, resp, *p.Client); err != nil {
		return nil, statusCode, err
	} else {
		return resp, statusCode, nil
	}
}

func (p *NumberPool) QueryNumberApplyRecordList(req *QueryNumberApplyRecordListRequest) (*QueryNumberApplyRecordListResponse, int, error) {
	resp := new(QueryNumberApplyRecordListResponse)
	if statusCode, err := handler("QueryNumberApplyRecordList", req, resp, *p.Client); err != nil {
		return nil, statusCode, err
	} else {
		return resp, statusCode, nil
	}
}

func (p *MercService) CreateNumberApplication(req *CreateNumberApplicationRequest) (*CreateNumberApplicationResponse, int, error) {
	resp := new(CreateNumberApplicationResponse)
	if statusCode, err := handlerJson("CreateNumberApplication", req, resp, *p.Client); err != nil {
		return nil, statusCode, err
	} else {
		return resp, statusCode, nil
	}
}

func (p *ConfigService) AddQualification(req *AddQualificationRequest) (*AddQualificationResponse, int, error) {
	resp := new(AddQualificationResponse)
	if statusCode, err := handlerJson("AddQualification", req, resp, *p.Client); err != nil {
		return nil, statusCode, err
	} else {
		return resp, statusCode, nil
	}
}

func (p *ConfigService) UpdateQualification(req *UpdateQualificationRequest) (*UpdateQualificationResponse, int, error) {
	resp := new(UpdateQualificationResponse)
	if statusCode, err := handlerJson("UpdateQualification", req, resp, *p.Client); err != nil {
		return nil, statusCode, err
	} else {
		return resp, statusCode, nil
	}
}

func (p *ConfigService) AddQualificationScene(req *AddQualificationSceneRequest) (*AddQualificationSceneResponse, int, error) {
	resp := new(AddQualificationSceneResponse)
	if statusCode, err := handlerJson("AddQualificationScene", req, resp, *p.Client); err != nil {
		return nil, statusCode, err
	} else {
		return resp, statusCode, nil
	}
}

func (p *ConfigService) UpdateQualificationScene(req *UpdateQualificationSceneRequest) (*UpdateQualificationSceneResponse, int, error) {
	resp := new(UpdateQualificationSceneResponse)
	if statusCode, err := handlerJson("UpdateQualificationScene", req, resp, *p.Client); err != nil {
		return nil, statusCode, err
	} else {
		return resp, statusCode, nil
	}
}

func (p *ConfigService) QueryQualification(req *QueryQualificationRequest) (*QueryQualificationResponse, int, error) {
	resp := new(QueryQualificationResponse)
	if statusCode, err := handlerJson("QueryQualification", req, resp, *p.Client); err != nil {
		return nil, statusCode, err
	} else {
		return resp, statusCode, nil
	}
}

func (p *SecretNumber) handler(api string, req interface{}, resp interface{}) (int, error) {
	form := base.ToUrlValues(req)
	respBody, statusCode, err := p.Client.Post(api, nil, form)
	if err != nil {
		return statusCode, err
	}
	if statusCode >= 500 {
		respBody, statusCode, err = p.Client.Post(api, nil, form)
		if err != nil {
			return statusCode, err
		}
	}

	if err := json.Unmarshal(respBody, resp); err != nil {
		return statusCode, err
	}
	return statusCode, nil
}

func (p *DataCenter) handler(api string, req interface{}, resp interface{}) (int, error) {
	form := base.ToUrlValues(req)
	respBody, statusCode, err := p.Client.Post(api, nil, form)
	if err != nil {
		return statusCode, err
	}
	if statusCode >= 500 {
		respBody, statusCode, err = p.Client.Post(api, nil, form)
		if err != nil {
			return statusCode, err
		}
	}

	if err := json.Unmarshal(respBody, resp); err != nil {
		return statusCode, err
	}
	return statusCode, nil
}

func handler(api string, req interface{}, resp interface{}, p base.Client) (int, error) {
	form := base.ToUrlValues(req)
	apiInfo := p.ApiInfoList[api]
	var respBody []byte
	var statusCode int
	var err error
	if http.MethodGet == apiInfo.Method {
		respBody, statusCode, err = p.Query(api, form)
	} else {
		respBody, statusCode, err = p.Post(api, nil, form)
	}
	if err != nil {
		return statusCode, err
	}
	if statusCode >= 500 {
		if http.MethodGet == apiInfo.Method {
			respBody, statusCode, err = p.Query(api, form)
		} else {
			respBody, statusCode, err = p.Post(api, nil, form)
		}
		if err != nil {
			return statusCode, err
		}
	}

	if err := json.Unmarshal(respBody, resp); err != nil {
		return statusCode, err
	}
	return statusCode, nil
}

func handlerJson(api string, req interface{}, resp interface{}, p base.Client) (int, error) {
	jsonBody, _ := json.Marshal(req)
	respBody, statusCode, err := p.Json(api, nil, string(jsonBody))
	if err != nil {
		return statusCode, err
	}
	if statusCode >= 500 {
		respBody, statusCode, err = p.Json(api, nil, string(jsonBody))
		if err != nil {
			return statusCode, err
		}
	}

	if err := json.Unmarshal(respBody, resp); err != nil {
		return statusCode, err
	}
	return statusCode, nil
}
