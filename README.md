##### How to get this project running

1. Correctly configure your go toolchain follow

   1. https://go.dev/doc/install

2. Add GOPATH to to your path

   1. e.g `export PATH=$PATH:$(go env GOPATH)/bin` OR 

   2. ```
      export PATH=$PATH:$(go env GOPATH)/bin <---- Confirm this line in you profile!!!
      ```

   3. Verify your path is updated with `echo $PATH` you should see the GOPATH in the PATH.

   4. References: https://livebook.manning.com/book/go-web-programming/appendix/11

3. Clone this repository

4. cd into the root directory

5. run `./scripts/start.sh` to run the server

   1. Currently uses https://dev.to/komfysach/go-live-reload-using-air-40ll

##### How to track third party dependecies

https://github.com/go-modules-by-example/index/blob/master/010_tools/README.md
https://github.com/golang/go/wiki/Modules#how-can-i-track-tool-dependencies-for-a-module
https://github.com/golang/go/wiki/Modules#how-can-i-track-tool-dependencies-for-a-module

https://github.com/golang/go/wiki/Modules#how-can-i-track-tool-dependencies-for-a-module

##### Need to add GOPATH to your path for third party tools to work