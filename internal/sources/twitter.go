package sources

import (
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/krau/favpics-helper/internal/db"
	"github.com/krau/favpics-helper/internal/models"
	"github.com/krau/favpics-helper/pkg/client"
	"github.com/krau/favpics-helper/pkg/config"
	"github.com/krau/favpics-helper/pkg/util"
)

type Twitter struct{}

func (t Twitter) NewFavPics() ([]models.Pic, error) {
	util.Log.Info("getting twitter new fav pics")
	client, err := client.Client()
	if err != nil {
		return nil, err
	}
	resp, err := client.Get(config.Conf.Sources.Twitter.RssURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, err
	}
	pics := make([]models.Pic, 0)
	doc.Find("item").Each(func(i int, s *goquery.Selection) {
		src, isExist := s.Find("img").Attr("src")
		if isExist {
			srcs := make([]string, 0)
			src = url.QueryEscape(src)
			srcs = append(srcs, src)
			//title := regexp.MustCompile(`CDATA\[(.*)\]\]>`).FindStringSubmatch(s.Find("title").Text())
			title := s.Find("guid").Text()
			description := strings.Split(s.Find("description").Text(), "]]>")[0]
			link := s.Find("guid").Text()
			pic := models.Pic{
				Title:       title,
				Link:        link,
				Srcs:        srcs,
				Description: description,
				Source:      "twitter",
			}
			if !db.IsPicExist(pic) {
				util.Log.Debug("new pic struct:", pic)
				pics = append(pics, pic)
			} else {
				util.Log.Debug("pic exist:", pic.Link)
			}
		}
	})
	util.Log.Info("get twitter new fav pics done")
	return pics, nil
}
