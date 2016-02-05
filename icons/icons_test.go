package icons

import (
    "strings"
    "testing"
)


func TestGetIconThemePaths(t *testing.T) {
    have := GetIconThemePaths()

    if len(have) > 0 {
        for i, p := range have {
            t.Logf("Path %d: %s", i, p)
        }
    }else{
        t.Errorf("Did not discover any icon theme paths")
    }
}


func TestLoadThemes(t *testing.T) {
    if themes, err := LoadThemes(); err == nil {
        if len(themes) > 0 {
            for i, p := range themes {
                t.Logf("Theme %02d: %s", i, p.Path)
                t.Logf("          IName:    %s", p.InternalName)
                t.Logf("          Name:     %s", p.Name)
                t.Logf("          Inherits: %s", strings.Join(p.Inherits, `, `))
                t.Logf("          Icons:    %d", len(p.Icons))
                t.Logf("\n")
            }
        }else{
            t.Errorf("Did not discover any icon themes")
        }
    }else{
        t.Errorf("Error loading icon themes: %v", err)
    }
}
