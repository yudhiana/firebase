package model

// Firebase ...
type Firebase interface {
	FindAll(reference string) ([]Profile, error)
	Set(data map[string]interface{}) error
	Update(data map[string]interface{}) error
}

// FirebaseModel ...
type FirebaseModel struct {
	// Conf
}

// NewFirebase ...
// func NewFirebase() Firebase {
// 	// return Firebase
// }
