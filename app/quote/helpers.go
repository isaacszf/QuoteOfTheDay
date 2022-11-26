package quote

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func getPageByUrlAndParse(url string) (*goquery.Document, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Accept", `application/json, text/javascript, */*; q=0.01`)
	req.Header.Add("User-Agent", `Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/107.0.0.0 Safari/537.36 Edg/107.0.1418.56`)
	req.Header.Add("Content-Type", `application/json`)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

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
