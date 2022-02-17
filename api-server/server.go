package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"rtuumsfsmga-server/core"
	"rtuumsfsmga-server/extension"
	"rtuumsfsmga-server/handler"
	"rtuumsfsmga-server/reddit"
	"rtuumsfsmga-server/tools"
	"rtuumsfsmga-server/utils"
	"sort"
	"time"
	"github.com/gorilla/context"
	"github.com/gorilla/sessions"
)

var (
  cookie_store *sessions.CookieStore
  id_list []string
  BLOCK_REQUEST_WAIT_INPUT, BLOCK_INPUT_WAIT_REQUEST bool
)

func main() {
  fs := http.FileServer(http.Dir("./build"))
  http.Handle("/", http.StripPrefix("/build", fs))
  http.HandleFunc("/api/login", handler.PermissionMiddleware(LoginHandler))
  http.HandleFunc("/api/callback", CallbackHandler)
	http.HandleFunc("/api/all", handler.PermissionMiddleware(AllHandler))
	http.HandleFunc("/api/fetch", handler.PermissionMiddleware(FetchHandler))
	http.HandleFunc("/api/convert", handler.PermissionMiddleware(ConvertHandler))
	http.HandleFunc("/api/unsave", handler.PermissionMiddleware(UnsaveHandler))
	http.HandleFunc("/api/sync", handler.PermissionMiddleware(SyncHandler))
  http.HandleFunc("/api/logout", handler.PermissionMiddleware(LogoutHandler))
  http.HandleFunc("/favicon.ico", doNothing)
  log.Fatal(http.ListenAndServe(utils.PORT, context.ClearHandler(http.DefaultServeMux)))
}

func AllHandler(w http.ResponseWriter, r *http.Request) {

  session, _ := cookie_store.Get(r, "r_Cookie")
  as := tools.GetSession(session)
  var read []*reddit.DataPost

  if as == nil {
    w.WriteHeader(http.StatusNotFound)
    fmt.Println("Session not found")
  } else {
    if _, err := os.Stat(as.SavedDataPath); err == nil {
      read = extension.ReadJSONData(as.SavedDataPath)
    } else if _, err := os.Stat(as.SavedFetchPath); err == nil {
      read = extension.ReadJSONData(as.SavedFetchPath)
    } else {
      w.WriteHeader(http.StatusOK)
      return
  }

  extension.Reverse(read)
  extension.SendJSONResponse(w, read)

  }
}

func FetchHandler(w http.ResponseWriter, r *http.Request) {
  
  session, _ := cookie_store.Get(r, "r_Cookie")
  
  reddit.CheckToken(session, w, r)

  as := tools.GetSession(session)
  
  start := time.Now()

  if _, err := os.Stat(as.SavedDataPath); err == nil {
    fmt.Println("Found BIG file, fetching partially..")
    old_file := extension.ReadJSONData(as.SavedDataPath) 
    new_file := reddit.GetSaved(as, false)
    extension.Reverse(new_file)
    merge := extension.Merge(old_file, new_file, false)
    extension.WriteJSONFile(as.SavedDataPath, merge)
  } else if _, err := os.Stat(as.SavedFetchPath); err == nil {
    fmt.Println("Found saved file, fetching partially..")
    old_file := extension.ReadJSONData(as.SavedFetchPath) 
    new_file := reddit.GetSaved(as, false)
    extension.Reverse(new_file)
    merge := extension.Merge(old_file, new_file, false)
    extension.WriteJSONFile(as.SavedFetchPath, merge)
  } else {
    fetch := reddit.GetSaved(as, true)
    extension.Reverse(fetch)
    extension.WriteJSONFile(as.SavedFetchPath, fetch)
  }

  fmt.Println(time.Since(start))

}

func ConvertHandler(w http.ResponseWriter, r *http.Request) {
  
  session, _ := cookie_store.Get(r, "r_Cookie")
  reddit.CheckToken(session, w, r)

  as := tools.GetSession(session)

  if _, err := os.Stat(as.SavedDataPath); err == nil {
    fmt.Println("Already converted")
    w.WriteHeader(http.StatusOK)
    return
  }

  if _, err := os.Stat(as.SourcePostPath); err == nil {
		if _, err := os.Stat(as.SourceCommentPath); err == nil {
			fmt.Println("Converting posts..")
			convertedPost := core.ConvertData(as, true)
			fmt.Println("Converting comments..")
			convertedComment := core.ConvertData(as, false)
			mergedSlice := extension.Merge(convertedPost, convertedComment, false)
			sort.Sort(reddit.ByTimestamp(mergedSlice))
			
      if _, err := os.Stat(as.SavedFetchPath); err == nil {
				fmt.Println("Found saved files, merging...")
        fetch := extension.ReadJSONData(as.SavedFetchPath)
        mergedFull := extension.Merge(fetch, mergedSlice, true)
				extension.WriteJSONFile(as.SavedDataPath, mergedFull)
        err = os.Remove(as.SavedFetchPath)
        utils.CheckError(err)
        w.WriteHeader(http.StatusOK)
			} else {
				extension.WriteJSONFile(as.SavedDataPath, mergedSlice)
        w.WriteHeader(http.StatusOK)
			}
		}
	} else {
    fmt.Println("Source files not found!")
    w.WriteHeader(http.StatusNoContent)
  }

}

