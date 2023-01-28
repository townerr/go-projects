# go-projects
Projects for learning Go

#### 1. Go-simple-hello-server
- Go server to basically display html files and handles simple hello world responses/requests

#### 2. Go-user-crud
- Go server using gin to perform CRUD operations on a user data set
- Reads users from a text file - No DB
  - Doesnt update the text file just the array of users in memory

#### 3. Go-sales-crud
- CRUD REST api using gin and gorm
- performs crud operations on fake sales records
- Uses a postgres db

#### 4. Go-discord-bot
- A discord bot made using [discordgo](https://github.com/bwmarrin/discordgo) library
- Has 3 commands:
  - ```.helloworld``` - replies Hello World
  - ```.time``` - replies the bots current time and date
  - ```.pokemon``` - fetches a random pokemon name and number from [poke api](https://pokeapi.co/docs/v2)
