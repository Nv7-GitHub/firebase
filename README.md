# Firebase
The firebase library supplies lightweight, easy-to-use bindings to the Firebase REST HTTP API. It was created because of a lack of all-in-one packages that allowed Authentication, and Database management.

## Installation
You can install the latest version of this library with 
```bash
go get -v -u github.com/Nv7-Github/firebase
```

## Q & A
### Why not use the firebase-admin-sdk?
The Firebase admin SDK, as said in the name, is an Admin SDK. It doesn't allow authenticating with passwords, only changing accounts. In addition, it is not very lightweight. On my laptop (Mac OS), as of when this was written, programs using Only the Firebase SDK were around 21.6 MB. On the other hand, programs written using this library are around 4.3 MB.
