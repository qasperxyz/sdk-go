package middlewares

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"strings"

	"github.com/go-http-utils/headers"
	"github.com/google/uuid"
	e "github.com/qasperxyz/sdk-go/errors"
	"github.com/qasperxyz/sdk-go/types"

	"github.com/gin-gonic/gin"
)

type AuthorizationPayload struct {
	Entity string `json:"entity"`
	Action string `json:"action"`
}

// get access token and verify user session
func AuthorizationMiddleware(c *gin.Context) {
	// get token from request
	token := c.Request.Header.Get(headers.Authorization)
	if strings.Compare(token, "") == 0 {
		err := errors.New("missing " + headers.Authorization + " header")
		e.ThrowBadRequest(c, err)
		return
	}
	env := os.Getenv("ENV")
	if env == "PRODUCTION" {
		r := regexp.MustCompile(`/api/v[1-9]/(\w+)/(\w+)`)
		result := r.FindStringSubmatch(c.FullPath())

		if len(result) >= 3 {
			entity := result[1]
			action := result[2]

			// request authorization to identity provider
			payload := []byte(fmt.Sprintf(`{"entity": %s, "action": %s}`, entity, action))
			request, err := http.NewRequest("POST", types.IDENTITY_PROVIDER_AUTHORIZATION_URL, bytes.NewBuffer(payload))
			if err != nil {
				e.ThrowInternalServerError(c, err)
				return
			}
			request.Header.Set("Authorization", c.Request.Header.Get(headers.Authorization))
			request.Header.Set("Content-Type", "application/json")

			client := http.DefaultClient

			response, err := client.Do(request)
			if err != nil {
				e.ThrowInternalServerError(c, err)
				return
			}
			defer response.Body.Close()

			// Read the response body
			jsonData, err := io.ReadAll(response.Body)
			if err != nil {
				e.ThrowInternalServerError(c, err)
				return
			}
			session := types.Session{}
			err = json.Unmarshal([]byte(jsonData), &session)
			if err != nil {
				e.ThrowInternalServerError(c, err)
				return
			}
			// return user session to next route
			c.Set(types.USER_SESSION_CONTEXT_KEY, session)
			c.Next()
		} else {
			e.ThrowBadRequest(c, errors.New("Path not valid"))
		}
	} else {
		// create dummy userSession
		userSession := types.Session{
			Id:       uuid.New().String(),
			Username: "admin",
		}
		// set token data to context
		c.Set(types.USER_SESSION_CONTEXT_KEY, userSession)
		c.Next()
	}
}
