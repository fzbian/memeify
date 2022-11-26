package services

import (
	"net/url"
)

type MemeServices interface {
	GenerateMeme(nameMeme string, queryParams url.Values) (string, error)
}
