package reddit

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"rtuumsfsmga-server/tools"
	"rtuumsfsmga-server/utils"
	"strings"
	"time"

	"github.com/gorilla/sessions"
)

type TokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func GetToken(code string) TokenResponse {
	c := http.Client{Timeout: time.Duration(20) * time.Second}
	
  data := url.Values{
    "grant_type":[]string{"authorization_code"},
    "code":[]string{code},
    "redirect_uri":[]string{utils.CALLLBACK_URI},
  }	
  reader := strings.NewReader(data.Encode())
  
  URL := "https://www.reddit.com/api/v1/access_token"
  
  req, err := http.NewRequest("POST", URL, reader)
  utils.CheckError(err)
  req.Header = http.Header{
     "User-Agent":[]string{utils.HEADER_USER},
  }
  req.SetBasicAuth(utils.CLIENT_ID, "")

	resp, err := c.Do(req)
  utils.CheckError(err)
  defer resp.Body.Close()
	
	var token TokenResponse
  err = json.NewDecoder(resp.Body).Decode(&token)
  utils.CheckError(err)

	return token
}

func GetNewToken(refresh_token string) string {
  c := http.Client{Timeout: time.Duration(20) * time.Second}
	
  data := url.Values{
    "grant_type":[]string{"refresh_token"},
    "refresh_token":[]string{refresh_token},
  }
	reader := strings.NewReader(data.Encode())
  
  URL := "https://www.reddit.com/api/v1/access_token"
  
  req, err := http.NewRequest("POST", URL, reader)
  utils.CheckError(err)
  req.Header = http.Header{
     "User-Agent":[]string{utils.HEADER_USER},
  }
  req.SetBasicAuth(utils.CLIENT_ID, "")
	
	resp, err := c.Do(req)
	utils.CheckError(err)
  defer resp.Body.Close()

	var token TokenResponse
  err = json.NewDecoder(resp.Body).Decode(&token)
  utils.CheckError(err)

	return token.AccessToken
}

func CheckToken(session *sessions.Session, w http.ResponseWriter, r *http.Request) {
  
  c := http.Client{Timeout: time.Duration(20) * time.Second}
  
  req, err := http.NewRequest("GET", "https://oauth.reddit.com/api/v1/me", nil)
  utils.CheckError(err)
  req.Header = http.Header{
     "User-Agent":[]string{utils.HEADER_USER},
     "Authorization":[]string{"bearer "+session.Values["ACCESS_TOKEN"].(string)},
  }

  resp, err := c.Do(req)
  utils.CheckError(err)

  if resp.StatusCode == 401 {
    fmt.Println("creating new token..")
		new_token := GetNewToken(session.Values["REFRESH_TOKEN"].(string))
    sobj := map[string]string{ "ACCESS_TOKEN":new_token}
    tools.SaveSession(session, w, r, sobj)
	}
  
}

