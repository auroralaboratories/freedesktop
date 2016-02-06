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

type Theme struct {
    Comment      string
    Contexts     map[string]IconContext
    Directories  []string
    Example      string
    Hidden       bool
    Icons        []*Icon
    IndexFile    string
    Inherits     []string
    InternalName string
    Name         string
    Path         string

    iconIndex    map[string]*Icon
}

func NewTheme(themePath string) *Theme {
    return &Theme{
        Contexts:  make(map[string]IconContext),
        Icons:     make([]*Icon, 0),
        IndexFile: DEFAULT_ICONTHEME_INDEX_FILE,
        Inherits:  make([]string, 0),
        Path:      themePath,

        iconIndex: make(map[string]*Icon),
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
                        context := IconContext{}

                        if v, err := stringutil.ConvertToInteger(contextConfig[`Size`]); err == nil {
                            context.Size = int(v)
                        }

                        if v, ok := contextConfig[`Context`]; ok {
                            context.Name = v
                        }else{
                            continue
                        }

                        switch strings.ToLower(contextConfig[`Type`]) {
                        case `fixed`:
                            context.Type = IconContextFixed

                            if v, err := stringutil.ConvertToInteger(contextConfig[`MinSize`]); err == nil {
                                context.MinSize = int(v)
                            }

                            if v, err := stringutil.ConvertToInteger(contextConfig[`MaxSize`]); err == nil {
                                context.MaxSize = int(v)
                            }

                        case `scalable`:
                            context.Type = IconContextScalable

                        default:
                            context.Type = IconContextThreshold

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
//  populate icons
    for _, directory := range self.Directories {
        iconBaseDir := path.Join(self.Path, directory)

        if stat, err := os.Stat(iconBaseDir); err == nil && stat.IsDir() {
            if files, err := filepath.Glob(path.Join(iconBaseDir, `*.*`)); err == nil {
                for _, iconFilename := range files {
                    switch filepath.Ext(iconFilename) {
                    case `.png`, `.xpm`, `.svg`:
                        icon := NewIcon(iconFilename, self)

                        if context, ok := self.Contexts[directory]; ok {
                            icon.Context = context
                        }

                        if err := icon.Refresh(); err == nil {
                        //  populate index keyed on name alone
                            if _, ok := self.iconIndex[icon.Name]; !ok {
                                // log.Printf("IDX %s -> %s", self.InternalName, icon.Name)
                                self.iconIndex[icon.Name] = icon
                            }

                        //  if an appropriate context was found...
                            if icon.Context.IsValid() {
                                indexKey := icon.Name

                                switch icon.Context.Type {
                                case IconContextScalable:
                                    indexKey = indexKey + `:scalable`
                                default:
                                    indexKey = fmt.Sprintf("%s:%d", indexKey, icon.Context.Size)
                                }

                            //  populate index keyed on name-size
                                if _, ok := self.iconIndex[indexKey]; !ok {
                                    // log.Printf("IDX %s -> %s", self.InternalName, indexKey)
                                    self.iconIndex[indexKey] = icon
                                }
                            }

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
func (self *Theme) FindIcon(name string, size int) (*Icon, bool) {
    var indexKey string

    if size <= 0 {
        indexKey = fmt.Sprintf("%s:scalable", name)
    }else{
        indexKey = fmt.Sprintf("%s:%d", name, size)
    }

    if icon, ok := self.iconIndex[indexKey]; ok {
        return icon, true
    }else{
        if icon, ok := self.iconIndex[name]; ok {
            return icon, true
        }
    }

    return nil, false
}