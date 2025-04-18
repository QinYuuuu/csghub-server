package component

import (
	"context"
	"errors"
	"fmt"
	"log/slog"

	"opencsg.com/csghub-server/builder/git"
	"opencsg.com/csghub-server/builder/git/gitserver"
	"opencsg.com/csghub-server/builder/store/database"
	"opencsg.com/csghub-server/common/config"
	"opencsg.com/csghub-server/common/types"
)

type OrganizationComponent interface {
	FixOrgData(ctx context.Context, org *database.Organization) (*database.Organization, error)
	Create(ctx context.Context, req *types.CreateOrgReq) (*types.Organization, error)
	Index(ctx context.Context, username, search string, per, page int) ([]types.Organization, int, error)
	Get(ctx context.Context, orgName string) (*types.Organization, error)
	Delete(ctx context.Context, req *types.DeleteOrgReq) error
	Update(ctx context.Context, req *types.EditOrgReq) (*database.Organization, error)
}

func NewOrganizationComponent(config *config.Config) (OrganizationComponent, error) {
	c := &organizationComponentImpl{}
	c.orgStore = database.NewOrgStore()
	c.nsStore = database.NewNamespaceStore()
	c.userStore = database.NewUserStore()
	var err error
	c.gs, err = git.NewGitServer(config)
	if err != nil {
		newError := fmt.Errorf("fail to create git server,error:%w", err)
		slog.Error(newError.Error())
		return nil, newError
	}
	c.msc, err = NewMemberComponent(config)
	if err != nil {
		newError := fmt.Errorf("fail to create membership component,error:%w", err)
		slog.Error(newError.Error())
		return nil, newError
	}
	return c, nil
}

type organizationComponentImpl struct {
	orgStore  database.OrgStore
	nsStore   database.NamespaceStore
	userStore database.UserStore
	gs        gitserver.GitServer

	msc MemberComponent
}

func (c *organizationComponentImpl) FixOrgData(ctx context.Context, org *database.Organization) (*database.Organization, error) {
	user := org.User
	req := new(types.CreateOrgReq)
	req.Name = org.Name
	req.Nickname = org.Nickname
	req.Username = org.User.Username
	req.Description = org.Description
	err := c.gs.FixOrganization(req, *user)
	if err != nil {
		slog.Error("fix git org data has error", slog.Any("error", err))
	}
	// need to create roles for a new org before adding members
	err = c.msc.InitRoles(ctx, org)
	if err != nil {
		slog.Error("fix organization role has error", slog.String("error", err.Error()))
	}
	// org creator defaults to be admin role
	err = c.msc.SetAdmin(ctx, org, user)
	return org, err
}

func (c *organizationComponentImpl) Create(ctx context.Context, req *types.CreateOrgReq) (*types.Organization, error) {
	user, err := c.userStore.FindByUsername(ctx, req.Username)
	if err != nil {
		return nil, fmt.Errorf("failed to find user, error: %w", err)
	}

	es, err := c.nsStore.Exists(ctx, req.Name)
	if err != nil {
		return nil, err
	}
	if es {
		return nil, errors.New("the name already exists")
	}

	dbOrg, err := c.gs.CreateOrganization(req, user)
	if err != nil {
		return nil, fmt.Errorf("failed create git organization, error: %w", err)
	}
	dbOrg.Homepage = req.Homepage
	dbOrg.Logo = req.Logo
	dbOrg.OrgType = req.OrgType
	dbOrg.Verified = req.Verified
	namespace := &database.Namespace{
		Path:   dbOrg.Name,
		UserID: user.ID,
	}
	err = c.orgStore.Create(ctx, dbOrg, namespace)
	if err != nil {
		return nil, fmt.Errorf("failed create database organization, error: %w", err)
	}
	// need to create roles for a new org before adding members
	err = c.msc.InitRoles(ctx, dbOrg)
	if err != nil {
		return nil, fmt.Errorf("failed init roles for organization, error: %w", err)
	}
	// org creator defaults to be admin role
	err = c.msc.SetAdmin(ctx, dbOrg, &user)
	if err != nil {
		return nil, fmt.Errorf("failed set admin role for organization, error: %w", err)
	}

	org := &types.Organization{
		Name:     dbOrg.Name,
		Nickname: dbOrg.Nickname,
		Homepage: dbOrg.Homepage,
		Logo:     dbOrg.Logo,
		OrgType:  dbOrg.OrgType,
		Verified: dbOrg.Verified,
	}
	return org, err
}

