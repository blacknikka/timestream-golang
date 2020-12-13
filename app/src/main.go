package main

import (
	"fmt"
	"net"
	"net/http"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/timestreamquery"
	"github.com/aws/aws-sdk-go/service/timestreamwrite"

	"github.com/blacknikka/timestream-golang/timestream"

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

	// write service
	writeSvc := timestreamwrite.New(sess)

	databaseName := "sampleDB"
	tableName := "IoT"

	now := time.Now()
	currentTimeInMilliSeconds := now.UnixNano()
	currentTimeInMilliSeconds = int64(currentTimeInMilliSeconds / (1000 * 1000))
	fmt.Println(currentTimeInMilliSeconds)

	writeRecordsInput := &timestreamwrite.WriteRecordsInput{
		DatabaseName: aws.String(databaseName),
		TableName:    aws.String(tableName),
		Records: []*timestreamwrite.Record{
			&timestreamwrite.Record{
				Dimensions: []*timestreamwrite.Dimension{
					&timestreamwrite.Dimension{
						Name:  aws.String("region"),
						Value: aws.String("us-east-1"),
					},
					&timestreamwrite.Dimension{
						Name:  aws.String("az"),
						Value: aws.String("az1"),
					},
					&timestreamwrite.Dimension{
						Name:  aws.String("hostname"),
						Value: aws.String("host1"),
					},
				},
				MeasureName:      aws.String("cpu_utilization"),
				MeasureValue:     aws.String("13.5"),
				MeasureValueType: aws.String("DOUBLE"),
				Time:             aws.String(strconv.FormatInt(currentTimeInMilliSeconds, 10)),
				TimeUnit:         aws.String(timestreamwrite.TimeUnitMilliseconds),
			},
			&timestreamwrite.Record{
				Dimensions: []*timestreamwrite.Dimension{
					&timestreamwrite.Dimension{
						Name:  aws.String("region"),
						Value: aws.String("us-east-1"),
					},
					&timestreamwrite.Dimension{
						Name:  aws.String("az"),
						Value: aws.String("az1"),
					},
					&timestreamwrite.Dimension{
						Name:  aws.String("hostname"),
						Value: aws.String("host1"),
					},
				},
				MeasureName:      aws.String("memory_utilization"),
				MeasureValue:     aws.String("40"),
				MeasureValueType: aws.String("DOUBLE"),
				Time:             aws.String(strconv.FormatInt(currentTimeInMilliSeconds, 10)),
				TimeUnit:         aws.String(timestreamwrite.TimeUnitMilliseconds),
			},
		},
	}

	_, err = writeSvc.WriteRecords(writeRecordsInput)

	if err != nil {
		fmt.Println("Error:")
		fmt.Println(err)
		return
	} else {
		fmt.Println("Write records is successful")
	}

	// read service
	querySvc := timestreamquery.New(sess)
	query := `SELECT * FROM sampleDB.IoT ORDER BY time DESC LIMIT 3`

	fmt.Println("Submitting a query:")
	tsQuery := timestream.TimestreamQuery{}
	queryOutput, err := tsQuery.Query(querySvc, query)

	if err != nil {
		fmt.Println("Error:")
		fmt.Println(err)
	}

	fmt.Println(queryOutput)

}
