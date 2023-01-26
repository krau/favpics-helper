package pixiv

import (
	"regexp"

	"github.com/PuerkitoBio/goquery"
	"github.com/krau/favpics-helper/internal/structs"
	"github.com/krau/favpics-helper/pkg/client"
	"github.com/krau/favpics-helper/pkg/config"
	"github.com/krau/favpics-helper/pkg/util"
)

type Pixiv struct {
}

func (p Pixiv) NewFavPics() ([]structs.Pic, error) {
	util.Log.Info("getting pixiv new fav urls")
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
	pics := make([]structs.Pic, 0)
	doc.Find("item").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the content
		title := regexp.MustCompile(`CDATA\[(.*)\]\]>`).FindStringSubmatch(s.Find("title").Text())[1]
		description := s.Find("description").Text()
		link := s.Find("guid").Text()
		srcs := make([]string, 0)
		s.Find("img").Each(func(i int, s *goquery.Selection) {
			src, _ := s.Attr("src")
			srcs = append(srcs, src)
		})
		pic := structs.Pic{
			Title:       title,
			Link:        link,
			Srcs:        srcs,
			Description: description,
			Source:      "pixiv",
		}
		pics = append(pics, pic)
		util.Log.Debug("new pic struct:", pic)
	})
	util.Log.Info("get pixiv new fav pics success")
	return pics, nil
}
