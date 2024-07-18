# Homework

## Question 1
### Intro
Assume we have a serials TCP stream, and we already split it and generate it to a series of frames.


### Require

1. You need parsing frames and dump visible characters.
2. You need complete tests codes without errors.
3. You need to think about time complexity.

### Frame rules

````
|layout| bytes | type |
|---| ---| --- |
| layer-1| 8 bytes| header |

````

### layer-1
* BigEndian([byte[0], byte[1]]) represents frame bytes length.
* (byte[2], byte[3]) represents data flags, you can ignore its.
* byte[4] represent frame contains SQL payload, you need dump it.

### Code

```go
package tests

import (
	"encoding/binary"
	"github.com/stretchr/testify/assert"
	"testing"
)

type Cases struct {
	Name      string
	Payload   []byte
	Result    string
	ResultRaw []byte
}

func Test_query(t *testing.T) {
	cases := []Cases{
		{
			Name:      "SQLPLUS tools connect packet with select sql",
			Payload:   []byte{1, 75, 0, 0, 6, 0, 0, 0, 0, 0, 3, 94, 6, 97, 128, 0, 0, 0, 0, 0, 0, 254, 255, 255, 255, 255, 255, 255, 255, 23, 1, 0, 0, 254, 255, 255, 255, 255, 255, 255, 255, 13, 0, 0, 0, 254, 255, 255, 255, 255, 255, 255, 255, 254, 255, 255, 255, 255, 255, 255, 255, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 254, 255, 255, 255, 255, 255, 255, 255, 0, 0, 0, 0, 0, 0, 0, 0, 254, 255, 255, 255, 255, 255, 255, 255, 254, 255, 255, 255, 255, 255, 255, 255, 254, 255, 255, 255, 255, 255, 255, 255, 0, 0, 0, 0, 0, 0, 0, 0, 254, 255, 255, 255, 255, 255, 255, 255, 254, 255, 255, 255, 255, 255, 255, 255, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 93, 83, 69, 76, 69, 67, 84, 32, 68, 69, 67, 79, 68, 69, 40, 85, 83, 69, 82, 44, 32, 39, 88, 83, 36, 78, 85, 76, 76, 39, 44, 32, 32, 88, 83, 95, 83, 89, 83, 95, 67, 79, 78, 84, 69, 88, 84, 40, 39, 88, 83, 36, 83, 69, 83, 83, 73, 79, 78, 39, 44, 39, 85, 83, 69, 82, 78, 65, 77, 69, 39, 41, 44, 32, 85, 83, 69, 82, 41, 32, 70, 82, 79, 77, 32, 83, 89, 83, 46, 68, 85, 65, 76, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 128, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			Result:    "SELECT DECODE(USER, 'XS$NULL',  XS_SYS_CONTEXT('XS$SESSION','USERNAME'), USER) FROM SYS.DUAL",
			ResultRaw: []byte("SELECT DECODE(USER, 'XS$NULL',  XS_SYS_CONTEXT('XS$SESSION','USERNAME'), USER) FROM SYS.DUAL"),
		},
		{
			Name:      "SQLPLUS custom select sql",
			Payload:   []byte{0, 254, 0, 0, 6, 0, 0, 0, 0, 0, 3, 94, 13, 97, 128, 0, 0, 0, 0, 0, 0, 254, 255, 255, 255, 255, 255, 255, 255, 48, 0, 0, 0, 254, 255, 255, 255, 255, 255, 255, 255, 13, 0, 0, 0, 254, 255, 255, 255, 255, 255, 255, 255, 254, 255, 255, 255, 255, 255, 255, 255, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 254, 255, 255, 255, 255, 255, 255, 255, 0, 0, 0, 0, 0, 0, 0, 0, 254, 255, 255, 255, 255, 255, 255, 255, 254, 255, 255, 255, 255, 255, 255, 255, 254, 255, 255, 255, 255, 255, 255, 255, 0, 0, 0, 0, 0, 0, 0, 0, 254, 255, 255, 255, 255, 255, 255, 255, 254, 255, 255, 255, 255, 255, 255, 255, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 16, 115, 101, 108, 101, 99, 116, 32, 42, 32, 102, 114, 111, 109, 32, 120, 101, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 128, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			Result:    "select * from xe",
			ResultRaw: []byte("select * from xe"),
		},
		{
			Name:      "Navicat 16 Alter ddl",
			Payload:   []byte{1, 42, 0, 0, 6, 0, 0, 0, 0, 0, 17, 105, 103, 254, 255, 255, 255, 255, 255, 255, 255, 1, 0, 0, 0, 2, 0, 0, 0, 3, 94, 104, 33, 129, 0, 0, 0, 0, 0, 0, 254, 255, 255, 255, 255, 255, 255, 255, 123, 0, 0, 0, 254, 255, 255, 255, 255, 255, 255, 255, 13, 0, 0, 0, 254, 255, 255, 255, 255, 255, 255, 255, 254, 255, 255, 255, 255, 255, 255, 255, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 254, 255, 255, 255, 255, 255, 255, 255, 0, 0, 0, 0, 0, 0, 0, 0, 254, 255, 255, 255, 255, 255, 255, 255, 254, 255, 255, 255, 255, 255, 255, 255, 254, 255, 255, 255, 255, 255, 255, 255, 0, 0, 0, 0, 0, 0, 0, 0, 254, 255, 255, 255, 255, 255, 255, 255, 254, 255, 255, 255, 255, 255, 255, 255, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 41, 65, 76, 84, 69, 82, 32, 83, 69, 83, 83, 73, 79, 78, 32, 83, 69, 84, 32, 67, 85, 82, 82, 69, 78, 84, 95, 83, 67, 72, 69, 77, 65, 32, 61, 32, 115, 121, 115, 116, 101, 109, 1, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 7, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			Result:    "ALTER SESSION SET CURRENT_SCHEMA = system",
			ResultRaw: []byte("ALTER SESSION SET CURRENT_SCHEMA = system"),
		},

		{
			Name:      "Navicat 16 select sql",
			Payload:   []byte{1, 177, 0, 0, 6, 0, 0, 0, 0, 0, 17, 105, 83, 254, 255, 255, 255, 255, 255, 255, 255, 1, 0, 0, 0, 3, 0, 0, 0, 3, 94, 84, 97, 129, 0, 0, 0, 0, 0, 0, 254, 255, 255, 255, 255, 255, 255, 255, 4, 2, 0, 0, 254, 255, 255, 255, 255, 255, 255, 255, 13, 0, 0, 0, 254, 255, 255, 255, 255, 255, 255, 255, 254, 255, 255, 255, 255, 255, 255, 255, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 254, 255, 255, 255, 255, 255, 255, 255, 0, 0, 0, 0, 0, 0, 0, 0, 254, 255, 255, 255, 255, 255, 255, 255, 254, 255, 255, 255, 255, 255, 255, 255, 254, 255, 255, 255, 255, 255, 255, 255, 0, 0, 0, 0, 0, 0, 0, 0, 254, 255, 255, 255, 255, 255, 255, 255, 254, 255, 255, 255, 255, 255, 255, 255, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 254, 64, 83, 69, 76, 69, 67, 84, 32, 67, 46, 84, 65, 66, 76, 69, 95, 78, 65, 77, 69, 44, 32, 67, 46, 67, 79, 76, 85, 77, 78, 95, 78, 65, 77, 69, 44, 32, 67, 46, 79, 87, 78, 69, 82, 32, 70, 82, 79, 77, 32, 34, 83, 89, 83, 34, 46, 34, 65, 76, 76, 95, 84, 65, 66, 95, 64, 67, 79, 76, 85, 77, 78, 83, 34, 32, 67, 32, 87, 72, 69, 82, 69, 32, 67, 46, 79, 87, 78, 69, 82, 32, 61, 32, 39, 67, 84, 88, 83, 89, 83, 39, 32, 65, 78, 68, 32, 67, 46, 84, 65, 66, 76, 69, 95, 78, 65, 77, 69, 32, 61, 32, 39, 68, 82, 36, 73, 78, 68, 69, 88, 44, 95, 83, 69, 84, 39, 32, 79, 82, 68, 69, 82, 32, 66, 89, 32, 67, 46, 84, 65, 66, 76, 69, 95, 78, 65, 77, 69, 44, 32, 67, 46, 67, 79, 76, 85, 77, 78, 95, 73, 68, 32, 65, 83, 67, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			Result:    `SELECT C.TABLE_NAME, C.COLUMN_NAME, C.OWNER FROM "SYS"."ALL_TAB_@COLUMNS" C WHERE C.OWNER = 'CTXSYS' AND C.TABLE_NAME = 'DR$INDEX,_SET' ORDER BY C.TABLE_NAME, C.COLUMN_ID ASC`,
			ResultRaw: []byte(`SELECT C.TABLE_NAME, C.COLUMN_NAME, C.OWNER FROM "SYS"."ALL_TAB_@COLUMNS" C WHERE C.OWNER = 'CTXSYS' AND C.TABLE_NAME = 'DR$INDEX,_SET' ORDER BY C.TABLE_NAME, C.COLUMN_ID ASC`),
		},
	}

	for _, c := range cases {
		t.Logf("Testing %s", c.Name)
		result, err := parser(c.Payload)
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, int(binary.BigEndian.Uint16(c.Payload[0:2])), len(c.Payload))
		assert.Equal(t, c.Result, string(result))
		assert.Equal(t, c.ResultRaw, result)
	}
}

func parser(payload []byte) ([]byte, error) {
	return nil, nil
}

```

