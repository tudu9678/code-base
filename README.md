# code-base

1. Set permission for sh file
   chmod +x ./deploy
2. Run Docker compose to init Postgres DB
   cd deploy
   make setup start
   make run_script
3. Run project
   make go_start

4. Test API :
   Test API Register :
   curl --location 'localhost:8080/v1/users' \
   --header 'Content-Type: application/json' \
   --data-raw '{
   "email" : "test3411@mail.com",
   "userName" : "hungtest3414",
   "phoneNumber" : "",
   "password" : "password1" ,
   "fullName" : "hung tien le " ,
   "dob" : "2020-10-10"
   }'

   Test API Login :
   curl --location 'localhost:8080/v1/users/login' \
    --header 'Content-Type: application/json' \
    --data '{
   "userName" : "hungtest3414",  
    "password" : "password1"
   }'
