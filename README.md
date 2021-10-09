# Instagram-Dupe Backend

RestFull APIs for MongoDB using GoLang

## To use it
Clone the code and set up MongoDB locally on your machine.
Run the main file with the command:
<br> _go run main.go_<br>

The database has 2 collections:
<br>1. users
<br>2. posts
<br><br>
Use postman to request APIs
### End points
POST:  /users :   <br>
_{  <br>
    "name": "jarun madhesh",<br>
    "email":"jarunmadhesh.gmail.com",<br>
    "password":"temporarypassword"<br>
}_<br> 
GET:  /users  :   <br>
GET:  /users/<userid> :   <br>
DELETE:  /users :  <br>
<br>
POST:  /posts :   <br>
_{<br>
    "accountid": "61618edfdb5a84faf07979a6",<br>
    "caption": "Sun is bright.",<br>
    "imageurl":"url1 goes here",<br>
    "posted_Timestamp" : "2012-10-31 15:50:13.793654 +0000 UTC"<br>
}_<br> 
GET:  /posts  :   <br>
GET:  /posts/<postId> :   <br>
GET:  /posts/<userId>/users :   <br>
DELETE:  /posts :  <br>
<br>



## packages used
1. mongo-driver <br>https://github.com/mongodb/mongo-go-driver
2. encoding/json
3. context
4. fmt
5. log
6. net/http
7. time
8. crypto
