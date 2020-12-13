package timestream

import (
	"github.com/aws/aws-sdk-go/service/timestreamwrite"
)

type TimestreamInsert struct{}

func (TimestreamInsert) Insert(writeSvc *timestreamwrite.TimestreamWrite, input *timestreamwrite.WriteRecordsInput) error {
	_, err := writeSvc.WriteRecords(input)
	return err
}
