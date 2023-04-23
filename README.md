# Text/diff

This is a copy of the Go text diffing package that [gopls uses] to generate
unified diffs. There is [some discussion] about moving the package out of an
internal package, but in the meantime this can be used.

This is identical to https://github.com/hexops/gotextdiff, but that hasn't been
updated with the newer API, so I just did the same thing.

The code is entirely written by the Go Authors.

No PRs will be accepted here - please contribute to the upstream source!

[gopls uses]: https://github.com/golang/tools/tree/master/internal/diff
[some discussion]: https://github.com/golang/go/issues/58893

## License

See https://github.com/golang/tools/blob/master/LICENSE
