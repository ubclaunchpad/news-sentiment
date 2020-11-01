# Contributing to News-sentiment 🎉💃
First off, thank you for taking the time to contribute!

The following is a set of guidelines for contributing to `news-sentiment`.

## Code of conduct 📖
This project and everyone participating in it is governed by the [UBC Launch Pad Code of Conduct](https://docs.ubclaunchpad.com/handbook/manifesto#code-of-conduct). By participating you are expected to uphold this code. Although the code's language may be directed towards UBC Launch Pad members, all contributers are expected to abide by it, specifically the following policy:
> Treat everyone with respect - we do not condone any form of discrimination or intolerance

## What should I know before I get started? 💻
This is a UBC Launch Pad project. [UBC Launch Pad](https://ubclaunchpad.com) is a Software Engineering Club at the University of British Columbia, and as such, all its members are students. Please be considerate of that fact as we are trying to juggle a full course load and maintaining this project.

This project is built using a [`go`](https://golang.org/) backend and a chrome extension as a frontend.

### Backend 🍑
The backend is written in `go` and uses go modules for dependency management.

#### Routing 🚗
We use `go`'s rich http support to handle routing. As of the time I'm writing this, we don't use any special libraries here, but we might use something like [`mux`](https://github.com/gorilla/mux).

#### Persistence 🥭
We use a `MongoDB` database for persistence, accessed from go.

### Frontend 😄

The Frontend is a chrome extension... Ya that's it at the time!


## How can I setup my environment? 🖥️
### Backend
To setup the backend you need `go` installed, you can find instructions on [go's download website](https://golang.org/dl/). Beyond that, you would need a `MongoDB` database, you can that working using a local `MongoDB` environment, easiest way is through [`docker-compose`](https://docs.docker.com/compose/) by running the following in the `backend` directory:
```bash
docker-compose up
```
Otherwise you should be ready to go! (get it? go? ok I'll let myself out)

### Frontend
You can run the extension locally by following the instructions in [the frontend directory](./frontend/README.md)

## Pull-request checklist ✅
First of all, you don't need an issue to file a pull-request! Although if the change is significant please let us know beforehand.
Before you open your pull-request check the following (the pull-request template will have the same!)
- ✔️ I have read the code of conduct and I recognize that I'm expected to uphold it
- ✔️ I have tested my code, and the pull request includes a description of what the tests were like, or a justification why tests were not needed
- ✔️ My change does not change any public APIs, or if it does, I made sure to change any documentation to reflect the changes
