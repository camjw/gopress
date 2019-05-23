# Gopress

Run cypress tests - but not all at once!

## Installation

A bit involved. First you must install Go:
```
brew install go
```
Next, clone this repo into your `GOPATH` and then run `Make` from the root to install.

I will be adding this to homebrew to make it a bit easier to use.

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
				"a regexp matching files which should trigger a retesting"
			]
		},
		.
		.
		.
	]
}
```

then just run `gopress` in the command line to run all of the matching tests.

## Improvements

Currently, the output piping is all in black and white - not the nice colouring cypress provides.

## License

MIT