func (c *organizationComponentImpl) Index(ctx context.Context, username, search string, per, page int) ([]types.Organization, int, error) {
	var (
		err    error
		total  int
		u      database.User
		dborgs []database.Organization
	)
	u, err = c.userStore.FindByUsername(ctx, username)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to find user, error: %w", err)
	}
	if u.CanAdmin() {
		dborgs, total, err = c.orgStore.Search(ctx, search, per, page)
		if err != nil {
			return nil, 0, fmt.Errorf("failed to get organizations for admin user, error: %w", err)
		}
	} else {
		dborgs, total, err = c.orgStore.GetUserOwnOrgs(ctx, username)
		if err != nil {
			return nil, 0, fmt.Errorf("failed to get organizations for owner, error: %w", err)
		}
	}
	var orgs []types.Organization
	for _, dborg := range dborgs {
		org := types.Organization{
			Name:     dborg.Name,
			Nickname: dborg.Nickname,
			Homepage: dborg.Homepage,
			Logo:     dborg.Logo,
			OrgType:  dborg.OrgType,
			Verified: dborg.Verified,
		}
		orgs = append(orgs, org)
	}
	return orgs, total, nil
}

func (c *organizationComponentImpl) Get(ctx context.Context, orgName string) (*types.Organization, error) {
	dborg, err := c.orgStore.FindByPath(ctx, orgName)
	if err != nil {
		return nil, fmt.Errorf("failed to get organizations by name, error: %w", err)
	}
	org := &types.Organization{
		Name:     dborg.Name,
		Nickname: dborg.Nickname,
		Homepage: dborg.Homepage,
		Logo:     dborg.Logo,
		OrgType:  dborg.OrgType,
		Verified: dborg.Verified,
	}
	return org, nil
}

func (c *organizationComponentImpl) Delete(ctx context.Context, req *types.DeleteOrgReq) error {
	r, err := c.msc.GetMemberRole(ctx, req.Name, req.CurrentUser)
	if err != nil {
		slog.Error("faild to get member role",
			slog.String("org", req.Name), slog.String("user", req.CurrentUser),
			slog.String("error", err.Error()))
	}
	if !r.CanAdmin() {
		return fmt.Errorf("current user does not have permission to edit the organization, current user: %s", req.CurrentUser)
	}
	err = c.gs.DeleteOrganization(req.Name)
	if err != nil {
		return fmt.Errorf("failed to delete git organizations, error: %w", err)
	}
	err = c.orgStore.Delete(ctx, req.Name)
	if err != nil {
		return fmt.Errorf("failed to delete database organizations, error: %w", err)
	}
	return nil
}

func (c *organizationComponentImpl) Update(ctx context.Context, req *types.EditOrgReq) (*database.Organization, error) {
	r, err := c.msc.GetMemberRole(ctx, req.Name, req.CurrentUser)
	if err != nil {
		slog.Error("faild to get member role",
			slog.String("org", req.Name), slog.String("user", req.CurrentUser),
			slog.String("error", err.Error()))
	}
	if !r.CanAdmin() {
		return nil, fmt.Errorf("current user does not have permission to edit the organization, current user: %s", req.CurrentUser)
	}
	org, err := c.orgStore.FindByPath(ctx, req.Name)
	if err != nil {
		return nil, fmt.Errorf("organization does not exists, error: %w", err)
	}

	if req.Nickname != nil {
		org.Nickname = *req.Nickname
	}
	if req.Logo != nil {
		org.Logo = *req.Logo
	}
	if req.Homepage != nil {
		org.Homepage = *req.Homepage
	}
	if req.Verified != nil {
		org.Verified = *req.Verified
	}
	if req.OrgType != nil {
		org.OrgType = *req.OrgType
	}
	err = c.orgStore.Update(ctx, &org)
	if err != nil {
		return nil, fmt.Errorf("failed to update database organization, error: %w", err)
	}

	//skip update git server
	if req.Nickname == nil && req.Description == nil {
		return &org, nil
	}
	var gitEditReq types.EditOrgReq
	gitEditReq.Name = org.Name
	gitEditReq.Nickname = &org.Nickname
	gitEditReq.Description = &org.Description
	_, err = c.gs.UpdateOrganization(&gitEditReq, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to update git organization, error: %w", err)
	}
	return &org, err
}
