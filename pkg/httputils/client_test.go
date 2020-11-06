/**
 * singsen
 * SNQU
 * laixingchen@sndo.com
 * 2020/10/26 15:02
 */
package httputils

import "testing"

func TestGet(t *testing.T) {
	re, err := Get("http://example.com")
	if err != nil {
		t.Error(err)
		return
	}
	if re.StatusCode != 200 {
		t.Fail()
	}
}

func TestPost(t *testing.T) {

}

func TestPostForm(t *testing.T) {

}

func TestPostJSON(t *testing.T) {

}

func TestBytes(t *testing.T) {

}

func TestJSON(t *testing.T) {

}
