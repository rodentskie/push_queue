package jwt

// import (
// 	testhelper "library/go/testHelper"
// 	"library/go/typings"
// 	"reflect"
// 	"testing"
// 	"time"
// )

// func TestGenerateToken(t *testing.T) {
// 	loginInfo, err := testhelper.HelperGenerateLoginInfo()

// 	if err != nil {
// 		t.Error("FAIL: Helper generate login info.")

// 	}

// 	token, err := GenerateToken(loginInfo)

// 	if err != nil {
// 		t.Error("FAIL: GenerateToken func.")

// 	}

// 	if reflect.TypeOf(token).Kind() != reflect.String {
// 		t.Errorf("FAIL: Unexpected token type. Expected: %s, Got: %s", reflect.String, reflect.TypeOf(token))
// 	}
// }

// func TestGenerateRefreshToken(t *testing.T) {
// 	loginInfo, err := testhelper.HelperGenerateLoginInfo()

// 	if err != nil {
// 		t.Error("FAIL: Helper generate login info.")

// 	}

// 	token, err := GenerateRefreshToken(loginInfo)

// 	if err != nil {
// 		t.Error("FAIL: GenerateRefreshToken func.")

// 	}

// 	if reflect.TypeOf(token).Kind() != reflect.String {
// 		t.Errorf("FAIL: Unexpected token type. Expected: %s, Got: %s", reflect.String, reflect.TypeOf(token))
// 	}
// }

// func TestParseToken(t *testing.T) {
// 	loginInfo, err := testhelper.HelperGenerateLoginInfo()
// 	if err != nil {
// 		t.Error("FAIL: Helper generate login info.")

// 	}

// 	token, err := GenerateToken(loginInfo)

// 	if err != nil {
// 		t.Error("FAIL: GenerateToken func.")

// 	}

// 	if reflect.TypeOf(token).Kind() != reflect.String {
// 		t.Errorf("FAIL: Unexpected token type. Expected: %s, Got: %s", reflect.String, reflect.TypeOf(token))
// 	}

// 	info, err := ParseToken(token)

// 	if err != nil {
// 		t.Error("FAIL: ParseToken func.")

// 	}

// 	expected := &typings.JwtSessionParse{
// 		AccountId:  loginInfo.AccountId,
// 		EmployeeId: loginInfo.EmployeeId,
// 		Username:   loginInfo.Username,
// 		Iat:        time.Now().UnixMilli(),
// 	}

// 	if !reflect.DeepEqual(info, expected) {
// 		t.Errorf("Unexpected result. Expected: %+v, Got: %+v", expected, info)
// 	}

// }
