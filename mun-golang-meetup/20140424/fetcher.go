type FetchOptions struct {
	BasicAuthUsername string
	BasicAuthPassword string
	SocksProxy        string
	HttpProxy         string
	Timeout           int32
	Retries           int32
	Cookies           []*http.Cookie
	Header            http.Header
	PostFormData      url.Values
}

type Fetcher struct {
	options FetchOptions
	client  *http.Client
}
