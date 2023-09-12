# HNG X
## Stage 2 - A simple CRUD Api written in GO.

### Description
This API is written in the go programming language, and uses mysql as the database.

It allows a user to create a person, find a person, update a person and also delete a person's data from the database.

### Outline
The project's entry point is `cmd/main.go`.

This is where the bootstrapping happens.
1. We set up the .env configuration.
2. We set up the database connection.
3. We migrate the tables that are required.
4. We set up the routers.
5. Start the application server.

The all requests into the app goes to the router, which can be found in the `src/routes` folder.

We have just one router, the person router. This handles all the CRUD requests on the person resource (model).

The router calls the controller to handle each endpoint, the controller in turn communicates with the models for
any database related activity.

Some helper functions for handling responses can be found in `src/utils`.


### API Documentation
The API is fully documented, and can be found [here](https://documenter.getpostman.com/view/4194134/2s9YC31ZgX).

### Testing
Unit tests/feature tests were not part of the requirements, but from the API docs, you will be able to test it on Postman.

The API is deployed to [hng.jameesjohn.com](https://hng.jameesjohn.com).

### Run Locally.
To run this project locally, you need to have your environment setup for go development with mysql.

Then copy `.env.example` to `.env` and update the variables as required.

In the root folder, run the following to download the required libraries.
```shell
go mod download
```

CD to the cmd directory

```shell
cd cmd
 ```

Compile the code

```shell
go build -o ../server
```

The compiled binary will be found in the root directory, and can then be run.

```shell
./server
```

Once you get the following, you are good to go!
```shell
2023/09/12 11:50:31 Config Loaded Successfully
2023/09/12 11:50:31 Http server running on port 8000
```


