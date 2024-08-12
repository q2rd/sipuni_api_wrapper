package sipuni_api_wrapper

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"golang.org/x/time/rate"
	"io"
	"net/http"
	"time"
)

const DefaultURL = "https://sipuni.com/api/statistic/export"

type Client struct {
	_Client     *http.Client
	BaseUrl     string
	Key, UserId string
	throttle    *rate.Limiter
	debug       bool
	ctx         context.Context
}

func NewClient(Key, UserId string) *Client {
	limit := rate.Every(time.Second / 8)
	return &Client{
		_Client:  http.DefaultClient,
		BaseUrl:  DefaultURL,
		Key:      Key,
		UserId:   UserId,
		throttle: rate.NewLimiter(limit, 1),
		debug:    false,
		ctx:      context.Background(),
	}
}
func (c *Client) WithContext(ctx context.Context) *Client {
	newClient := *c
	newClient.ctx = ctx
	return &newClient
}

func (c *Client) Throttle() {
	if !c.debug {
		err := c.throttle.Wait(c.ctx)
		if err != nil {
			return
		}
	}
}

func (c *Client) Post(path string, args Arguments) []Record {
	c.Throttle()

	params := args.ToURLValuesAndHashMd5()
	url := fmt.Sprintf("%s%s", c.BaseUrl, path)
	urlWithParams := fmt.Sprintf("%s?%s", url, params.Encode())

	request, err := http.NewRequest("POST", urlWithParams, nil)
	fmt.Println(urlWithParams)
	if err != nil {
		fmt.Println(errors.Wrapf(err, "Invalid POST request %s", urlWithParams))
	}

	return c.do(request, urlWithParams)
}

func (c *Client) do(req *http.Request, _ string) []Record {
	resp, err := c._Client.Do(req)
	if err != nil {
		panic(err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(resp.Body)

	if resp.StatusCode < http.StatusOK || resp.StatusCode >= http.StatusMultipleChoices {
		fmt.Printf("Failed to make request |%s| %s| %s|\n", req.Method, req.URL, resp.Status)
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Failed to read response body |%s| %s|\n", req.Method, req.URL)
	}
	res, err := parseCSVResponse(bodyBytes)
	if err != nil {
		panic(err)
	}
	return res
}
