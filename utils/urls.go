package utils

import (
	"github.com/tanyoAH/tanyo-server/config"
	"strings"
)

func GetFullUrlPathWithConfig(path string) string {
	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}
	return config.Conf.PreferredLinkProtocol + "://" + config.Conf.Home + path
}
