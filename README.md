# Fund trading challenge

A small challenge application for fund trading written in Go, using a SQL 
database with preparation for deployment with containers and pipeline.

## Description

To implement this fund trading service, based on the requirements I would have, at least:
- service/application to deploy the http API that provide the features (which is the focus on the repo)
- use a code structure with what on this repo
- create a external repo to contain packages that could be use throughout other code bases, like managing cash and unit data type,
database setup tools, leaving the code here

The structure I choose to use is very in-step with Go recommendations with some specific choices:
- `cmd` - in here it is the `main.go` and `Dockerfile` that relates to the application
- `internal` - packages split in the blocks that constitute the logic of the app:
  - This as a handler for the server requests with is API routes
  - The processor that contains only the flow to process the request, good to have methods to handle the specific routes step by step, in a very tickety way 
  - The models to have the application type of data, like a DTO, make use of struct tags to encode/decode json
  - The database to have the mapping and handling of use cases for each operation based on route request 
  - Could include other packages that would interact with other services
- `tools` - this is very useful to have some scripts or small go apps to help out with development or tests.
- `tests` - While not available, it very useful to setup the integration tests in here.

For no go code, my experience is to have a `deployments` folder to store and organise the k8s config files, using kustomize yaml descriptions, an a `build` folder here I add files to help with the app, example, smaller makefiles, or pipeline config files.

### Code

Looking over the sample code, the application as the `main.go` will just instantiate the several parts of the internal services and make it run.

We have the database with a Config struct that will have read from the environment variables the values that then would be use on start the DB `New()` so we have a connection.

This is a excellent candidate to have a lib package, like `appconfig` so we can simplify the reading and use default values, that way, across database, http service, messaging we just pick the package and just need to set the variable names, making easy and standard way to name them.

We pass the DB connection to the processor, handling the responsability to access database to it, in there is where the business logic and flow would be managed,
like having the validation rules call upon, interact with database, I would consider the sending of messaging to other services internally, and mapping data between the HTTP data and database.

Finally the handler will prepare the API routes handling just doing some validation about headers and authentication and provide the mapping to send to the processor.

This type of code organisation is very useful to read and separate concerns between different parts of the code, while it may make some methods big, they can be split between files/packages more easily, and for new members is easier to understand.

### Database 

I choose `gorm` to use to interact to SQL database as I have experience in it and provides some easy methodology to interact with a database.
The set up is easy, handles transactions already and provide some nice handling for model representation and automate some parts, like using model struct tags
to perform some  "magic", example the created at and already likes using UUID that postgresql have some support.

About overhead, I don't find it so big in experience, however it as a con to use it, some error handling is counter intuitive, check the example about `FindOrderByCorrelation`, the "not found" is a bit hidden and I suffer from the past from it.

### Unit tests 

A bit like the sample provided, unit tests will tend to follow closely the code and use a table drive test format, is more readable and easy to catch test cases,
and AI code generation works better that way.

I would choose to use `mockery` go package, installed to generate all mocks for the interfaces and assist in generate mocks. the reason being, I wrote mocks by
hand and are a pain to keep update, the process I advise and follow is:
1. keep a `interface.go` in packages or the interface definition on the top of package main file.
2. Run mockery through a make file entry like this:

```make
.PHONY: generate-mocks
generate-mocks:
  rm -rf mocks/*
  mockery --all --output=mocks --outpkg=mocks
```

### Integration tests

I have some experience in writing integration tests but most I learned from a QA, and this organisation I think is useful, using a `tests` folder,
where the QA and Devs can write the integration tests, where the preparations with test data and setup can be made and the files for the tests can be placed.

I have experience using `ginkgo` and `gomega` to write the tests and is very useful to use: it facilitates automation, parallel run and exclusion and TAG tests. This is useful to run local as well in a pipeline.

## Left to do

Almost everything! Tried to keep in the 2 hours limit but planned to do around 4 to 6hours, but that was way limited:
- wanted to provide a skeleton of the repo with code organisation and knowledge of my experience
- write more comments and samples through the repo

I experienced the fact that I have no personal hardware that allow me to keep up and most of my code that I wrote and experiment was done in company laptop, which I dont have anymore.

But big things I would do:
- Would probably move to a two services setup:

1. HTTP server: provides the API and for read requests, will access to the database but for write access would publish a message with data to another service.
2. Fund processing: reads the messages, from something like AWS SQS queue with FIFO, process everything and notifies through a Notification event stream or with a more complex system where it responds to the HTTP service.

This two services setup can be made using the same repo where in `cmd` will have:
- `cmd/api-server/{main.go,Dockerfile}`
- `cmd/fund-processor/{main.go,Dockerfile}`

For deploument k8s configuration will be following this rule.

- Add check for authentication tokens in the API 
- Add a Config setup to HTTP service with a `config.go` where it loads env vars and use it to configure the http server with timeouts, endpoints and such so it can be using env vars instead of hardcoded values.
- Complete the `CreateOrders` flow through the application. It would be a good sample of the project.
- review the SQL definition, as I used sqlite to write it, and using Postegresl would allow better data type definition and use the UUID I grow fond of.
- Prepare a bit of the audit or notification idea to be clear about it, I think a centralised system and indepedent would be better in an asynchronous world.
- Do better investigation about the cash and units requirement, I check some resources and found that using Integer is common, which I was not aware. Would write a package for that, common to app and org, maybe a Go struct to handle it and match it to database data type, allowing to be a UINT64 and decimal units accordingly and convert to/from strings.
- I wrote the Dockerfile from my knowledge and checking the docs, but my laptop does not run Docker so is untested, so no way I could try a docker-compose file valid:
  - I work with docker compose like 3 to 5 years ago, so I would know to setup a service, a volume and expose the app ports to enable it to start the system
  - Been working with skaffold that allows to use the k8s kustomization files and localstack to mock some 3rd party services, which is nice.
  - K8s: would write the base config files, with resources organised with:
    - `deployment.yaml`: configure the application, service, healtz endpoint, environment mappings and credentials
    - `namespace.yaml`: setup the namespace where the services would run
    - `secrets.yaml`: the mapping and unseal of secrets to send related to password and user for database, for example 

  - Pipeline, I worked mainly worked on a Jenkins pipeline setup, where each stage of the pipeline was driven from Makefiles. It is a good experience to have this type of stagings:
    1. Linting
    2. Sonarqube with static code analysis
    3. Unit test run
    4. Build image with code app
    5. Deploy to env setup for PR or main branch
    6. Run integrations tests again running pod.
    7. Manual production deploy

  Looking into gitlab ci I never used, but looking at the syntax, it looks kind of similar to in a way to implement this kind of stage by stage. Note: secrets and some linting should be done locally, preferable on 
  on a git hook (example `ggshield` to look for secrets and run `golangci-lint` there is better than find a commit there, failures made me choose this).

:q
