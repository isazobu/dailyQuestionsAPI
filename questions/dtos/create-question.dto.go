package dto

type CreateQuestion struct {
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Difficulty  string   `json:"difficulty"`
	Category    string   `json:"category"`
	Tags        []string `json:"tags"`
}
