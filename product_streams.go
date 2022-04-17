package main

import (
	"os"

	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/imageservice/v2/images"

	"github.com/gophercloud/gophercloud/pagination"
	"github.com/gophercloud/utils/openstack/clientconfig"
)

type server struct {
	regionName string
	endpoint   string

	glanceClient *gophercloud.ServiceClient
}

func getRegionName() string {
	if r := os.Getenv("OS_REGION_NAME"); r != "" {
		return r
	}
	if cloud, err := clientconfig.GetCloudFromYAML(&clientconfig.ClientOpts{Cloud: os.Getenv("OS_CLOUD")}); err == nil && cloud.RegionName != "" {
		return cloud.RegionName
	}

	return ""
}

func newServer() (*server, error) {
	// Get a Provider Client
	ao, err := clientconfig.AuthOptions(&clientconfig.ClientOpts{})
	if err != nil {
		return nil, err
	}
	provider, err := openstack.AuthenticatedClient(*ao)
	if err != nil {
		return nil, err
	}

	client, err := openstack.NewImageServiceV2(provider, gophercloud.EndpointOpts{})
	if err != nil {
		return nil, err
	}

	return &server{
		regionName:   getRegionName(),
		endpoint:     ao.IdentityEndpoint,
		glanceClient: client,
	}, nil
}

func (s *server) getUbuntuImages() (map[string]images.Image, error) {
	imageMap := make(map[string]images.Image)
	if err := images.List(s.glanceClient, images.ListOpts{}).EachPage(func(page pagination.Page) (bool, error) {
		list, err := images.ExtractImages(page)
		if err != nil {
			panic(err)
		}
		for _, image := range list {
			if image.Properties["os_distro"] == "ubuntu" {
				if ver := image.Properties["os_version"].(string); ver != "" {
					if existing, ok := imageMap[ver]; !ok || existing.UpdatedAt.Before(image.UpdatedAt) {
						imageMap[ver] = image
					}
				}
			}
		}
		return true, nil
	}); err != nil {
		return nil, err
	}
	return imageMap, nil
}
