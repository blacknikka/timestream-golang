package main

import (
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/timestreamwrite"

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
	writeSvc := timestreamwrite.New(sess)

	// 事前に作ってあるdatabase名を取得する
	// Describe database.
	describeDatabaseInput := &timestreamwrite.DescribeDatabaseInput{
		DatabaseName: aws.String("mysample"),
	}
	describeDatabaseOutput, err := writeSvc.DescribeDatabase(describeDatabaseInput)

	if err != nil {
		fmt.Println("Error:")
		fmt.Println(err)
	} else {
		fmt.Println("Describe database is successful, below is the output:")
		fmt.Println(describeDatabaseOutput)
	}
}
