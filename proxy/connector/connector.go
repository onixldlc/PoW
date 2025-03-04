package connector

import (
	_ "embed"
	"fmt"
)

//go:embed web/index.html
var IndexHtml string

type ConnectorPage struct {
	index string
}

func (cp ConnectorPage) GetPage(page string) (string, error) {
	if page == "index" {
		return cp.index, nil
	}
	return "", fmt.Errorf("invalid page")
}

func New() ConnectorPage {
	return ConnectorPage{
		index: IndexHtml,
	}
}
