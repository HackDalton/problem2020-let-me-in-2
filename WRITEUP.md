# HackDalton: Let Me In 2 (Writeup)

> Warning! There are spoilers ahead

Upon loading the website you are presented with a simple login screen. 

You can take a SQL Injection approach to this problem. If we type `'` into the username field, we get the following error message:

> An error occured executing query:
> SELECT id FROM users WHERE username = ''' AND password = '';
> 
> unrecognized token: "''' AND password = '';"

We can use this error message to figure out the SQL being used to check the login:

```sql
SELECT id FROM users WHERE username = '' AND password = '';
```

We can then begin to craft a username that would allow us to view the account.

We know that the username is `admin`, so the username must begin with that; however, we don't want the password to be checked, so we should comment out that section of the query. We can use this username:

```
admin';--
```

This would be then be injected:
```sql
SELECT id FROM users WHERE username = 'admin';-- AND password = '';
```

SQL uses `--` as a comment, so nothing after the username field would be parsed, and therefore the password would not be checked.