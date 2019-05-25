# Gopress

Run cypress tests - but not all at once!

## Installation

This project requires Go and Go modules to be enabled:
```
brew install go
echo 'export GO111MODULES=on' >> ~/.bash_profile
source ~/.bash_profile
```
Next, clone this repo and then run `make` from the root of the repo to install.

## Usage

Create a `gopress.json` file at the root of the repo with the following structure:

```
{
	"directory": "the directory your test files live in i.e. cypress/integration",
	"extension": "extension for the test files i.e. .feature",
	"basebranch": "the branch you want to check for diffs against i.e. origin/develop",
	"tests": [
		{
			"testfile": "the name of your test file i.e. account_page",
			"regexes": [
				"a regexp matching files which should trigger a retesting",
				.
				.
			]
		},
		.
		.
		.
	]
}
```

then just run `gopress` in the command line to run all of the matching tests.

You can add more than one regexp, just so that you don't have to write long gnarly rexexps.

## Improvements

Currently, the output piping is all in black and white - not the nice colouring cypress provides.

## License

MIT
