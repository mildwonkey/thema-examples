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
			datasource: string
		},
	},
	{
		version: [1, 0]
		schema:
		{
			lies_source: string
		},
	},
	{
		version: [2, 0]
		schema:
		{
			rando_source: string
		},
	},
]

lenses:[
	{
		from: [0, 0]
		to: [1, 0]
		input: _
		result: {}
	},
	{
		from: [1, 0]
		to: [2, 0]
		input: _
		result: {}
	},
]
