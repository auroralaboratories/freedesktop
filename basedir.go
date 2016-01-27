// Provides definitions and helper functions for working with XDG
// data, config, and cache files in standard locations
//
// See also: http://standards.freedesktop.org/basedir-spec/basedir-spec-0.6.html
//
package freedesktop

import (
    "fmt"
    "os"
    "github.com/auroralaboratories/freedesktop/util"
)

// the base directory relative to which user specific data files should be stored
var XdgDataHome   string = util.Getenv(`XDG_DATA_HOME`,   os.ExpandEnv("${HOME}/.local/share"))

// the preference-ordered set of base directories to search for data files in addition to the XdgDataHome base directory
var XdgDataDirs   string = util.Getenv(`XDG_DATA_DIRS`,   `/usr/local/share/:/usr/share/`)

// the base directory relative to which user specific configuration files should be stored
var XdgConfigHome string = util.Getenv(`XDG_CONFIG_HOME`, os.ExpandEnv("${HOME}/.config"))

// the preference-ordered set of base directories to search for configuration files in addition to the XdgConfigHome base directory
var XdgConfigDirs string = util.Getenv(`XDG_CONFIG_DIRS`, `/etc/xdg`)

// the base directory relative to which user specific non-essential data files should be stored
var XdgCacheHome  string = util.Getenv(`XDG_CACHE_HOME`,  os.ExpandEnv("${HOME}/.cache"))


// Returns the filename of a data file located in a standard XDG location.
// The name should be specified relative to whichever root it may live in.
// For example, if a file is expected to be at $HOME/.local/share/my-app/data.ini,
// you would supply "my-app/data.ini" as the first argument.
//
// Directories will be search in preference order, starting with XdgDataHome, then
// proceeding through each directory listed in the colon-separated XdgDataDirs list.
//
// An error will be returned if a file could not be located or is not readable.
//
func GetDataFilename(name string) (string, error) {
    return ``, fmt.Errorf("Not implemented")
}

// Returns the filename of a config file located in a standard XDG location.
// The name should be specified relative to whichever root it may live in.
// For example, if a file is expected to be at $HOME/.config/my-app/config.ini,
// you would supply "my-app/config.ini" as the first argument.
//
// Directories will be search in preference order, starting with XdgConfigHome, then
// proceeding through each directory listed in the colon-separated XdgConfigDirs list.
//
// An error will be returned if a file could not be located or is not readable.
//
func GetConfigFilename(name string) (string, error) {
    return ``, fmt.Errorf("Not implemented")
}

// Returns the filename of a cache file located in a standard XDG location.
// The name should be specified relative to whichever root it may live in.
// For example, if a file is expected to be at $HOME/.cache/my-app/cache.dat,
// you would supply "my-app/cache.dat" as the first argument.
//
// The directory at XdgCacheHome will be searched.
//
// An error will be returned if a file could not be located or is not readable.
//
func GetCacheFilename(name string) (string, error) {
    return ``, fmt.Errorf("Not implemented")
}