# Golang Coding Challenge

#What was changed?

Hey Rackspace, I spent the time to revise some of the code. I didn't spend a ton of time doing revisions but hopefully you can see I understand how to use golang and packages to organize code, and not repeat myself. There is plenty more I could of did but I didn't go to overboard because I know this is just a simple coding challenge.

* Used Mysql (just more comfortable with it).
* I consolidated the responses into a centralized package (that would need to be revised based on needs).
* I seperated out db logic and handler logic, the handler should call the database functions and return the models.
* I passed in a context that way we don't have to repeat ourselves in every database functions.
* I added simple tests to make sure everything is running properly.
* I changed the router setup to use the web.Router (this is just a personal preference that i've gotten used to).
* Added comments and documentation throughout the application for go doc.
* Added a middleware (to setup the context, and also to show that I understand middleware and their use).
* Uncluttered main.go that way we can have room for other important things.
* Added a get route to the router.
* Added an update route to the router.




#Instructions

Hi! Welcome to the Golang coding challenge. Below is a set of instructions that must attempt to complete within 3 days. Fork the repo when you're ready and good luck! 😀

Within this repository, you will find a hastily thrown together application. It's a very basic, a simple To-Do API with the ability to create and list your to-dos.

Here are some things we need help with.

1. We have create and list features but lack the ability to update. Introduce a new PUT endpoint at `/todos/{todoID}` that receives a JSON body containing title and status. The feature should update the existing record and return a JSON body representing the _new_ state of the todo item.  
You may notice the lack of tests in the repo, maybe set a good example and add tests to your method if you have time. That way the other devs can copy-paste from your good example.  
Once done, go ahead and open a pull request again the repo.

2. As mentioned before, we have create and list already in place. The dev team was super excited because they knocked this out faster than anyone thought, maybe too fast.  
Feel free to open a github issue and point out some of those shortcuts. If you're feeling bold, Pull Requests are always welcome 😀.

# Setup

Within the repo you will find a docker-compose.yaml file. If you're familar with docker and docker-compose, great! You can get started by simply running `docker-compose up` and that will create an API and Postgres container for you.

If you are not familar with those tools, feel free to setup whatever environment you are comfortable working within. At the very least you will need a go environment to run your API and a SQL database. There is a todo_schema.sql file that will create a basic table and sequence to get you started. To run the API, simply run `go run main.go`. You will need a few variables within the app, so please make sure those are provided.

Environment Variables

* DB_USER
* DB_PASSWORD
* DB_NAME

Also, regardless of which environment you use, you'll need to install dependencies. For this repo, we're using [Dep](https://golang.github.io/dep/). A Gopkg toml and lock file has been provided, simply run `dep ensure`.
#   G o C o d i n g C h a l l e n g e 
 
 