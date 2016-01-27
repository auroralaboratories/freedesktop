package util

import (
    "os"
)

func Getenv(name string, fallback string) string {
    if rv := os.Getenv(name); rv == `` {
        return fallback
    }else{
        return rv
    }
}

func FileExistsAndIsReadable(name string) bool {
    file, err := os.Open(name)
    defer file.Close()
    return (err == nil)
}