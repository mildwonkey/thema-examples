testing codegen with grafana-app-sdk

I've been using `thema lineage gen` in my local thema experimentation, but that's not the canonical method for grafana kinds (any kinds - core, custom, etc). I want to try out the canonical sdk method - at the moment it only supports custom kinds, but that's more than fine for my purposes.

Resolved: needed a newer version of the customkind.cue file (grab from main); that one was from the release (still weird that we managed to reach an "unreachable" panic, but not my problem)

Issue:
running `grafana-app-sdk generate -c . ` results in an absolutely lovely `panic: unreachable...?` error. I followed the stacktrace to this line: https://github.com/grafana/cuetsy/blob/main/generator.go#L1242

The customkind.cue file is copied from https://github.com/grafana/grafana-app-sdk/blob/ff0c8b6bc9ed309ff4c74ca732e0a65912b9dd33/codegen/testing/cue/customkind.cue.


