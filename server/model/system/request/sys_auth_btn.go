package request

type SysAuthBtnReq struct {
	MenuID   uint   `json:"menuID"`
	AuthId   uint   `json:"authId"`
	Selected []uint `json:"selected"`
}
