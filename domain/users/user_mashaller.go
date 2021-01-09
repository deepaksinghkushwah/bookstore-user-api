package users

import "encoding/json"

// Marshal return user based on request type
func (p *User) Marshal(isPublic bool) interface{} {
	if isPublic == true {
		return PublicUser{
			ID:          p.ID,
			DateCreated: p.DateCreated,
			Status:      p.Status,
		}
	}

	userJSON, _ := json.Marshal(p)
	var privateUser PrivateUser
	json.Unmarshal(userJSON, &privateUser)
	return privateUser
}

// Users is [] of users
type Users []User

// Marshal to multiple users
func (p Users) Marshal(isPublic bool) []interface{} {
	result := make([]interface{}, len(p))
	for index, user := range p {
		result[index] = user.Marshal(isPublic)
	}
	return result
}
