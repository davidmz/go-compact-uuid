package cuuid_test

import (
	"fmt"

	cuuid "github.com/davidmz/go-compact-uuid"
	"github.com/gofrs/uuid"
)

func ExampleUUID_String() {
	u := uuid.Must(uuid.FromString("3867b6f3-dbb0-4ef5-8078-364897154fd9"))

	compact := cuuid.UUID(u).String()
	fmt.Println(compact)
	// Output: 3gsxpyfdv0kqn80y1p92bhakys
}

func ExampleFromString() {
	cu, err := cuuid.FromString("3gsxpyfdv0kqn80y1p92bhakys")
	fmt.Println(cu, err)
	fmt.Println(uuid.UUID(cu))
	// Output:
	// 3gsxpyfdv0kqn80y1p92bhakys <nil>
	// 3867b6f3-dbb0-4ef5-8078-364897154fd9
}

func ExampleFromString_mixedCase() {
	cu, err := cuuid.FromString("3GSXPYFDV0Kqn80y1p92BHAkys")
	fmt.Println(cu, err)
	// Output:
	// 3gsxpyfdv0kqn80y1p92bhakys <nil>
}

func ExampleFromString_error() {
	cu, err := cuuid.FromString("3gsxpyfdv0kqn80y1p92bhak##")
	fmt.Println(cu, err)
	// Output:
	// 00000000000000000000000000 uuid: incorrect UUID format 3gsxpyfdv0kqn80y1p92bhak##
}
