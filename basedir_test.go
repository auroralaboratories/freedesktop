package freedesktop

import (
    "testing"
)

func TestGetDataFilename(t *testing.T) {
    data1 := `testing/xdg-basedir/home/gotest/.local/share/my-app/data1.file`
    data2   := `testing/xdg-basedir/usr/local/share/my-app/data2.file`
    data3   := `testing/xdg-basedir/usr/share/my-app/data3.file`

//  mock home directory
    XdgDataHome = `testing/xdg-basedir/home/gotest/.local/share`

//  mock data dirs
    XdgDataDirs = `testing/xdg-basedir/usr/local/share/:testing/xdg-basedir/usr/share`

//  test for a file that exists in all directories
    if v, err := GetDataFilename(`my-app/data1.file`); err == nil {
        if v != data1 {
            t.Errorf("data1: expected '%s', got '%s", data1, v)
        }
    }else{
        t.Errorf("data1: %v", err)
    }

//  test for a file that exists only in global directories
    if v, err := GetDataFilename(`my-app/data2.file`); err == nil {
        if v != data2 {
            t.Errorf("data2: expected '%s', got '%s", data2, v)
        }
    }else{
        t.Errorf("data2: %v", err)
    }

//  test for a file that exists only in the last global directory
    if v, err := GetDataFilename(`my-app/data3.file`); err == nil {
        if v != data3 {
            t.Errorf("data3: expected '%s', got '%s", data3, v)
        }
    }else{
        t.Errorf("data3: %v", err)
    }

//  test for a file that does not exist
    if v, err := GetDataFilename(`nonexistent-dir/nothing.file`); err == nil {
        t.Errorf("File exists, but should not: %s", v)
    }
}



func TestGetConfigFilename(t *testing.T) {
    config1 := `testing/xdg-basedir/home/gotest/.config/my-app/config1.txt`
    config2 := `testing/xdg-basedir/etc/xdg/my-app/config2.txt`

//  mock home directory
    XdgConfigHome = `testing/xdg-basedir/home/gotest/.config`

//  mock config dirs
    XdgConfigDirs = `testing/xdg-basedir/etc/xdg`

//  test for a file that exists in all directories
    if v, err := GetConfigFilename(`my-app/config1.txt`); err == nil {
        if v != config1 {
            t.Errorf("config1: expected '%s', got '%s", config1, v)
        }
    }else{
        t.Errorf("config1: %v", err)
    }

//  test for a file that exists only in global directories
    if v, err := GetConfigFilename(`my-app/config2.txt`); err == nil {
        if v != config2 {
            t.Errorf("config2: expected '%s', got '%s", config2, v)
        }
    }else{
        t.Errorf("config2: %v", err)
    }

//  test for a file that does not exist
    if v, err := GetConfigFilename(`nonexistent-dir/nothing.file`); err == nil {
        t.Errorf("File exists, but should not: %s", v)
    }
}

