package lib

type LoginRespInfo struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Desc  string `json:"desc"`
	Token string `json:"token"`
}
