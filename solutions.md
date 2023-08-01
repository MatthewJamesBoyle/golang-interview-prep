# Solutions

- The DB Credentials are hardcoded and also committed. 
- Connection string is hardcoded with credentials and also committed to the repo.
- The SQL query is susceptible to SQL injection (https://tip.golang.org/doc/database/sql-injection)
- DB connections are established in multiple places.
- Passwords are stored in plain text in the DB.
- Currently, usernames don't have a NULL constraint. This could be problematic, depending on your business requirements.
- The error handling could be much better. Sometimes we return the error, sometimes we panic.
- There's no test coverage at all.
- There's no authentication at all. Anybody can add a user.
