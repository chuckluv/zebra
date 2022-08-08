package main

import (
	"log"
	"net/http"

	"github.com/project-safari/zebra/auth"
	"github.com/spf13/cobra"
)

func NewRegistration() *cobra.Command {
	registerCmd := &cobra.Command{ //nolint:exhaustivestruct,exhaustruct
		Use:          "registration",
		Short:        "register for zebra",
		RunE:         RegisterReq,
		SilenceUsage: true,
		Args:         cobra.ExactArgs(1),
	}

	return registerCmd
}

func RegisterReq(cmd *cobra.Command, arg []string) error {
	cfgFile := cmd.Flag("config").Value.String()
	log.Default().Printf("Flag Value: %v", cfgFile)

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
