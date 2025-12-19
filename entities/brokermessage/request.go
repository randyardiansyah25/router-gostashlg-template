package brokermessage

type Metadata[T any] struct {
	Action       string `json:"action"`
	Description  string `json:"description"`
	Identity     string `json:"identity"`
	RegisteredId string `json:"registered_id"`
	Channel      string `json:"channel"`
	Data         T      `json:"data"`
}
