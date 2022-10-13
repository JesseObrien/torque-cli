package gen

import (
	"fmt"
	"strings"
	"unicode"

	"github.com/apex/log"
	"github.com/jesseobrien/torque-cli/internal/scaffold"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
)

var (
	ControllerCmd = &cobra.Command{
		Use:   "controller [name]",
		Short: "Generate a new controller for a Torque app. ie: torque gen controller users",
		Long:  "",
		Run:   executeController,
	}
)

func executeController(cmd *cobra.Command, args []string) {
	if len(args) < 1 {
		log.Error("Required `name` arg is missing.")
		return
	}

	controllerName := args[0]

	err := ValidateControllerName(controllerName)
	if err != nil {
		log.Errorf(err.Error())
		return
	}

	controllerName = fmt.Sprintf("%sController", controllerName)

	log.Info(fmt.Sprintf("🔨 Generating New Controller -- %s", controllerName))

	cpf, ok := scaffold.ProjectFiles[scaffold.CONTROLLER]
	if !ok {
		log.Errorf("could not find controller project file")
		return
	}

	err = cpf.ScaffoldCustomTemplate(
		afero.NewOsFs(),
		fmt.Sprintf("internal/http/%s_controller.go", controllerName),
		struct{ ControllerName string }{ControllerName: strings.Title(controllerName)},
	)

	if err != nil {
		log.Errorf(err.Error())
	}
}

func ValidateControllerName(controllerName string) error {
	for _, l := range controllerName {
		isLetter := unicode.IsLetter(l)
		isNumber := unicode.IsNumber(l)
		if !isLetter && !isNumber {
			return fmt.Errorf("error: controller name `%s` contains invalid character: %q", controllerName, l)
		}
	}

	return nil
}
