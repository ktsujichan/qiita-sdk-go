package qiita

// Represents a group on Qiita:Team
type Group struct {
	CreatedAt string `json:"created_at"`
	Id        uint   `json:"id"`
	Name      string `json:"name"`
	Private   bool   `json:"private"`
	UpdatedAt string `json:"updated_at"`
	UrlName   string `json:"url_name"`
}
