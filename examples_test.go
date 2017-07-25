package main

/*
TODO error type for model
TODO error type for Unprocessable Entity
TODO testing store
TODO testing controller user
TODO testing controller book
TODO testing controller ownership
TODO testing controller auth
TODO testing controller validate
TODO set consumes in designs


# User list
curl -v "http://localhost:8089/users"
curl -v "http://localhost:8089/users?nickname="

# User create

-- same email
curl -d "email=admin%40myinventory.com&nickname=NBR41&password=favosaga" "http://localhost:8089/users"

## AUTHENTICATE

curl -v -X POST -d {''} "http://localhost:8089/authenticate"
curl -v -X POST -d '{"login":"NBR41","password":"caca"}' "http://localhost:8089/authenticate"

curl -v -d '{"email":"foo@bar.com","nickname":"NBR41","password":"foobar"}' "http://localhost:8089/users"
curl -v -X POST -d '{"login":"NBR41","password":"foobar"}' "http://localhost:8089/authenticate"

# USer update
curl -v -X PUT -d '' "http://localhost:8089/users/4"
curl -v -H "Authorization: BEARER eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1MDExODg1OTMsImlzcyI6InRlc3QiLCJ1c2VyX2lkIjo0LCJpc19hZG1pbiI6ZmFsc2V9.PUmcG_4Ww8XCygn4WTC59LXBJYQIlQSiDQS6wmaGYfA" -X PUT -d '{"nickname":"admin"}' "http://localhost:8089/users/4"




## VALIDATE
-X
-H
*/
