package internal

func (user *User) exitRoom() {
	userID := user.userID()
	if room, ok := userIDRooms[userID]; ok {
		room.Exit(userID)
	}
}