package users

type UserUpdate struct {
	Name     string `json:"name"`
	Address  string `json:"address"`
	Email    string `json:"email"`
	PhotoURL string `json:"photo_url"`
}
