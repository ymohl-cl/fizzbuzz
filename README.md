# fizzbuzz

fizzbuzz is a test for recruiting

## Describe

The original fizz-buzz consists in writing all numbers from 1 to 100, and just replacing all multiples of 3 by “fizz”, all multiples of 5 by “buzz”, and all multiples of 15 by “fizzbuzz”. The output would look like this: “1,2,fizz,4,buzz,fizz,7,8,fizz,buzz,11,fizz,13,14,fizzbuzz,16,...”.

### The goal is to implement a web server that will expose a REST API endpoint that

* Accepts five parameters : three integers int1, int2 and limit, and two strings str1 and str2.
* Returns a list of strings with numbers from 1 to limit, where: all multiples of int1 are replaced by str1, all multiples of int2 are replaced by str2, all multiples of int1 and int2 are replaced by str1str2.

### The server needs to be

* Ready for production
* Easy to maintain by other developers

### Bonus question

Add a statistics endpoint allowing users to know what the most frequent request has been. This endpoint should:

* Accept no parameter
* Return the parameters corresponding to the most used request, as well as the number of hits for this request

## Usage

``` bash
./fizzbuzz -help
Usage ./fizzbuzz
  -appName string
        your application name (default "fizzbuzz")
  -help
        show app configuration

This application is configured via the environment

To configure the server the following environment variables can be used:

            KEY                         TYPE             DEFAULT    REQUIRED    DESCRIPTION
            FIZZBUZZ_SSL_ENABLE         True or False               true
            FIZZBUZZ_SSL_CERTIFICATE    String
            FIZZBUZZ_SSL_KEY            String
            FIZZBUZZ_PORT               String                      true

To configure postgres the following environment variables can be used:

            KEY                           TYPE      DEFAULT    REQUIRED    DESCRIPTION
            FIZZBUZZ_POSTGRES_HOST        String               true
            FIZZBUZZ_POSTGRES_PORT        String               true
            FIZZBUZZ_POSTGRES_USER        String               true
            FIZZBUZZ_POSTGRES_PASSWORD    String               true
            FIZZBUZZ_POSTGRES_DB          String               true

Enjoy !
```

## Docker

The project is dockerised so, you can make

``` bash
docker-compose up --build
```

On the process build, golint and go test are execute

## Endpoint list

### localhost:4242/fizzbuzz

Method POST on http protocol

``` JSON
{
    "limit": 16,
    "int1": 3,
    "int2": 5,
    "str1": "fizz",
    "str2": "buzz"
}
```

### localhost:4242/stats

Method GET on http protocol

## Enjoy to test :)