# Simple blogging server by Gin

## design
1. user register with user_name & password
2. user log-in with user_name & password
3. user log-out
4. logged in user can post blogs
5. display user's blogs on home page
6. display single article in its page

## routing
| usage         | routing       |   request_type |
| ------------- |:-------------:|:-------------: |
| homepage      | `/` | GET |
| user log-in   | `/user/login`  | POST |
| user log-out  | `/user/logout` | GET |
| user register  | `/user/logout` | GET |
| create user register  | `/user/register` | POST |
| edit a new article | `/article/create` | GET |
| create article | `/article/create` | POST |
| display article | `/article/view/:article_id`| GET|
