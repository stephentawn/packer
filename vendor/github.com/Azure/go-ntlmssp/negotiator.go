package ntlmssp

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

// GetDomain : parse domain name from based on slashes in the input
func GetDomain(user string) (string, string) {
	domain := ""

	if strings.Contains(user, "\\") {
		ucomponents := strings.SplitN(user, "\\", 2)
		domain = ucomponents[0]
		user = ucomponents[1]
	}
	return user, domain
}

//Negotiator is a http.Roundtripper decorator that automatically
//converts basic authentication to NTLM/Negotiate authentication when appropriate.
type Negotiator struct{ http.RoundTripper }

//RoundTrip sends the request to the server, handling any authentication
//re-sends as needed.
func (l Negotiator) RoundTrip(req *http.Request) (res *http.Response, err error) {
	// Use default round tripper if not provided
	rt := l.RoundTripper
	log.Printf("PACKER_NTLM_DEBUG 1")
	if rt == nil {
		log.Printf("PACKER_NTLM_DEBUG 2; using default http transpoort as roundtripper")
		rt = http.DefaultTransport
	}
	// If it is not basic auth, just round trip the request as usual
	reqauth := authheader(req.Header.Get("Authorization"))
	if !reqauth.IsBasic() {
		log.Printf("PACKER_NTLM_DEBUG 3")
		res, err := rt.RoundTrip(req)
		log.Printf(fmt.Sprintf("res is: %#v and err is: %#v ", res, err))
		return res, err
	}
	// Save request body
	body := bytes.Buffer{}
	log.Printf("PACKER_NTLM_DEBUG 4")
	if req.Body != nil {
		log.Printf("PACKER_NTLM_DEBUG 5")
		_, err = body.ReadFrom(req.Body)
		if err != nil {
			log.Printf("PACKER_NTLM_DEBUG 6")
			log.Printf(fmt.Sprintf("res is nil and err is: %#v ", err))
			return nil, err
		}
		log.Printf("PACKER_NTLM_DEBUG 7")

		req.Body.Close()
		log.Printf("PACKER_NTLM_DEBUG 8")
		req.Body = ioutil.NopCloser(bytes.NewReader(body.Bytes()))
		log.Printf("PACKER_NTLM_DEBUG 9")
	}
	// first try anonymous, in case the server still finds us
	// authenticated from previous traffic
	log.Printf("PACKER_NTLM_DEBUG 10")
	req.Header.Del("Authorization")
	res, err = rt.RoundTrip(req)
	log.Printf("PACKER_NTLM_DEBUG 11")
	if err != nil {
		log.Printf("PACKER_NTLM_DEBUG 12")
		log.Printf(fmt.Sprintf("res is: %#v and err is: %#v ", nil, err))
		return nil, err
	}
	log.Printf("PACKER_NTLM_DEBUG 13")
	if res.StatusCode != http.StatusUnauthorized {
		log.Printf("PACKER_NTLM_DEBUG 14")
		log.Printf(fmt.Sprintf("res is: %#v and err is: %#v ", res, err))
		return res, err
	}

	resauth := authheader(res.Header.Get("Www-Authenticate"))
	log.Printf("PACKER_NTLM_DEBUG 15")
	if !resauth.IsNegotiate() && !resauth.IsNTLM() {
		log.Printf("PACKER_NTLM_DEBUG 16")
		// Unauthorized, Negotiate not requested, let's try with basic auth
		req.Header.Set("Authorization", string(reqauth))
		io.Copy(ioutil.Discard, res.Body)
		res.Body.Close()
		req.Body = ioutil.NopCloser(bytes.NewReader(body.Bytes()))

		res, err = rt.RoundTrip(req)
		log.Printf("PACKER_NTLM_DEBUG 17")
		if err != nil {
			log.Printf("PACKER_NTLM_DEBUG 18")
			log.Printf(fmt.Sprintf("res is: %#v and err is: %#v ", nil, err))
			return nil, err
		}
		log.Printf("PACKER_NTLM_DEBUG 19")
		if res.StatusCode != http.StatusUnauthorized {
			log.Printf("PACKER_NTLM_DEBUG 20")
			log.Printf(fmt.Sprintf("res is: %#v and err is: %#v ", res, err))
			return res, err
		}
		resauth = authheader(res.Header.Get("Www-Authenticate"))
	}
	log.Printf("PACKER_NTLM_DEBUG 21")
	if resauth.IsNegotiate() || resauth.IsNTLM() {
		log.Printf("PACKER_NTLM_DEBUG 22")
		// 401 with request:Basic and response:Negotiate
		io.Copy(ioutil.Discard, res.Body)
		res.Body.Close()

		// recycle credentials
		u, p, err := reqauth.GetBasicCreds()
		log.Printf("PACKER_NTLM_DEBUG 23")
		if err != nil {
			log.Printf("PACKER_NTLM_DEBUG 24")
			log.Printf(fmt.Sprintf("res is: %#v and err is: %#v ", nil, err))
			return nil, err
		}
		log.Printf("PACKER_NTLM_DEBUG 25")
		// get domain from username
		domain := ""
		u, domain = GetDomain(u)

		// send negotiate
		negotiateMessage, err := NewNegotiateMessage(domain, "")
		log.Printf("PACKER_NTLM_DEBUG 26")
		if err != nil {
			log.Printf("PACKER_NTLM_DEBUG 27")
			log.Printf(fmt.Sprintf("res is: %#v and err is: %#v ", nil, err))
			return nil, err
		}
		log.Printf("PACKER_NTLM_DEBUG 28")
		if resauth.IsNTLM() {
			req.Header.Set("Authorization", "NTLM "+base64.StdEncoding.EncodeToString(negotiateMessage))
		} else {
			req.Header.Set("Authorization", "Negotiate "+base64.StdEncoding.EncodeToString(negotiateMessage))
		}

		req.Body = ioutil.NopCloser(bytes.NewReader(body.Bytes()))

		res, err = rt.RoundTrip(req)
		log.Printf("PACKER_NTLM_DEBUG 29")
		if err != nil {
			log.Printf("PACKER_NTLM_DEBUG 30")
			log.Printf(fmt.Sprintf("res is: %#v and err is: %#v ", nil, err))
			return nil, err
		}

		// receive challenge?
		resauth = authheader(res.Header.Get("Www-Authenticate"))
		challengeMessage, err := resauth.GetData()
		log.Printf("PACKER_NTLM_DEBUG 31")
		if err != nil {
			log.Printf("PACKER_NTLM_DEBUG 32")
			log.Printf(fmt.Sprintf("res is: %#v and err is: %#v ", nil, err))
			return nil, err
		}
		if !(resauth.IsNegotiate() || resauth.IsNTLM()) || len(challengeMessage) == 0 {
			// Negotiation failed, let client deal with response
			log.Printf("PACKER_NTLM_DEBUG 33")
			log.Printf(fmt.Sprintf("res is: %#v and err is: %#v ", res, err))
			return res, nil
		}
		io.Copy(ioutil.Discard, res.Body)
		res.Body.Close()

		// send authenticate
		authenticateMessage, err := ProcessChallenge(challengeMessage, u, p)
		if err != nil {
			log.Printf("PACKER_NTLM_DEBUG 34")
			log.Printf(fmt.Sprintf("res is: %#v and err is: %#v ", nil, err))
			return nil, err
		}
		if resauth.IsNTLM() {
			req.Header.Set("Authorization", "NTLM "+base64.StdEncoding.EncodeToString(authenticateMessage))
		} else {
			req.Header.Set("Authorization", "Negotiate "+base64.StdEncoding.EncodeToString(authenticateMessage))
		}

		req.Body = ioutil.NopCloser(bytes.NewReader(body.Bytes()))

		res, err = rt.RoundTrip(req)
		log.Printf("PACKER_NTLM_DEBUG 35")
	}
	log.Printf("PACKER_NTLM_DEBUG 36")
	log.Printf(fmt.Sprintf("res is: %#v and err is: %#v ", res, err))
	return res, err
}
