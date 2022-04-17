package main

import (
	"fmt"
	"time"

	"github.com/gophercloud/gophercloud/openstack/imageservice/v2/images"
)

type metadata struct {
	ContentID string             `json:"content_id"`
	Format    string             `json:"format"`
	Products  map[string]product `json:"products"`
}

type product struct {
	Architecture string                    `json:"arch"`
	Version      string                    `json:"version"`
	Versions     map[string]productVersion `json:"versions"`
}

type productVersion struct {
	Items map[string]productVersionItem `json:"items"`
}

type productVersionItem struct {
	Endpoint string `json:"endpoint"`
	Region   string `json:"region"`
	ID       string `json:"id"`
}

func (s *server) makeMetadataFromImages(imageMap map[string]images.Image) metadata {
	now := time.Now().UTC()
	dateVersion := fmt.Sprintf("%4d%02d%02d", now.Year(), now.Month(), now.Day())

	products := make(map[string]product, len(imageMap))
	for ver, image := range imageMap {
		products[fmt.Sprintf("com.ubuntu.cloud:server:%s:amd64", ver)] = product{
			Architecture: "amd64",
			Version:      ver,
			Versions: map[string]productVersion{
				dateVersion: {
					Items: map[string]productVersionItem{
						image.ID: {
							Endpoint: s.endpoint,
							Region:   s.regionName,
							ID:       image.ID,
						},
					},
				},
			},
		}
	}
	return metadata{
		ContentID: "com.ubuntu.cloud:custom",
		Format:    "products:1.0",
		Products:  products,
	}
}
