package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/NBR41/go-testgoa/internal/model"
	"github.com/NBR41/go-testgoa/internal/security"
	jwt "github.com/dgrijalva/jwt-go"
)

const domain = "http://localhost:8089"

func callWs(req *http.Request) (int, string, string, error) {
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return 0, "", "", err
	}
	defer resp.Body.Close()
	ret, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, "", "", err
	}
	return resp.StatusCode, resp.Header.Get("Location"), string(ret), nil
}

func getAuthTokenValues(raw string) (string, error) {
	token, err := jwt.ParseWithClaims(raw, jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		return security.JWTAuthKey, nil
	})
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("user_id:%d is_admin:%t", int(token.Claims.(jwt.MapClaims)["user_id"].(float64)), token.Claims.(jwt.MapClaims)["is_admin"]), nil
}

func test(method, token, path string, body io.Reader, expCode int, expBody, expLocation string) error {
	req, err := http.NewRequest(method, domain+path, body)
	if err != nil {
		return err
	}
	if token != "" {
		req.Header.Set("Authorization", "Bearer "+token)
	}

	code, location, resp, err := callWs(req)
	if err != nil {
		return err
	}
	if code != expCode {
		return fmt.Errorf("unexpected code, exp %d got %d", expCode, code)
	}

	switch path {
	case "/token/access_token":
		if code == http.StatusOK {
			token := struct {
				Token string
			}{}
			if err = json.Unmarshal([]byte(resp), &token); err != nil {
				return err
			}
			v, err := getAuthTokenValues(token.Token)
			if err != nil {
				return err
			}
			if v != expBody {
				return fmt.Errorf("unexpected acces token, exp\n%s\ngot\n%s", expBody, v)
			}
		}
	case "/token/auth":
		fallthrough
	case "/authenticate":
		parts := strings.Split(expBody, "$")
		if code == http.StatusOK {
			auths := struct {
				Access  string          `json:"access_token"`
				Refresh string          `json:"refresh_token"`
				User    json.RawMessage `json:"user"`
			}{}
			if err = json.Unmarshal([]byte(resp), &auths); err != nil {
				return err
			}
			var v string
			v, err = getAuthTokenValues(auths.Access)
			if err != nil {
				return err
			}
			if v != parts[0] {
				return fmt.Errorf("unexpected acces token, exp\n%s\ngot\n%s", parts[0], v)
			}
			v, err = getAuthTokenValues(auths.Refresh)
			if err != nil {
				return err
			}
			if v != parts[1] {
				return fmt.Errorf("unexpected refresh token, exp\n%s\ngot\n%s", parts[1], v)
			}

			if string(auths.User) != parts[2] {
				return fmt.Errorf("unexpected body, exp\n%s\ngot\n%s", parts[2], string(auths.User))
			}
		}
	default:
		if expCode == http.StatusUnprocessableEntity || expCode == http.StatusUnauthorized {
			if resp != "" || expBody != "" {
				detail := struct {
					Detail string
				}{}
				if err = json.Unmarshal([]byte(resp), &detail); err != nil {
					return err
				}
				if detail.Detail != expBody {
					return fmt.Errorf("unexpected error detail body, exp\n%s\ngot\n%s", expBody, detail.Detail)
				}
			}
		} else {
			if resp != expBody {
				return fmt.Errorf("unexpected body, exp\n%s\ngot\n%s", expBody, resp)
			}
		}
	}

	if location != expLocation {
		return fmt.Errorf("unexpected location, exp\n%s\ngot\n%s", expLocation, location)
	}
	return nil
}

