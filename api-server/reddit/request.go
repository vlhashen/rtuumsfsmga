package reddit

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"rtuumsfsmga-server/tools"
	"rtuumsfsmga-server/utils"
	"time"
)

func GetUsername(access_token string) string {
	c := http.Client{Timeout: time.Duration(20) * time.Second}
	req, _ := http.NewRequest("GET", "https://oauth.reddit.com/api/v1/me", nil)
	req.Header.Set("User-Agent", utils.HEADER_USER)
	req.Header.Set("Authorization", "bearer "+access_token)

	resp, err := c.Do(req)
	utils.CheckError(err)
	defer resp.Body.Close()

	var res map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&res)

	username := res["name"].(string)
	return username
}

func GetData(URL string, header http.Header, query url.Values, batch, keep_link bool) []*DataPost {

	var (
		after string
		res   Response
		data  []*DataPost
	)

	req, err := http.NewRequest("GET", URL, nil)
	utils.CheckError(err)
	req.Header = header
	req.URL.RawQuery = query.Encode()

	c := http.Client{Timeout: time.Duration(60) * time.Second}

	if batch {

		for i := 0; i < 10; i++ {
			resp, err := c.Do(req)
			utils.CheckError(err)
			defer resp.Body.Close()

			json.NewDecoder(resp.Body).Decode(&res)

			temp := res.ParseData(keep_link)
			data = append(data, temp...)

			after = data[len(data)-1].Fullname
			q := req.URL.Query()
			q.Set("after", after)
			req.URL.RawQuery = q.Encode()

		}

	} else {

		resp, err := c.Do(req)
		utils.CheckError(err)
		defer resp.Body.Close()

		json.NewDecoder(resp.Body).Decode(&res)

		data = res.ParseData(keep_link)

	}

	return data

}

func GetSaved(as *tools.AppSession, batch bool) []*DataPost {

	header := http.Header{
		"User-Agent":    []string{utils.HEADER_USER},
		"Authorization": []string{"bearer " + as.AccessToken},
	}

	query := url.Values{
		"limit":    []string{"100"},
		"raw_json": []string{"1"},
		"after":    []string{""},
	}

	URL := "https://oauth.reddit.com/user/" + as.Username + "/saved"

	data := GetData(URL, header, query, batch, false)

	return data

}

func GetInfo(s *tools.AppSession, chunks [][]string, keep_link bool) []*DataPost {

	var info []*DataPost

	header := http.Header{
		"User-Agent":    []string{utils.HEADER_USER},
		"Authorization": []string{"bearer " + s.AccessToken},
	}

	query := url.Values{
		"raw_json": []string{"1"},
		"id":       []string{},
	}

	for _, chunk := range chunks {
		URL := utils.URL_REDDIT_INFO
		var id string

		for _, v := range chunk {
			id = id + v + `,`
		}

		query["id"] = []string{id}

		data := GetData(URL, header, query, false, keep_link)
		info = append(info, data...)
	}

	return info
}

func Unsave(as *tools.AppSession, w http.ResponseWriter, r *http.Request, link_id string) {

	c := http.Client{Timeout: time.Duration(20) * time.Second}

	req, _ := http.NewRequest("POST", "https://oauth.reddit.com/api/unsave", nil)
	q := req.URL.Query()
	q.Add("id", link_id)
	req.URL.RawQuery = q.Encode()

	req.Header = http.Header{
		"User-Agent":    []string{utils.HEADER_USER},
		"Authorization": []string{"bearer " + as.AccessToken},
	}

	resp, err := c.Do(req)
	utils.CheckError(err)
	defer resp.Body.Close()

	var json_resp map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&json_resp)
	utils.CheckError(err)

	if json_resp["error"] != nil {
		fmt.Printf("%v: %v\n", json_resp["error"], json_resp["message"])
	}

}
