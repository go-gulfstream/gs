package pkgdev

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

const versionSelector = ".js-versionLink"
const totalScan = 6

var ErrLoadPage = errors.New("load page")

func Version(mod string) ([]string, string, error) {
	url := makeURL(mod)
	doc, err := loadHTML(url)
	if err != nil {
		return nil, url, fmt.Errorf("%v %w", err, ErrLoadPage)
	}
	versions := make([]string, 0, totalScan)
	doc.Find(versionSelector).Each(func(i int, s *goquery.Selection) {
		if i >= totalScan {
			return
		}
		if len(s.Text()) > 0 {
			versions = append(versions, s.Text())
		}
	})
	return versions, url, nil
}

func loadHTML(url string) (*goquery.Document, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		return nil, fmt.Errorf("pkgdev: status code error: %d %s",
			res.StatusCode, res.Status)
	}
	return goquery.NewDocumentFromReader(res.Body)
}

func makeURL(mod string) string {
	return "https://pkg.go.dev/" + mod + "?tab=versions"
}
