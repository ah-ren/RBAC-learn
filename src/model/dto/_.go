package dto

// @Model         Result
// @Description   global response model
// @Property      code    integer true "status code"
// @Property      message string  true "status message"

//////////////////////////////////////////////////////////////////////////////////////
// Result

// @Model         Result<TokenDto>
// @Description   token response
// @Property      code    integer            true "status code"
// @Property      message string             true "status message"
// @Property      data    object(#_TokenDto) true "response data"

// @Model         Result<UserDto>
// @Description   user response
// @Property      code    integer           true "status code"
// @Property      message string            true "status message"
// @Property      data    object(#_UserDto) true "response data"

// @Model         Result<LoginDto>
// @Description   login response
// @Property      code    integer            true "status code"
// @Property      message string             true "status message"
// @Property      data    object(#_LoginDto) true "response data"

//////////////////////////////////////////////////////////////////////////////////////
// _Page

// @Model         Page<UserDto>
// @Description   user page
// @Property      total  integer          true "data count"
// @Property      page   integer          true "current page"
// @Property      limit  integer          true "page size"
// @Property      data   array(#_UserDto) true "response data"

// @Model         Page<PolicyDto>
// @Description   policy page
// @Property      total  integer            true "data count"
// @Property      page   integer            true "current page"
// @Property      limit  integer            true "page size"
// @Property      data   array(#_PolicyDto) true "response data"

//////////////////////////////////////////////////////////////////////////////////////
// PageResult

// @Model         Result<Page<UserDto>>
// @Description   user page response
// @Property      code      integer                true "status code"
// @Property      message   string                 true "status message"
// @Property      data      object(#Page<UserDto>) true "response data"

// @Model         Result<Page<PolicyDto>>
// @Description   policy page response
// @Property      code      integer                  true "status code"
// @Property      message   string                   true "status message"
// @Property      data      object(#Page<PolicyDto>) true "response data"
