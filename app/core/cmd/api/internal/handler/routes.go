// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	configdict "ark-admin-zero/app/core/cmd/api/internal/handler/config/dict"
	jobs "ark-admin-zero/app/core/cmd/api/internal/handler/jobs"
	loglogin "ark-admin-zero/app/core/cmd/api/internal/handler/log/login"
	sysdept "ark-admin-zero/app/core/cmd/api/internal/handler/sys/dept"
	sysjob "ark-admin-zero/app/core/cmd/api/internal/handler/sys/job"
	sysmenu "ark-admin-zero/app/core/cmd/api/internal/handler/sys/menu"
	sysprofession "ark-admin-zero/app/core/cmd/api/internal/handler/sys/profession"
	sysrole "ark-admin-zero/app/core/cmd/api/internal/handler/sys/role"
	sysuser "ark-admin-zero/app/core/cmd/api/internal/handler/sys/user"
	user "ark-admin-zero/app/core/cmd/api/internal/handler/user"
	"ark-admin-zero/app/core/cmd/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/login",
				Handler: user.LoginHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/login/captcha",
				Handler: user.GetLoginCaptchaHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/logout",
				Handler: user.LogoutHandler(serverCtx),
			},
		},
		rest.WithPrefix("/admin/user"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/info",
				Handler: user.GetUserInfoHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/permmenu",
				Handler: user.GetUserPermMenuHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/profile/info",
				Handler: user.GetUserProfileInfoHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/profile/update",
				Handler: user.UpdateUserProfileHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/password/update",
				Handler: user.UpdateUserPasswordHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/avatar/generate",
				Handler: user.GetGenerateAvatarHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.JwtAuth.AccessSecret),
		rest.WithPrefix("/admin/user"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.PermMenuAuth},
			[]rest.Route{
				{
					Method:  http.MethodGet,
					Path:    "/list",
					Handler: sysmenu.GetSysPermMenuListHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/add",
					Handler: sysmenu.AddSysPermMenuHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/delete",
					Handler: sysmenu.DeleteSysPermMenuHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/update",
					Handler: sysmenu.UpdateSysPermMenuHandler(serverCtx),
				},
			}...,
		),
		rest.WithJwt(serverCtx.Config.JwtAuth.AccessSecret),
		rest.WithPrefix("/admin/sys/perm/menu"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.PermMenuAuth},
			[]rest.Route{
				{
					Method:  http.MethodGet,
					Path:    "/list",
					Handler: sysrole.GetSysRoleListHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/add",
					Handler: sysrole.AddSysRoleHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/delete",
					Handler: sysrole.DeleteSysRoleHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/update",
					Handler: sysrole.UpdateSysRoleHandler(serverCtx),
				},
			}...,
		),
		rest.WithJwt(serverCtx.Config.JwtAuth.AccessSecret),
		rest.WithPrefix("/admin/sys/role"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.PermMenuAuth},
			[]rest.Route{
				{
					Method:  http.MethodGet,
					Path:    "/list",
					Handler: sysdept.GetSysDeptListHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/add",
					Handler: sysdept.AddSysDeptHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/delete",
					Handler: sysdept.DeleteSysDeptHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/update",
					Handler: sysdept.UpdateSysDeptHandler(serverCtx),
				},
			}...,
		),
		rest.WithJwt(serverCtx.Config.JwtAuth.AccessSecret),
		rest.WithPrefix("/admin/sys/dept"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.PermMenuAuth},
			[]rest.Route{
				{
					Method:  http.MethodGet,
					Path:    "/page",
					Handler: sysjob.GetSysJobPageHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/add",
					Handler: sysjob.AddSysJobHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/delete",
					Handler: sysjob.DeleteSysJobHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/update",
					Handler: sysjob.UpdateSysJobHandler(serverCtx),
				},
			}...,
		),
		rest.WithJwt(serverCtx.Config.JwtAuth.AccessSecret),
		rest.WithPrefix("/admin/sys/job"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.PermMenuAuth},
			[]rest.Route{
				{
					Method:  http.MethodGet,
					Path:    "/page",
					Handler: sysprofession.GetSysProfessionPageHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/add",
					Handler: sysprofession.AddSysProfessionHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/delete",
					Handler: sysprofession.DeleteSysProfessionHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/update",
					Handler: sysprofession.UpdateSysProfessionHandler(serverCtx),
				},
			}...,
		),
		rest.WithJwt(serverCtx.Config.JwtAuth.AccessSecret),
		rest.WithPrefix("/admin/sys/profession"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.PermMenuAuth},
			[]rest.Route{
				{
					Method:  http.MethodGet,
					Path:    "/page",
					Handler: sysuser.GetSysUserPageHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/add",
					Handler: sysuser.AddSysUserHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/delete",
					Handler: sysuser.DeleteSysUserHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/update",
					Handler: sysuser.UpdateSysUserHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/password/update",
					Handler: sysuser.UpdateSysUserPasswordHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/rdpj/info",
					Handler: sysuser.GetSysUserRdpjInfoHandler(serverCtx),
				},
			}...,
		),
		rest.WithJwt(serverCtx.Config.JwtAuth.AccessSecret),
		rest.WithPrefix("/admin/sys/user"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.PermMenuAuth},
			[]rest.Route{
				{
					Method:  http.MethodGet,
					Path:    "/list",
					Handler: configdict.GetConfigDictListHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/data/page",
					Handler: configdict.GetConfigDictPageHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/add",
					Handler: configdict.AddConfigDictHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/delete",
					Handler: configdict.DeleteConfigDictHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/update",
					Handler: configdict.UpdateConfigDictHandler(serverCtx),
				},
			}...,
		),
		rest.WithJwt(serverCtx.Config.JwtAuth.AccessSecret),
		rest.WithPrefix("/admin/config/dict"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.PermMenuAuth},
			[]rest.Route{
				{
					Method:  http.MethodGet,
					Path:    "/page",
					Handler: loglogin.GetLogLoginPageHandler(serverCtx),
				},
			}...,
		),
		rest.WithJwt(serverCtx.Config.JwtAuth.AccessSecret),
		rest.WithPrefix("/admin/log/login"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.PermMenuAuth},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/queue/addqueue",
					Handler: jobs.SendHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/get/:id",
					Handler: jobs.GetHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/add",
					Handler: jobs.CreateHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/update",
					Handler: jobs.UpdateHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/delete",
					Handler: jobs.DeleteHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/search",
					Handler: jobs.SearchHandler(serverCtx),
				},
			}...,
		),
		rest.WithJwt(serverCtx.Config.JwtAuth.AccessSecret),
		rest.WithPrefix("/admin/jobs"),
	)
}
