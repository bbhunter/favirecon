/*
favirecon - Use favicon.ico to improve your target recon phase. Quickly detect technologies, WAF, exposed panels, known services.

This repository is under MIT License https://github.com/edoardottt/favirecon/blob/main/LICENSE
*/

package favirecon

import (
	"bytes"
	"crypto/tls"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/edoardottt/favirecon/pkg/input"
	"github.com/projectdiscovery/gologger"
	"github.com/projectdiscovery/mapcidr"
	"github.com/twmb/murmur3"
)

const (
	TLSHandshakeTimeout = 10
	KeepAlive           = 30
	MinURLLength        = 4
)

var (
	ErrMalformedURL  = errors.New("malformed input URL")
	ErrCidrBadFormat = errors.New("malformed input CIDR")
)

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}

	return false
}

func getFavicon(url, ua string, client *http.Client) (string, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return "", err
	}

	gologger.Debug().Msgf("Checking favicon for %s", url)

	req.Header.Add("User-Agent", ua)

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return GetFaviconHash(body), nil
}

// PrepareURL takes as input a string and prepares
// the input URL in order to get the favicon icon.
func PrepareURL(input string) (string, error) {
	if len(input) < MinURLLength {
		return "", ErrMalformedURL
	}

	if !strings.Contains(input, "://") {
		input = "http://" + input
	}

	u, err := url.Parse(input)
	if err != nil {
		return "", err
	}

	if !(len(u.Path) > 3 && u.Path[len(u.Path)-4:] == ".ico") {
		if len(u.Path) == 0 || u.Path[len(u.Path)-1:] != "/" {
			u.Path += "/"
		}

		u.Path += "favicon.ico"
	}

	return u.Scheme + "://" + u.Host + u.Path, nil
}

// base64Content : RFC2045.
func base64Content(input []byte) []byte {
	inputEncoded := base64.StdEncoding.EncodeToString(input)
	buffer := bytes.Buffer{}

	for i := 0; i < len(inputEncoded); i++ {
		ch := inputEncoded[i]
		buffer.WriteByte(ch)

		if (i+1)%76 == 0 { // 76 bytes.
			buffer.WriteByte('\n')
		}
	}

	buffer.WriteByte('\n')

	return buffer.Bytes()
}

// GetFaviconHash computes the murmur3 hash.
func GetFaviconHash(input []byte) string {
	b64 := base64Content(input)
	return fmt.Sprint(int32(murmur3.Sum32(b64)))
}

func customClient(options *input.Options) (*http.Client, error) {
	transport := http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		Proxy:           http.ProxyFromEnvironment,
		Dial: (&net.Dialer{
			Timeout:   time.Duration(options.Timeout) * time.Second,
			KeepAlive: KeepAlive * time.Second,
		}).Dial,
		TLSHandshakeTimeout: TLSHandshakeTimeout * time.Second,
	}

	if options.Proxy != "" {
		u, err := url.Parse(options.Proxy)
		if err != nil {
			return nil, err
		}

		transport.Proxy = http.ProxyURL(u)

		if options.Verbose {
			gologger.Debug().Msgf("Using Proxy %s", options.Proxy)
		}
	}

	client := http.Client{
		Transport: &transport,
		Timeout:   time.Duration(options.Timeout) * time.Second,
	}

	return &client, nil
}

func handleCidrInput(inputCidr string) ([]string, error) {
	if !isCidr(inputCidr) {
		return nil, ErrCidrBadFormat
	}

	ips, err := mapcidr.IPAddresses(inputCidr)
	if err != nil {
		return nil, err
	}

	return ips, nil
}

// isCidr determines if the given ip is a cidr range.
func isCidr(inputCidr string) bool {
	_, _, err := net.ParseCIDR(inputCidr)
	return err == nil
}
