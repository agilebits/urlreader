package urlreader

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

const (
	prefixHTTPS = "https://"
	prefixS3    = "s3://"
	prefixFile  = "file://"
)

// Open https, s3, or file URLs. If the url value does not begin with a scheme then function will try to open a local file.
// S3 is accessed using environmnent credentials or EC2 role if available.
func Open(url string) (io.ReadCloser, error) {
	if strings.HasPrefix(url, prefixHTTPS) {
		return openHTTPS(url)
	}

	if strings.HasPrefix(url, prefixS3) {
		return openS3(url)
	}

	if strings.HasPrefix(url, prefixFile) {
		return openFile(url)
	}

	return openFile(url)
}

func openHTTPS(url string) (io.ReadCloser, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	return resp.Body, nil
}

func openS3(url string) (io.ReadCloser, error) {
	if strings.HasPrefix(url, prefixS3) {
		url = url[len(prefixS3):]
	}

	parts := strings.SplitN(url, "/", 2)
	if len(parts) < 2 {
		return nil, fmt.Errorf("invalid S3 URL %q", url)
	}

	session := session.New(getAWSConfig("", nil))
	svc := s3.New(session)
	data, err := svc.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(parts[0]),
		Key:    aws.String(parts[1]),
	})

	if err != nil {
		return nil, err
	}

	return data.Body, nil
}

func openFile(path string) (io.ReadCloser, error) {
	if strings.HasPrefix(path, prefixFile) {
		path = path[len(prefixFile):]
	}

	return os.Open(path)
}
