### Build

Changing the version in .docker-hub-tag will cause a new docker image to be published by the CI/CD pipeline, once commited to the master branch.

All builds that doesn't have changes to the .docker-hub-tag will also be built and deployed to the [ews-demo.r59q.com](https://ews-demo.r59q.com/swagger/index.html) site, regardless of versioning.