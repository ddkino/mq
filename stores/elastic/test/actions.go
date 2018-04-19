package mapping

import "time"

type Cha struct {
	C       string    `json:"c"`
	H       string    `json:"h"`
	C_alt   string    `json:"c_alt,omitempty"`
	H_alt   string    `json:"h_alt,omitempty"`
	Created time.Time `json:"created,omitempty"`
	User    `json:"user"`
}

type User struct {
	Name string `json:"name"`
	Id   string `json:"id"`
}
