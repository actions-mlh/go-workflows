# Usage

1. Run `go run cmd/serve/main.go -port 5007`.  (port 5007 is hard-coded in `client/src/extension.ts` but can be changed there.)
2. Open the `client/` directory in VSCode (`code client` on the command line).
3. In VSCode, click `Run and Debug` (left side) -> `Launch Extension` (top, green play button).
4. In the new VSCode instance, create or open a YAML file.  You should see any errors underlined.  In the terminal running the Go server, you should see various outputs as well.

# Details

VSCode uses the Language Server Protocol for rendering errors, meaning that it sends information to a separate server through JSON and receives information that it uses to render the errors.  The client-side code is contained in `client/` using Typescript; use `npm run watch` or `npm run compile` to generate the Javascript code in `client/out/`.  The server code is split between `cmd/serve/main/go` and `server/`; the bulk of the logic is in the former, with the latter containing some structs and initialization-specific code.

For future development:
1. A more streamlined development environment and robust testing procedure.
2. Organize and refactor the Go server code in `cmd/` and `server/`.
3. More informative print statements in server.
4. Figure out how to deploy the extension to the VSCode store or whatever it is.  (i use emacs)
5. Nice coloring and formatting in `client/`.
