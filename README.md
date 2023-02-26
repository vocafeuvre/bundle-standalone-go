# Standalone bundle using Go
PoC on how to bundle apps with their dependencies into a Go standalone binary, for cross-platform deployment.

## How apps or dependencies are stored in the bundle
1. Before starting, all apps and dependencies to be bundled should be in binary form.
2. In the build step, get all the apps and deps binaries and pass them thru go-bindata.
3. The final binary should have a reference to these apps and deps so that they can be included in it.

## How the bundle runs apps or dependencies
1. Extract binary file saved in binary and place it in a temp folder (with placeholder folder).
2. Execute it from there, but the current working directory should be where the binary was run.

## Usage
First, make the `gui` module ready for use by running these commands inside it:
````
npm ci
npm run build-all
````

It should output a `bin` directory inside. This `bin` directory contains the sample binary that we will bundle.

After this, go inside `standalone`, and do ``go install``. This should install the `go-bindata` command.

Generate the Go code that bundles the `gui` binary through this command: ``go-bindata -nocompress -o static.go ../gui/bin``

To test, do ``go run .``. This should generate an `output` executable file.