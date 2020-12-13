package main

import (
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/timestreamquery"

	"golang.org/x/net/http2"
)

func main() {
	tr := &http.Transport{
		ResponseHeaderTimeout: 20 * time.Second,
		// Using DefaultTransport values for other parameters: https://golang.org/pkg/net/http/#RoundTripper
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			KeepAlive: 30 * time.Second,
			DualStack: true,
			Timeout:   30 * time.Second,
		}).DialContext,
		MaxIdleConns:          100,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	}

	// So client makes HTTP/2 requests
	http2.ConfigureTransport(tr)

	sess, err := session.NewSession(
		&aws.Config{
			Region:     aws.String("us-east-1"),
			MaxRetries: aws.Int(2),
			HTTPClient: &http.Client{Transport: tr},
		},
	)
	if err != nil {
		panic(err)
	}

	// read service
	querySvc := timestreamquery.New(sess)
	query := `SELECT * FROM sampleDB.IoT limit 5`
	queryInput := &timestreamquery.QueryInput{
		QueryString: aws.String(query),
	}

	fmt.Println("Submitting a query:")
	fmt.Println(queryInput)
	// submit the query
	queryOutput, err := querySvc.Query(queryInput)

	if err != nil {
		fmt.Println("Error:")
		fmt.Println(err)
	}

	fmt.Println(queryOutput)
}
