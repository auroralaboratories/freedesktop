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