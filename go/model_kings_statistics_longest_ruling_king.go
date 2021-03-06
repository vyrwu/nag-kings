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

type KingsStatisticsLongestRulingKing struct {
	King []King `json:"king,omitempty"`

	YearsRuled int `json:"yearsRuled,omitempty"`
}
