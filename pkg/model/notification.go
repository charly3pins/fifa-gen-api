package model

type Notification struct {
	FriendRequests []User `json:"friendRequests"`
}
