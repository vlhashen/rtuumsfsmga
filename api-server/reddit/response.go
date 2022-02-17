package reddit

type Response struct {
	Kind string `json:"kind"`
	Data Data   `json:"data"`
}

type Data struct {
	After    string     `json:"after"`
	Children []Children `json:"children"`
}

type Children struct {
	Kind     string   `json:"kind"`
	DataPost DataPost `json:"data"`
}

type DataPost struct {
	Id          string  `json:"id"`
	Subreddit   string  `json:"subreddit"`
	Author      string  `json:"author"`
	Title       string  `json:"title"`
	Created_Utc float64 `json:"created_utc"`
	Saved       bool    `json:"saved"`
	Selftext    string  `json:"selftext"`
	Body        string  `json:"body,omitempty"`
	Url         string  `json:"url,omitempty"`
	Fullname    string  `json:"name"`
	Permalink   string  `json:"permalink"`
	Link_Title  string  `json:"link_title,omitempty"`
	Link_Id     string  `json:"link_id,omitempty"`
}

func (res *Response) ParseData(keep_link bool) []*DataPost{

  var data []*DataPost
	col := res.Data.Children
  for _, v := range col {
		dp := v.DataPost
		dp.Permalink = "https://www.reddit.com" + dp.Permalink
		if len(dp.Selftext) == 0 && len(dp.Body) > 0 {
      dp.Selftext = dp.Body
      dp.Body = ""
    }
    if len(dp.Selftext) >= 500 { dp.Selftext = dp.Selftext[:500] }
		if len(dp.Title) == 0 {
			dp.Title = dp.Link_Title
			dp.Link_Title = ""
			if !keep_link { dp.Link_Id = "" }
		}
		data = append(data, &dp)
	}
	return data

}
