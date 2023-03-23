# Awesome Service
### Stack: Go, MySQL

Start the service with a port parameter (port).

Endpoints:
```GET /profile```

Headers:
```Content-type: application/json```

Create a test database and import the structure and data.

```data/scheme.sql``` - structure
```data/data.sql``` - data

Add constraints to the tables.

Write authentication middleware, checking the "Api-key" header in the auth table. In case of an incorrect "Api-key", return a 403 error.

```GET /profile``` - return data for all users. If the "username" parameter is present, return one object.

The output object should contain: id, username, first_name, last_name, city, school taken from the user, user_profile, user_data tables.
