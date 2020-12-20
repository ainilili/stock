/**
2 * @Author: Nico
3 * @Date: 2020/12/21 3:30
4 */
package http

import "testing"

func TestDecode(t *testing.T) {
	t.Log(Decode("Hello \xb3\xa3\xd3\xc3\x87\xf8\xd7\xd6\x98\xcb\x9c\xca\xd7\xd6\xf3\x77\xb1\xed", "gb18030"))
}


