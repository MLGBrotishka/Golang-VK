/*
 * Фильмотека API
 *
 * API Фильмотеки
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type SortBy string

// List of SortBy
const (
	TITLE SortBy = "title"
	RATING SortBy = "rating"
	RELEASE_DATE SortBy = "release_date"
)
