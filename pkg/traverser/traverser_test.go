package traverser_test

import (
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

	conversion_list := traverser.GetConversionList(fs, "source", "dest")

	if len(conversion_list) != 4 {
		t.Errorf("got %d files, wanted %d files", len(conversion_list), 4)
	}
}
