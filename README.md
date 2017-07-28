# urlreader

Single function that support opening file, https, and s3 URLs. It returns `io.ReadCloser`:

```
	import "github.com/agilebits/urlreader"

	...

	// url := "https://twitter.com/1Password"
	url := "file://./main.go"

	reader, err := urlreader.Open(url)
	if err != nil {
		panic(err)
	}

	defer reader.Close()
	result, err := ioutil.ReadAll(reader)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(result))
```
