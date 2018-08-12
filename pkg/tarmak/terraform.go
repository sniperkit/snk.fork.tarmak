/*
Sniperkit-Bot
- Status: analyzed
*/

// Copyright Jetstack Ltd. See LICENSE for details.
package tarmak

import (
	"context"
	"errors"
	"fmt"

	"github.com/sniperkit/snk.fork.tarmak/pkg/tarmak/interfaces"
)

func (t *Tarmak) Terraform() interfaces.Terraform {
	return t.terraform
}

func (t *Tarmak) CmdTerraformPlan(args []string, ctx context.Context) error {
	t.cluster.Log().Info("validate steps")
	if err := t.Validate(); err != nil {
		return fmt.Errorf("failed to validate tarmak: %s", err)
	}

	t.cluster.Log().Info("verify steps")
	if err := t.Verify(); err != nil {
		return err
	}

	t.cluster.Log().Info("write SSH config")
	if err := t.writeSSHConfigForClusterHosts(); err != nil {
		return err
	}

	t.cluster.Log().Info("running plan")
	err := t.terraform.Plan(t.Cluster())
	if err != nil {
		return err
	}

	return nil
}

func (t *Tarmak) CmdTerraformApply(args []string, ctx context.Context) error {
	t.cluster.Log().Info("validate steps")
	if err := t.Validate(); err != nil {
		return fmt.Errorf("failed to validate tarmak: %s", err)
	}

	t.cluster.Log().Info("verify steps")
	if err := t.Verify(); err != nil {
		return err
	}

	t.cluster.Log().Info("write SSH config")
	if err := t.writeSSHConfigForClusterHosts(); err != nil {
		return err
	}

	t.cluster.Log().Info("running apply")
	// run terraform apply always, do not run it when in configuration only mode
	if !t.flags.Cluster.Apply.ConfigurationOnly {
		err := t.terraform.Apply(t.Cluster())
		if err != nil {
			return err
		}
	}

	// upload tar gz only if terraform hasn't uploaded it yet
	if t.flags.Cluster.Apply.ConfigurationOnly {
		err := t.Cluster().UploadConfiguration()
		if err != nil {
			return err
		}
	}

	// reapply config expect if we are in infrastructure only
	if !t.flags.Cluster.Apply.InfrastructureOnly {
		err := t.Cluster().ReapplyConfiguration()
		if err != nil {
			return err
		}
	}

	// wait for convergance in every mode
	err := t.Cluster().WaitForConvergance()
	if err != nil {
		return err
	}

	return nil
}

func (t *Tarmak) CmdTerraformDestroy(args []string, ctx context.Context) error {
	t.cluster.Log().Info("validate steps")
	if err := t.Validate(); err != nil {
		return fmt.Errorf("failed to validate tarmak: %s", err)
	}

	t.cluster.Log().Info("verify steps")
	if err := t.Verify(); err != nil {
		return err
	}

	t.cluster.Log().Info("write SSH config")
	if err := t.writeSSHConfigForClusterHosts(); err != nil {
		return err
	}

	t.cluster.Log().Info("running destroy")

	err := t.terraform.Destroy(t.Cluster())
	if err != nil {
		return err
	}
	return nil
}

func (t *Tarmak) CmdTerraformShell(args []string) error {
	if err := t.writeSSHConfigForClusterHosts(); err != nil {
		return err
	}

	err := t.terraform.Shell(t.Cluster())
	if err != nil {
		return err
	}
	return nil
}

func (t *Tarmak) verifyImageExists() error {
	images, err := t.Packer().List()
	if err != nil {
		return err
	}

	if len(images) == 0 {
		return errors.New("no images found")
	}

	return nil
}
