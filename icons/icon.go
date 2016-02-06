package icons

import (
    "path/filepath"
    "strings"

    "github.com/vaughan0/go-ini"
)

type Icon struct {
    Context               IconContext
    Filename              string
    Name                  string
    Theme                 *Theme
    DisplayName           string
    EmbeddedTextRectangle *Rectangle
    AttachPoints          []Point


    dataFilename          string
}

func NewIcon(filename string, theme *Theme) *Icon {
    rv := &Icon{
        Filename:     filename,
        Name:         strings.TrimSuffix(filepath.Base(filename), filepath.Ext(filename)),
        Theme:        theme,
        AttachPoints: make([]Point, 0),
    }

    rv.DisplayName = rv.Name

    return rv
}


func (self *Icon) Size() int {
    if self.Context.IsValid() {
        return self.Context.Size
    }else{
        return -1
    }
}

func (self *Icon) HasDataFile() bool {
    return (self.dataFilename != ``)
}

func (self *Icon) Refresh() error {
    dataFilename := strings.TrimSuffix(self.Filename, filepath.Ext(self.Filename)) + `.icon`

    if iconData, err := ini.LoadFile(dataFilename); err == nil {
        self.dataFilename = dataFilename

        if d, ok := iconData[`Icon Data`]; ok {
            if v, ok := d[`DisplayName`]; ok {
                self.DisplayName = v
            }

            if v, ok := d[`EmbeddedTextRectangle`]; ok {
                if rect, err := CreateRectangleFromString(v); err == nil {
                    self.EmbeddedTextRectangle = rect
                }
            }

            if v, ok := d[`AttachPoints`]; ok {
                self.AttachPoints = CreatePointsFromString(v)
            }
        }
    }

    return nil
}

//  Icon Directories
//      $HOME/.icons
//      freedesktop.GetDataFilename()
//      /usr/share/pixmaps



//  Lookup an icon by name and desired size from a specific theme
//
func FindIconInTheme(name string, size int, themeName string) (Icon, bool) {
    // Psuedocode from http://standards.freedesktop.org/icon-theme-spec/icon-theme-spec-latest.html
    //
    // for each subdir in $(theme subdir list) {
    // for each directory in $(basename list) {
    //   for extension in ("png", "svg", "xpm") {
    //     if DirectoryMatchesSize(subdir, size) {
    //       filename = directory/$(themename)/subdir/iconname.extension
    //       if exist filename
    //     return filename
    //     }
    //   }
    // }
    // }
    // minimal_size = MAXINT
    // for each subdir in $(theme subdir list) {
    // for each directory in $(basename list) {
    //   for extension in ("png", "svg", "xpm") {
    //     filename = directory/$(themename)/subdir/iconname.extension
    //     if exist filename and DirectorySizeDistance(subdir, size) < minimal_size {
    //    closest_filename = filename
    //    minimal_size = DirectorySizeDistance(subdir, size)
    //     }
    //   }
    // }
    // }
    // if closest_filename set
    //  return closest_filename
    // return none
    return Icon{}, false
}