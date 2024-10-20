package settings

import (
	"bytes"
	"crypto/tls"
	"net/http"
	"time"
)

type HttpService struct {
	Host        string
	Token       string
	Certificate tls.Certificate
	HttpClient  *http.Client
}

func NewHttpService(uri string) *HttpService {
	return &HttpService{
		Host:  uri,
		Token: "",
		HttpClient: &http.Client{
			Timeout: 50 * time.Second,
		},
	}
}

func (s *HttpService) SetPort(port string) {
	s.Host = s.Host + ":" + port
}

func (s *HttpService) Get(path string, headers map[string]string) (*http.Response, error) {
	req, err := http.NewRequest("GET", s.Host+path, nil)
	if err != nil {
		return nil, err
	}

	for k, v := range headers {
		req.Header.Add(k, v)
	}

	return s.HttpClient.Do(req)
}

func (s *HttpService) Post(path string, body []byte, headers map[string]string) (*http.Response, error) {
	reader := bytes.NewReader(body)
	req, err := http.NewRequest("POST", s.Host+path, reader)

	if err != nil {
		return nil, err
	}

	for k, v := range headers {
		req.Header.Add(k, v)
	}

	return s.HttpClient.Do(req)
}

//Formdata Api Calling
/*
func (s *HttpService) PostFormNoFile(path string, formFields map[string]interface{}, authorizationDetails *AuthDetails) (*http.Response, error) {
	data := url.Values{}

	u, err := url.ParseRequestURI(s.Host)

	if err != nil {
		return nil, err
	}

	u.Path = path

	urlString := u.String()

	for key, value := range formFields {
		data.Set(key, fmt.Sprintf("%v", value))
	}

	r, err := http.NewRequest(http.MethodPost, urlString, strings.NewReader(data.Encode()))

	if err != nil {
		return nil, err
	}

	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	if authorizationDetails.IsBasicAuth() {
		r.SetBasicAuth(authorizationDetails.Username, authorizationDetails.Password)
	}

	if authorizationDetails.IsBearerAuth() {
		r.Header.Add("Authorization", "Bearer "+authorizationDetails.Token)
	}

	return s.HttpClient.Do(r)
}

func (s *HttpService) PostFormData(
	path, fileKey string,
	filePath string,
	formFields map[string]string,
	authorizationDetails *AuthDetails,
) (*http.Response, error) {

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	if fileKey != "" && filePath != "" {
		file, err := os.Open(filePath)

		if err != nil {
			return nil, err
		}

		defer file.Close()

		part, err := writer.CreateFormFile(fileKey, filePath)

		if err != nil {
			return nil, err
		}

		if _, err := io.Copy(part, file); err != nil {
			return nil, err
		}
	}

	for key, value := range formFields {
		if err := writer.WriteField(key, value); err != nil {
			return nil, err
		}
	}

	writer.Close()

	req, err := http.NewRequest("POST", s.Host+path, body)
	if err != nil {
		fmt.Println("ERR ", err)
		return nil, err
	}

	if authorizationDetails.IsBasicAuth() {
		req.SetBasicAuth(authorizationDetails.Username, authorizationDetails.Password)
	}

	if authorizationDetails.IsBearerAuth() {
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", authorizationDetails.Token))
	}

	if authorizationDetails.IsNoAuth() {
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", authorizationDetails.Token))
	}

	req.Header.Add("Content-Type", writer.FormDataContentType())

	return s.HttpClient.Do(req)
}
*/
func (s *HttpService) Put(path string, body []byte, headers map[string]string) (*http.Response, error) {

	reader := bytes.NewReader(body)
	req, err := http.NewRequest("PUT", s.Host+path, reader)

	if err != nil {
		return nil, err
	}

	for k, v := range headers {
		req.Header.Add(k, v)
	}

	return s.HttpClient.Do(req)
}

func (s *HttpService) Delete(path string, headers map[string]string) (*http.Response, error) {
	req, err := http.NewRequest("DELETE", s.Host+path, nil)
	if err != nil {
		return nil, err
	}

	for k, v := range headers {
		req.Header.Add(k, v)
	}

	return s.HttpClient.Do(req)
}
