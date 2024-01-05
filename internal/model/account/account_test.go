package account

import (
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"jlowell000.github.io/budgeting/internal/model/bookentry"
)

const (
	TEST_ID   = "16cfd708-db6d-42fd-8ad1-55316690520c"
	TEST_NAME = "test name"
	TEST_FLAG = false
	TEST_TIME = "2006-01-23T15:04:05Z"
)

var TEST_AMOUNT = decimal.NewFromFloat(100.99)

func TestAccountToJSON(t *testing.T) {
	id := getUUID(TEST_ID)
	timestamp := getTime(TEST_TIME)

	expected := getTestJson(TEST_ID, TEST_NAME, TEST_FLAG, TEST_AMOUNT, TEST_TIME)
	periodicFlow := Account{
		Id:         id,
		Name:       TEST_NAME,
		Excludable: TEST_FLAG,
		Book: []*bookentry.BookEntry{
			{
				Id:        id,
				Amount:    TEST_AMOUNT,
				Timestamp: timestamp,
			},
		},
		UpdatedTimestamp: timestamp,
	}
	actual := string(periodicFlow.ToJSON())

	assert.Equal(t, expected, actual)
}

func TestAccountFromJSON_data_there(t *testing.T) {
	id := getUUID(TEST_ID)
	timestamp := getTime(TEST_TIME)

	expected := Account{
		Id:         id,
		Name:       TEST_NAME,
		Excludable: TEST_FLAG,
		Book: []*bookentry.BookEntry{
			{
				Id:        id,
				Amount:    TEST_AMOUNT,
				Timestamp: timestamp,
			},
		},
		UpdatedTimestamp: timestamp,
	}
	actual := FromJSON([]byte(
		getTestJson(TEST_ID, TEST_NAME, TEST_FLAG, TEST_AMOUNT, TEST_TIME),
	))

	assert.Equal(t, expected, actual)
}

func TestAccountFromJSON_partial_data_there(t *testing.T) {
	id := getUUID(TEST_ID)
	timestamp := getTime(TEST_TIME)

	expected := Account{
		Id: id,
		Book: []*bookentry.BookEntry{
			{
				Id:        id,
				Amount:    TEST_AMOUNT,
				Timestamp: timestamp,
			},
		},
		UpdatedTimestamp: timestamp,
	}
	actual := FromJSON([]byte(
		getTestJson(TEST_ID, "", TEST_FLAG, TEST_AMOUNT, TEST_TIME),
	))

	assert.Equal(t, expected, actual)
}

func TestAccountFromJSON_no_data(t *testing.T) {
	expected := Account{}
	actual := FromJSON([]byte(""))

	assert.Equal(t, expected, actual)
}

func Test_GetLatestBookEntry(t *testing.T) {
	expected := &bookentry.BookEntry{
		Id:        uuid.New(),
		Amount:    decimal.NewFromFloat(666.6),
		Timestamp: time.Now(),
	}
	account := Account{
		Id: uuid.New(),
		Book: []*bookentry.BookEntry{
			{
				Id:        uuid.New(),
				Amount:    TEST_AMOUNT,
				Timestamp: time.UnixMilli(1),
			},
			{
				Id:        uuid.New(),
				Amount:    TEST_AMOUNT,
				Timestamp: time.UnixMilli(1),
			},
			expected,
			{
				Id:        uuid.New(),
				Amount:    TEST_AMOUNT,
				Timestamp: time.UnixMilli(1),
			},
			{
				Id:        uuid.New(),
				Amount:    TEST_AMOUNT,
				Timestamp: time.UnixMilli(1),
			},
		},
		UpdatedTimestamp: time.Now(),
	}
	actual := account.GetLatestBookEntry()

	assert.Equal(t, expected, actual)
}

func Test_GetEarliestBookEntry(t *testing.T) {
	expected := &bookentry.BookEntry{
		Id:        uuid.New(),
		Timestamp: time.UnixMilli(100000),
	}
	account := Account{
		Id: uuid.New(),
		Book: []*bookentry.BookEntry{
			{
				Id:        uuid.New(),
				Timestamp: time.UnixMilli(500000),
			},
			{
				Id:        uuid.New(),
				Timestamp: time.UnixMilli(500000),
			},
			expected,
			{
				Id:        uuid.New(),
				Timestamp: time.UnixMilli(500000),
			},
			{
				Id:        uuid.New(),
				Timestamp: time.UnixMilli(500000),
			},
		},
		UpdatedTimestamp: time.Now(),
	}
	actual := account.GetEarliestBookEntry()

	assert.Equal(t, expected, actual)
}

func Test_GetBookEndEntries(t *testing.T) {
	expected1 := &bookentry.BookEntry{
		Id:        uuid.New(),
		Timestamp: time.UnixMilli(100000),
	}

	expected2 := &bookentry.BookEntry{
		Id:        uuid.New(),
		Timestamp: time.UnixMilli(700000),
	}
	account := Account{
		Id: uuid.New(),
		Book: []*bookentry.BookEntry{
			{
				Id:        uuid.New(),
				Timestamp: time.UnixMilli(500000),
			},
			{
				Id:        uuid.New(),
				Timestamp: time.UnixMilli(500000),
			},
			expected2,
			{
				Id:        uuid.New(),
				Timestamp: time.UnixMilli(500000),
			},
			expected1,
			{
				Id:        uuid.New(),
				Timestamp: time.UnixMilli(500000),
			},
		},
		UpdatedTimestamp: time.Now(),
	}
	actual1, actual2 := account.GetBookEndEntries()

	assert.Equal(t, expected1, actual1)
	assert.Equal(t, expected2, actual2)

}

