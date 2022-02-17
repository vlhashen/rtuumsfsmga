package main

import (
	"encoding/json"
	"fmt"
	"os"
	"rtuumsfsmga-server/extension"
	"rtuumsfsmga-server/tools"
	"rtuumsfsmga-server/utils"
	"github.com/gorilla/sessions"
)

func init() {
  if _, err := os.Stat("config.json"); err == nil {
    
    fmt.Println("Reading config..")
    var read map[string]interface{}
    file, err := os.ReadFile("config.json")
	  utils.CheckError(err) 
    json.Unmarshal(file, &read)
    
    secret_key := read["SECRET_KEY"].(string)
    fmt.Println(secret_key)
    cookie_store = sessions.NewCookieStore([]byte(secret_key))
  
  } else {
	  
    fmt.Println("Creating config..")
    secret_key := tools.GenerateSecretKey()
    var s = map[string]string{"SECRET_KEY":secret_key}
    extension.WriteJSONFile("config.json", s)
    fmt.Println(secret_key)
    cookie_store = sessions.NewCookieStore([]byte(secret_key))

  }

  fmt.Printf("Open: %v\n", utils.URL_HOMEPAGE)

}
