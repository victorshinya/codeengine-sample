# Sample code for IBM Cloud Code Engine

This repository has sample codes in Golang and Node.js to be used as starting point when you need to learn how each strategy works in IBM Cloud Code Engine: [Dockerfile](https://docs.docker.com/engine/reference/builder) and [Paketo Buildpacks](https://paketo.io). Most samples will be relatively small and you can improve them to learn more during your journey.

## Getting started

You can download the repository and access each sample in its respective folder.

```sh
git clone https://github.com/victorshinya/go-codeengine-sample.git
```

## Samples

The list of samples are grouped by the category of functionality that is demonstrating.

### Source-to-Image

- [s2i-dockerfile](s2i-dockerfile)
One of the Code Engine strategy to build an application from a git repository is by building a Dockerfile and then push it to a private repository (IBM Cloud Container Registry), and then deploy the application using that image.

- [s2i-paketo-buildpacks](s2i-paketo-buildpacks)
One of the Code Engine strategy to build an application from a git repository is by building using a Paketo Buildpacks - to identify the source code's language and build with the correct language in a pre-built environment -, push it to a private repository (IBM Cloud Container Registry), and then deploy the application using that image.

## LICENSE

Copyright 2021 Victor Shinya

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