func main() {
	jwtHelper := security.JWTHelper{}
	adminToken, err := jwtHelper.GetAccessToken(1, true)
	if err != nil {
		log.Fatal(err)
	}

	userToken, err := jwtHelper.GetAccessToken(5, false)
	if err != nil {
		log.Fatal(err)
	}

	refreshToken, err := jwtHelper.GetRefreshToken(5, false)
	if err != nil {
		log.Fatal(err)
	}

	validationToken, err := jwtHelper.GetValidationToken(5, "foo@bar.com")
	if err != nil {
		log.Fatal(err)
	}

	passwordToken, err := jwtHelper.GetPasswordToken(5, "foo@bar.com")
	if err != nil {
		log.Fatal(err)
	}

	tests := []struct {
		label       string
		method      string
		token       string
		path        string
		body        io.Reader
		expCode     int
		expBody     string
		expLocation string
	}{
		{
			label:       "createUser",
			method:      http.MethodPost,
			path:        "/users",
			body:        strings.NewReader(`{"email": "foo@bar.com","nickname":"foo","password":"footest"}`),
			expCode:     http.StatusCreated,
			expLocation: "/users/5",
		},
		{
			label:   "createUserDuplicateKey",
			method:  http.MethodPost,
			path:    "/users",
			body:    strings.NewReader(`{"email": "foo@bar.com","nickname":"foo","password":"footest"}`),
			expCode: http.StatusUnprocessableEntity,
			expBody: model.ErrDuplicateKey.Error(),
		},
		{
			label:   "createUserDuplicateEmail",
			method:  http.MethodPost,
			path:    "/users",
			body:    strings.NewReader(`{"email": "foo@bar.com","nickname":"baz","password":"footest"}`),
			expCode: http.StatusUnprocessableEntity,
			expBody: model.ErrDuplicateEmail.Error(),
		},
		{
			label:   "createUserDuplicateNickname",
			method:  http.MethodPost,
			path:    "/users",
			body:    strings.NewReader(`{"email": "baz@bar.com","nickname":"foo","password":"footest"}`),
			expCode: http.StatusUnprocessableEntity,
			expBody: model.ErrDuplicateNickname.Error(),
		},

		{
			label:   "updateUserAsUserUnauthorized",
			method:  http.MethodPut,
			token:   userToken,
			path:    "/users/4",
			body:    strings.NewReader(`{"nickname":"foo"}`),
			expCode: http.StatusUnauthorized,
			expBody: "unauthorized",
		},
		{
			label:   "updateUserDuplicateNickname",
			method:  http.MethodPut,
			token:   adminToken,
			path:    "/users/4",
			body:    strings.NewReader(`{"nickname":"foo"}`),
			expCode: http.StatusUnprocessableEntity,
			expBody: model.ErrDuplicateNickname.Error(),
		},
		{
			label:   "updateUserNotFound",
			method:  http.MethodPut,
			token:   adminToken,
			path:    "/users/999",
			body:    strings.NewReader(`{"nickname":"baz"}`),
			expCode: http.StatusNotFound,
		},
		{
			label:   "updateUserAsSameUser",
			method:  http.MethodPut,
			token:   userToken,
			path:    "/users/5",
			body:    strings.NewReader(`{"nickname":"baz"}`),
			expCode: http.StatusNoContent,
		},
		{
			label:   "updateUserAsAdmin",
			method:  http.MethodPut,
			token:   adminToken,
			path:    "/users/5",
			body:    strings.NewReader(`{"nickname":"foo"}`),
			expCode: http.StatusNoContent,
		},
		{
			label:   "showUserOKasUser",
			method:  http.MethodGet,
			token:   userToken,
			path:    "/users/5",
			expCode: http.StatusOK,
			expBody: `{"email":"foo@bar.com","href":"/users/5","is_admin":false,"is_validated":false,"nickname":"foo","user_id":5}` + "\n",
		},
		{
			label:   "showUserNotFound",
			method:  http.MethodGet,
			token:   adminToken,
			path:    "/users/999",
			expCode: http.StatusNotFound,
		},
		{
			label:   "showUserAsUserUnauthorized",
			method:  http.MethodGet,
			token:   userToken,
			path:    "/users/999",
			expCode: http.StatusUnauthorized,
			expBody: "unauthorized",
		},
		{
			label:   "listUserUnauthorized",
			method:  http.MethodGet,
			token:   userToken,
			path:    "/users?email=foo@bar.com",
			expCode: http.StatusUnauthorized,
			expBody: `unauthorized`,
		},
		{
			label:   "listUserByEmail",
			method:  http.MethodGet,
			token:   adminToken,
			path:    "/users?email=foo@bar.com",
			expCode: http.StatusOK,
			expBody: `[{"email":"foo@bar.com","href":"/users/5","is_admin":false,"is_validated":false,"nickname":"foo","user_id":5}]` + "\n",
		},
		{
			label:   "listUserByNickname",
			method:  http.MethodGet,
			token:   adminToken,
			path:    "/users?nickname=foo",
			expCode: http.StatusOK,
			expBody: `[{"email":"foo@bar.com","href":"/users/5","is_admin":false,"is_validated":false,"nickname":"foo","user_id":5}]` + "\n",
		},
		{
			label:   "listUserByEmailNickname",
			method:  http.MethodGet,
			token:   adminToken,
			path:    "/users?email=foo@bar.com&nickname=foo",
			expCode: http.StatusOK,
			expBody: `[{"email":"foo@bar.com","href":"/users/5","is_admin":false,"is_validated":false,"nickname":"foo","user_id":5}]` + "\n",
		},
		{
			label:   "listUserByEmailNickname",
			method:  http.MethodGet,
			token:   adminToken,
			path:    "/users?email=foo@baz.com&nickname=baz",
			expCode: http.StatusOK,
			expBody: `[]` + "\n",
		},
		//auth
		{
			label:   "authEmailNotFound",
			method:  http.MethodPost,
			path:    "/authenticate",
			body:    strings.NewReader(`{"login": "foo@baz.com","password":"footest"}`),
			expCode: http.StatusUnprocessableEntity,
		},
		{
			label:   "authNicknameNotFound",
			method:  http.MethodPost,
			path:    "/authenticate",
			body:    strings.NewReader(`{"login": "baz","password":"footest"}`),
			expCode: http.StatusUnprocessableEntity,
		},
		{
			label:   "authEmail",
			method:  http.MethodPost,
			path:    "/authenticate",
			body:    strings.NewReader(`{"login": "foo@bar.com","password":"footest"}`),
			expCode: http.StatusOK,
			expBody: `user_id:5 is_admin:false$user_id:5 is_admin:false${"email":"foo@bar.com","href":"/users/5","is_admin":false,"is_validated":false,"nickname":"foo","user_id":5}`,
		},
		{
			label:   "authNickname",
			method:  http.MethodPost,
			path:    "/authenticate",
			body:    strings.NewReader(`{"login": "foo","password":"footest"}`),
			expCode: http.StatusOK,
			expBody: `user_id:5 is_admin:false$user_id:5 is_admin:false${"email":"foo@bar.com","href":"/users/5","is_admin":false,"is_validated":false,"nickname":"foo","user_id":5}`,
		},
		//validate
		{
			label:   "getValidationTokenUserUnauthorized",
			method:  http.MethodGet,
			token:   userToken,
			path:    "/validation/4",
			expCode: http.StatusUnauthorized,
			expBody: `unauthorized`,
		},
		{
			label:   "getValidationTokenUserNotFound",
			method:  http.MethodGet,
			token:   adminToken,
			path:    "/validation/999",
			expCode: http.StatusNotFound,
		},
		{
			label:   "getValidationAdmin",
			method:  http.MethodGet,
			token:   adminToken,
			path:    "/validation/5",
			expCode: http.StatusNoContent,
		},
		{
			label:   "getValidationUser",
			method:  http.MethodGet,
			token:   userToken,
			path:    "/validation/5",
			expCode: http.StatusNoContent,
		},
		{
			label:   "validateWrongToken",
			method:  http.MethodPost,
			path:    "/validation",
			body:    strings.NewReader(fmt.Sprintf(`{"token": "%s"}`, passwordToken)),
			expCode: http.StatusUnprocessableEntity,
		},
		{
			label:   "validateOK",
			method:  http.MethodPost,
			path:    "/validation",
			body:    strings.NewReader(fmt.Sprintf(`{"token": "%s"}`, validationToken)),
			expCode: http.StatusNoContent,
		},
		//password
		{
			label:   "getPasswordNotFound",
			method:  http.MethodGet,
			path:    "/password?email=bar@foo.com",
			expCode: http.StatusUnprocessableEntity,
		},
		{
			label:   "getPasswordOK",
			method:  http.MethodGet,
			path:    "/password?email=foo@bar.com",
			expCode: http.StatusNoContent,
		},
		{
			label:   "resetPasswordWrongToken",
			method:  http.MethodPost,
			path:    "/password",
			body:    strings.NewReader(fmt.Sprintf(`{"token": "%s", "password": "bazquux"}`, validationToken)),
			expCode: http.StatusUnprocessableEntity,
		},
		{
			label:   "resetPasswordOK",
			method:  http.MethodPost,
			path:    "/password",
			body:    strings.NewReader(fmt.Sprintf(`{"token": "%s", "password": "bazquux"}`, passwordToken)),
			expCode: http.StatusNoContent,
		},
		//token
		{
			label:   "getAccessTokenWithAccessToken",
			method:  http.MethodGet,
			path:    "/token/access_token",
			token:   userToken,
			expCode: http.StatusUnauthorized,
		},
		{
			label:   "getAccessToken",
			method:  http.MethodGet,
			path:    "/token/access_token",
			token:   refreshToken,
			expCode: http.StatusOK,
			expBody: "user_id:5 is_admin:false",
		},
		{
			label:   "authTokenWithAccessToken",
			method:  http.MethodGet,
			path:    "/token/auth",
			token:   userToken,
			expCode: http.StatusUnauthorized,
		},
		{
			label:   "authTokenWithRefreshToken",
			method:  http.MethodGet,
			path:    "/token/auth",
			token:   refreshToken,
			expCode: http.StatusOK,
			expBody: `user_id:5 is_admin:false$user_id:5 is_admin:false${"email":"foo@bar.com","href":"/users/5","is_admin":false,"is_validated":true,"nickname":"foo","user_id":5}`,
		},
		{
			label:   "deleteUserUnauthorized",
			method:  http.MethodDelete,
			token:   userToken,
			path:    "/users/4",
			expCode: http.StatusUnauthorized,
			expBody: `unauthorized`,
		},
		{
			label:   "deleteUserNotFound",
			method:  http.MethodDelete,
			token:   adminToken,
			path:    "/users/999",
			expCode: http.StatusNotFound,
		},
		{
			label:   "deleteUserOK",
			method:  http.MethodDelete,
			token:   adminToken,
			path:    "/users/5",
			expCode: http.StatusNoContent,
		},
	}

	for i := range tests {
		if err := test(tests[i].method, tests[i].token, tests[i].path, tests[i].body, tests[i].expCode, tests[i].expBody, tests[i].expLocation); err != nil {
			log.Println(tests[i].label, err)
		}
	}
}
