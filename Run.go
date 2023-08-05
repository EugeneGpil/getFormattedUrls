package getFormattedUrls

import "github.com/EugeneGpil/getFormattedUrls/app"

// Format url for "net/http.ServeMux.HandleFunc"
// Both returned urls should be provided to HandleFunc function

// This package was created for preventing "forget trailing slash" error
// For example "http://test.local/some/route" instead of "http://test.local/some/route/"
// and "http://test.local/some/route/" instead of "http://test.local/some/route"

// And to prevent "forget first slash" when providing url for HandlerFunc function
// For example "some/route/" instead of "/some/route/"

// Input url may be with or without trailing slash "/"
// Input url may be with or without first slash "/"

// Returns 2 urls

// Both with first slash "/"

// First without trailing slash "/"
// And second with trailing slash "/"

// Examples
// Input "some/route".		Output "/some/route", "/some/route/"
// Input "/some/route".		Output "/some/route", "/some/route/"
// Input "some/route/".		Output "/some/route", "/some/route/"
// Input "/some/route/".	Output "/some/route", "/some/route/"

func Run(url string) (string, string) {
	return app.Run(url)
}
