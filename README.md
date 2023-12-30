# Request Counter

A simple web service that on each request responds with a counter of the total number of requests that it has received
during the previous 60 seconds (moving window).

### Quick Start

1. **Build the Web Service with Docker:**

Use the following command to build the web service using Docker:

```shell
make build
```

2. **Run the Web Service:**

Execute the command below to run the web service and serve it on localhost:8080:

```shell
make run
```

3. **Use the Web Service:**

To use the web service you can call the endpoint below:

`GET`: http://localhost:8080/count

this will return the count of the requests in the last 60 seconds

4. **Run Local Tests:**

To run tests on your local machine, simply execute:

```shell
make test
```