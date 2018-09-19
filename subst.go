// Package subst tests if export substitution occured.
package subst

const (
	fixture      = "$Format:%h$"
	expectedData = "\x24\x46\x6f\x72\x6d\x61\x74\x3a\x25\x68\x24"
)

// Enabled returns true if export-subst was enabled when the zip file was
// created for this module.
func Enabled() bool {
	return fixture != expectedData
}
