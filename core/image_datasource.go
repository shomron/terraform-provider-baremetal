// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package core

import (
	"github.com/hashicorp/terraform/helper/schema"

	"github.com/oracle/terraform-provider-baremetal/client"
	"github.com/oracle/terraform-provider-baremetal/crud"
)

func ImageDatasource() *schema.Resource {
	return &schema.Resource{
		Read: readImages,
		Schema: map[string]*schema.Schema{
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"limit": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"page": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"operating_system": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"operating_system_version": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"images": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     ImageResource(),
			},
		},
	}
}

func readImages(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(client.BareMetalClient)
	sync := &ImageDatasourceCrud{}
	sync.D = d
	sync.Client = client
	return crud.ReadResource(sync)
}
