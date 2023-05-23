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
			title: string
		},
	},
	{
		version: [0, 1]
		schema:
		{
			title: string
			description?: string
		},
	},
	{
		version: [1, 0]
		schema:
		{
			header: string
			description?: string
		},
	},
]
lenses: [
{
	to: [1, 0]
    from: [0, 1]
    input: _
    result: {
      header: input.title
    }
}, 
{
	to: [0, 1]
    from: [1, 0]
    input: _
    result: {
      title: input.header
    }
},
	// things work when the 0.1->0.0 lens is uncommented:
	// {
	// 	to: [0, 0]
	// 	from: [0, 1]
	// 	input: _
	// 	result: {}
	// }, 
	//
	// adding the 0.0->0.1 lens breaks causes another panic, with or without the above lens.
	// 	{
	// 	to: [0, 1]
	// 	from: [0, 0]
	// 	input: _
	// 	result: {}
	// }, 
]
