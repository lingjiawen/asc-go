package asc

import (
	"context"
	"fmt"
)

// BetaLicenseAgreement defines model for BetaLicenseAgreement.
//
// https://developer.apple.com/documentation/appstoreconnectapi/betalicenseagreement
type BetaLicenseAgreement struct {
	Attributes *struct {
		AgreementText *string `json:"agreementText,omitempty"`
	} `json:"attributes,omitempty"`
	ID            string        `json:"id"`
	Links         ResourceLinks `json:"links"`
	Relationships *struct {
		App *struct {
			Data  *RelationshipsData  `json:"data,omitempty"`
			Links *RelationshipsLinks `json:"links,omitempty"`
		} `json:"app,omitempty"`
	} `json:"relationships,omitempty"`
	Type string `json:"type"`
}

// BetaLicenseAgreementUpdateRequest defines model for BetaLicenseAgreementUpdateRequest.
//
// https://developer.apple.com/documentation/appstoreconnectapi/betalicenseagreementupdaterequest
type BetaLicenseAgreementUpdateRequest struct {
	Attributes *BetaLicenseAgreementUpdateRequestAttributes `json:"attributes,omitempty"`
	ID         string                                       `json:"id"`
	Type       string                                       `json:"type"`
}

// BetaLicenseAgreementUpdateRequestAttributes are attributes for BetaLicenseAgreementUpdateRequest
//
// https://developer.apple.com/documentation/appstoreconnectapi/betalicenseagreementupdaterequest/data/attributes
type BetaLicenseAgreementUpdateRequestAttributes struct {
	AgreementText *string `json:"agreementText,omitempty"`
}

// BetaLicenseAgreementsResponse defines model for BetaLicenseAgreementsResponse.
//
// https://developer.apple.com/documentation/appstoreconnectapi/betalicenseagreementsresponse
type BetaLicenseAgreementsResponse struct {
	Data     []BetaLicenseAgreement `json:"data"`
	Included *[]App                 `json:"included,omitempty"`
	Links    PagedDocumentLinks     `json:"links"`
	Meta     *PagingInformation     `json:"meta,omitempty"`
}

// BetaLicenseAgreementResponse defines model for BetaLicenseAgreementResponse.
//
// https://developer.apple.com/documentation/appstoreconnectapi/betalicenseagreementresponse
type BetaLicenseAgreementResponse struct {
	Data     BetaLicenseAgreement `json:"data"`
	Included *[]App               `json:"included,omitempty"`
	Links    DocumentLinks        `json:"links"`
}

// ListBetaLicenseAgreementsQuery defines model for ListBetaLicenseAgreements
//
// https://developer.apple.com/documentation/appstoreconnectapi/list_beta_license_agreements
type ListBetaLicenseAgreementsQuery struct {
	FieldsApps                  []string `url:"fields[apps],omitempty"`
	FieldsBetaLicenseAgreements []string `url:"fields[betaLicenseAgreements],omitempty"`
	FilterApp                   []string `url:"filter[app],omitempty"`
	Include                     []string `url:"include,omitempty"`
	Limit                       int      `url:"limit,omitempty"`
	Cursor                      string   `url:"cursor,omitempty"`
}

// GetBetaLicenseAgreementQuery defines model for GetBetaLicenseAgreement
//
// https://developer.apple.com/documentation/appstoreconnectapi/read_beta_license_agreement_information
type GetBetaLicenseAgreementQuery struct {
	FieldsApps                  []string `url:"fields[apps],omitempty"`
	FieldsBetaLicenseAgreements []string `url:"fields[betaLicenseAgreements],omitempty"`
	Include                     []string `url:"include,omitempty"`
}

// GetAppForBetaLicenseAgreementQuery defines model for GetAppForBetaLicenseAgreement
//
// https://developer.apple.com/documentation/appstoreconnectapi/read_the_app_information_of_a_beta_license_agreement
type GetAppForBetaLicenseAgreementQuery struct {
	FieldsApps []string `url:"fields[apps],omitempty"`
}

// GetBetaLicenseAgreementForAppQuery defines model for GetBetaLicenseAgreementForApp
//
// https://developer.apple.com/documentation/appstoreconnectapi/read_the_beta_license_agreement_of_an_app
type GetBetaLicenseAgreementForAppQuery struct {
	FieldsBetaLicenseAgreements []string `url:"fields[betaLicenseAgreements],omitempty"`
}

// ListBetaLicenseAgreements finds and lists beta license agreements for all apps.
//
// https://developer.apple.com/documentation/appstoreconnectapi/list_beta_license_agreements
func (s *TestflightService) ListBetaLicenseAgreements(ctx context.Context, params *ListBetaLicenseAgreementsQuery) (*BetaLicenseAgreementsResponse, *Response, error) {
	res := new(BetaLicenseAgreementsResponse)
	resp, err := s.client.get(ctx, "betaLicenseAgreements", params, res)
	return res, resp, err
}

// GetBetaLicenseAgreement gets a specific beta license agreement.
//
// https://developer.apple.com/documentation/appstoreconnectapi/read_beta_license_agreement_information
func (s *TestflightService) GetBetaLicenseAgreement(ctx context.Context, id string, params *GetBetaLicenseAgreementQuery) (*BetaLicenseAgreementResponse, *Response, error) {
	url := fmt.Sprintf("betaLicenseAgreements/%s", id)
	res := new(BetaLicenseAgreementResponse)
	resp, err := s.client.get(ctx, url, params, res)
	return res, resp, err
}

// GetAppForBetaLicenseAgreement gets the app information for a specific beta license agreement.
//
// https://developer.apple.com/documentation/appstoreconnectapi/read_the_app_information_of_a_beta_license_agreement
func (s *TestflightService) GetAppForBetaLicenseAgreement(ctx context.Context, id string, params *GetAppForBetaLicenseAgreementQuery) (*AppResponse, *Response, error) {
	url := fmt.Sprintf("betaLicenseAgreements/%s/app", id)
	res := new(AppResponse)
	resp, err := s.client.get(ctx, url, params, res)
	return res, resp, err
}

// GetBetaLicenseAgreementForApp gets the beta license agreement for a specific app.
//
// https://developer.apple.com/documentation/appstoreconnectapi/read_the_beta_license_agreement_of_an_app
func (s *TestflightService) GetBetaLicenseAgreementForApp(ctx context.Context, id string, params *GetBetaLicenseAgreementForAppQuery) (*BetaLicenseAgreementResponse, *Response, error) {
	url := fmt.Sprintf("apps/%s/betaLicenseAgreement", id)
	res := new(BetaLicenseAgreementResponse)
	resp, err := s.client.get(ctx, url, params, res)
	return res, resp, err
}

// UpdateBetaLicenseAgreement updates the text for your beta license agreement.
//
// https://developer.apple.com/documentation/appstoreconnectapi/modify_a_beta_license_agreement
func (s *TestflightService) UpdateBetaLicenseAgreement(ctx context.Context, id string, body *BetaLicenseAgreementUpdateRequest) (*BetaLicenseAgreementResponse, *Response, error) {
	url := fmt.Sprintf("betaLicenseAgreements/%s", id)
	res := new(BetaLicenseAgreementResponse)
	resp, err := s.client.patch(ctx, url, body, res)
	return res, resp, err
}
