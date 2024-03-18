/*
 * Фильмотека API
 *
 * API Фильмотеки
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type Actor struct {

	Id string `json:"id,omitempty"`

	Name *FullName `json:"name,omitempty"`

	Gender *Gender `json:"gender,omitempty"`

	Movies []string `json:"movies,omitempty"`

	BirthDate string `json:"birthDate,omitempty"`
}