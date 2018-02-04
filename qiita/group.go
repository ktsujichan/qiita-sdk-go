package qiita

// Represents a group on Qiita:Team
type Group struct {
	CreatedAt string `json:"created_at"`
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	Private   bool   `json:"private"`
	UpdatedAt string `json:"updated_at"`
	URLName   string `json:"url_name"`
}
