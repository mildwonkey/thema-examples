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
			datasource: { // datasource is now an object
				type?: string // this is a required field, which is probably nonsense!
				uid?:  string
			}
		},
	},
]

lenses:[
	{
		from: [0, 0]
		to: [1, 0]
		input: _
		result: {}
	}
]
