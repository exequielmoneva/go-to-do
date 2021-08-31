package views

type TodoResponse struct {
	// ` == Alt+96
	Todo string `json:"todo"`
	Name string `json:"name"`
	Id   string
}
