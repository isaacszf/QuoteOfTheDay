package quote

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func makeRequest(url string) (*http.Request, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	headers := map[string]string{
		"Accept":           `application/json, text/javascript, */*; q=0.01`,
		"User-Agent":       `Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/107.0.0.0 Safari/537.36 Edg/107.0.1418.56`,
		"Content-Type":     `text/javascript;charset=UTF-8;application/json`,
		"Content-Length":   `355`,
		"Sec-Ch-Ua":        `Microsoft Edge";v="107", "Chromium";v="107", "Not=A?Brand";v="24`,
		"x-requested-with": `XMLHttpRequest`,
	}

	for key, value := range headers {
		req.Header.Add(key, value)
	}

	return req, nil
}

func getPageByUrlAndParse(req *http.Request, url string) (*goquery.Document, error) {
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode == 403 {
		log.Println("⚠️ ST 403 | Trying again..")

		req, err := makeRequest(url)
		if err != nil {
			return nil, err
		}

		return getPageByUrlAndParse(req, url)
	}

	if res.StatusCode != 200 {
		body, err := io.ReadAll(res.Body)
		if err != nil {
			return nil, err
		}

		t := fmt.Sprintf(
			"Status Code error: %d | Body: %s", res.StatusCode, body)
		return nil, errors.New(t)
	}

	document, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, err
	}

	return document, nil
}

func getFirstElementByClass(doc *goquery.Document, class string) string {
	data := doc.Find(class).Get(0).FirstChild.Data
	return strings.TrimSpace(data)
}
