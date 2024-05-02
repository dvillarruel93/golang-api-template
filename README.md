# Golang-API-Template

## About

Golang-API-Template is a clean template with all you need to start a new API.

The template was built in [GO](https://github.com/golang/go) programming language in 1.21 version, using the [Echo Context](https://github.com/labstack/echo) web framework.  

It is following a [Domain-Driven Design](https://en.wikipedia.org/wiki/Domain-driven_design) architecture divided in four layers:
* Handler: First layer, where the request begins its journey. You can find the bindings of the ULR and query parameters and their validations. Located in `delivery/rest` directory. 
* Usecase: Second layer, where you can see business logic. Located in `usecase` directory.
* Repository: Third layer who is in charge of connect and get/save data from/to the database. Located in `repository` directory.
* Model: Fourth and last layer where all models should be defined, with their validation and ORM tags.

### Authentication

Handler layer is using a middleware to authenticate every call. It is in `delivery/middleware` directory. This is just an example of how to use a middleware.    

### Validator

Project was designed to validate any struct with [Go Package Validator](https://pkg.go.dev/github.com/go-playground/validator/v10)
which basically implements value validations for structs and individual fields based on tags.  
You can see an example on how to define tags in `models/person.go` inside Person struct and how to call validator in `delivery/rest/example_handler.go` inside `AddPerson()` function.  

### Tests

Handler and Usecase layers are prepared for unit tests thanks to their interface definitions. 
Project uses a mock generator called [Mockery](https://vektra.github.io/mockery/latest/) which you must install if you want to use it.
After its installation you can generate all mocks running `make generate-mocks` from console in root directory.
It will generate all the mocks for every function defined in `interfaces` directory.  

In addition to unit tests, project is also prepared for integration tests thanks to the Repository layer.
You can check an example on how to do it in `repository/example_repository_test.go`  

### Database and ORM

[MYSQL](https://dev.mysql.com/doc/) is used in this project due to its popularity through [GORM](https://gorm.io/docs/index.html), which provides all the necessary tools to handle data from and to DB.

### Error handler

There is a little error handler inside `api_error` which just wrapper the generation of the errors.
It's simple and useful to return any error code and message you want to the client.

## How to

Project has a short example of a few API calls to check all work as expected.  
First of all, you must have MySQL installed and running in your computer.
In my case, my user and password for my local DB are root and example and MySQL process runs in port 3307.
You have to set up them as you wish, changing them in `repository/db.go` inside `connectionString` constant.  

After previous step was completed, you have to execute the following command in root directory:
```
go run ./cmd/app
```
and will see the following message in console 
```
http server started on [::]:8080
```

Once you see that, you are ready to execute the following cURL's:

1. This is just to test API is working (like a ping but connecting all layers)
```
curl --location 'http://localhost:8080/v1/test' \
--header 'auth-token: valid-token'
```

2. Post a person in DB
```
curl --location 'http://localhost:8080/v1/person' \
--header 'auth-token: valid-token' \
--header 'Content-Type: application/json' \
--data '{
    "first_name": "a first name",
    "last_name": "a last name",
    "address": "an address"
}'
```
You can play changing the values in the body to check if validator returns an error (hint: send an empty string in any of the three values)  

3. GET a person from DB
```
curl --location 'http://localhost:8080/v1/person/:personID' \
--header 'auth-token: valid-token'
```
You must replace `:personID` with the ID os the person in DB.  
If you've executed step 2 successfully, you can take the id that you received in response and use it here. 