# HTML Link Parser

A Go package that parses HTML files and extracts all links (`<a>` tags) along with their text content. This exercise helps understand HTML parsing and DOM traversal in Go using the `golang.org/x/net/html` package.

## Description

This program provides functionality to:
- Parse HTML files and extract all hyperlinks
- Handle nested HTML elements within links
- Clean and format link text by removing extra whitespace
- Process multiple HTML files with different structures

## Features

- Extracts both the `href` attribute and text content from `<a>` tags
- Handles nested HTML elements (like `<strong>` tags) within links
- Removes unnecessary whitespace and newlines from link text
- Provides clean struct-based output


The parser returns a slice of `Link` structures:

## Testing

The package includes comprehensive tests covering various HTML structures and edge cases. Run the tests using:

```
go test
```

## Dependencies

- golang.org/x/net/html - For HTML parsing functionality

## Notes

This exercise is part of the Gophercises series, designed to help developers learn Go through practical exercises.

