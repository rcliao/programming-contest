# Model Design in Programming Contest

This note contains my thought toward on the design for models used in programming
contest.

A programming contest should have the following models:

* Contests (containing multiple teams, test questions, which year/term is it for)
* Questions (containing description and some sample input output)
* Submissions (user submitted program to a question)
* Teams (containing multiple users, which school they presented)
* Users (containing user email and some basic user information?)

## Models in detail (with attributes)

### Contest

Should contain:

* List of teams
* List of test questions
* Year/term
* Name
* Hosting location
* Description
* [Optional] Location

### Question

Should contain:

* Description
* Sample input
* Sample output
* [Optional] Attachments

### Submission

Should contain:

* Which team/person submitted
* Submitted program (source code file)
* Which question is it for
* Result (success, compile time error, runtime error, not correct answer ... etc.)
* [Optional] Feedback from Judge

### Team

Should contain:

* School name
* Team members (a list of user)
* Team name
* Description
* Contact information (email and/or phone)

### User

Should contain:

* Name
* Email
* Password
* User type (judge, parcipants, admin)
