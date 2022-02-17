package core

import (
	"encoding/csv"
	"os"
	"rtuumsfsmga-server/extension"
	"rtuumsfsmga-server/reddit"
	"rtuumsfsmga-server/tools"
	"rtuumsfsmga-server/utils"
)

func ConvertData(as *tools.AppSession, submission bool) []*reddit.DataPost {

	var (
		source_slice  []string
		source_chunks [][]string
		info_prefix   string
		source_file   *os.File
		err           error
		keep_link     = false
	)

	if submission {

		source_file, err = os.Open(as.SourcePostPath)
		utils.CheckError(err)

		info_prefix = `t3_`

	} else {

		source_file, err = os.Open(as.SourceCommentPath)
		utils.CheckError(err)

		info_prefix = `t1_`
		keep_link = true

	}

	defer source_file.Close()
	reader := csv.NewReader(source_file)
	read_csv, _ := reader.ReadAll()

	for _, v := range read_csv[1:] {
		source_slice = append(source_slice, info_prefix+v[0])
	}

	source_chunks = extension.DivideSlice(source_slice)

	var converted []*reddit.DataPost
	converted = reddit.GetInfo(as, source_chunks, keep_link)

	if !submission {
		var link_ids []string
		var parents []*reddit.DataPost

		for _, v := range converted {
			link_ids = append(link_ids, v.Link_Id)
		}

		id_chunks := extension.DivideSlice(link_ids)
		parents = reddit.GetInfo(as, id_chunks, keep_link)

		for i, v := range converted {
			v.Title = parents[i].Title
			v.Link_Id = ""
		}
	}

	return converted

}

func DeleteDataFromFile(comparison_sets map[string]struct{}, filename string) {

	var substracted_slice []*reddit.DataPost
	read := extension.ReadJSONData(filename)

	for _, v := range read {
		if _, ok := comparison_sets[v.Fullname]; !ok {
			substracted_slice = append(substracted_slice, v)
		}
	}

	extension.WriteJSONFile(filename, substracted_slice)

}
