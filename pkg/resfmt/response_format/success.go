package response_format

type Success struct {
	StatusCode int                    `json:"-"`
	Status     string                 `json:"status" binding:"required"`
	Data       interface{}            `json:"data,omitempty"`
	Metadata   map[string]interface{} `json:"metadata,omitempty"`
}

func NewSuccess(data interface{}) *Success {
	return &Success{StatusCode: 200, Status: "success", Data: data}
}

func NewSuccessWithMetadata(data interface{}, metadata map[string]interface{}) *Success {
	return &Success{StatusCode: 200, Status: "success", Data: data, Metadata: metadata}
}
