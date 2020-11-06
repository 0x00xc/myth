/**
 * singsen
 * SNQU
 * laixingchen@sndo.com
 * 2020/10/23 16:39
 */
package httputils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

var c = &http.Client{Timeout: time.Second * 15}
var tracker Tracker

func SetClient(cli *http.Client) {
	c = cli
}

func SetTracker(t Tracker) {
	tracker = newWrapper(t, 16)
}

func Do(method string, addr string, body io.Reader, header ...url.Values) (*http.Response, error) {
	start := time.Now()
	req, err := http.NewRequest(method, addr, body)
	if err != nil {
		return nil, err
	}
	if len(header) > 0 {
		for k := range header[0] {
			req.Header.Set(k, header[0].Get(k))
		}
	}
	re, err := c.Do(req)
	if tracker != nil {
		track := &Track{
			Timestamp: start,
			Method:    method,
			Address:   addr,
			Error:     err,
			Latency:   time.Since(start),
		}
		if re != nil {
			track.Status = re.Status
			track.StatusCode = re.StatusCode
			track.ContentLength = re.ContentLength
		}
		tracker.Trace(track)
	}
	return re, err
}

func Get(addr string, header ...url.Values) (*http.Response, error) {
	return Do(http.MethodGet, addr, nil, header...)
}

func Post(addr string, body io.Reader, header ...url.Values) (*http.Response, error) {
	return Do(http.MethodPost, addr, body, header...)
}

func PostJSON(addr string, v interface{}, header ...url.Values) (*http.Response, error) {
	var h url.Values
	if len(header) > 0 {
		h = header[0]
	} else {
		h = url.Values{}
	}
	h.Set("Content-Type", "application/json")
	buf := bytes.NewBuffer([]byte{})
	err := json.NewEncoder(buf).Encode(v)
	if err != nil {
		return nil, err
	}
	return Post(addr, buf, h)
}

func PostForm(addr string, form url.Values, header ...url.Values) (*http.Response, error) {
	var h url.Values
	if len(header) > 0 {
		h = header[0]
	} else {
		h = url.Values{}
	}
	h.Set("Content-Type", "application/x-www-form-urlencoded")
	return Post(addr, strings.NewReader(form.Encode()), h)
}

func Bytes(re *http.Response, err error) ([]byte, error) {
	if err != nil {
		return nil, err
	}
	if re.Body != nil {
		defer re.Body.Close()
	}
	var b []byte
	b, err = ioutil.ReadAll(re.Body)
	if re.StatusCode != http.StatusOK {
		err = fmt.Errorf("%d %s", re.StatusCode, re.Status)
	}
	return b, err
}

type jsonDecoder struct {
	decoder *json.Decoder
	err     error
}

func (d *jsonDecoder) Decode(v interface{}) error {
	if d.err != nil {
		return d.err
	}
	return d.decoder.Decode(v)
}

func JSON(re *http.Response, err error) *jsonDecoder {
	b, err := Bytes(re, err)
	return &jsonDecoder{decoder: json.NewDecoder(bytes.NewReader(b)), err: err}
}
