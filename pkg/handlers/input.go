package handlers

type CreateResumeInput struct {
	Title string `json:"title"`
	Description string `json:"description"`
}

type UpdateResumeInput struct {
	Field string `json:"field"`
	Value string `json:"value"`
}