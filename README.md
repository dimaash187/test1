# GoLANG test1 BACKEND

This project is a backend server to run and expose some endpoint to process transactions and sort them:
It exposes 3 endpoints:

    1. /transactions [ GET request ]
    2. /transactions_by_timestamp [ GET request ]
    3. /transactions_by_timestamp [ POST request ], send a payload of json data ( see data.json file in the project).
    
## Build

Run `make build` to build the project. Main executable will be generated in the project root. 


## Running unit tests

Run `make test` to execute the unit tests, WHICH WE DONT'T HAVE yet...

