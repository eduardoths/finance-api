package http

type Response[T any] struct {
	Data   T      `json:"data"`
	Errors *Error `json:"errors,omitempty"`
}

type Error struct {
	Message string
}
