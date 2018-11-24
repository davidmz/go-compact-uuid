# cuuid

[![GoDoc](https://godoc.org/github.com/davidmz/go-compact-uuid?status.svg)](https://godoc.org/github.com/davidmz/go-compact-uuid) ![Version](https://img.shields.io/github/tag/davidmz/go-compact-uuid.svg?label=version)

The `go-compact-uuid` (imports as `cuuid`) is a Go package that encodes/decodes UUIDs to/from a more compact form that a regular UUID format. For example, the UUID `3867b6f3-dbb0-4ef5-8078-364897154fd9` (36 chars) can be encoded as `3gsxpyfdv0kqn80y1p92bhakys` (26 chars).

This package was inspired by Clojure [compact-uuids](https://github.com/tonsky/compact-uuids) library and I recommend to read its README to learn about the benefits of this format.

## API

```go
// UUID is an array to represent UUID value. This type
// is compatible with the most of other UUID packages.
type UUID [16]byte
```

```go
// String returns a compact string representation of UUID,
// the 26-character string of '0123456789abcdefghjkmnpqrstvwxyz' alphabet.
func (u UUID) String() string
```

```go
// FromString parses text and returns UUID or error. It expects the 26-character
// string of '0123456789abcdefghjkmnpqrstvwxyz' alphabet (case-insensitive).
func FromString(text string) (UUID, error)
```
