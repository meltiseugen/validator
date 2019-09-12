package validator_test

import (
	"fmt"
	"github.com/meltiseugen/validator"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"testing"
	"time"
)

func TestValidator_ValidateAndInit2(t *testing.T) {
	type MyStruct struct {
		A string `validate:"required" datakey:"a"`
		B int    `validate:"int" datakey:"b"`
	}
	m := map[string]string{
		"a": "123",
		"b": "1",
	}

	s := MyStruct{}
	v := validator.New()
	err := v.ValidateAndInit(m, &s)
	if err != nil {
		t.Error()
	}
	if s.A != "123" || s.B != 1 {
		t.Error()
	}
}

func TestValidator_ValidateAndInit3(t *testing.T) {
	type InnerStruct struct {
		C int `datakey:"c" validate:"required,nothing,int"`
		D string `datakey:"d" validate:"nothing"`
	}
	type MyStruct struct {
		A time.Time `validate:"required,nothing,time" datakey:"a"`
		B int `validate:"required,int" datakey:"b"`
		IS InnerStruct
	}

	s := MyStruct{}
	m := map[string]string{
		"a": "2019-08-21T09:00:00Z",
		"b": "1",
		"c": "1",
		"d": "qwerty",
	}

	v := validator.New()
	_ = v.RegisterRule("nothing", func(mapKey string, m map[string]string, params ...string) error {
		return nil
	})
	err := v.ValidateAndInit(m, &s)
	if err != nil {
		t.Error()
	}

	time_, _ := time.Parse(time.RFC3339, "2019-08-21T09:00:00Z")
	if !s.A.Equal(time_) || s.B != 1 || s.IS.C != 1 || s.IS.D != "qwerty" {
		t.Error()
	}
}

func TestValidator_ValidateAndInit5(t *testing.T) {
	type StatsData struct {
		MetricValue uint `bson:"metric_value" json:"metric_value"`
		Entered     uint `bson:"entered" json:"entered"`
		Exited      uint `bson:"exited" json:"exited"`
		MaxCount    uint `bson:"max_count" json:"max_count"`
	}

	type Stats struct {
		ID         primitive.ObjectID `bson:"-" json:"id" validate:"required,id" datakey:"a"`
		Location   string             `bson:"-" json:"location" validate:"required" datakey:"b"`
		Start      time.Time          `bson:"-" json:"start" validate:"required,time" datakey:"c"`
		End        time.Time          `bson:"-" json:"end" validate:"required,time" datakey:"d"`
		MetricType string             `bson:"metric_type" json:"metric_type" validate:"required,interval" datakey:"e"`
		Data       []StatsData        `bson:"-" json:"data"`
	}

	v := validator.New()
	_ = v.RegisterRule("id", func(mapKey string, m map[string]string, params ...string) error {
		if mapValue, ok := m[mapKey]; ok {
			_, err := primitive.ObjectIDFromHex(mapValue)
			if err != nil {
				return err
			}
		}
		return nil
	})

	_ = v.RegisterRule("interval", func(mapKey string, m map[string]string, params ...string) error {
		if mapValue, ok := m[mapKey]; ok {
			if mapValue != "minute" && mapValue != "hour" && mapValue != "day" &&
				mapValue != "month" && mapValue != "year" {
					return fmt.Errorf("map key '%s' with value '%s' does not match to " +
						"any of (minute hour day month year)", mapKey, mapValue)
			}
		}
		return nil
	})

	_ = v.RegisterConverter("primitive.ObjectID", func(value string, params ...string) (i interface{}, e error) {
		oid, err := primitive.ObjectIDFromHex(value)
		if err != nil {
			return nil, err
		}
		return oid, nil
	})

	data := map[string]string{
		"a": "000000000000000000000001",
		"b": "42.23,54.56",
		"c": "2019-08-21T09:00:00Z",
		"d": "2019-08-24T09:00:00Z",
		"e": "day",
	}

	s := Stats{}
	err := v.ValidateAndInit(data, &s)
	if err != nil {
		t.Error()
	}
	if s.ID.String() != "ObjectID(\"000000000000000000000001\")" ||
		s.Start.String() != "2019-08-21 09:00:00 +0000 UTC" ||
		s.End.String() != "2019-08-24 09:00:00 +0000 UTC" ||
		s.Location != "42.23,54.56" ||
		s.MetricType != "day" {

		t.Error()
	}
}