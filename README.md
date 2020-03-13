# urpv
User Resource Protection &amp; Verification

This library contains all the parts to build a User Auth Service.
The following 3 Services are required to be able to do a full-cycle program;

## 1. Auth Server
This service will be required to authenticate client applications, and direct to login if required
Should be able to;
* Authorize Clients (/authorize)


## 2. Client Application



## 3. Resource Server
This service contains the data the client application would like to access
Should be able to;
* Validate Bearer Tokens (Middleware)
* Get Claims from Token (/info)



