/*
 * Kings API
 *
 * This is a sample Kings Server based on the OpenAPI 3.0 specification.
 *
 * API version: 0.1.0
 * Contact: dev.anowak@gmail.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type King struct {
	ID int64 `json:"id,omitempty"`

	Name string `json:"name,omitempty"`

	City string `json:"city,omitempty"`

	House string `json:"house,omitempty"`
	// Range of years the king rule over
	Years string `json:"years,omitempty"`
}

type KingRaw struct {
	ID int64 `json:"id,omitempty"`

	Nm string `json:"nm,omitempty"`

	Cty string `json:"cty,omitempty"`

	Hse string `json:"hse,omitempty"`
	// Range of years the king rule over
	Yrs string `json:"yrs,omitempty"`
}
