## Context in go

- ## What is context?

  - In large applications, it can sometimes be helpful to get context of the request that's being sent
  - For example, if we have a web server serving http requests and it became quite popular, it would be helpful to know if a client disconnects before getting the response
    - if we didn't know about the disconnect, we might spend time downstream trying to get the resources and report back
    - but if we had context and knew about the client's disconnect when it happens, we can potentially save on resources

- https://www.digitalocean.com/community/tutorials/how-to-use-contexts-in-go

## Common use case

- imagine we have a third party api that might take some time to get a response back
- we might not want to wait x seconds - only 100ms or something like that
- we can set up a context to only wait for the ms, and return an error otherwise
