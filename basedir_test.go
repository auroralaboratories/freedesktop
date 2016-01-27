package freedesktop

import (
    "testing"
)

func TestGetDataFilename(t *testing.T) {
    homeLocal := `./test/testing-home/local/share/test1/data.file`
    global1   := `./test/testing-dir1/test2/data.file`
    global2   := `./test/testing-dir2/test3/data.file`

//  mock home directory
    XdgDataHome = `test/testing-home`

//  mock data dirs
    XdgDataDirs = `test/testing-dir1/:test/testing-dir2/`

//  test for a file that exists in all directories
    if v, err := GetDataFilename(`test1/data.file`); err == nil {
        if v != homeLocal {
            t.Errorf("Invalid data filename: expected '%s', got '%s", homeLocal, v)
        }
    }else{
        t.Errorf("Failed to retrieve XDG data file: %v", err)
    }

//  test for a file that exists only in global directories
    if v, err := GetDataFilename(`test2/data.file`); err == nil {
        if v != global1 {
            t.Errorf("Invalid data filename: expected '%s', got '%s", homeLocal, v)
        }
    }else{
        t.Errorf("Failed to retrieve XDG data file: %v", err)
    }

//  test for a file that exists only in the last global directory
    if v, err := GetDataFilename(`test3/data.file`); err == nil {
        if v != global2 {
            t.Errorf("Invalid data filename: expected '%s', got '%s", homeLocal, v)
        }
    }else{
        t.Errorf("Failed to retrieve XDG data file: %v", err)
    }
}