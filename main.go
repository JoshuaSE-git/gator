package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/JoshuaSE-git/gator/internal/config"
	"github.com/JoshuaSE-git/gator/internal/database"
	_ "github.com/lib/pq"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		panic(err)
	}

	db, err := sql.Open("postgres", cfg.DbURL)
	if err != nil {
		panic(err)
	}

	dbQueries := database.New(db)
	state := &State{cfg: cfg, db: dbQueries}

	commands := &Commands{commandMap: map[string]func(state *State, cmd Command) error{}}
	commands.register("login", handlerLogin)
	commands.register("register", handlerRegister)
	commands.register("reset", handlerReset)
	commands.register("users", handlerUsers)
	commands.register("agg", handlerAgg)
	commands.register("addfeed", middlewareLoggedIn(handlerAddfeed))
	commands.register("feeds", handlerFeeds)
	commands.register("follow", middlewareLoggedIn(handlerFollow))
	commands.register("following", middlewareLoggedIn(handlerFollowing))
	commands.register("unfollow", middlewareLoggedIn(handlerUnfollow))
	commands.register("browse", middlewareLoggedIn(handlerBrowse))

	if len(os.Args) < 2 {
		fmt.Println("too few arguments")
		os.Exit(1)
	}

	commandName := os.Args[1]
	commandArguments := os.Args[2:]
	command := Command{name: commandName, args: commandArguments}

	err = commands.run(state, command)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
