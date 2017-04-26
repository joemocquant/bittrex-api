// Bittrex public API implementation.
//
// API Doc: https://bittrex.com/Home/Api
package bittrex

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/Sirupsen/logrus"
)

var (
	conf   *configuration
	logger *logrus.Entry
)

type PublicClient struct {
	httpClient *http.Client
	throttle   <-chan time.Time
}

type APIResult struct {
	Success bool            `json:"success"`
	Message string          `json:"message"`
	Result  json.RawMessage `json:"result"`
}

type configuration struct {
	PublicAPIConf `json:"bittrex_public_api"`
}

type PublicAPIConf struct {
	APIUrl               string `json:"api_url"`
	HTTPClientTimeoutSec int    `json:"httpclient_timeout_sec"`
	MaxRequestsSec       int    `json:"max_requests_sec"`
	LogLevel             string `json:"log_level"`
}

func init() {

	customFormatter := new(logrus.TextFormatter)
	customFormatter.FullTimestamp = true
	logrus.SetFormatter(customFormatter)

	logger = logrus.WithField("context", "[bittrex:publicapi]")

	content, err := ioutil.ReadFile("conf.json")

	if err != nil {
		logger.WithField("error", err).Fatal("loading configuration")
	}

	if err := json.Unmarshal(content, &conf); err != nil {
		logger.WithField("error", err).Fatal("loading configuration")
	}

	switch conf.LogLevel {
	case "debug":
		logrus.SetLevel(logrus.DebugLevel)
	case "info":
		logrus.SetLevel(logrus.InfoLevel)
	case "warn":
		logrus.SetLevel(logrus.WarnLevel)
	case "error":
		logrus.SetLevel(logrus.ErrorLevel)
	case "fatal":
		logrus.SetLevel(logrus.FatalLevel)
	case "panic":
		logrus.SetLevel(logrus.PanicLevel)
	default:
		logrus.SetLevel(logrus.WarnLevel)
	}
}

// NewPublicClient returns a newly configured client
func NewPublicClient() *PublicClient {

	reqInterval := 1000 * time.Millisecond / time.Duration(conf.MaxRequestsSec)

	client := http.Client{
		Timeout: time.Duration(conf.HTTPClientTimeoutSec) * time.Second,
	}

	return &PublicClient{&client, time.Tick(reqInterval)}
}

// Do prepares and executes api call requests.
func (c *PublicClient) do(endpoint string, params map[string]string) ([]byte, error) {

	url := buildUrl(endpoint, params)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("http.NewRequest: %v (API endpoint: %s)",
			err, endpoint)
	}

	req.Header.Add("Accept", "application/json")

	type result struct {
		resp *http.Response
		err  error
	}

	done := make(chan result)
	go func() {
		<-c.throttle
		resp, err := c.httpClient.Do(req)
		done <- result{resp, err}
	}()
	res := <-done

	if res.err != nil {
		return nil, fmt.Errorf("http.Client.Do: %v", res.err)
	}

	defer res.resp.Body.Close()

	body, err := ioutil.ReadAll(res.resp.Body)
	if err != nil {
		return body, fmt.Errorf("ioutil.readAll: %v", err)
	}

	if res.resp.StatusCode != 200 {
		return body, fmt.Errorf("status code: %s (API endpoint: %s)",
			res.resp.Status, endpoint)
	}

	ae := APIResult{}
	if err := json.Unmarshal(body, &ae); err != nil {
		return nil, fmt.Errorf("API error: %s", ae.Message)
	}

	if !ae.Success {
		return nil, errors.New(ae.Message)
	}

	return ae.Result, nil
}

func buildUrl(endpoint string, params map[string]string) string {

	u := conf.APIUrl + "/" + endpoint + "?"

	var parameters []string
	for k, v := range params {
		parameters = append(parameters, k+"="+url.QueryEscape(v))
	}

	return u + strings.Join(parameters, "&")
}
