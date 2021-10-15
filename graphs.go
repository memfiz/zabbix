package zabbix

import "github.com/AlekSi/reflector"

// type (
// 	AvailableType int
// 	StatusType    int
// )

// const (
// 	Available   AvailableType = 1
// 	Unavailable AvailableType = 2

// 	Monitored   StatusType = 0
// 	Unmonitored StatusType = 1
// )

// https://www.zabbix.com/documentation/3.2/manual/appendix/api/map/definitions
type Graph struct {
	GraphID string `json:"graphid,omitempty"`
	Error   string `json:"error"`
	Name    string `json:"name,omitempty"`
}

type Graphs []Graph

// Wrapper for host.get: https://www.zabbix.com/documentation/3.2/manual/appendix/api/map/get
func (api *API) GraphsGet(params Params) (res Graphs, err error) {
	if _, present := params["output"]; !present {
		params["output"] = "extend"
	}
	response, err := api.CallWithError("graph.get", params)
	if err != nil {
		return
	}
	reflector.MapsToStructs2(response.Result.([]interface{}), &res, reflector.Strconv, "json")
	return
}
