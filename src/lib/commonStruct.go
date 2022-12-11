package lib

type AuthorBaseInfo struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	Desc string `json:"desc"`
}

type LoginRespInfo struct {
	*AuthorBaseInfo
	Token string `json:"token"`
}
