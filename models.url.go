package main

// just store in a map for the exercise
var urls = make(map[string]string)

func retrieveURL(shortURL string) string {
	return urls[shortURL]
}

func addURL(shortURL, fullURL string) {
	urls[shortURL] = fullURL
}
