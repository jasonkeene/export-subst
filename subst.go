// Package subst will fail to build if export substitution occurred.
package subst

type got = [len("$Format:%h$")]byte
type want = [len("Format:%h") + 2]byte

var check got = want{}
