package traverser_test

import (
	"reflect"
	"testing"
	"testing/fstest"

	traverser "github.com/SteveHan-233/MTS-to-mp4/pkg/traverser"
)

func TestGetConversionList(t *testing.T) {
	fs := fstest.MapFS{
		"source/video1.MTS":           {},
		"source/dir1/video2.MTS":      {},
		"source/dir1/not_video.txt":   {},
		"source/dir2/video3.MTS":      {},
		"source/dir2/dir3/video4.MTS": {},
	}

	got := traverser.GetConversionList(fs, "source", "dest")

	want := []traverser.SourceDestPair{
		{
			Source: "source/video1.MTS",
			Dest:   "dest/video1.mp4",
		},
		{
			Source: "source/dir1/video2.MTS",
			Dest:   "dest/dir1/video2.mp4",
		},
		{
			Source: "source/dir2/video3.MTS",
			Dest:   "dest/dir2/video3.mp4",
		},
		{
			Source: "source/dir2/dir3/video4.MTS",
			Dest:   "dest/dir2/dir3/video4.mp4",
		},
	}

	if len(got) != 4 {
		t.Errorf("got %d files, wanted %d files", len(got), 4)
	}

	if !isEqual(got, want) {
		t.Errorf("got %+v, want %+v", got, want)
	}
}

func isEqual(aa, bb []traverser.SourceDestPair) bool {
	eqCtr := 0
	for _, a := range aa {
		for _, b := range bb {
			if reflect.DeepEqual(a, b) {
				eqCtr++
			}
		}
	}
	if eqCtr != len(bb) || len(aa) != len(bb) {
		return false
	}
	return true
}
