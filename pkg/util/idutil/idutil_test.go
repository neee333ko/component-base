package idutil

import (
	"fmt"
	"testing"
)

func TestIDutil(t *testing.T) {
	id1 := GetNextID()
	fmt.Printf("id1 = %d\n", id1)

	id2 := GetNextID()
	fmt.Printf("id2 = %d\n", id2)

	instanceID1 := GetInstanceID(1, "secret-")
	fmt.Printf("instanceID1 = %s\n", instanceID1)

	instanceID2 := GetInstanceID(2, "user-")
	fmt.Printf("instanceID2 = %s\n", instanceID2)

	uuid36 := GetInstanceID(1, "uuid")
	fmt.Printf("uuid36 = %s\n", uuid36)

	sid1 := NewSecretID()
	sid2 := NewSecretID()

	skey1 := NewSecretKey()
	skey2 := NewSecretKey()

	fmt.Printf("sid1 = %s, sid2 = %s\n", sid1, sid2)
	fmt.Printf("skey1 = %s, skey2 = %s\n", skey1, skey2)
}
