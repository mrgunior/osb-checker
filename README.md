[![Build Status](https://travis-ci.org/openservicebrokerapi/osb-checker.svg?branch=master)](https://travis-ci.org/openservicebrokerapi/osb-checker "Travis")

# osb-checker

An automatic checker to verify an Open Service Broker API implementation against the [specification](https://github.com/openservicebrokerapi/servicebroker).

# Project Status

This project should be considered **experimental**. You should validate the results against the released [specification](https://github.com/openservicebrokerapi/servicebroker). In the case of any discrepancy, the specification should be considered correct.

# Usage

## Test
* Firstly, u need to deploy your own service broker to be tested.
* Modify the `config_mock.yaml` under test/configs folder.
* Just run `go test -v ./test/` to start the test job.

## Generate model
All the model will be generated from [swagger.yaml](https://raw.githubusercontent.com/openservicebrokerapi/servicebroker/master/swagger.yaml) automatically. Here are some steps for developers to generate the model:
* First developers should download `go-swagger` tool from [here](https://github.com/go-swagger/go-swagger/releases)
* Then run these command below:

	```shell
	cd ./autogenerated && wget https://raw.githubusercontent.com/openservicebrokerapi/servicebroker/master/swagger.yaml
	go-swagger generate model ./swagger.yaml
	```

