# Instagram-Dupe Backend

RestFull API for MongoDB using GoLang

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
POST:  /users :  To create an user <br>
_{  <br>
    "name": "jarun madhesh",<br>
    "email":"jarunmadhesh.gmail.com",<br>
    "password":"temporarypassword"<br>
}_<br> 
GET:  /users  : To list all the users. _Used for production and debugging._  <br>
GET:  /users/:userid :  To fetch the user by userID <br>
DELETE:  /users : To delete all the users._Used for production and debugging._ <br>
<br>
POST:  /posts : To create a post  <br>
_{<br>
    "accountid": "61618edfdb5a84faf07979a6",<br>
    "caption": "Sun is bright.",<br>
    "imageurl":"url1 goes here",<br>
    "posted_Timestamp" : "2012-10-31 15:50:13.793654 +0000 UTC"<br>
}_<br> 
GET:  /posts  : To get all the posts. _Used for production and debugging._  <br>
GET:  /posts/:postId :  To fetch the post by post ID  <br>
GET:  /posts/:userId/users :  To fetch all the posts posted by a particular user of the given userID  <br>
DELETE:  /posts : To delete all the posts by all the users.. _Used for production and debugging._ <br>
<br>



## packages used
1. context
2. fmt
3. log
4. net/http
5. time
6. crypto
7. crypto/aes
8. crypto/cipher
9. crypto/rand
10. encoding/base64
11. encoding/json
12. errors
13. io
14. github.com/julienschmidt/httprouter
15. go.mongodb.org/mongo-driver/bson
16.	go.mongodb.org/mongo-driver/bson/primitive
17.	go.mongodb.org/mongo-driver/mongo
18.	go.mongodb.org/mongo-driver/mongo/options
