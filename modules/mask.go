package modules

type Mask struct {
	List []One `json:"list"`
}

type One struct {
	Id   string `json:"id"`
	Name string `json:"Name"`
}
