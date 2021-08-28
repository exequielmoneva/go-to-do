package views

type Response struct {
	// ` == Alt+96
	Body interface{} `json:"body"`
}

type Todo struct {
	Todo string `json:"todo"`
	Name string `json:"name"`
	Id   string
}
