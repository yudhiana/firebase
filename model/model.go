package model

// Profile ...
type Profile struct {
	Name    string `json:"name"`
	Address string `json:"address"`
	Phone   string `json:"phone"`
	Email   string `json:"email"`
	Bio     string `json:"bio"`
}
