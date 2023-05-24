testing codegen with grafana-app-sdk

I've been using `thema lineage gen` in my local thema experimentation, but that's not the canonical method for grafana kinds (any kinds - core, custom, etc). I want to try out the canonical sdk method - at the moment it only supports custom kinds, but that's more than fine for my purposes.


Issue:
running `grafana-app-sdk generate -c . ` results in an absolutely lovely `panic: unreachable...?` error. I followed the stacktrace to this line: https://github.com/grafana/cuetsy/blob/main/generator.go#L1242
