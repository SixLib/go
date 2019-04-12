package models

type SysInfo struct {
	Key   string `json:"key"`
	Name  string `json:"name"`
	Value string `json:"value"`
}
type SysInfos []SysInfo
