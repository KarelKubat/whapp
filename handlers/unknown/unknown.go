package unknown

import "fmt"

type unknown struct {
	fingerprint string
}

func New(fp string) *unknown {
	return &unknown{
		fingerprint: fp,
	}
}

func (u *unknown) String() string {
	return fmt.Sprintf("unknown: [%s]", u.fingerprint)
}
