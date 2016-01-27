package freedesktop

import (
    "testing"
)

func TestGetDataFilename(t *testing.T) {
    homeLocal := `./test/testing-home/local/share/test1/data.file`
    // global1   := `./test/testing-dir1/test2/data.file`
    // global2   := `./test/testing-dir1/test3/data.file`

    if v, err := GetDataFilename(`test1/data.file`); err == nil {
        if v != homeLocal {
            t.Errorf("Invalid data filename: expected '%s', got '%s", homeLocal, v)
        }
    }else{
        t.Errorf("Failed to retrieve XDG data file: %v", err)
    }
}