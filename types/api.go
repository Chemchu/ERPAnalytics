package types

type APIResponse struct {
	Message    string `json:"message"`
	Successful bool   `json:"successful"`
	Data       string `json:"data,omitempty"`
}

type APIData struct {
	Message    string `json:"message,omitempty"`
	Successful bool   `json:"successful,omitempty"`
}
