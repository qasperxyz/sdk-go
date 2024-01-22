package managers

import "github.com/qasperxyz/sdk-go/types"

func SessionToMap(session types.Session) map[string]interface{} {
	sessionMap := make(map[string]interface{})
	sessionMap["id"] = session.Id
	sessionMap["userId"] = session.UserId
	sessionMap["email"] = session.Email
	return sessionMap
}

func SessionFromMap(sessionMap map[string]string) types.Session {
	session := types.Session{}
	session.Id = sessionMap["id"]
	session.UserId = sessionMap["userId"]
	session.Email = sessionMap["email"]
	return session
}