func UnsaveHandler(w http.ResponseWriter, r *http.Request) {
  
  r.Header.Set("Content-Type", "application/x-www-form-urlencoded")

  var empty struct{}
  not_saved_sets := make(map[string]struct{})
  
  session, _ := cookie_store.Get(r, "r_Cookie")
  as := tools.GetSession(session)

  id := r.URL.Query().Get("name")

  if (BLOCK_INPUT_WAIT_REQUEST) { return }
  if (BLOCK_REQUEST_WAIT_INPUT && !BLOCK_INPUT_WAIT_REQUEST) {
    id_list = append(id_list, id)
    fmt.Println(id_list)
  } else {
    id_list = append(id_list, id)
    BLOCK_REQUEST_WAIT_INPUT = true
    fmt.Println("waiting for more input...")
    time.Sleep(60*time.Second)
    fmt.Println("will start sending request...")

    start1 := time.Now()
    BLOCK_INPUT_WAIT_REQUEST = true
    reddit.CheckToken(session, w, r)
    for _, id := range id_list {
      reddit.Unsave(as, w, r, id)
      not_saved_sets[id] = empty
    }
    fmt.Println(time.Since(start1))

    if _, err := os.Stat(as.SavedFetchPath); err == nil {
      core.DeleteDataFromFile(not_saved_sets, as.SavedFetchPath)
    } else if _, err := os.Stat(as.SavedDataPath); err == nil {
      core.DeleteDataFromFile(not_saved_sets, as.SavedDataPath)
    } else {
      utils.CheckError(err)
    }

    BLOCK_REQUEST_WAIT_INPUT = false
    BLOCK_INPUT_WAIT_REQUEST = false
    id_list = nil
  } 

}

func SyncHandler(w http.ResponseWriter, r *http.Request) {
  session, _ := cookie_store.Get(r, "r_Cookie")
  as := tools.GetSession(session)

  var (
    source_file, slice_to_compare []*reddit.DataPost
    filename string
    id_list []string
    empty struct{}
  )

  not_saved_sets := make(map[string]struct{})
  
  if _, err := os.Stat(as.SavedFetchPath); err == nil {
    source_file = extension.ReadJSONData(as.SavedFetchPath)
    filename = as.SavedFetchPath 
  } else if _, err := os.Stat(as.SavedDataPath); err == nil {
    source_file = extension.ReadJSONData(as.SavedDataPath)
    filename = as.SavedDataPath 
  } else {
    utils.CheckError(err)
  }

  for _, v := range source_file{
    id_list = append(id_list, v.Fullname)
  }

  id_chunks := extension.DivideSlice(id_list)

  reddit.CheckToken(session, w, r)
  
  start := time.Now()
  slice_to_compare = reddit.GetInfo(as, id_chunks, false)
  fmt.Println(time.Since(start))
  
  for _, v := range slice_to_compare {
    if (v.Saved == false) {
      not_saved_sets[v.Id] = empty
    }
  }

  core.DeleteDataFromFile(not_saved_sets, filename) 

}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	
  URL := fmt.Sprintf("https://www.reddit.com/api/v1/authorize?client_id=%s"+
		                 "&response_type=code&state=%s&redirect_uri=%s&"+
		                 "duration=permanent&scope=%s",
		                 utils.CLIENT_ID, utils.STATE, utils.CALLLBACK_URI, utils.SCOPE)

  var json_resp = map[string]string{"auth_url":URL}
  extension.SendJSONResponse(w, json_resp)

}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
  
  session, _ := cookie_store.Get(r, "r_Cookie")
  session.Options.MaxAge = -1
  session.Save(r, w)

}

func CallbackHandler(w http.ResponseWriter, r *http.Request) {
  
  code := r.URL.Query().Get("code")
	token := reddit.GetToken(code)

  session, _ := cookie_store.Get(r, "r_Cookie")
  username := reddit.GetUsername(token.AccessToken)
  
  var sobj = map[string]string{
    "REFRESH_TOKEN":token.RefreshToken,
    "ACCESS_TOKEN":token.AccessToken,
    "USERNAME":username,
  }

  tools.SaveSession(session, w, r, sobj)
  http.Redirect(w, r, utils.URL_HOMEPAGE, http.StatusSeeOther)

}

func doNothing(w http.ResponseWriter, r *http.Request) {}
