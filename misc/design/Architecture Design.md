	Components and Design Principles
# 1. Technology
- HTMX [https://hypermedia.systems/book/contents/]
- GoLang
	- backend interfacing
- Air
	- live reload
- Docker
	- microservices
# 2. Frontend
- Templates
- Interactivity
	- HTMX
	- REACT
- 
# 3. Backend
- Logging
- Error Handling
- 
# 4. Database

# 5. API
- We'll use [TMDb](https://www.themoviedb.org/) for our API calls
	- [API](https://developer.themoviedb.org/docs/faq)
		- [Getting Started](https://developer.themoviedb.org/docs/getting-started)
		- [API Reference](https://developer.themoviedb.org/reference/intro/getting-started)
		- Useful Pages:
			- [Movie Details](https://developer.themoviedb.org/reference/movie-details)
			- [Append to Response](https://developer.themoviedb.org/docs/append-to-response)
			- Security:
			- [Examples]
				- Use [Postman] for testing requests locally
			- [Request Endpoints]
			- [Reference]
		- Constraints:
			- 50 API requests per second ([source](https://developer.themoviedb.org/docs/rate-limiting))
- A [TMDb Go library](https://github.com/ryanbradynd05/go-tmdb/) has already been made
	- [Common workflow](https://developer.themoviedb.org/docs/search-and-query-for-details) to query data
- [Daily Data Updates](https://developer.themoviedb.org/docs/daily-id-exports) that will be useful for a cron job to do.
# 6. Version Control
# 7. Build Projects
- Set up
	- Github
	- Server
	- database
	- APIs
	- schemas
	- Style Skeleton
- Create Post 


# EX. 1 - Playd goals
- Playd: We'll use [IGDb](https://www.igdb.com/) for the video game database.
	  - [API](https://www.igdb.com/api)
		  - [Getting Started](https://api-docs.igdb.com/#getting-started)
		  - Useful Pages:
			  - [Webhooks](https://api-docs.igdb.com/#webhooks)
			  - Security: [CORS](https://api-docs.igdb.com/#cors-proxy)
			  - [Examples](https://api-docs.igdb.com/#examples)
				  - Use [Postman](https://www.postman.com/) for testing requests locally
			  - [Request Endpoints](https://api-docs.igdb.com/#endpoints)
			  - Relational database preferred: [Collection Relation](https://api-docs.igdb.com/#collection-relation)
				  - [PostgreSQL](https://www.postgresql.org/)
					  - [Docs](https://www.postgresql.org/docs/)
				  - [SQLite](https://www.sqlite.org/)
					  - [Docs](https://www.sqlite.org/docs.html)
		  - [Reference](https://api-docs.igdb.com/#reference)