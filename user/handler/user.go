package handler

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
	"go.temporal.io/sdk/client"
	"opencsg.com/csghub-server/api/httpbase"
	"opencsg.com/csghub-server/common/config"
	"opencsg.com/csghub-server/common/types"
	"opencsg.com/csghub-server/common/utils/common"
	apicomponent "opencsg.com/csghub-server/component"
	"opencsg.com/csghub-server/user/component"
	"opencsg.com/csghub-server/user/workflow"
	workflowCommon "opencsg.com/csghub-server/user/workflow/common"
)

type UserHandler struct {
	c                        component.UserComponent
	sc                       apicomponent.SensitiveComponent
	publicDomain             string
	EnableHTTPS              bool
	signinSuccessRedirectURL string
	config                   *config.Config
}

func NewUserHandler(config *config.Config) (*UserHandler, error) {
	h := &UserHandler{}
	var err error
	h.c, err = component.NewUserComponent(config)
	if err != nil {
		return nil, fmt.Errorf("failed to create user component: %w", err)
	}
	sc, err := apicomponent.NewSensitiveComponent(config)
	if err != nil {
		return nil, fmt.Errorf("error creating sensitive component:%w", err)
	}
	h.sc = sc
	domainParsedUrl, err := url.Parse(config.APIServer.PublicDomain)
	if err != nil {
		return nil, fmt.Errorf("failed to parse public domain '%s': %w", config.APIServer.PublicDomain, err)
	}
	h.publicDomain = domainParsedUrl.Hostname()
	h.EnableHTTPS = config.EnableHTTPS
	h.signinSuccessRedirectURL = config.User.SigninSuccessRedirectURL
	h.config = config
	return h, err
}

// CreateUser godoc
// @Security     ApiKey
// @Summary      Create a new user
// @Description  create a new user
// @Tags         User
// @Accept       json
// @Produce      json
// @Param        body   body  types.CreateUserRequest true "body"
// @Success      200  {object}  types.Response{data=database.User} "OK"
// @Failure      400  {object}  types.APIBadRequest "Bad request"
// @Failure      500  {object}  types.APIInternalServerError "Internal server error"
// @Router       /users [post]
// func (h *UserHandler) Create(ctx *gin.Context) {
// 	var req *types.CreateUserRequest
// 	if err := ctx.ShouldBindJSON(&req); err != nil {
// 		slog.Error("Bad request format", "error", err)
// 		httpbase.BadRequest(ctx, err.Error())
// 		return
// 	}

// 	slog.Debug("Creating user", slog.Any("req", req))
// 	user, err := h.c.Create(ctx, req)
// 	if err != nil {
// 		slog.Error("Failed to create user", slog.Any("error", err))
// 		httpbase.ServerError(ctx, err)
// 		return
// 	}

// 	slog.Info("Create user succeed", slog.String("user", user.Username))
// 	httpbase.OK(ctx, user)
// }

// UpdateUser godoc
// @Security     ApiKey
// @Summary      Update user. If change user name, should only send 'new_username' in the request body.
// @Description  update user
// @Tags         User
// @Accept       json
// @Produce      json
// @Param        id path string true "user identifier, could be username(depricated) or uuid"
// @Param        current_user  query  string true "current user"
// @Param        type query string false "type of identifier, uuid or username, default is username" Enums(uuid, username)
// @Param        body   body  types.UpdateUserRequest true "body"
// @Success      200  {object}  types.Response{} "OK"
// @Failure      400  {object}  types.APIBadRequest "Bad request"
// @Failure      500  {object}  types.APIInternalServerError "Internal server error"
// @Router       /user/{id} [put]
func (h *UserHandler) Update(ctx *gin.Context) {
	currentUser := httpbase.GetCurrentUser(ctx)
	var req *types.UpdateUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		slog.Error("Bad request format", "error", err)
		httpbase.BadRequest(ctx, err.Error())
		return
	}

	var err error
	_, err = h.sc.CheckRequestV2(ctx, req)
	if err != nil {
		slog.Error("failed to check sensitive request", slog.Any("error", err))
		httpbase.BadRequest(ctx, fmt.Errorf("sensitive check failed: %w", err).Error())
		return
	}

	id := ctx.Param("id")
	req.UUID = &id
	req.OpUser = currentUser
	err = h.c.UpdateByUUID(ctx.Request.Context(), req)
	if err != nil {
		slog.Error("Failed to update user by uuid", slog.Any("error", err), slog.String("uuid", *req.UUID), slog.String("current_user", currentUser), slog.Any("req", *req))
		httpbase.ServerError(ctx, err)
		return
	}

	slog.Info("Update user by uuid succeed", slog.String("uuid", *req.UUID), slog.String("current_user", currentUser))
	httpbase.OK(ctx, nil)
}

