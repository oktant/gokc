package base

import (
	"resty.dev/v3"
	"time"
)

type Base struct {
	Path    string
	restyClient *resty.Client

	Config      struct {
		CertsInvalidateTime time.Duration
		authAdminRealms     string
		authRealms          string
		tokenEndpoint       string
		revokeEndpoint      string
		logoutEndpoint      string
		openIDConnect       string
		attackDetection     string
		version             string
	}
}