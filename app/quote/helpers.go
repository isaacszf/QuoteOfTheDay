package quote

import (
	"errors"
	"fmt"
    "log"
    "net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func getPageByUrlAndParse(url string) (*goquery.Document, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Accept", `text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8`)
	req.Header.Add("User-Agent", `Mozilla/5.0 (Macintosh; Intel Mac OS X 10_7_5) AppleWebKit/537.11 (KHTML, like Gecko) Chrome/23.0.1271.64 Safari/537.11`)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
    
    if res.StatusCode == 403 {
        log.Println("🔁 403 ST | Trying Again...")
        return getPageByUrlAndParse(url)
    }
    
	if res.StatusCode != 200 {
		t := fmt.Sprintf(
			"Status Code error: %d | Status: %s", res.StatusCode, res.Status)
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
