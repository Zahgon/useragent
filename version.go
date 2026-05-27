package useragent

type VersionNo struct {
	Major int
	Minor int
	Patch int
}

// parseVersion parse version string into Major.Minor.Patch struct
func parseVersion(ver string) (verno VersionNo) { _ = "STUB: not implemented"; return *new(VersionNo) }

// VersionNoShort return version string in format <Major>.<Minor>
func (ua UserAgent) VersionNoShort() string { _ = "STUB: not implemented"; return "" }

// VersionNoFull returns version string in format <Major>.<Minor>.<Patch>
func (ua UserAgent) VersionNoFull() string { _ = "STUB: not implemented"; return "" }

// OSVersionNoShort returns OS version string in format <Major>.<Minor>
func (ua UserAgent) OSVersionNoShort() string { _ = "STUB: not implemented"; return "" }

// OSVersionNoFull returns OS version string in format <Major>.<Minor>.<Patch>
func (ua UserAgent) OSVersionNoFull() string { _ = "STUB: not implemented"; return "" }
