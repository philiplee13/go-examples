## Rest API

- uses docker to bring up psql instance
- psql
  - connection - done
  - pooling connection
  - crud - done
  - middlware - done
  - error handling - done
  - auth
    - authn
    - authz
- restful api - done
  - get
  - post
  - delete
  - put
- http requests to 3rd party services
- unit tests
- integration tests

## Structure

/api - contains the request/response structures for the API. Used by the HTTP handlers. Can be imported by other applications who want to make requests against the service.

/client - if I implement a go client library for my API then I implement it here

/cmd/service/main.go - main entrypoint for the service

/cmd/cli/main.go - main entrypoint for the optional cli tool I build that provides a cli tool against my API. Uses the client I expose in /client.

/internal - contains all internal application logic

/internal/rest - implements the REST API, using echo or fiber or std lib http handlers etc. whatever. I'm not the boss of you

/internal/model - I define all the database handling stuff in here. Gorm models, database bootstrap stuff, etc.

### Test Commands

GET: `curl -X GET http://localhost:8080/users | jq`
POST: `curl -X POST http://localhost:8080/users --data '{"name": "test-user","age":10}'`
DELETE: `curl -X DELETE http://localhost:8080/users/7 | jq`
PUT: `curl -X PUT http://localhost:8080/users/6 --data '{"name": "updated-name", "age": 100}'`
