package structs

type Response struct {
	// ` se escribe con Alt+96
	Code int `json:"code"`
	Body interface{} `json:"body"`
}