func TestRateOfChange(t *testing.T) {
	testSize := 100
	testSizeSlice := make([]bool, testSize)
	a := bookentry.BookEntry{
		Amount:    decimal.NewFromInt(0),
		Timestamp: time.UnixMilli(0),
	}

	for i := range testSizeSlice {
		for j := range testSizeSlice {
			di := decimal.NewFromInt(int64(i + 1))
			dj := int64(j + 1)
			expected := di.Div(decimal.NewFromInt(dj))
			account := Account{
				Id: uuid.New(),
				Book: []*bookentry.BookEntry{
					&a,
					{
						Amount:    di,
						Timestamp: time.UnixMilli(dj),
					},
				},
				UpdatedTimestamp: time.Now(),
			}

			assert.True(t, expected.Equals(account.RateOfChange()))
		}
	}

	for i := range testSizeSlice {
		for j := range testSizeSlice {
			di := decimal.NewFromInt(int64(i + 1))
			dj := int64(testSize - j + 1)
			expected := di.Div(decimal.NewFromInt(dj))
			account := Account{
				Id: uuid.New(),
				Book: []*bookentry.BookEntry{
					&a,
					{
						Amount:    di,
						Timestamp: time.UnixMilli(dj),
					},
				},
				UpdatedTimestamp: time.Now(),
			}

			assert.True(t, expected.Equals(account.RateOfChange()))
		}
	}

	expected := decimal.NewFromInt(0)
	account := Account{
		Id:               uuid.New(),
		Book:             []*bookentry.BookEntry{&a},
		UpdatedTimestamp: time.Now(),
	}
	assert.True(t, expected.Equals(account.RateOfChange()))

	account = Account{
		Id:               uuid.New(),
		Book:             []*bookentry.BookEntry{},
		UpdatedTimestamp: time.Now(),
	}
	assert.True(t, expected.Equals(account.RateOfChange()))

	account = Account{
		Id:               uuid.New(),
		UpdatedTimestamp: time.Now(),
	}
	assert.True(t, expected.Equals(account.RateOfChange()))
}

func Test_Sum_accounts(t *testing.T) {
	amount := decimal.NewFromFloat(666.66)
	testSize := 100
	testSizeSlice := make([]int, testSize)
	var accounts []Account
	for i := range testSizeSlice {
		testSizeSlice[i] = i
		accounts = append(accounts, createAccount(amount, false))
	}
	expected := amount.Mul(decimal.NewFromInt(int64(testSize)))
	actual := Sum(accounts)

	assert.Equal(t, expected, actual)
}

func Test_SumExclusion_accounts(t *testing.T) {
	amount := decimal.NewFromFloat(666.66)
	testSize := 100
	testSizeSlice := make([]int, testSize)
	var accounts []Account
	for i := range testSizeSlice {
		testSizeSlice[i] = i
		third := i%4 == 0
		accounts = append(accounts, createAccount(amount, !third))
	}
	expected := amount.Mul(decimal.NewFromFloat(float64(testSize) / 4))
	actual := SumExclusion(accounts)

	assert.Equal(t, expected, actual)
}

func createAccount(amount decimal.Decimal, excludable bool) Account {
	return Account{
		Id:         uuid.New(),
		Excludable: excludable,
		Book: []*bookentry.BookEntry{
			{
				Id:        uuid.New(),
				Amount:    amount,
				Timestamp: time.Now(),
			},
		},
		UpdatedTimestamp: time.Now(),
	}
}

func getPFParsedValues() (uuid.UUID, time.Time) {
	id, err1 := uuid.Parse(TEST_ID)
	if err1 != nil {
		log.Fatal(err1)
	}
	timestamp, err2 := time.Parse(time.RFC3339, TEST_TIME)
	if err2 != nil {
		log.Fatal(err2)
	}
	return id, timestamp
}

func getTestJson(
	id string,
	name string,
	excludable bool,
	amount decimal.Decimal,
	time string,
) string {
	return "{\"id\":\"" + id + "\"," +
		"\"name\":\"" + name + "\"," +
		"\"excludable\":" + fmt.Sprintf("%t", excludable) + "," +
		"\"book\":[" +
		"{\"id\":\"" + id +
		"\",\"amount\":\"" + amount.String() + "\"," +
		"\"timestamp\":\"" + time + "\"}" +
		"]," +
		"\"updated_timestamp\":\"" + time + "\"}"
}

func getUUID(s string) uuid.UUID {
	id, err1 := uuid.Parse(s)
	if err1 != nil {
		log.Fatal(err1)
	}
	return id
}

func getTime(s string) time.Time {
	timestamp, err2 := time.Parse(time.RFC3339, s)
	if err2 != nil {
		log.Fatal(err2)
	}
	return timestamp
}
