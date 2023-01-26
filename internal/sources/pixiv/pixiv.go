package pixiv

import (
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/krau/favpics-helper/internal/db"
	"github.com/krau/favpics-helper/internal/interfaces"
	"github.com/krau/favpics-helper/internal/models"
	"github.com/krau/favpics-helper/pkg/client"
	"github.com/krau/favpics-helper/pkg/config"
	"github.com/krau/favpics-helper/pkg/util"
)

type Pixiv struct {
	interfaces.Fav
}

func (p Pixiv) NewFavPics() ([]models.Pic, error) {
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
	pics := make([]models.Pic, 0)
	doc.Find("item").Each(func(i int, s *goquery.Selection) {
		title := regexp.MustCompile(`CDATA\[(.*)\]\]>`).FindStringSubmatch(s.Find("title").Text())[1]
		description := strings.Split(s.Find("description").Text(), "]]>")[0]
		link := s.Find("guid").Text()
		srcs := make([]string, 0)
		s.Find("img").Each(func(i int, s *goquery.Selection) {
			src, _ := s.Attr("src")
			srcs = append(srcs, src)
		})
		pic := models.Pic{
			Title:       title,
			Link:        link,
			Srcs:        srcs,
			Description: description,
			Source:      "pixiv",
		}
		if !db.IsPicExist(pic) {
			util.Log.Debug("new pic struct:", pic)
			pics = append(pics, pic)
		} else {
			util.Log.Debug("pic exist:", pic)
		}
	})
	util.Log.Info("get pixiv new fav pics success")
	return pics, nil
}
