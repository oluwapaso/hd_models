package models

type RepoResponse struct {
	Success        bool   `json:"success"`
	Num_of_results int    `json:"num_of_results"`
	Error_message  string `json:"error_message"`
	//Data           ImportedDealFiles `json:"data"` //ImportedDealFiles //[]map[string]interface{}
}
type ScannedQryResp struct {
	Resp RepoResponse
	Data [][]interface{}
}
type ScannedSingleQryResp struct {
	Resp RepoResponse
	Data []interface{}
}
