//go:build ignore
// +build ignore

package p

import (
	refs "go.mindeco.de/ssb-refs"
)

func before(r refs.FeedRef) string {
	return r.Ref()
}

func after(r refs.FeedRef) string {
	return r.Sigil()
}
