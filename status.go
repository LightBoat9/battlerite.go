// Package battleritego provides access to the battlerite game data service
package battleritego

// Status contains information about the state of the Gamelocker API
// See: https://battlerite-docs.readthedocs.io/en/latest/status/status.html
type Status struct {
	Type    string
	ID      string
	Release string
	Version string
}
