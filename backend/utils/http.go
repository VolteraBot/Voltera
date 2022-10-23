package utils

import (
	"bytes"
	"fmt"
	"github.com/VolteraBot/Voltera/backend/types"
	"io"
	"net/http"
	"time"
)

func Get(requestURL string, headers types.Header) string {
	req, err := http.NewRequest(http.MethodGet, requestURL, nil)
	if err != nil {
		fmt.Printf("client: could not create request: %s\n", err)
	}

	for key, val := range headers {
		req.Header.Set(key, val)
	}

	client := http.Client{
		Timeout: 30 * time.Second,
	}

	res, err := client.Do(req)
	if err != nil {
		fmt.Printf("client: error making http request: %s\n", err)
	}
	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("client: could not read response body: %s\n", err)
	}
	return bytes.NewBuffer(resBody).String()
}
func Patch(requestURL string, headers types.Header, body io.Reader) string {
	req, err := http.NewRequest(http.MethodPatch, requestURL, body)
	if err != nil {
		fmt.Printf("client: could not create request: %s\n", err)
	}

	for key, val := range headers {
		req.Header.Set(key, val)
	}

	client := http.Client{
		Timeout: 30 * time.Second,
	}

	res, err := client.Do(req)
	if err != nil {
		fmt.Printf("client: error making http request: %s\n", err)
	}
	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("client: could not read response body: %s\n", err)
	}
	return bytes.NewBuffer(resBody).String()
}
func Put(requestURL string, headers types.Header, body io.Reader) string {
	req, err := http.NewRequest(http.MethodPut, requestURL, body)
	if err != nil {
		fmt.Printf("client: could not create request: %s\n", err)
	}

	for key, val := range headers {
		req.Header.Set(key, val)
	}

	client := http.Client{
		Timeout: 30 * time.Second,
	}

	res, err := client.Do(req)
	if err != nil {
		fmt.Printf("client: error making http request: %s\n", err)
	}
	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("client: could not read response body: %s\n", err)
	}
	return bytes.NewBuffer(resBody).String()
}
func Post(requestURL string, headers types.Header, body io.Reader) string {
	req, err := http.NewRequest(http.MethodPost, requestURL, body)
	if err != nil {
		fmt.Printf("client: could not create request: %s\n", err)
	}

	for key, val := range headers {
		req.Header.Set(key, val)
	}

	client := http.Client{
		Timeout: 30 * time.Second,
	}

	res, err := client.Do(req)
	if err != nil {
		fmt.Printf("client: error making http request: %s\n", err)
	}
	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("client: could not read response body: %s\n", err)
	}
	return bytes.NewBuffer(resBody).String()
}
func Delete(requestURL string, headers types.Header, body io.Reader) string {
	req, err := http.NewRequest(http.MethodDelete, requestURL, body)
	if err != nil {
		fmt.Printf("client: could not create request: %s\n", err)
	}

	for key, val := range headers {
		req.Header.Set(key, val)
	}

	client := http.Client{
		Timeout: 30 * time.Second,
	}

	res, err := client.Do(req)
	if err != nil {
		fmt.Printf("client: error making http request: %s\n", err)
	}
	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("client: could not read response body: %s\n", err)
	}
	return bytes.NewBuffer(resBody).String()
}
