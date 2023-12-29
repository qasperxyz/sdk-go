package types

type Session struct {
	Id       string `json:"id"`
	Username string `json:"username"`
}

func (h *Session) fromMap(m map[string]string) {
	h.Id = m["id"]
	h.Username = m["username"]
}

func (h *Session) toMap() map[string]string {
	return map[string]string{
		"id":       h.Id,
		"username": h.Username,
	}
}
