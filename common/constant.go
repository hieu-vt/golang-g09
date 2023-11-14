package common

import "fmt"

const (
	DELETED = "Deleted"
	DOING   = "Doing"
	DONE    = "Done"
)

const CurrentUser = "CurrentUser"

type DbType int

const (
	DbTypeItem DbType = 1
	DbTypeUser DbType = 2
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

type Requester interface {
	GetUserId() int
	GetEmail() string
	GetRole() string
}

func IsAdmin(requester Requester) bool {
	return requester.GetRole() == "admin" || requester.GetRole() == "mod"
}

const (
	PluginDBMain      = "PluginDBMain"
	PluginJwtProvider = "PluginJwtProvider"
	PluginPubSub      = "PluginPubSub"

	TopicUserLikeItem   = "TopicUserLikeItem"
	TopicUserUnlikeItem = "TopicUserUnlikeItem"
)

func Recovery() {
	if r := recover(); r != nil {
		fmt.Println("Recovered:", r)
	}
}
