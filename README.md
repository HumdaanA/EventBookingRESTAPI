This is my Personal Project that I did. This is an Event Booking REST API made with GoLang.

It has many different features and technologies that went into building this.

First thing I did was set up the routes, as you can see in the routes.go file in the routes package
  I used a package on github called gin, that is an HTTP request package you can use to handle requests like GET, POST, PUT, etc.

There are mutliple routes that I used:

  GET, list of available events
  
  GET, list of available single event
  
  POST, create a new bookable event
  
  PUT, updating an event
  
  DELETE, deleting an event
  
  POST, registering a user for an event
  
  DELETE, deleting a registration for a user

  POST, creating a new user
  
  POST, login a user, with Authentication
  

- the routes 3-7 use an authentication method to secure them, since we dont want anybody, for example, deleting an event.
  so had them in a server group to group them together in the "authenticated" group which uses a middleware to authenticate
  the user
  

- For authentication, I used another golang package called golang-jwt. This isnt really available in the standard library, but
it is upheld up GO themselves.

  with this package, you can generate a token for a user when they login, and certain actions can only be completed with a 
  verified token, for example, using the POST route to create a new event.

  furthermore, using the DELETE route to delete an event can only be used if the token is verified, AND it matches
  the token that was used to create it, so only the creator of the event can delete said event.
  

- The signup process contains a password as well, so we needed some security measures for that.

  I didnt want the passwords to be kept in plain text, so I used another package,
  crypto/bcrypt to hash the passwords, as you can see in the hash.go file in the utils package
  this is another package that is not part of the standard library.
  

- And finally, I also added a SQL database to ensure all the information is kept, even if the program is not running.
  for this, I used the go-sqlite3 package, which is not available in the standard, library. this needs to be used with the
  database/sql package which is in the standard library.
  

  I first initialized the database, and then created the tables that I needed to store info, as you can see in the db.go file
  in the db package. I needed to create a users tables, events table, and registrations table to store info. I also added error
  handling to make sure the tables get created.















  
