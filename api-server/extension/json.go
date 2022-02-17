package extension

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"rtuumsfsmga-server/reddit"
	"rtuumsfsmga-server/utils"
)

func ReadJSONData(filepath string) []*reddit.DataPost {
	file, err := os.ReadFile(filepath)
	utils.CheckError(err)

	var read []*reddit.DataPost
	json.Unmarshal(file, &read)

	return read
}

func WriteJSONFile(filename string, v interface{}) {
	file, _ := json.Marshal(v)

	if filepath.Dir(filename) != "." {
		if _, err := os.Stat(filepath.Dir(filename)); os.IsNotExist(err) {
			err := os.Mkdir(filepath.Dir(filename), 0777)
			utils.CheckError(err)
		}
	}

	err := ioutil.WriteFile(filename, file, 0644)
	utils.CheckError(err)
}

func SendJSONResponse(w http.ResponseWriter, i interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(i)
}
