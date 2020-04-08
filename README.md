## mir-fiber

This is the source code of the demo is explain how to use mir and fiber to develop web application. 

#### Prerequisites
* Go > go1.12
* MongoDB with a database called `todos`

#### Usage
```bash
% make generate
% make run
```

#### Custom connection string

Use the `MONGODB_URI` environment variable to specify a custom MongoDB connection string:

```sh
MONGODB_URI=mongodb://user:pass@host:27017 go run .
```