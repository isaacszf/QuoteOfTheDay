package quote

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

func getPageByUrlAndParse(url string) (*goquery.Document, error) {
	time.Sleep(time.Second * 5)

	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return nil, err
		}

		t := fmt.Sprintf(
			"Status Code error: %d | Body: %s", res.StatusCode, string(body))
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
