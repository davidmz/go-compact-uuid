package cuuid // import "github.com/davidmz/go-compact-uuid"
import (
	"encoding/binary"
	"fmt"
)

// UUID is an array to represent UUID value. This type
// is compatible with the most of other UUID packages.
type UUID [16]byte

// String returns a compact string representation of UUID,
// the 26-character string of '0123456789abcdefghjkmnpqrstvwxyz' alphabet.
func (u UUID) String() string {
	result := make([]byte, 0, 26)
	longs := []uint64{
		binary.BigEndian.Uint64(u[0:]),
		binary.BigEndian.Uint64(u[8:]),
	}

	for _, long := range longs {
		result = append(result, alphabet[(long>>60)&0x0f])
		for i := 0; i < 12; i++ {
			off := uint(55 - i*5)
			result = append(result, alphabet[(long>>off)&0x1f])
		}
	}
	return string(result)
}

// FromString parses text and returns UUID or error. It expects the 26-character
// string of '0123456789abcdefghjkmnpqrstvwxyz' alphabet (case-insensitive).
func FromString(text string) (UUID, error) {
	u := UUID{}
	if len(text) != 26 {
		return u, fmt.Errorf("uuid: incorrect UUID length: %s", text)
	}

	for n := 0; n < 2; n++ {
		var long uint64
		for i, b := range []byte(text)[n*13:] {
			var c byte = 255
			if '0' <= b && b <= '9' {
				c = b - '0'
			} else {
				b |= 0x20 // lowcase in ASCII
				for j, a := range alphabet {
					if a == b {
						c = byte(j)
						break
					}
				}
				if c > 31 || i == 0 && c > 15 {
					return u, fmt.Errorf("uuid: incorrect UUID format %s", text)
				}
			}
			if i == 0 {
				long |= uint64(c) << 60
			} else {
				long |= uint64(c) << uint(60-5*i)
			}
		}
		binary.BigEndian.PutUint64(u[n*8:], long)
	}

	return u, nil
}

var alphabet = []byte("0123456789abcdefghjkmnpqrstvwxyz")
