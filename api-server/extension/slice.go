package extension

import "rtuumsfsmga-server/reddit"

func Reverse(slice []*reddit.DataPost) {
	if len(slice) > 2 {
		for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
			slice[i], slice[j] = slice[j], slice[i]
		}
	}
}

func Merge(first, second []*reddit.DataPost, prepend bool) []*reddit.DataPost {
	if len(first) > 0 && len(second) > 0 {

		var empty struct{}
		sets := make(map[string]struct{})
		for _, v := range first {
			sets[v.Id] = empty
		}

		if prepend {
			Reverse(second)
		}

		for _, v := range second {
			if _, ok := sets[v.Id]; !ok {
				if prepend {
					first = append([]*reddit.DataPost{v}, first...)
				} else {
					first = append(first, v)
				}
			}
		}
	}

	return first
}

func DivideSlice(s []string) [][]string {
	var chunks [][]string
	for i := 0; i < len(s); i += 100 {
		last := i + 100

		if last > len(s) {
			last = len(s)
		}

		chunks = append(chunks, s[i:last])

	}
	return chunks
}
