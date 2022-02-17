package tools

import (
	"github.com/gorilla/sessions"
	"net/http"
	"rtuumsfsmga-server/utils"
)

type AppSession struct {
	Username          string
	RefreshToken      string
	AccessToken       string
	SavedDataPath     string
	SavedFetchPath    string
	SourcePostPath    string
	SourceCommentPath string
}

func SaveSession(session *sessions.Session, w http.ResponseWriter, r *http.Request, sobj map[string]string) {

	for k, v := range sobj {
		session.Values[k] = v
	}
	session.Options.MaxAge = 31536000
	session.Save(r, w)

}

func GetSession(session *sessions.Session) *AppSession {

	if session.Values["USERNAME"] != nil {

		username := session.Values["USERNAME"].(string)

		return &AppSession{
			Username:          username,
			RefreshToken:      session.Values["REFRESH_TOKEN"].(string),
			AccessToken:       session.Values["ACCESS_TOKEN"].(string),
			SavedDataPath:     username + "/" + utils.SAVED_DATA,
			SavedFetchPath:    username + "/" + utils.SAVED_FETCH,
			SourcePostPath:    username + "/" + utils.SOURCE_POST,
			SourceCommentPath: username + "/" + utils.SOURCE_COMMENT,
		}

	} else {
		return nil
	}

}
