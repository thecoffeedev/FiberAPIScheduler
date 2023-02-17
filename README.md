# Scheduler 

## Prerequisites

Docker must be installed on your system. If you don't have Docker installed, follow the instructions for your operating system here.

You should have basic knowledge of Docker and its concepts, such as containers and images.

## Installation

 - Clone the repository containing the Docker file
 - There are two parts to the scheduler, "main.go" and "scheduler.go"
 - Both should be connected to the same Redis Server, And run both the files

    go run main.go
    
    go run ./scheduler/scheduler.go

## How to send request


FiberAPI accepts only  POST and the json format for the payload should be as follows

       {
        "seconds":5,
        "url": "http://127.0.0.1:5000/",
        "payload": "{'seconds':'5', 'url':'http://127.0.0.1:3000/', 'payload':'woah','type':'GET'}",
        "type": "GET"
        }

Note : Make sure the pay load is enclosed in double quotes and inside is enclosed in single quotes
