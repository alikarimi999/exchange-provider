package utils

func JoinSafe(url, path string) string {

	if url[len(url)-1] == '/' {
		url = url[:len(url)-1]
	}

	if path[0] != '/' {
		path = "/" + path
	}

	return url + path
}
