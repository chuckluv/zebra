package main

import (
	"net/http"

	"github.com/project-safari/zebra/auth"
	"github.com/spf13/cobra"
)

func NewRegistration() *cobra.Command {
	registerCmd := &cobra.Command{
		Use:          "registration",
		Short:        "register for zebra",
		RunE:         registerReq,
		SilenceUsage: true,
		Args:         cobra.ExactArgs(1),
	}

	return registerCmd
}

func registerReq(cmd *cobra.Command, arg []string) error {
	cfgFile := cmd.Flag("config").Value.String()

	cfg, err := Load(cfgFile)
	if err != nil {
		return err
	}

	client, err := NewClient(cfg)
	if err != nil {
		return err
	}

	user := auth.UserType().Constructor()
	reqBody := &struct {
		Name     string            `json:"name"`
		Password string            `json:"password"`
		Email    string            `json:"email"`
		Key      *auth.RsaIdentity `json:"key"`
	}{
		Name:     cfg.User,
		Password: "",
		Email:    cfg.Email,
		Key:      cfg.Key,
	}

	resCode, err := client.Post("/register", reqBody, user)
	if resCode != http.StatusOK {
		return err
	}

	return nil
}
