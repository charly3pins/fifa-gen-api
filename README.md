# fifa-gen-api

API for manage the FIFA tournaments generator. Used Go for the backend and Flutter for the app. You can find the app repo [here](https://github.com/charly3pins/fifa_gen)


## How to run


Firs of all you need to have installed the [dep](https://github.com/golang/dep) for manage the Go dependencies. And then type:

```
dep ensure
```

After that, you can start the application following these steps:


- First start the Postgres database:
```
docker-compose up
```

- Second run all the migrations on it:
```
make migration-run dir=up
```

- Third start up the application:

```
make run
```
