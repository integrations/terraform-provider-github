package github

import "fmt"

func projectV2OwnerMetadata(meta any) *Owner {
	owner, ok := meta.(*Owner)
	if !ok {
		panic(fmt.Sprintf("Projects V2 provider metadata has unexpected type %T", meta))
	}
	return owner
}
