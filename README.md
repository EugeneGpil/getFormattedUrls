# Format url

Format url for "net/http.ServeMux.HandleFunc"

Both returned urls should be provided to HandleFunc function

You should use this package to prevent "forget trailing slash" error
when you trying to access your app route
For example
```
            no trailing slash 
                            V
http://test.local/some/route
```
instead of
```
               trailing slash
                            V
http://test.local/some/route/
```
and
```
               trailing slash
                            V
http://test.local/some/route/
```
instead of
```
            no trailing slash
                            V
http://test.local/some/route
```

And use it to prevent "forget first slash" when providing url for HandlerFunc function
For example
```
no slash
V
some/route/
```
instead of
```
slash
V
/some/route/
```

Input url may be with or without trailing slash `/`

Input url may be with or without first slash `/`


Returns 2 urls

Both with first slash `/`

First without trailing slash `/`
and second with trailing slash `/`

For 4 possible inputs will be the same output

### Inputs
```
some/route
/some/route
some/route/
/some/route/
```

### Output
```
/some/route
/some/route/
```

## Dummy example

```
package main

import (
	"net/http"

	"github.com/EugeneGpil/getFormattedUrls"
)

func main() {
	// create handler for route
	handler := func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("Hello World!\n"))
	}

	mux := http.NewServeMux()

	// should be slashes here?
	//     V               V
	url := "get/hello/world"

	url1, url2 := getFormattedUrls.Run(url)

	//                                           V
	// for "http://localhost:8080/get/hello/world"
	mux.HandleFunc(url1, handler)

	//                                           V
	// for "http://localhost:8080/get/hello/world/"
	mux.HandleFunc(url2, handler)

	http.ListenAndServe(":8080", mux)
}
```
