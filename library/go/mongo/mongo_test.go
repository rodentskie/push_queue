package mongo

import (
	"testing"
)

func TestInitDbConnection(t *testing.T) {
	testCases := []struct {
		uri      string
		db       string
		expected DatabaseInfo
	}{
		{
			uri:      "mongodb://localhost:27017",
			db:       "mydatabase",
			expected: DatabaseInfo{Uri: "mongodb://localhost:27017", Database: "mydatabase"},
		},
	}

	for _, tc := range testCases {
		actual := InitDbConnection(tc.uri, tc.db)

		if actual != tc.expected {
			t.Errorf("InitDbConnection(%q, %q) = %v; want %v", tc.uri, tc.db, actual, tc.expected)
		}
	}
}
