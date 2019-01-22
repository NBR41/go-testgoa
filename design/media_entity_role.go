package design

// package "design"
import (
	// Use . imports to enable the DSL
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var (
	rolePath     = "/roles"
	roleIDPath   = "/:role_id"
	attrRoleID   = func() { Attribute("role_id", Integer, "Unique Role ID", defIDConstraint) }
	attrRoleName = func() { Attribute("role_name", String, "Role Name (Author/Scenarist/Cartoonist)", defStringConstraint) }

	fRoleList = func() {
		// ok
		Response(OK, CollectionOf(RoleMedia))
		// class not found
		Response(NotFound)
		// Errors
		Response(InternalServerError)
		Response(ServiceUnavailable)
	}
)

//RoleMedia defines the media type used to render roles.
var RoleMedia = MediaType("application/vnd.role+json", func() {
	Description("A Role")

	Attributes(func() {
		attrRoleID()
		attrRoleName()
		attrHref()
		Required("role_id", "role_name", "href")
	})

	View("default", func() {
		Attribute("role_id")
		Attribute("role_name")
		Attribute("href")
	})

	View("link", func() {
		Attribute("role_id")
		Attribute("href")
	})
})

var _ = Resource("roles", func() {
	BasePath(rolePath)
	DefaultMedia(RoleMedia)

	Action("list", func() {
		Description("List roles")
		Routing(GET(""))
		// ok
		Response(OK, CollectionOf(RoleMedia))
		// Errors
		Response(InternalServerError)
		Response(ServiceUnavailable)
	})

	Action("show", func() {
		Description("Get role by id")
		Routing(GET(roleIDPath))
		Params(attrRoleID)
		// ok
		Response(OK)
		// role not found
		Response(NotFound)
		// Errors
		Response(InternalServerError)
		Response(ServiceUnavailable)
		Response(BadRequest, ErrorMedia)
	})

	Action("create", func() {
		Description("Create new role")
		Routing(POST(""))
		Payload(func() {
			attrRoleName()
			Required("role_name")
		})
		Security(JWTAuth)
		// unauthorized
		Response(Unauthorized)
		// OK
		Response(Created, "/roles/[0-9]+")
		// App error
		Response(UnprocessableEntity)
		// Errors
		Response(InternalServerError)
		Response(ServiceUnavailable)
		Response(BadRequest, ErrorMedia)
	})

	Action("update", func() {
		Description("Update role by id")
		Routing(PUT(roleIDPath))
		Params(attrRoleID)
		Payload(func() {
			attrRoleName()
			Required("role_name")
		})
		Security(JWTAuth)
		// Unauthorized
		Response(Unauthorized)
		// OK
		Response(NoContent)
		// NotFound
		Response(NotFound)
		// App error
		Response(UnprocessableEntity)
		// Errors
		Response(InternalServerError)
		Response(ServiceUnavailable)
		Response(BadRequest, ErrorMedia)
	})

	Action("delete", func() {
		Description("delete role by id")
		Routing(DELETE(roleIDPath))
		Params(attrRoleID)
		Security(JWTAuth)
		// Unauthorized
		Response(Unauthorized)
		// OK
		Response(NoContent)
		// NotFound
		Response(NotFound)
		// Errors
		Response(InternalServerError)
		Response(ServiceUnavailable)
		Response(BadRequest, ErrorMedia)
	})
})
