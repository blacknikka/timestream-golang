### What's this
- This is an example for AWS Timestream by Go language.
  - Query and insert data sample.

### How to use
```bash
$ docker-compose up -d
$ docker-compose exec app ash
# In the container, run below
$ go run main.go
```

### Result
```
$ go run main.go
Submitting a query:
{
  QueryString: "SELECT * FROM sampleDB.IoT limit 5"
}
{
  ColumnInfo: [
    {
      Name: "fleet",
      Type: {
        ScalarType: "VARCHAR"
      }
    },
    {
      Name: "truck_id",
      Type: {
        ScalarType: "VARCHAR"
      }
    },
    {
      Name: "fuel_capacity",
      Type: {
        ScalarType: "VARCHAR"
      }
    },
    {
      Name: "model",
      Type: {
        ScalarType: "VARCHAR"
      }
    },
    {
      Name: "load_capacity",
      Type: {
        ScalarType: "VARCHAR"
      }
    },
    {
      Name: "make",
      Type: {
        ScalarType: "VARCHAR"
      }
    },
    {
      Name: "measure_value::double",
      Type: {
        ScalarType: "DOUBLE"
      }
    },
    {
      Name: "measure_value::varchar",
      Type: {
        ScalarType: "VARCHAR"
      }
    },
    {
      Name: "measure_name",
      Type: {
        ScalarType: "VARCHAR"
      }
    },
    {
      Name: "time",
      Type: {
        ScalarType: "TIMESTAMP"
      }
    }
  ],
  QueryId: "AEBQEAMY4RJ7YO6QPLLVJWZB3XVD6YNMELX22XF73MNMID4WG7VDDFQIRBRHAMY",
  Rows: [
    {
      Data: [
        {
          ScalarValue: "Alpha"
        },
        {
          ScalarValue: "368680024"
        },
        {
          ScalarValue: "100"
        },
        {
          ScalarValue: "359"
        },
        {
          ScalarValue: "1000"
        },
        {
          ScalarValue: "Peterbilt"
        },
        {
          ScalarValue: "309.0"
        },
        {
          NullValue: true
        },
        {
          ScalarValue: "load"
        },
        {
          ScalarValue: "2020-12-12 02:13:08.917000000"
        }
      ]
    },
    {
      Data: [
        {
          ScalarValue: "Alpha"
        },
        {
          ScalarValue: "368680024"
        },
        {
          ScalarValue: "100"
        },
        {
          ScalarValue: "359"
        },
        {
          ScalarValue: "1000"
        },
        {
          ScalarValue: "Peterbilt"
        },
        {
          ScalarValue: "134.0"
        },
        {
          NullValue: true
        },
        {
          ScalarValue: "load"
        },
        {
          ScalarValue: "2020-12-12 02:24:02.498000000"
        }
      ]
    },
    {
      Data: [
        {
          ScalarValue: "Alpha"
        },
        {
          ScalarValue: "368680024"
        },
        {
          ScalarValue: "100"
        },
        {
          ScalarValue: "359"
        },
        {
          ScalarValue: "1000"
        },
        {
          ScalarValue: "Peterbilt"
        },
        {
          ScalarValue: "881.0"
        },
        {
          NullValue: true
        },
        {
          ScalarValue: "load"
        },
        {
          ScalarValue: "2020-12-12 02:37:07.044000000"
        }
      ]
    },
    {
      Data: [
        {
          ScalarValue: "Alpha"
        },
        {
          ScalarValue: "368680024"
        },
        {
          ScalarValue: "100"
        },
        {
          ScalarValue: "359"
        },
        {
          ScalarValue: "1000"
        },
        {
          ScalarValue: "Peterbilt"
        },
        {
          ScalarValue: "358.0"
        },
        {
          NullValue: true
        },
        {
          ScalarValue: "load"
        },
        {
          ScalarValue: "2020-12-12 03:12:14.639000000"
        }
      ]
    },
    {
      Data: [
        {
          ScalarValue: "Alpha"
        },
        {
          ScalarValue: "368680024"
        },
        {
          ScalarValue: "100"
        },
        {
          ScalarValue: "359"
        },
        {
          ScalarValue: "1000"
        },
        {
          ScalarValue: "Peterbilt"
        },
        {
          ScalarValue: "359.0"
        },
        {
          NullValue: true
        },
        {
          ScalarValue: "load"
        },
        {
          ScalarValue: "2020-12-12 03:23:55.952000000"
        }
      ]
    }
  ]
}
```
