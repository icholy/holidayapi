package holidayapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

const TimeFormat = "2006-01-02 15:04:05"

type Time struct {
	time.Time
}

func (t *Time) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	t0, err := time.Parse(TimeFormat, s)
	if err != nil {
		return err
	}
	t.Time = t0
	return nil
}

const DateFormat = "2006-01-02"

type Date struct {
	time.Time
}

func (d *Date) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	t0, err := time.Parse(DateFormat, s)
	if err != nil {
		return err
	}
	d.Time = t0
	return nil
}

type Requests struct {
	Used      int  `json:"used"`
	Available int  `json:"available"`
	Resets    Time `json:"resets"`
}

type Response struct {
	Status    int       `json:"status"`
	Requests  Requests  `json:"requests"`
	Holidays  []Holiday `json:"holidays"`
	Countries []Country `json:"countries"`
}

func (r Response) Err() error {
	if r.Status != http.StatusOK {
		return errors.New(http.StatusText(r.Status))
	}
	return nil
}

type Holiday struct {
	Name     string `json:"name"`
	Date     Date   `json:"date"`
	Observed Date   `json:"observed"`
	Public   bool   `json:"public"`
	Country  string `json:"country"`
}

func (h Holiday) String() string {
	return fmt.Sprintf("%s: %s", h.Name, h.Date.Format(DateFormat))
}

type Country struct {
	Code         string        `json:"code"`
	Name         string        `json:"name"`
	Languages    []string      `json:"languages"`
	Flag         string        `json:"flag"`
	Subdivisions []Subdivision `json:"subdivisions"`
}

func (c Country) String() string {
	return fmt.Sprintf("%s: %s", c.Code, c.Name)
}

type Subdivision struct {
	Code      string   `json:"code"`
	Name      string   `json:"name"`
	Languages []string `json:"languages"`
}

func (s Subdivision) String() string {
	return fmt.Sprintf("%s: %s", s.Code, s.Name)
}
