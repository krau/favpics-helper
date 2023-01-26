package pixiv

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/krau/favpics-helper/pkg/client"
	"github.com/krau/favpics-helper/pkg/config"
	"github.com/krau/favpics-helper/pkg/util"
)

func NewFavURLs() ([]string, error) {
	util.Log.Info("get pixiv new fav urls")
	client, err := client.Client()
	if err != nil {
		return nil, err
	}
	resp, err := client.Get(config.Conf.Sources.Pixiv.RssURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, err
	}
	urls := make([]string, 0)
	doc.Find("item").Each(func(i int, s *goquery.Selection) {
		s.Find("img").Each(func(i int, s *goquery.Selection) {
			url, _ := s.Attr("src")
			util.Log.Debugf("get new fav url: %s", url)
			urls = append(urls, url)
		})
	})
	util.Log.Info("get pixiv new fav urls success")
	return urls, nil
}
