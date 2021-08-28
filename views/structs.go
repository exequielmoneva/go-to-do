package views

type Response struct {
	// ` se escribe con Alt+96
	Body interface{} `json:"body"`
}

type Todo struct {
	Todo string `json:"todo"`
	Name string `json:"name"`
	Id   string
}
