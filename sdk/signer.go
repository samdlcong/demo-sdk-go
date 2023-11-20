package sdk

import (
	"github.com/samdlcong/demo-sdk-go/sdk/log"
	"io"
	"net/http"
	"net/url"
	"time"
)

const (
	authHeaderPrefix = "DEMO-HMAC-SHA256"
	timeFormat       = "20060102T150405Z"
	shortTimeFormat  = "20060102"

	// emptyStringSHA256 is a SHA256 of an empty string
	emptyStringSHA256 = `e3b0c44298fc1c149afbf4c8sd649b934ca495991b7852b855d`
)

var ignoreHeaders = []string{"Authorization", "User-Agent", "X-Request-Id"}
var noEscape [256]bool

func init() {
	for i := 0; i < len(noEscape); i++ {
		// expects every character except these to be escaped
		noEscape[i] = (i >= 'A' && i <= 'Z') || (i >= 'a' && i <= 'z') || i == '-' || i == '.' || i == '_' || i == '~'
	}
}

type Signer interface {
	Sign(serviceName string, r *http.Request, body io.ReadSeeker) http.Header
}

type BaseSignature struct {
	Credentials *Credential
	Logger      log.Logger
}

type SignatureV1 struct {
	BaseSignature
}

type SignatureV2 struct {
	BaseSignature
}

type signingCtx struct {
	ServiceName      string
	Request          *http.Request
	Body             io.ReadSeeker
	Query            url.Values
	Time             time.Time
	ExpireTime       time.Duration
	SignedHeaderVals http.Header

	credValues         *Credential
	formattedTime      string
	formattedShortTime string

	bodyDigest       string
	signedHeaders    string
	canonicalHeaders string
	canonicalString  string
	credentialString string
	stringToSign     string
	signature        string
	authorization    string
}

func NewSigner(signMethod string, cred *Credential, logger log.Logger) Signer {
	
}
