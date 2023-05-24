Schema changes

This example models a real-ish grafana dashboard schema change: datasources were strings, now they are objects. I've been playing with a variation on the actual schema, where "type" is required; it's option in the current dashboard_kind.cue file. 

```
datasource: string  // schema 0.0

datasource: {
    type: string
    uid?:  string
}                   // schema 1.0
```

I'm curious about expected behavior here! There's *absolutely* an error in my schema - a lens exists, but it's a no-op - and at the very least I expected some kind of error from Translate() (a panic or nil return value) since a required field wasn't set, but instead it happily returned an instance with an empty string for that field.

Should this have resulted in an error? Maybe, maybe not. I'd think Translate() could have flagged this as an error, and at a higher level thema could potentially validate lenses and call out issues like this ("you are adding a required field, but it's not dealt with in the lens"), but I'm not even sure if those things make sense. Having said that I'm also not sure how we would debug similar issues in a real grafana instance - if there are no errors expected, how long might we spend tracking down this bug? (especially since it's only coming out of translate, so anything conforming to the newer schema would be just fine).

Since the entire datasource field is required should thema have returned an error or refused to Translate() (translate doesn't return errors, but it may return a nil instance? Should thema refuse to generate the lineage if breaking changes aren't directly addressed? Or does thema only verify that the lineage _exists_ without validing correctness? There's a good argument for the latter, so perhaps the SDK could validate if the lens-correctness is up to Grafana standards.