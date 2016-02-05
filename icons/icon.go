package icons

type Icon struct {
    Context  *ThemeContext
    Filename string
    Theme    *Theme
}

func NewIcon(filename string, theme *Theme) *Icon {
    return &Icon{
        Filename: filename,
        Theme:    theme,
    }
}


func (self *Icon) Refresh() error {
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