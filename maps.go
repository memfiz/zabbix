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
type Map struct {
	SysMapID  string       `json:"sysmapid,omitempty"`
	Error     string       `json:"error"`
	Selements MapSelements `json:"selements,omitempty"`
}

type MapSelement struct {
	SelementID string `json:"selementid,omitempty"`
	ElementID  string `json:"elementid,omitempty"`
}

type Maps []Map
type MapSelements []MapSelement

// Wrapper for host.get: https://www.zabbix.com/documentation/3.2/manual/appendix/api/map/get
func (api *API) HostOnMapGet(params Params) (res Maps, err error) {
	if _, present := params["output"]; !present {
		params["output"] = "extend"
	}
	response, err := api.CallWithError("map.get", params)
	if err != nil {
		return
	}
	reflector.MapsToStructs2(response.Result.([]interface{}), &res, reflector.Strconv, "json")
	return
}
