package views

type Response struct {
	// ` == Alt+96
	Body interface{} `json:"body"`
}

type TodoResponse struct {
	Todo string `json:"todo"`
	Name string `json:"name"`
	Id   string
}
