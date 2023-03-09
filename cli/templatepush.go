package cli

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"

	"github.com/briandowns/spinner"
	"golang.org/x/xerrors"

	"github.com/coder/coder/cli/clibase"
	"github.com/coder/coder/cli/cliui"
	"github.com/coder/coder/coderd/database"
	"github.com/coder/coder/codersdk"
	"github.com/coder/coder/provisionersdk"
)

// templateUploadFlags is shared by `templates create` and `templates push`.
type templateUploadFlags struct {
	directory string
}

func (pf *templateUploadFlags) option() clibase.Option {
	currentDirectory, _ := os.Getwd()
	return clibase.Option{
		Name:          "directory",
		Flag:          "directory",
		FlagShorthand: "d",
		Description:   "Specify the directory to create from, use '-' to read tar from stdin",
		Default:       currentDirectory,
		Value:         clibase.StringOf(&pf.directory),
	}
}

func (pf *templateUploadFlags) stdin() bool {
	return pf.directory == "-"
}

func (pf *templateUploadFlags) upload(inv *clibase.Invokation, client *codersdk.Client) (*codersdk.UploadResponse, error) {
	var content io.Reader
	if pf.stdin() {
		content = inv.Stdin
	} else {
		prettyDir := prettyDirectoryPath(pf.directory)
		_, err := cliui.Prompt(inv, cliui.PromptOptions{
			Text:      fmt.Sprintf("Upload %q?", prettyDir),
			IsConfirm: true,
			Default:   cliui.ConfirmYes,
		})
		if err != nil {
			return nil, err
		}

		pipeReader, pipeWriter := io.Pipe()
		go func() {
			err := provisionersdk.Tar(pipeWriter, pf.directory, provisionersdk.TemplateArchiveLimit)
			_ = pipeWriter.CloseWithError(err)
		}()
		defer pipeReader.Close()
		content = pipeReader
	}

	spin := spinner.New(spinner.CharSets[5], 100*time.Millisecond)
	spin.Writer = inv.Stdout
	spin.Suffix = cliui.Styles.Keyword.Render(" Uploading directory...")
	spin.Start()
	defer spin.Stop()

	resp, err := client.Upload(inv.Context(), codersdk.ContentTypeTar, bufio.NewReader(content))
	if err != nil {
		return nil, xerrors.Errorf("upload: %w", err)
	}
	return &resp, nil
}

func (pf *templateUploadFlags) templateName(args []string) (string, error) {
	if pf.stdin() {
		// Can't infer name from directory if none provided.
		if len(args) == 0 {
			return "", xerrors.New("template name argument must be provided")
		}
		return args[0], nil
	}

	name := filepath.Base(args[0])
	if len(args) > 0 {
		name = args[0]
	}
	return name, nil
}

func (r *RootCmd) templatePush() *clibase.Cmd {
	var (
		versionName     string
		provisioner     string
		parameterFile   string
		variablesFile   string
		variables       []string
		alwaysPrompt    bool
		provisionerTags []string
		uploadFlags     templateUploadFlags
	)
	client := new(codersdk.Client)
	cmd := &clibase.Cmd{
		Use:   "push [template]",
		Short: "Push a new template version from the current directory or as specified by flag",
		Middleware: clibase.Chain(
			clibase.RequireRangeArgs(0, 1),
			r.UseClient(client),
		),
		Handler: func(inv *clibase.Invokation) error {
			organization, err := CurrentOrganization(inv, client)
			if err != nil {
				return err
			}

			name, err := uploadFlags.templateName(inv.Args)
			if err != nil {
				return err
			}

			template, err := client.TemplateByName(inv.Context(), organization.ID, name)
			if err != nil {
				return err
			}

			resp, err := uploadFlags.upload(inv, client)
			if err != nil {
				return err
			}

			tags, err := ParseProvisionerTags(provisionerTags)
			if err != nil {
				return err
			}

			job, _, err := createValidTemplateVersion(inv, createValidTemplateVersionArgs{
				Name:            versionName,
				Client:          client,
				Organization:    organization,
				Provisioner:     database.ProvisionerType(provisioner),
				FileID:          resp.ID,
				ParameterFile:   parameterFile,
				VariablesFile:   variablesFile,
				Variables:       variables,
				Template:        &template,
				ReuseParameters: !alwaysPrompt,
				ProvisionerTags: tags,
			})
			if err != nil {
				return err
			}

			if job.Job.Status != codersdk.ProvisionerJobSucceeded {
				return xerrors.Errorf("job failed: %s", job.Job.Status)
			}

			err = client.UpdateActiveTemplateVersion(inv.Context(), template.ID, codersdk.UpdateActiveTemplateVersion{
				ID: job.ID,
			})
			if err != nil {
				return err
			}

			_, _ = fmt.Fprintf(inv.Stdout, "Updated version at %s!\n", cliui.Styles.DateTimeStamp.Render(time.Now().Format(time.Stamp)))
			return nil
		},
	}

	cmd.Options = []clibase.Option{
		{
			Name:          "provisioner",
			Flag:          "test.provisioner",
			FlagShorthand: "p",
			Description:   "Customize the provisioner backend",
			Default:       "terraform",
			Value:         clibase.StringOf(&provisioner),
			// This is for testing!
			Hidden: true,
		},
		{
			Name:          "parameter-file",
			Flag:          "parameter-file",
			FlagShorthand: "f",
			Description:   "Specify a file path with parameter values.",
			Value:         clibase.StringOf(&parameterFile),
		},
		{
			Name:          "variables-file",
			Flag:          "variables-file",
			FlagShorthand: "f",
			Description:   "Specify a file path with values for Terraform-managed variables.",
			Value:         clibase.StringOf(&variablesFile),
		},
		{
			Name:        "variable",
			Flag:        "variable",
			Description: "Specify a set of values for Terraform-managed variables.",
			Value:       clibase.StringsOf(&variables),
		},
		{
			Name:          "provisioner-tag",
			Flag:          "provisioner-tag",
			FlagShorthand: "t",
			Description:   "Specify a set of tags to target provisioner daemons.",
			Value:         clibase.StringsOf(&provisionerTags),
		},
		{
			Name:        "name",
			Flag:        "name",
			Description: "Specify a name for the new template version. It will be automatically generated if not provided.",
			Value:       clibase.StringOf(&versionName),
		},
		{
			Name:        "always-prompt",
			Flag:        "always-prompt",
			Description: "Always prompt all parameters. Does not pull parameter values from active template version",
			Value:       clibase.BoolOf(&alwaysPrompt),
		},
		cliui.SkipPromptOption(),
		uploadFlags.option(),
	}
	return cmd
}
