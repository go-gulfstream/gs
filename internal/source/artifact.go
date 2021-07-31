package source

import (
	"errors"
	"fmt"
	"go/ast"
	"strings"

	"github.com/go-gulfstream/gs/internal/strutil"
)

const (
	commandMutationInterfaceName      = "CommandMutation"
	commandMakeCommandControllersName = "MakeCommandControllers"
	eventMutationInterfaceName        = "EventMutation"
)

type Artifact struct {
	snapshot *Snapshot

	commandMutations       map[string]*commandMutation
	commandController      map[string]*commandController
	makeCommandControllers *makeControllers

	eventMutations map[string]*eventMutation
}

var parseBreak = errors.New("parseBreak")

func NewArtifact(s *Snapshot) *Artifact {
	return &Artifact{
		snapshot:          s,
		commandMutations:  make(map[string]*commandMutation),
		commandController: make(map[string]*commandController),
		eventMutations:    make(map[string]*eventMutation),
	}
}

func (a *Artifact) Scan() error {
	for _, parser := range []func() error{
		a.parseCommandMutationInterface,
		a.parseCommandController,
		a.parseMakeCommandControllers,
	} {
		if err := parser(); err != nil && err != parseBreak {
			return err
		}
	}
	return nil
}

func (a *Artifact) parseMakeCommandControllers() error {
	return a.snapshot.Walk(func(info FileInfo) (wer error) {
		info.Inspect(func(node ast.Node) bool {
			funcDecl, ok := node.(*ast.FuncDecl)
			if !ok {
				return true
			}
			if funcDecl.Name.Name == commandMakeCommandControllersName {
				a.makeCommandControllers = &makeControllers{
					filepath: info.Path(),
				}
				wer = parseBreak
				return false
			}
			return true
		})
		return
	})
}

func (a *Artifact) parseCommandController() error {
	return a.snapshot.Walk(func(info FileInfo) (wer error) {
		info.Inspect(func(node ast.Node) bool {
			funcDecl, ok := node.(*ast.FuncDecl)
			if !ok {
				return true
			}
			if len(funcDecl.Type.Params.List) != 1 {
				return true
			}
			ident := funcDecl.Type.Params.List[0].Type.(*ast.Ident)
			if ident.Name != commandMutationInterfaceName {
				return true
			}
			fn := strings.Replace(funcDecl.Name.Name, "CommandController", "", -1)
			fn = strutil.UcFirst(fn)
			spec, found := a.commandMutations[fn]
			if found {
				ctrl := &commandController{
					filepath: info.Path(),
					name:     fn,
				}
				spec.hasController = true
				spec.controller = ctrl
				a.commandController[fn] = ctrl
				return false
			}
			return true
		})
		return
	})
}

func (a *Artifact) parseCommandMutationInterface() error {
	return a.snapshot.Walk(func(info FileInfo) (wer error) {
		info.Inspect(func(node ast.Node) bool {
			typeSpec, ok := node.(*ast.TypeSpec)
			if !ok {
				return true
			}
			if typeSpec.Name.Name == commandMutationInterfaceName {
				ifaceSpec, ok := typeSpec.Type.(*ast.InterfaceType)
				if !ok {
					return true
				}
				for _, method := range ifaceSpec.Methods.List {
					fnType, ok := method.Type.(*ast.FuncType)
					if !ok {
						wer = fmt.Errorf("Artifact.parseCommandMutationInterface(%s) => assertion error filed.Type.(*ast.FuncType)", info.Path())
						return false
					}
					lenArgs := len(fnType.Params.List)
					lenRes := len(fnType.Results.List)
					if lenArgs < 4 {
						wer = fmt.Errorf("Artifact.parseCommandMutationInterface(%s) => invalid number of arguments for the command mutation %s. got %d, expected 4 or 5",
							info.Path(), method.Names[0].Name, lenArgs)
						return false
					}
					if lenRes == 0 {
						wer = fmt.Errorf("Artifact.parseCommandMutationInterface(%s) => invalid number of results for the command mutation %s. got %d, expected > 0",
							info.Path(), method.Names[0].Name, lenRes)
					}
					m := &commandMutation{
						filepath:          info.Path(),
						mutation:          method.Names[0].Name,
						hasCommandPayload: lenArgs == 5,
						hasEventPayload:   lenRes == 2,
					}
					a.commandMutations[method.Names[0].Name] = m
				}
				wer = parseBreak
				return false
			}
			return true
		})
		return
	})
}

type commandMutation struct {
	filepath          string
	mutation          string
	hasEventPayload   bool
	hasCommandPayload bool
	hasController     bool
	controller        *commandController
	makeControllers   *makeControllers
}

type commandController struct {
	filepath string
	name     string
}

type makeControllers struct {
	filepath string
}

type eventMutation struct {
	filepath           string
	mutation           string
	hasInEventPayload  bool
	hasOutEventPayload bool
}

type Command struct {
	Name    string
	Payload string
}

type Event struct {
	Name    string
	Payload string
}
