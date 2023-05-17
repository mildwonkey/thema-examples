package main

import (
	"github.com/grafana/thema"
)

lineage: thema.#Lineage
lineage: {
	name: "example"
	schemas: [
		{
			version: [0, 0]
			schema:
			{
				title: string
			},
		},
		{
			version: [0, 1]
			schema:
			{
				title: string
				header?: string // new, optional field ->  no error
			},
		},
	]
}
