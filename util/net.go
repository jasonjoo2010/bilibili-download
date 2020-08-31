package util

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"strings"

	"github.com/jasonjoo2010/bilibili-download/types"
	"golang.org/x/time/rate"
)

var (
	REGEXP_PLAYINFO = regexp.MustCompile("__playinfo__\\s*=\\s*([\\s\\S]+?)<\\/script>")
	LOG_LIMITER     = rate.NewLimiter(1, 1)
)

func readBody(r io.Reader) string {
	buf := make([]byte, 1024)
	b := strings.Builder{}
	for {
		n, err := r.Read(buf)
		if n > 0 {
			b.Write(buf[:n])
		}
		if err != nil {
			break
		}
	}
	return b.String()
}

func newRequest(url string) (*http.Request, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Referer", "https://www.bilibili.com/")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_6) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/13.1.2 Safari/605.1.15")
	return req, nil
}

func GetPlayerInfo(url string) (*types.PlayInfo, error) {
	req, err := newRequest(url)
	if err != nil {
		return nil, err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body := readBody(resp.Body)
	result := REGEXP_PLAYINFO.FindStringSubmatch(body)
	if len(result) < 2 {
		return nil, errors.New("Could not find play infomation")
	}
	obj := types.Response{}
	json.Unmarshal([]byte(result[1]), &obj)
	if obj.Code != 0 {
		return nil, errors.New(obj.Msg)
	}

	return &obj.Data, nil
}

func writeFile(f *os.File, data []byte) error {
	pos := 0
	for pos < len(data) {
		n, err := f.Write(data[pos:])
		if n >= 0 {
			pos += n
		}
		if err != nil {
			return err
		}
	}
	return nil
}

func Download(url, path string) error {
	req, err := newRequest(url)
	if err != nil {
		return err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	total := resp.ContentLength
	f, err := os.Create(path)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Create destination file failed:", err.Error())
	}

	buf := make([]byte, 1024*500)
	bytes := 0
	for {
		n, err := resp.Body.Read(buf)
		if n > 0 {
			if write_err := writeFile(f, buf[:n]); err != nil {
				return write_err
			}
			bytes += n
			if LOG_LIMITER.Allow() {
				fmt.Printf("Downloading: %d%% %d/%d\r", int64(bytes)*100/total, bytes, total)
			}
		}
		if err != nil {
			break
		}
	}
	fmt.Println()
	fmt.Println("Downloaded")
	return nil
}
