# Sample book store

A sample Golang application which utilizes Gin to have a Web API for managing a book store

We will take some inspiration of project structure from sample buffalo projects

# Running the application

To test run the application, run the following command:

```
go run *.go
```

To create the final binary for deployment:

```
go build
```

# Application Decisions

Notes on some application decisions
- Instead of having database related stuff within models; extract them out -> They will handled on a high level construct (probably the controllers or sth)
- Instead of having database related stuff within controllers; extract them out into a service layer. -> This layer will need to be extremely lightweight (Does very little processing as possible)
- So in summary:
  - Models: Contain the business logic and structs that would returned to user
  - Controllers: Contains the API logic that takes in interfaces; this would allow it to be able to switch out for fake services to do API level testing. API level testing is important in order to ensure a consistent API for frontend users
  - Services: Is the intermediate layer that would contain intergration work etc. This would include the interfacing with DB or even APIs
- IMPORTANT: Although its tempting to just provide a DB interface to the handlers functions, that would mean exposing some of the implementation of DB functionality to the handlers -> Making it slightly fragile, e.g. The DB implementation becomes depreciated and not in use. By spliting the DB implementation from handler, we can then have the handler call the services via interfaces.

# References
- Database structure: Dependency injection thru structs. http://www.alexedwards.net/blog/organising-database-access