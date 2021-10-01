package cmd

import (
	"encoding/gob"
	"fmt"
	"os"

	"github.com/i582/CodeQuery/pkg/models"
	"github.com/i582/CodeQuery/pkg/shell"
	"github.com/i582/cfmt"
	"github.com/urfave/cli/v2"
)

type AppContext struct {
	Database *models.Database
}

var AppCtx = AppContext{}

type CreateCommandFlags struct {
	Name string
}

type OpenCommandFlags struct {
	Name string
}

type ExecCommandFlags struct {
	Name    string
	Command string
}

func Run() {
	// 	os.Args = []string{os.Args[0], "exec", "--name", "vkcom", "--command",
	// 		strings.ReplaceAll(`run "
	// 	/*SELECT * FROM calls WHERE
	// 		call.func().name() = 'inc_getAdsUnionIdNew' AND
	// 		call.args() > 1 AND call.arg(1).isConstant() AND call.arg(0).string().contains('$union[')
	// 	LIMIT 20 ORDER BY args_count DESC
	//
	// 	SELECT deps FROM (
	// 	   SELECT * FROM funcs WHERE func.name() = 'inc_getAdsPriceLists'
	// 	) WHERE
	// 	   path.contains(
	// 	      SELECT * FROM funcs
	// 	      WHERE func.className() = ''
	// 	        AND func.name() = 'rpcMcGet'
	// 	      LIMIT 1
	// 	   )*/
	//
	// 	/* SELECT COUNT(*) FROM funcs WHERE func.className() = '' LIMIT 50000000 ORDER BY uses DESC*/
	//
	// SELECT deps FROM (SELECT * FROM funcs WHERE func.name() = 'inc_getAdsPriceLists' ) WHERE path.length() = 4
	//
	// 	/*
	// 	SELECT * FROM funcs WHERE
	// 		func.name() = 'debugServerLog'
	// 	LIMIT 20 ORDER BY uses DESC
	// 	*/
	// 	"`, "\n", "")}

	MainShell := shell.NewShell()
	MainShell.Prefix = "#generic> "

	MainShell.AddExecutor(About())
	MainShell.AddExecutor(RunQuery())

	createFlags := CreateCommandFlags{}
	openFlags := OpenCommandFlags{}
	execFlags := ExecCommandFlags{}

	app := &cli.App{
		Name:        "CodeQuery",
		Version:     "v0.0.1",
		Usage:       "Tool for searching and aggregating data for PHP",
		Description: "Tool for searching and aggregating data for PHP",
		Commands: []*cli.Command{
			{
				Name:  "exec",
				Usage: "Exec command for database",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:        "name",
						Usage:       "database name",
						Value:       "generic",
						Destination: &execFlags.Name,
						Required:    true,
					},
					&cli.StringFlag{
						Name:        "command",
						Usage:       "command to execute",
						Value:       "",
						Destination: &execFlags.Command,
						Required:    true,
					},
				},
				Action: func(c *cli.Context) error {
					if c.NArg() != 0 {
						return fmt.Errorf("'exec' command takes no arguments\n")
					}

					name := execFlags.Name + ".db"

					f, err := os.OpenFile(name, os.O_RDONLY, 0777)
					if err != nil {
						return fmt.Errorf("can't opening file with database named %s.db: %v", execFlags.Name, err)
					}
					defer f.Close()

					db := &models.Database{}

					dec := gob.NewDecoder(f)
					err = dec.Decode(db)
					if err != nil {
						return fmt.Errorf("can't decode database, data is corrupted, please regenerate the database")
					}

					AppCtx.Database = db

					MainShell.RunSingleCommand(execFlags.Command)

					return nil
				},
			},
			{
				Name:  "open",
				Usage: "Open database",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:        "name",
						Usage:       "database name",
						Value:       "generic",
						Destination: &openFlags.Name,
						Required:    true,
					},
				},
				Action: func(c *cli.Context) error {
					if c.NArg() != 0 {
						return fmt.Errorf("'open' command takes no arguments\n")
					}

					MainShell.Prefix = fmt.Sprintf("#%s> ", openFlags.Name)

					name := openFlags.Name + ".db"

					f, err := os.OpenFile(name, os.O_RDONLY, 0777)
					if err != nil {
						return fmt.Errorf("can't opening file with database named %s.db: %v", openFlags.Name, err)
					}
					defer f.Close()

					db := &models.Database{}

					dec := gob.NewDecoder(f)
					err = dec.Decode(db)
					if err != nil {
						return fmt.Errorf("can't decode database, data is corrupted, please regenerate the database")
					}

					AppCtx.Database = db

					MainShell.Run()

					return nil
				},
			},
			{
				Name:  "create",
				Usage: "Create database from source",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:        "name",
						Usage:       "database name",
						Value:       "generic",
						Destination: &createFlags.Name,
					},
				},
				Action: func(c *cli.Context) error {
					if c.NArg() == 0 {
						return fmt.Errorf("'create' command needs arguments specifying folders or files for which the database will be generated\n")
					}

					paths := c.Args().Slice()

					db := RunLinterForPaths(paths)

					f, err := os.OpenFile(createFlags.Name+".db", os.O_CREATE|os.O_WRONLY, 0777)
					if err != nil {
						return fmt.Errorf("can't creating file with database named %s.db: %v", createFlags.Name, err)
					}
					defer f.Close()

					enc := gob.NewEncoder(f)
					err = enc.Encode(db)
					if err != nil {
						return fmt.Errorf("can't encode database, data is corrupted, please regenerate the database")
					}

					cfmt.Println("{{Database has been successfully created}}::green")
					fmt.Printf("Use './CodeQuery open --name %s' command to start an interactive shell\n", createFlags.Name)

					return nil
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		cfmt.Fatalf("{{Error:}}::bold %v", err)
	}
}
