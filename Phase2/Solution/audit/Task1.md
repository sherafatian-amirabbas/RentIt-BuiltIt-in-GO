
# Task 1 review
**Task done by:** Amirabbas Sherafatian, Masud Rana, Lauri Leiten, Einar Linde \
**Reviewer:** Einar Linde \
**Was the stated problem/task solved in an acceptable manner?** \
The task was to create backend api application for rentit. It needed to have postgres, mongo and redis for storage. And it should have http, websocket and gPRC interfaces for retrieving data. Tested all those requirements and it worked. So task is solved. It does exactly that what requirements document said and because of that it is ready for potential client.  \
**What has been done well and why?** \
File structure is done well. Every file is in suitable folder. There isn't excess files. There is only one Dockerfile and docker-compose file. In last homework we had multiple dockerfiles and it was hassle to work with that because you had to exlicitly say what docker-compose or dockerfile you wanted to run. Packaging is good. There is separate package for each service. OOP good. Also there is used dependency injection to inject repository and services. Application uses environment variable for database connection that is good because then it is easy to change databases when needed.  \
**What is not well implemented and why?** \
There isn't mongoDb init script. But because it is not required to have some initial values in db everything is good.
There is some places in code that hides errors. For example in file pkg/cache/redis_client.go function GetPlant just disregards error. Because of that it is hard to debug that. It should log error into console or smth. 
Also there is a lot of duplication especially when handling errors. Don't have many suggestions for that. 
Because we use 2 separate db there should be some unified place for business logic. For example query that checks if plant is available is duplicated multiple times. \
**How easy is it to understand?** \
It is easy to understand . Code is pretty much self explanatory. There some comments that help to understand optimization/tricky parts. \
**Are there any recommendations to simplify some of this task that the reviewer would like to share?** \
Do not overthink \
**Anything else notable, encouraging, or funny?** \
Dockerfiles looks excepional. It looks better than it was in hw1. One dockerfile and docker-compose file for everything. 
