package menus

import (
	"fmt"
	"log"

	"github.com/dixonwille/wmenu/v5"
	auth_ui "github.com/needsomesleeptd/annotarer-tech-ui/tech_ui/utils/auth"
	"github.com/needsomesleeptd/annotater-core/models"
	"github.com/needsomesleeptd/annotater-core/service"
	auth_utils_adapter "github.com/needsomesleeptd/annotater-utils/pkg/authUtils"
	response "github.com/needsomesleeptd/http-server/lib/api"
)

func (m *Menu) SignInMenu(opt wmenu.Opt) error {
	client, ok := opt.Value.(ClientEntity)
	if !ok {
		log.Fatal("Could not cast option's value to ClientEntity")
	}
	var login string
	var passwd string
	fmt.Println("Enter login:")
	fmt.Scan(&login)
	fmt.Println("Enter password:")
	fmt.Scan(&passwd)
	jwt, err := auth_ui.SignIn(client.Client, login, passwd)
	if err != nil {
		return err
	}
	m.jwt = jwt
	payload, err := auth_utils_adapter.JWTTokenHandler{}.ParseToken(jwt, service.SECRET)
	if err != nil {
		return err
	}
	m.ID = payload.ID
	m.role = payload.Role
	switch m.role {
	case models.Sender:
		m.RunUserMenu(client.Client)
	case models.Controller:
		m.RunControllerMenu(client.Client)
	case models.Admin:
		m.RunAdminMenu(client.Client)
	}
	fmt.Println(response.StatusOK)
	return nil
}

func (m *Menu) SignUpMenu(opt wmenu.Opt) error {
	client, ok := opt.Value.(ClientEntity)
	if !ok {
		log.Fatal("Could not cast option's value to ClientEntity")
	}
	var login string
	var passwd string
	fmt.Println("Enter login:")
	fmt.Scan(&login)
	fmt.Println("Enter password:")
	fmt.Scan(&passwd)
	_, err := auth_ui.SignUp(client.Client, login, passwd)

	if err != nil {
		return err
	}
	fmt.Println(response.StatusOK)
	return nil
}
