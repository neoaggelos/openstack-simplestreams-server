package main

import (
	"fmt"
	"time"

	"github.com/gophercloud/gophercloud/openstack/imageservice/v2/images"
)

type cloud struct {
	Region   string `json:"region"`
	Endpoint string `json:"endpoint"`
}

type entry struct {
	Updated   time.Time `json:"updated"`
	Format    string    `json:"format"`
	DataType  string    `json:"datatype"`
	CloudName string    `json:"cloudname"`
	Clouds    []cloud   `json:"clouds"`
	Path      string    `json:"path"`
	Products  []string  `json:"products"`
}

type indexEntries struct {
	UbuntuCloudCustom entry `json:"com.ubuntu.cloud:custom"`
}

type index struct {
	Index   indexEntries `json:"index"`
	Updated time.Time    `json:"updated"`
	Format  string       `json:"format"`
}

func (s *server) makeIndexFromImages(imageMap map[string]images.Image) index {
	products := make([]string, 0, len(imageMap))
	for ver := range imageMap {
		products = append(products, fmt.Sprintf("com.ubuntu.cloud:server:%s:amd64", ver))
	}

	updated := time.Now().UTC()

	return index{
		Index: indexEntries{
			UbuntuCloudCustom: entry{
				Updated:   updated,
				Format:    "products:1.0",
				DataType:  "image-ids",
				CloudName: "custom",
				Clouds: []cloud{
					{
						Region:   s.regionName,
						Endpoint: s.endpoint,
					},
				},
				Path:     "streams/v1/com.ubuntu.cloud-released-imagemetadata.json",
				Products: products,
			},
		},
		Updated: updated,
		Format:  "index:1.0",
	}
}
