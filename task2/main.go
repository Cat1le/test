package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
)

var (
	Url   = "https://api.coingecko.com/api/v3/coins/markets?vs_currency=usd&order=market_cap_desc&per_page=250&page=1"
	Value string
)

func init() {
	value := flag.String("value", "", "")
	flag.Parse()
	Value = strings.ToLower(*value)
	if Value == "" {
		panic("invalid 'value' passed")
	}
	fmt.Printf("Searching value '%s'\n", Value)
}

func main() {
	courses, err := GetCourses()
	if err != nil {
		panic(err)
	}
	for _, course := range courses {
		if course.IsMatch() {
			fmt.Println("Current price:", course.CurrentPrice)
			return
		}
	}
	fmt.Printf("Not found")
}

type ErrStatusCode struct {
	Code int
}

func (e ErrStatusCode) Error() string {
	return "invalid status code: " + strconv.Itoa(e.Code)
}

type Course struct {
	Id           string  `json:"id"`
	Symbol       string  `json:"symbol"`
	Name         string  `json:"name"`
	CurrentPrice float32 `json:"current_price"`
}

func (c *Course) IsMatch() bool {
	return strings.ToLower(c.Id) == Value ||
		strings.ToLower(c.Name) == Value ||
		strings.ToLower(c.Symbol) == Value
}

func GetCourses() ([]Course, error) {
	get, err := http.Get(Url)
	if err != nil {
		return nil, err
	}
	if get.StatusCode != 200 {
		return nil, &ErrStatusCode{Code: get.StatusCode}
	}
	bytes, err := io.ReadAll(get.Body)
	if err != nil {
		return nil, err
	}
	var courses []Course
	return courses, json.Unmarshal(bytes, &courses)
}
