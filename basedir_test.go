package freedesktop

import (
    "testing"
)

func TestGetDataFilename(t *testing.T) {
    homeLocal := `testing/testing-home/.local/share/my-app/data1.file`
    global1   := `testing/testing-global1/my-app/data2.file`
    global2   := `testing/testing-global2/my-app/data3.file`

//  mock home directory
    XdgDataHome = `testing/testing-home/.local/share`

//  mock data dirs
    XdgDataDirs = `testing/testing-dir1/:test/testing-dir2/`

//  test for a file that exists in all directories
    if v, err := GetDataFilename(`my-app/data1.file`); err == nil {
        if v != homeLocal {
            t.Errorf("Invalid data filename: expected '%s', got '%s", homeLocal, v)
        }
    }else{
        t.Errorf("data1: %v", err)
    }

//  test for a file that exists only in global directories
    if v, err := GetDataFilename(`my-app/data2.file`); err == nil {
        if v != global1 {
            t.Errorf("Invalid data filename: expected '%s', got '%s", homeLocal, v)
        }
    }else{
        t.Errorf("data2: %v", err)
    }

//  test for a file that exists only in the last global directory
    if v, err := GetDataFilename(`my-app/data3.file`); err == nil {
        if v != global2 {
            t.Errorf("Invalid data filename: expected '%s', got '%s", homeLocal, v)
        }
    }else{
        t.Errorf("data3: %v", err)
    }
}