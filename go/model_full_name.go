/*
 * Фильмотека API
 *
 * API Фильмотеки
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// ФИО актера
type FullName struct {

	FirstName string `json:"first_name"`

	SecondName string `json:"second_name"`

	LastName string `json:"last_name,omitempty"`
}
