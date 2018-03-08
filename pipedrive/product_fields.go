package pipedrive

import (
	"fmt"
	"net/http"
)

type ProductFieldsService service

type ProductField struct {
	ID                 int         `json:"id"`
	Key                string      `json:"key"`
	Name               string      `json:"name"`
	OrderNr            int         `json:"order_nr"`
	PicklistData       interface{} `json:"picklist_data,omitempty"`
	FieldType          string      `json:"field_type"`
	AddTime            string      `json:"add_time"`
	UpdateTime         string      `json:"update_time"`
	ActiveFlag         bool        `json:"active_flag"`
	EditFlag           bool        `json:"edit_flag"`
	IndexVisibleFlag   bool        `json:"index_visible_flag"`
	DetailsVisibleFlag bool        `json:"details_visible_flag"`
	AddVisibleFlag     bool        `json:"add_visible_flag"`
	ImportantFlag      bool        `json:"important_flag"`
	BulkEditAllowed    bool        `json:"bulk_edit_allowed"`
	SearchableFlag     bool        `json:"searchable_flag"`
	FilteringAllowed   bool        `json:"filtering_allowed"`
	SortableFlag       bool        `json:"sortable_flag"`
	UseField           string      `json:"use_field,omitempty"`
	Link               string      `json:"link,omitempty"`
	MandatoryFlag      bool        `json:"mandatory_flag"`
	DisplayField       string      `json:"display_field,omitempty"`
	Options            []struct {
		ID    bool   `json:"id"`
		Label string `json:"label"`
	} `json:"options,omitempty"`
}

type ProductFieldsResponse struct {
	Success        bool           `json:"success"`
	Data           []ProductField `json:"data"`
	AdditionalData AdditionalData `json:"additional_data"`
}

type ProductFieldResponse struct {
	Success        bool           `json:"success"`
	Data           ProductField   `json:"data"`
	AdditionalData AdditionalData `json:"additional_data"`
}

type ProductFieldCreateOptions struct {
	Name      string    `url:"name"`
	FieldType FieldType `url:"field_type"`
	Options   string    `url:"options"`
}

type ProductFieldUpdateOptions struct {
	Name    string `url:"name"`
	Options string `url:"options"`
}

// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/ProductFields/get_productFields
func (s *ProductFieldsService) List() (*ProductFieldsResponse, *Response, error) {
	req, err := s.client.NewRequest(http.MethodGet, "/productFields", nil, nil)

	if err != nil {
		return nil, nil, err
	}

	var record *ProductFieldsResponse

	resp, err := s.client.Do(req, &record)

	if err != nil {
		return nil, resp, err
	}

	return record, resp, nil
}

// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/ProductFields/get_productFields_id
func (s *ProductFieldsService) GetById(id int) (*ProductFieldResponse, *Response, error) {
	uri := fmt.Sprintf("/productFields/%v", id)
	req, err := s.client.NewRequest(http.MethodGet, uri, nil, nil)

	if err != nil {
		return nil, nil, err
	}

	var record *ProductFieldResponse

	resp, err := s.client.Do(req, &record)

	if err != nil {
		return nil, resp, err
	}

	return record, resp, nil
}

// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/ProductFields/post_productFields
func (s *ProductFieldsService) Create(opt *ProductFieldCreateOptions) (*ProductFieldResponse, *Response, error) {
	req, err := s.client.NewRequest(http.MethodPost, "/productFields", nil, opt)

	if err != nil {
		return nil, nil, err
	}

	var record *ProductFieldResponse

	resp, err := s.client.Do(req, &record)

	if err != nil {
		return nil, resp, err
	}

	return record, resp, nil
}

// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/ProductFields/put_productFields_id
func (s *ProductFieldsService) Update(id int, opt *ProductFieldUpdateOptions) (*ProductFieldResponse, *Response, error) {
	uri := fmt.Sprintf("/productFields/%v", id)
	req, err := s.client.NewRequest(http.MethodPut, uri, nil, opt)

	if err != nil {
		return nil, nil, err
	}

	var record *ProductFieldResponse

	resp, err := s.client.Do(req, &record)

	if err != nil {
		return nil, resp, err
	}

	return record, resp, nil
}

// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/ProductFields/delete_productFields
func (s *ProductFieldsService) DeleteMultiple(ids []int) (*Response, error) {
	req, err := s.client.NewRequest(http.MethodDelete, "/productFields", &DeleteMultipleOptions{
		Ids: arrayToString(ids, ","),
	}, nil)

	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}

// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/ProductFields/delete_productFields_id
func (s *ProductFieldsService) Delete(id uint) (*Response, error) {
	uri := fmt.Sprintf("/productFields/%v", id)
	req, err := s.client.NewRequest(http.MethodDelete, uri, nil, nil)

	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}