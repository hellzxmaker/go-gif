# go-gif

## Description

Backend microservice supporting GoGifUrself.

## Dev Process

1. Pull the latest code from Github
2. Update the go deps `go get`
3. Run all local tests to ensure a good base has been pulled `go test`
4. Make any code changes. Please ensure test cases are added for any new code.
5. Run all local tests to ensure the code is working, again... `go test`

### Locally Run

From the root of the project:

`go get`

`go test`

`go run main.go`

### Containerized Execution (nope)

> We need to get an external repository to push to, so we can run this
> bish in the cloud somewhere.

1. Build the container: `docker build . -f dockerfile -t go-gif`
2. Run the container locally `docker run -p 8080:8080 go-gif`

## Endpoints Supported (well not really)

- GET /players/[id] # Get a player by ID
- GET /players/ # Get a collection of players
- POST /game_sessions # Create a game session
- GET /game_sessions/[id] # Get a game session by ID
- DELETE /game_sessions/[id] # Delete a game session by ID
- GET /game_sessions/[id]/players # Get players in a game session
- PUT /game_sessions/[id]/players # Add a player to a game session

## Additions Required ASAP

- use make or some build automation
- linting of code
- automatic `go test` reports
- determine how this will run in the cloud
- different place for actual tests
- DB type and hosting needs to be confirmed
- instrument the running to allow for better monitoring
- document the dev practices like branching etc


## NOTE: EXTRACT ALL OF THIS OUT OF CODE BASE

## GoGifUrself

Number of players: 2-8 (must be even number of players for pairing)

Players are paired off in a round-robin fashion and asked a series of questions which they must respond with a gif of their choice.
At the end of each round, the other players vote on their favorite answers. Highest score at the end of three rounds wins!

The three rounds: first blood, middle game, and the revelation.

First Blood: Players compete in pairs. Other players in game vote for their favorite answer.
Middle Game: Players compete in groups of total players / 2. Other players in game vote for their favorite answer.
Revelation: Players are all asked a single prompt. Players can vote on anyone other than themselves.

### Roles for Players

- Game Creator/Host
- Player

#### Game Creator/Host User Story: Create a New Game

1. The game creator opens the app on their client
2. The game creator creates a game session
3. The game creator can start the game
4. The game creator can end the game

#### Player User Story: Join an Existing Game

1. The player opens the app on their client
2. The player can join a game session via the join code
3. The player is prompted to input their player name

#### Both: Play the Game

1. The game starts and all users see the intro screen and rules
2. First blood starts
3. 