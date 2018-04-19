package elastic

type Cha struct {
	Id      string `json:"id"`
	C       string `json:"c"`
	H       string `json:"h"`
	C_alt   string `json:"c_alt,omitempty"`
	H_alt   string `json:"h_alt,omitempty"`
	Lang    string `json:"lang"`
	Created string `json:"created,omitempty"`
	User    `json:"user"`
}

type User struct {
	Id string `json:"id"`
}

type Index struct {
	Name string
	Type string
}

/**
index
*/
var IndexByCha = Index{
	Name: "cha.bycha",
	Type: "cha",
}
