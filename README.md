# LetterCloned

A homebrewed webapp inspired by Letterboxd
=======
# Front-end
# Backend

## Goal of README
- Layout the idea
- What we need to do for the idea
- What technologies we need/learn for the idea
	- MongoDB
	- htmx
- Time line to execute
    - End Date: May 1st, 2024
- How to distribute the work


## Blueprint
##### C Notes 4/22
- We'll use a MERN-like stack
	- MongoDB
	- Express.js
	- React
	- Node.js
- EXCEPT, we'll use Go for the backend
	- An MRG stack
      - MongoDB
      - React
      - Go

Database will be The Movie Database (TMDb), since that's where Letterboxd gets their data from

- Unit Test
- Deployment
The MVP will be a barebones application, where the main goal is to connect to the database and output information to the website frontend.
![Alt](https://repobeats.axiom.co/api/embed/4f6f093ec6298465efd1ec9625de28ffd2f39964.svg "Repobeats analytics image")

## Future Goals

### Playd

We'll use [IGDb](https://www.igdb.com/) for the video game database.
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