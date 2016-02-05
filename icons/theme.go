package icons

import (
    "fmt"
    "os"
    "path"
    "path/filepath"
    "strings"
    // "log"

    "github.com/shutterstock/go-stockutil/stringutil"
    "github.com/vaughan0/go-ini"
)

const (
    DEFAULT_ICONTHEME_INDEX_FILE = `index.theme`
    DEFAULT_ICONTHEME_INHERIT    = `hicolor`
)

type ThemeContextType int
const (
    ThemeContextThreshold ThemeContextType = 0
    ThemeContextFixed                      = 1
    ThemeContextScalable                   = 2
)

type ThemeContext struct {
    Size      int
    Context   string
    Type      ThemeContextType
    MaxSize   int               // Type=ThemeContextScalable
    MinSize   int               // Type=ThemeContextScalable
    Threshold int               // Type=ThemeContextThreshold
}

type Theme struct {
    Comment      string
    Contexts     map[string]ThemeContext
    Directories  []string
    Example      string
    Hidden       bool
    Icons        []*Icon
    IndexFile    string
    Inherits     []string
    InternalName string
    Name         string
    Path         string
}

func NewTheme(themePath string) *Theme {
    return &Theme{
        Contexts:  make(map[string]ThemeContext),
        Icons:     make([]*Icon, 0),
        IndexFile: DEFAULT_ICONTHEME_INDEX_FILE,
        Inherits:  make([]string, 0),
        Path:      themePath,
    }
}


func (self *Theme) Refresh() error {
    if self.Path == `` {
        return fmt.Errorf("Cannot refresh theme without a given path")
    }

    if err := self.refreshThemeDefinition(); err != nil {
        return err
    }

    if err := self.refreshIcons(); err != nil {
        return err
    }

    return nil
}

func (self *Theme) refreshThemeDefinition() error {
    themeIndexFilename := path.Join(self.Path, self.IndexFile)

    if themeIndex, err := ini.LoadFile(themeIndexFilename); err == nil {
        if config, ok := themeIndex[`Icon Theme`]; ok {
            self.InternalName = path.Base(path.Dir(themeIndexFilename))

            if v, ok := config[`Name`]; ok {
                self.Name = v
            }

            self.Inherits = strings.Split(config[`Inherits`], `,`)

            if len(self.Inherits) == 0 {
                self.Inherits = []string{ DEFAULT_ICONTHEME_INHERIT }
            }

            self.Comment  = config[`Comment`]
            self.Example  = config[`Example`]

            if v, ok := config[`Hidden`]; ok {
                self.Hidden = (v == `true`)
            }

            if v, ok := config[`Directories`]; ok {
                self.Directories = strings.Split(v, `,`)

                for _, directory := range self.Directories {
                    if contextConfig, ok := themeIndex[directory]; ok {
                        context := ThemeContext{}

                        if v, err := stringutil.ConvertToInteger(contextConfig[`Size`]); err == nil {
                            context.Size = int(v)
                        }

                        if v, ok := contextConfig[`Context`]; ok {
                            context.Context = v
                        }else{
                            continue
                        }

                        switch strings.ToLower(contextConfig[`Type`]) {
                        case `fixed`:
                            context.Type = ThemeContextFixed

                            if v, err := stringutil.ConvertToInteger(contextConfig[`MinSize`]); err == nil {
                                context.MinSize = int(v)
                            }

                            if v, err := stringutil.ConvertToInteger(contextConfig[`MaxSize`]); err == nil {
                                context.MaxSize = int(v)
                            }

                        case `scalable`:
                            context.Type = ThemeContextScalable

                        default:
                            context.Type = ThemeContextThreshold

                            if v, err := stringutil.ConvertToInteger(contextConfig[`Threshold`]); err == nil {
                                context.Threshold = int(v)
                            }
                        }

                        self.Contexts[directory] = context
                    }
                }
            }

        }else{
            return fmt.Errorf("Cannot load theme at %s: missing [Icon Theme] section", themeIndexFilename)
        }
    }else{
        return err
    }

    return nil
}

func (self *Theme) refreshIcons() error {
    for _, directory := range self.Directories {
        iconBaseDir := path.Join(self.Path, directory)

        if stat, err := os.Stat(iconBaseDir); err == nil && stat.IsDir() {
            if files, err := filepath.Glob(path.Join(iconBaseDir, `*.*`)); err == nil {
                for _, iconFilename := range files {
                    switch filepath.Ext(iconFilename) {
                    case `.png`, `.xpm`, `.svg`:
                        icon := NewIcon(iconFilename, self)

                        if context, ok := self.Contexts[directory]; ok {
                            icon.Context = &context
                        }

                        if err := icon.Refresh(); err == nil {
                            self.Icons = append(self.Icons, icon)
                        }else{
                            return err
                        }
                    }
                }
            }
        }
    }

    return nil
}

//  Lookup an icon by name and desired size
//
//  The icon lookup mechanism has two global settings, the list of base directories and the internal name of the
//  current theme. Given these we need to specify how to look up an icon file from the icon name and the nominal size.
//
//  The lookup is done first in the current theme, and then recursively in each of the current theme's parents, and
//  finally in the default theme called "hicolor" (implementations may add more default themes before "hicolor", but
//  "hicolor" must be last). As soon as there is an icon of any size that matches in a theme, the search is stopped.
//  Even if there may be an icon with a size closer to the correct one in an inherited theme, we don't want to use it.
//  Doing so may generate an inconsistant change in an icon when you change icon sizes (e.g. zoom in).
//
//  The lookup inside a theme is done in three phases. First all the directories are scanned for an exact match, e.g.
//  one where the allowed size of the icon files match what was looked up. Then all the directories are scanned for any
//  icon that matches the name. If that fails we finally fall back on unthemed icons. If we fail to find any icon at
//  all it is up to the application to pick a good fallback, as the correct choice depends on the context.
//
// func (self *Theme) FindIcon(name string, size int) (Icon, bool) {
//     if , err := filepath.Glob(path.Join(self.Path, `*`, fmt.Sprintf())); err == nil {
// }