// DeleteUser godoc
// @Security     ApiKey
// @Summary      Delete user
// @Tags         User
// @Accept       json
// @Produce      json
// @Param        username path string true "username"
// @Param        current_user  query  string true "current user"
// @Param        body   body  types.UpdateUserRequest true "body"
// @Success      200  {object}  types.Response{} "OK"
// @Failure      400  {object}  types.APIBadRequest "Bad request"
// @Failure      500  {object}  types.APIInternalServerError "Internal server error"
// @Router       /user/{username} [delete]
func (h *UserHandler) Delete(ctx *gin.Context) {
	operator := httpbase.GetCurrentUser(ctx)
	userName := ctx.Param("username")

	// Check if operator can delete user
	isServerErr, err := h.c.CheckOperatorAndUser(ctx, operator, userName)
	if err != nil && isServerErr {
		httpbase.ServerError(ctx, fmt.Errorf("user cannot be deleted: %w", err))
		return
	}
	if err != nil && !isServerErr {
		httpbase.BadRequest(ctx, err.Error())
		return
	}

	// Check if user has organizations
	hasOrgs, err := h.c.CheckIfUserHasOrgs(ctx, userName)
	if err != nil {
		httpbase.ServerError(ctx, fmt.Errorf("failed to check if user has organzitions, error: %w", err))
		return
	}
	if hasOrgs {
		httpbase.BadRequest(ctx, "users who own organizations cannot be deleted")
		return
	}
	// Check if user has running or building deployments
	hasDeployments, err := h.c.CheckIffUserHasRunningOrBuildingDeployments(ctx, userName)
	if err != nil {
		httpbase.ServerError(ctx, fmt.Errorf("failed to check if user has deployments, error: %w", err))
		return
	}
	if hasDeployments {
		httpbase.BadRequest(ctx, "users who own deployments cannot be deleted")
		return
	}

	// Check if user has bills, Saas only
	hasBills, err := h.c.CheckIfUserHasBills(ctx, userName)
	if err != nil {
		httpbase.ServerError(ctx, fmt.Errorf("failed to check if user has bills, error: %w", err))
		return
	}
	if hasBills {
		httpbase.BadRequest(ctx, "users who own bills cannot be deleted")
		return
	}

	//start workflow to delete user
	workflowClient := workflow.GetWorkflowClient()
	workflowOptions := client.StartWorkflowOptions{
		TaskQueue: workflow.WorkflowUserDeletionQueueName,
	}

	we, err := workflowClient.ExecuteWorkflow(context.Background(), workflowOptions, workflow.UserDeletionWorkflow,
		workflowCommon.User{
			Username: userName,
			Operator: operator,
		},
		h.config,
	)
	if err != nil {
		httpbase.ServerError(ctx, fmt.Errorf("failed to start user deletion workflow, error: %w", err))
		return
	}

	slog.Info("start user deletion workflow", slog.String("workflow_id", we.GetID()), slog.String("userName", userName), slog.String("operator", operator))
	httpbase.OK(ctx, nil)
}

// GetUser godoc
// @Security     ApiKey
// @Summary      Get user info. Admin and the user self can see full info, other users can only see basic info.
// @Tags         User
// @Accept       json
// @Produce      json
// @Param        username path string true "username or uuid, defined by the query string 'type'"
// @Param        current_user  query  string false "current user"
// @Param 		 type query string false "path param is usernam or uuid, default to username" Enums(username, uuid)
// @Success      200  {object}  types.Response{data=types.User} "OK"
// @Failure      400  {object}  types.APIBadRequest "Bad request"
// @Failure      500  {object}  types.APIInternalServerError "Internal server error"
// @Router       /user/{username} [get]
func (h *UserHandler) Get(ctx *gin.Context) {
	visitorName := httpbase.GetCurrentUser(ctx)
	authType := httpbase.GetAuthType(ctx)
	userNameOrUUID := ctx.Param("username")
	useUUID := ctx.Query("type") == "uuid"
	var user *types.User
	var err error
	if authType == httpbase.AuthTypeApiKey {
		user, err = h.c.GetInternal(ctx, userNameOrUUID, useUUID)
	} else {
		user, err = h.c.Get(ctx, userNameOrUUID, visitorName, useUUID)
	}
	if err != nil {
		slog.Error("Failed to get user", slog.Any("error", err))
		httpbase.ServerError(ctx, err)
		return
	}

	slog.Info("Get user succeed", slog.String("userName", userNameOrUUID))
	httpbase.OK(ctx, user)
}

