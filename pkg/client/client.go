package client

import (
	"net/http"
	"net/url"

	"github.com/krau/favpics-helper/pkg/config"
	"github.com/krau/favpics-helper/pkg/util"
)

func Client() (http.Client, error) {
	util.Log.Debug("get http client")
	if !config.Conf.Proxy.Enabled {
		return http.Client{}, nil
	}
	proxyURL, err := url.Parse(config.Conf.Proxy.Addr)
	if err != nil {
		return http.Client{}, err
	}
	client := http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyURL(proxyURL),
		},
	}
	util.Log.Debug("get http client success")
	return client, nil
}
