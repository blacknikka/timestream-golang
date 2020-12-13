package main

import (
	"fmt"
	"math/rand"
	"net"
	"net/http"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/timestreamwrite"
	"github.com/blacknikka/timestream-golang/timestream"

	"golang.org/x/net/http2"
)

func main() {
	rand.Seed(time.Now().UnixNano())

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
	tableName := "utilization"

	now := time.Now()
	baseTimeInMilliSeconds := now.UnixNano()
	baseTimeInMilliSeconds = int64(baseTimeInMilliSeconds / (1000 * 1000))

	// insert 50 records in every 5 second.
	t := time.NewTicker(5 * time.Second)
	defer t.Stop()
	for {
		select {
		case <-t.C:
			loopMax := 50

			// create 2ms granularly array
			insertData := struct {
				Timestamps []int64
				Values     []float64
			}{}

			// create 2ms granularly array
			for i := 0; i < loopMax; i++ {
				insertData.Timestamps = append(insertData.Timestamps, baseTimeInMilliSeconds)
				baseTimeInMilliSeconds += 100
			}

			// create random value (0 <= value < 1)
			for i := 0; i < loopMax; i++ {
				insertData.Values = append(insertData.Values, rand.Float64())
			}

			var records []*timestreamwrite.Record
			for i := 0; i < loopMax; i++ {
				r := []*timestreamwrite.Record{
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
						MeasureValue:     aws.String(strconv.FormatFloat(insertData.Values[i], 'f', -1, 64)),
						MeasureValueType: aws.String(timestreamwrite.MeasureValueTypeDouble),
						Time:             aws.String(strconv.FormatInt(insertData.Timestamps[i], 10)),
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
						MeasureValue:     aws.String(strconv.FormatFloat(insertData.Values[i]+5, 'f', -1, 64)),
						MeasureValueType: aws.String(timestreamwrite.MeasureValueTypeDouble),
						Time:             aws.String(strconv.FormatInt(insertData.Timestamps[i], 10)),
						TimeUnit:         aws.String(timestreamwrite.TimeUnitMilliseconds),
					},
				}
				records = append(records, r...)
			}
			writeRecordsInput := &timestreamwrite.WriteRecordsInput{
				DatabaseName: aws.String(databaseName),
				TableName:    aws.String(tableName),
				Records:      records,
			}

			insert := timestream.TimestreamInsert{}
			err = insert.Insert(writeSvc, writeRecordsInput)

			if err != nil {
				fmt.Println("Error:")
				fmt.Println(err)
				return
			} else {
				fmt.Println("Write records is successful")
			}
		}
	}
}
