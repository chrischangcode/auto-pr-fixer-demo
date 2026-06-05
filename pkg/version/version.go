package version

import "runtime/debug"
import "runtime"

// gitTag is set via ldflags at build time.
var gitTag string

const unknownVersionInfo = "unknown"

// Version holds build metadata.
type Version struct {
	Version   string
	GoVersion string
	BuildTime string
	Platform  string
}

// Get returns the current build version info.
func Get() Version {
	v := Version{
		Version:   unknownVersionInfo,
		GoVersion: unknownVersionInfo,
		BuildTime: unknownVersionInfo,
		Platform:  runtime.GOOS + "/" + runtime.GOARCH,
	}
	if gitTag != "" {
		v.Version = gitTag
	}
	info, ok := debug.ReadBuildInfo()
	if ok && info.GoVersion != "" {
		v.GoVersion = info.GoVersion
	}
	return v
}