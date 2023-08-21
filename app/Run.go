package app

func Run(url string) []string {
	if url == "" || url == "/" {
		return []string{
			"/",
		}
	}

	urlWithFirsSlash := addFirstSlash(url)

	urlWithoutLastSlash := removeLastSlash(urlWithFirsSlash)

	urlWithSlash := urlWithoutLastSlash + "/"

	return []string{
		urlWithoutLastSlash,
		urlWithSlash,
	}
}

func addFirstSlash(url string) string {
	firstChar := url[0]

	if firstChar == '/' {
		return url
	}

	return "/" + url
}

func removeLastSlash(url string) string {
	lastChar := url[len(url)-1]

	if lastChar == '/' {
		return url[:len(url)-1]
	}

	return url
}
