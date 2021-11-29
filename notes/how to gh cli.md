# For when it's time to incorporate this project into the gh cli tool

https://github.com/jjayeon/cli/commit/c03dd40769622389827bca6c46440cdec7d75bbf

The above commit shows an example of how we can incorporate the linter into the gh cli.  As a hack, I just copied the existing codebase into the directory with all the workflow commands; we should look elsewhere into submodules or imports for a more modular way to merge the codebases.

The key files are `pkg/cmd/workflow/lint/lint.go` and `pkg/cmd/workflow/workflow.go`.  The former contains the code that actually makes up the `lint` command; it's mostly copied from the `view` command, then hacked to use our linting function instead (line 153 onward).  If this code is used directly, this file would definitely require some cleanup from someone more familiar with Cobra than I am.  The latter is the controller for the workflow command as a whole, which just has two added lines to incorporate the `lint` command.

The other files are just copied directly from the codebase.  Do note that these files are from a previous version and don't reflect our most recent work; they're just there as an example of how the final merge might work.  Please use a more intelligent linking mechanism before anything hits the gh cli codebase!
