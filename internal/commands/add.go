package commands

import (
	"fmt"

	"github.com/go-gulfstream/gs/internal/schema"

	"github.com/manifoldco/promptui"

	"github.com/spf13/cobra"
)

func addCommand() *cobra.Command {
	command := &cobra.Command{
		Use:   "add [PATH]",
		Short: "Add mutation manager",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return fmt.Errorf("invalid number of arguments. got 0, expected 1")
			}
			drawBanner()
			return runAddCommand(args[0])
		},
	}
	return command
}

func runAddCommand(path string) error {
	mutations, importEvents, err := makeMutationsUI()
	if err != nil {
		return err
	}

	_ = mutations
	_ = importEvents

	return nil
}

const (
	commandMutations = iota
	eventMutations
)

func makeMutationsUI() (*schema.Mutations, []string, error) {
	pakages := make(map[string]struct{})
	mutations := &schema.Mutations{
		Commands: make([]schema.CommandMutation, 0),
		Events:   make([]schema.EventMutation, 0),
	}
	for {
		mutationID, err := selectMutationTypeUI()
		if err != nil {
			return nil, nil, err
		}
		switch mutationID {
		case commandMutations:
			cm, err := makeCommandMutationUI()
			if err != nil {
				return nil, nil, err
			}
			mutations.Commands = append(mutations.Commands, cm)
		case eventMutations:
			em, pkg, err := makeEventMutationUI()
			if err != nil {
				return nil, nil, err
			}
			mutations.Events = append(mutations.Events, em)
			pakages[pkg] = struct{}{}
		}
		if !nextStepUI() {
			break
		}
		fmt.Printf("\n\n")
	}
	importEvents := make([]string, 0, len(pakages))
	for pkg := range pakages {
		importEvents = append(importEvents, pkg)
	}
	return mutations, importEvents, nil
}

func makeCommandMutationUI() (schema.CommandMutation, error) {
	mutation := schema.CommandMutation{}

	mutationName, err := inputUI("MutationName", false)
	if err != nil {
		return mutation, err
	}
	mutationName = schema.SanitizeName(mutationName)

	commandName, err := inputUI(fmt.Sprintf("=> %s.InCommand.Name", mutationName), false)
	if err != nil {
		return mutation, err
	}
	commandName = schema.SanitizeName(commandName)

	commandPayload, err := inputUI(fmt.Sprintf("=> %s.InCommand.Payload", mutationName), true)
	if err != nil {
		return mutation, err
	}
	commandPayload = schema.SanitizeName(commandPayload)

	eventName, err := inputUI(fmt.Sprintf("=> %s.OutEvent.Name", mutationName), false)
	if err != nil {
		return mutation, err
	}
	eventName = schema.SanitizeName(eventName)

	eventPayload, err := inputUI(fmt.Sprintf("=> %s.OutEvent.Payload", mutationName), true)
	if err != nil {
		return mutation, err
	}
	eventPayload = schema.SanitizeName(eventPayload)

	opt, err := selectOptionsUI()
	if err != nil {
		return mutation, err
	}

	mutation.Command.Name = commandName
	mutation.Command.Payload = commandPayload
	mutation.Event.Name = eventName
	mutation.Event.Payload = eventPayload

	if opt == 1 {
		mutation.Create = schema.YesOp
	}
	if opt == 2 {
		mutation.Delete = schema.YesOp
	}

	return mutation, nil
}

func makeEventMutationUI() (schema.EventMutation, string, error) {
	mutation := schema.EventMutation{}

	mutationName, err := inputUI("MutationName", false)
	if err != nil {
		return mutation, "", err
	}
	mutationName = schema.SanitizeName(mutationName)

	inEventName, err := inputUI(fmt.Sprintf("=> %s.InEvent.Name", mutationName), false)
	if err != nil {
		return mutation, "", err
	}
	inEventName = schema.SanitizeName(inEventName)

	pkg, err := inputUI(fmt.Sprintf("=> %s.InEvent.Package", mutationName), false)
	if err != nil {
		return mutation, "", err
	}
	pkg = schema.SanitizePackageName(pkg)

	inEventPayload, err := inputUI(fmt.Sprintf("=> %s.InEvent.Payload", mutationName), true)
	if err != nil {
		return mutation, "", err
	}
	inEventPayload = schema.SanitizeName(inEventPayload)

	outEventName, err := inputUI(fmt.Sprintf("=> %s.OutEvent.Name", mutationName), false)
	if err != nil {
		return mutation, "", err
	}
	outEventName = schema.SanitizeName(outEventName)

	outEventPayload, err := inputUI(fmt.Sprintf("=> %s.OutEvent.Payload", mutationName), true)
	if err != nil {
		return mutation, "", err
	}
	outEventPayload = schema.SanitizeName(outEventPayload)

	opt, err := selectOptionsUI()
	if err != nil {
		return mutation, "", err
	}

	mutation.InEvent.Name = inEventName
	mutation.InEvent.Payload = inEventPayload
	mutation.OutEvent.Name = outEventName
	mutation.OutEvent.Payload = outEventPayload

	if opt == 1 {
		mutation.Create = schema.YesOp
	}
	if opt == 2 {
		mutation.Delete = schema.YesOp
	}

	return mutation, pkg, nil
}

func inputUI(label string, empty bool) (string, error) {
	prompt := promptui.Prompt{
		Label: label,
		Validate: func(s string) error {
			if empty {
				return nil
			}
			if len(s) < 3 {
				return fmt.Errorf("%s - too short", label)
			}
			return nil
		},
	}
	res, err := prompt.Run()
	if err != nil {
		return "", err
	}
	return res, nil
}

func selectMutationTypeUI() (int, error) {
	prompt := promptui.Select{
		Label: "Select the type of stream mutation",
		Items: []string{"CommandMutation", "EventMutation"},
	}
	idx, _, err := prompt.Run()
	if err != nil {
		return -1, err
	}
	return idx, nil
}

func selectOptionsUI() (int, error) {
	prompt := promptui.Select{
		Label: "Options",
		Items: []string{"None", "AllowCreate", "AllowDelete"},
	}
	idx, _, err := prompt.Run()
	if err != nil {
		return -1, err
	}
	return idx, nil
}

func nextStepUI() bool {
	prompt := promptui.Prompt{
		Label:     "Add more?",
		IsConfirm: true,
	}
	v, err := prompt.Run()
	if err != nil {
		return false
	}
	return v == "y"
}
