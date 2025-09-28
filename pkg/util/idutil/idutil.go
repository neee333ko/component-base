package idutil

import (
	"crypto/rand"
	"strconv"
	"strings"

	"github.com/neee333ko/component-base/pkg/util/iputil"
	"github.com/neee333ko/component-base/pkg/util/stringutil"
	"github.com/sony/sonyflake/v2"
	"github.com/speps/go-hashids/v2"
)

const (
	Alphabet62 string = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	Alphabet36 string = "abcdefghijklmnopqrstuvwxyz1234567890"
)

var sf *sonyflake.Sonyflake

func init() {
	setting := sonyflake.Settings{}

	setting.MachineID = func() (int, error) {
		ip := iputil.GetLocalIP()

		parts := strings.Split(ip, ".")

		high, _ := strconv.Atoi(parts[2])
		low, _ := strconv.Atoi(parts[3])

		return high<<8 + low, nil
	}

	var err error

	sf, err = sonyflake.New(setting)

	if err != nil {
		panic(err)
	}
}

func GetNextID() int64 {
	id, err := sf.NextID()

	if err != nil {
		panic(err)
	}

	return id
}

func GetInstanceID(uid int64, prefix string) string {
	hsdata := hashids.NewData()
	hsdata.Alphabet = Alphabet36
	hsdata.MinLength = 6
	hsdata.Salt = "neee333ko"

	hs, err := hashids.NewWithData(hsdata)
	if err != nil {
		panic(err)
	}

	hsstring, err := hs.EncodeInt64([]int64{uid})
	if err != nil {
		panic(err)
	}

	return prefix + stringutil.Reverse(hsstring)
}

func GetUUID36(prefix string) string {
	hsdata := hashids.NewData()
	hsdata.Alphabet = Alphabet36

	id := GetNextID()
	hs, err := hashids.NewWithData(hsdata)
	if err != nil {
		panic(err)
	}

	hsstring, err := hs.EncodeInt64([]int64{id})
	if err != nil {
		panic(err)
	}

	return prefix + stringutil.Reverse(hsstring)
}

func RandString(sample string, n int) string {
	result := make([]byte, n)

	randomNum := make([]byte, n)

	rand.Read(randomNum)

	for i := range result {
		rn := randomNum[i]
		pos := int(rn) % n

		result[i] = sample[pos]
	}

	return string(result)
}

func NewSecretID() string {
	return RandString(Alphabet62, 36)
}

func NewSecretKey() string {
	return RandString(Alphabet36, 32)
}
