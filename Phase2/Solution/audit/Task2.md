
# Task 1 review
**Task done by:** Amirabbas Sherafatian, Masud Rana, Einar Linde \
**Reviewer:** Lauri Leiten \
**Was the stated problem/task solved in an acceptable manner?** \
The task was to create one unit test on the service layer and external tests for each functionality that is run with docker-compose.
The one unit test on the service layer that tests the functionality of a service layer component. The chosen component to test was Redis. 
So the external tests were understood as testing HTTP, Websocket and gRPC APIs.
All of the tests were ran using the autograder command  - for me the tests did not pass, BUT that probably has something to do with me using Windows. The autograder passes on Github and for my teammate Einar the tests also passed on his Linux machine. \
**What has been done well and why?** \
The tests have been well written, I especially liked that the external data (URLs for example) were read in from the environment variables, making testing in different environments and docker easier.
The tests for a certain service or interfaces are in separate folders and a newcomer to the project can clearly find the wanted tests. \
**What is not well implemented and why?**\
Didn't really see anything bad - maybe there should have also been tests for Mongo and Postgres services, but the task was solved in an acceptable manner, so that is just a suggestion. \
**How easy is it to understand?** \
Extremely easy, I was able to understand the tests after a few seconds of looking at them. \
**Are there any recommendations to simplify some of this task that the reviewer would like to share?** \
No, didn't really see any. \
**Anything else notable, encouraging, or funny?** \
As I mentioned earlier, loved the environment variable usage.
