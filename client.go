package holidayapi

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"strconv"
)

type Client struct {
	Key        string
	BaseURL    string
	HTTPClient *http.Client
	Requests   Requests
}

func New(key string) *Client {
	return &Client{
		Key:        key,
		BaseURL:    "https://holidayapi.com",
		HTTPClient: http.DefaultClient,
	}
}

func (c *Client) Get(url string) (*Response, error) {
	resp, err := c.HTTPClient.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(resp.Status)
	}
	var r Response
	if err := json.NewDecoder(resp.Body).Decode(&r); err != nil {
		return nil, err
	}
	return &r, nil
}

func (c *Client) URL(path string, query url.Values) (string, error) {
	u, err := url.Parse(c.BaseURL)
	if err != nil {
		return "", err
	}
	u.Path = path
	if query == nil {
		query = url.Values{}
	}
	query.Set("key", c.Key)
	u.RawQuery = query.Encode()
	return u.String(), nil
}

func (c *Client) Countries() ([]Country, error) {
	url, err := c.URL("/v1/countries", nil)
	if err != nil {
		return nil, err
	}
	resp, err := c.Get(url)
	if err != nil {
		return nil, err
	}
	if err := resp.Err(); err != nil {
		return nil, err
	}
	c.Requests = resp.Requests
	return resp.Countries, nil
}

func (c *Client) Holidays(p Params) ([]Holiday, error) {
	url, err := c.URL("/v1/holidays", p.Values())
	if err != nil {
		return nil, err
	}
	resp, err := c.Get(url)
	if err != nil {
		return nil, err
	}
	if err := resp.Err(); err != nil {
		return nil, err
	}
	c.Requests = resp.Requests
	return resp.Holidays, nil
}

type Params struct {
	Country  string
	Year     int
	Month    int
	Day      int
	Previous bool
	Upcoming bool
	Public   bool
	Language string
}

func (p Params) Values() url.Values {
	v := url.Values{}
	v.Set("country", p.Country)
	v.Set("year", strconv.Itoa(p.Year))
	if p.Month > 0 {
		v.Set("month", strconv.Itoa(p.Month))
	}
	if p.Day > 0 {
		v.Set("day", strconv.Itoa(p.Day))
	}
	if p.Previous {
		v.Set("previous", "true")
	}
	if p.Upcoming {
		v.Set("upcoming", "true")
	}
	if p.Public {
		v.Set("public", "true")
	}
	if p.Language != "" {
		v.Set("language", p.Language)
	}
	return v
}
