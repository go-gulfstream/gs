package uiwizard

import (
	"fmt"
	"strings"

	"github.com/go-gulfstream/gs/internal/schema"

	"github.com/manifoldco/promptui"
)

func selectTemplate(label string) *promptui.SelectTemplates {
	return &promptui.SelectTemplates{
		Label:    "=> {{ . }}: ",
		Active:   "=> {{ .Name  | bold | green }} ({{ .Help }})",
		Inactive: "- {{ .Name }} ({{ .Help | white }})",
		Selected: fmt.Sprintf("=> %s: {{ .Name }}", label),
	}
}

func inputTemplate(label string, defVal string) *promptui.PromptTemplates {
	if len(defVal) > 0 {
		defVal = " (" + defVal + ")"
	}
	funcMap := promptui.FuncMap
	funcMap["fix"] = func(s string) string {
		return ""
	}
	return &promptui.PromptTemplates{
		FuncMap: funcMap,
		Prompt:  fmt.Sprintf("=> {{ .}}%s: ", defVal),
		Valid:   fmt.Sprintf("=> {{ . | green }}%s: ", defVal),
		Invalid: fmt.Sprintf("=> {{ . | red }}%s: ", defVal),
		Success: fmt.Sprintf("=> %s: {{ . | fix }}", label),
	}
}

func confirmTemplate() *promptui.PromptTemplates {
	return &promptui.PromptTemplates{
		Confirm: "=> {{.}}: [N/y]? ",
		Success: "=> {{.}}: [N/y]? ",
		Invalid: "=> {{.}}: N",
	}
}

type selectItem struct {
	ID   string
	Name string
	Help string
}

func sectionControl(label string) {
	label = strings.ReplaceAll(label, " ", "/")
	fmt.Printf("%s:\n", strings.ToUpper(label))
}

func confirmControl(
	label string,
) (bool, error) {
	control := &promptui.Prompt{
		Label:     label,
		Templates: confirmTemplate(),
		IsConfirm: true,
	}
	val, err := control.Run()
	if err != nil {
		return false, nil
	}
	return val == "y", nil
}

func inputControl(
	label string,
	defValue string,
	validate bool,
) (string, error) {
	control := &promptui.Prompt{
		Label:     label,
		Templates: inputTemplate(label, defValue),
		AllowEdit: true,
		Validate: func(s string) error {
			if !validate {
				return nil
			}
			if len(s) < 3 {
				return fmt.Errorf("%s value too short. got %d, expected > 3",
					s, len(s))
			}
			if schema.CheckUnique(s) {
				return fmt.Errorf("already exists")
			}
			return nil
		},
	}
	return control.Run()
}

func selectControl(
	label string,
	values []selectItem,
) (selectItem, error) {
	control := &promptui.Select{
		Label:     label,
		CursorPos: 0,
		HideHelp:  true,
		Items:     values,
		Templates: selectTemplate(label),
		Searcher: func(input string, index int) bool {
			val := values[index]
			name := strings.Replace(strings.ToLower(val.Name), " ", "", -1)
			input = strings.Replace(strings.ToLower(input), " ", "", -1)
			return strings.Contains(name, input)
		},
	}
	index, _, err := control.Run()
	if err != nil {
		return selectItem{}, err
	}
	return values[index], nil
}
