package common

const (
	DELETED = "Deleted"
	DOING   = "Doing"
	DONE    = "Done"
)

type TokenPayload struct {
	UId   int    `json:"user_id"`
	URole string `json:"role"`
}

func (p TokenPayload) UserId() int {
	return p.UId
}

func (p TokenPayload) Role() string {
	return p.URole
}
