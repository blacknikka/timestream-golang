package timestream

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/timestreamquery"
)

type TimestreamQuery struct{}

func (TimestreamQuery) Query(tsQuery *timestreamquery.TimestreamQuery, query string) (*timestreamquery.QueryOutput, error) {
	queryInput := &timestreamquery.QueryInput{
		QueryString: aws.String(query),
	}

	queryOutput, err := tsQuery.Query(queryInput)
	if err != nil {
		return nil, err
	}

	return queryOutput, nil
}
