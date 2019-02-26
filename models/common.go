package models

type LoginVM struct {
	Email   string `json:"email"`
	FBToken string `json:"fbToken"`
	FBId    string `json:"fbId"`
}

type FaceBookUser struct {
	Name      string  `json:"name"`
	FirstName string  `json:"first_name"`
	LastName  string  `json:"last_name"`
	Email     string  `json:"email"`
	Picture   Picture `json:"picture"`
}

type Picture struct {
	Data Data `json:"data"`
}

type Data struct {
	Width        int64  `json:"width"`
	Height       int64  `json:"height"`
	IsSilhouette bool   `json:"is_silhouette"`
	Url          string `json:"url"`
}
