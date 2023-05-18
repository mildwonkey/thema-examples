package main

import (
	"github.com/grafana/thema"
)

thema.#Lineage
name: "example"
schemas: [
	{
		version: [0, 0]
		schema:
		{
			title?: string
		},
	},
	{
		version: [0, 1]
		schema:
		{
			title?: string
			header: string // new required field
		},
	},
	{
		version: [1, 0]
		schema:
		{
			title?: string
		},
	},
]
