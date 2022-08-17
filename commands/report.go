package commands

import (
	"errors"

	"github.com/jfrog/jfrog-cli-core/v2/plugins/components"
	"github.com/jfrog/jfrog-client-go/utils/log"
)

func GetReportCommand() components.Command {
	return components.Command{
		Name:        "report",
		Description: "Provides metrics regarding Artifactory repositories",
		Aliases:     []string{"rep"},
		Arguments:   getReportArguments(),
		Flags:       getReportFlags(),
		Action: func(c *components.Context) error {
			return reportCmd(c)
		},
	}
}

func getReportArguments() []components.Argument {
	return []components.Argument{
		{
			Name:        "truc",
			Description: "test tests",
		},
	}
}

func getReportFlags() []components.Flag {
	return []components.Flag{
		components.StringFlag{
			Name:         "file",
			Description:  "[Optional] JSON provided by Storage Summary Artifactory Rest API.",
			DefaultValue: "",
		},
		components.StringFlag{
			Name:         "repoType",
			Description:  "{all, local, remote, virtual, fed, build}",
			DefaultValue: "all",
		},
	}
}

type reportConfiguration struct {
	repoType string
	file     string
}

func reportCmd(c *components.Context) error {
	if len(c.Arguments) == 0 {
		message := "Good no argument"
		// You log messages using the following log levels.
		log.Output(message)
		log.Debug(message)
		log.Info(message)
		log.Warn(message)
		log.Error(message)
	}

	if len(c.Arguments) > 0 {
		return errors.New("too many arguments received. Now run the command again, with one argument only")
	}

	var conf = new(reportConfiguration)
	conf.repoType = c.GetStringFlagValue("repoType")
	conf.file = c.GetStringFlagValue("file")

	doReport(conf)

	return nil
}

func doReport(c *reportConfiguration) {

	// retrieve JSON (file or API call)
	if c.file == "" {
		log.Info("No JSON file, will use the storage Summary API of that RT")
	} else {
		log.Info("Filepath :" + c.file)
	}
	log.Info("RepoType: " + c.repoType)

	// store JSON into an array
	// create Map from array
	// 		1. filter based on repoType flag
	// 		2. map[repoType]{Metric1, Metric2, ..., MetricN}
	// display report
}