// GetUsers godoc
// @Security     ApiKey
// @Summary      Get users info. Only Admin
// @Tags         User
// @Accept       json
// @Produce      json
// @Param        search  query  string true "search"
// @Success      200  {object}  types.Response{data=[]types.User,total=int} "OK"
// @Failure      400  {object}  types.APIBadRequest "Bad request"
// @Failure      500  {object}  types.APIInternalServerError "Internal server error"
// @Router       /users [get]
func (h *UserHandler) Index(ctx *gin.Context) {
	visitorName := httpbase.GetCurrentUser(ctx)
	search := ctx.Query("search")
	per, page, err := common.GetPerAndPageFromContext(ctx)
	if err != nil {
		slog.Error("Failed to get per and page", slog.Any("error", err))
		httpbase.BadRequest(ctx, err.Error())
		return
	}
	users, count, err := h.c.Index(ctx, visitorName, search, per, page)
	if err != nil {
		slog.Error("Failed to get user", slog.Any("error", err))
		httpbase.ServerError(ctx, err)
		return
	}
	respData := gin.H{
		"data":  users,
		"total": count,
	}

	slog.Info("Get users succeed")
	httpbase.OK(ctx, respData)
}

func (h *UserHandler) Casdoor(ctx *gin.Context) {
	code := ctx.Query("code")
	state := ctx.Query("state")
	slog.Debug("get casdoor callback", slog.String("code", code), slog.String("state", state))

	jwtToken, signed, err := h.c.Signin(ctx.Request.Context(), code, state)
	if err != nil {
		slog.Error("Failed to signin", slog.Any("error", err), slog.String("code", code), slog.String("state", state))
		httpbase.ServerError(ctx, fmt.Errorf("failed to signin: %w", err))
		return
	}
	expire := jwtToken.ExpiresAt
	targetUrl := fmt.Sprintf("%s?jwt=%s&expire=%d", h.signinSuccessRedirectURL, signed, expire.Unix())
	ctx.Redirect(http.StatusMovedPermanently, targetUrl)
}

// GetEmailsInternal godoc
// @Security     ApiKey
// @Summary      Get all user emails for internal services
// @Description  Retrieve all user email addresses for internal services
// @Tags         User
// @Accept       json
// @Produce      json
// @Param        per query int false "per" default(50)
// @Param        page query int false "per page" default(1)
// @Success      200  {object}  types.Response{data=[]string,total=int} "OK"
// @Failure      400  {object}  types.APIBadRequest "Bad request"
// @Failure      500  {object}  types.APIInternalServerError "Internal server error"
// @Router       /internal/user/emails [get]
func (h *UserHandler) GetEmailsInternal(ctx *gin.Context) {
	per, page, err := common.GetPerAndPageFromContext(ctx)
	if err != nil {
		httpbase.BadRequest(ctx, err.Error())
		return
	}

	emails, count, err := h.c.GetEmailsInternal(ctx, per, page)
	if err != nil {
		slog.Error("Failed to get all user emails", slog.Any("error", err))
		httpbase.ServerError(ctx, err)
		return
	}

	httpbase.OKWithTotal(ctx, emails, count)
}

// GetUserUUIDs godoc
// @Security     ApiKey
// @Summary      Get user UUIDs
// @Description  Get user UUIDs
// @Tags         User
// @Accept       json
// @Produce      json
// @Param        per query int false "per" default(20)
// @Param        page query int false "per page" default(1)
// @Success      200  {object}  types.Response{data=[]string,total=int} "OK"
// @Failure      400  {object}  types.APIBadRequest "Bad request"
// @Failure 	 401  {object}  types.APIUnauthorized "Unauthorized"
// @Failure      500  {object}  types.APIInternalServerError "Internal server error"
// @Router       /user/user_uuids [get]
func (h *UserHandler) GetUserUUIDs(ctx *gin.Context) {
	per, page, err := common.GetPerAndPageFromContext(ctx)
	if err != nil {
		httpbase.BadRequest(ctx, err.Error())
		return
	}
	userUUIDs, total, err := h.c.GetUserUUIDs(ctx, per, page)
	if err != nil {
		httpbase.ServerError(ctx, err)
		return
	}
	respData := gin.H{
		"data":  userUUIDs,
		"total": total,
	}
	httpbase.OK(ctx, respData)
}
