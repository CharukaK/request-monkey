# Request Files and Document Format Specification

## File Type Support
This specification applies to files with the `.req` extension.

## Document Structure

A request file consists of the following components:

- `*` signifies zero or more occurrences.
- `?` indicates an optional element.

The general document structure is as follows:

```
.
|
+-- variable-decl*
+-- request-decl*
```

### Comments
Comments can appear anywhere in the document. Each comment begins with a `#` symbol, and any text following `#` on the 
same line is ignored by the parser.

### Variable Declarations

**Format:**
```
@identifier = value
```

- **Identifiers:** An identifier must begin with an alphabetical character (A-Z, a-z) and may not start with numbers 
or special symbols.
- **Values:** Values can consist of any sequence of characters.

### Request Declaration

**Format:**
```
method url protocol-version?
headers*
payload
```

- **Method** specifies the HTTP request method (e.g., GET, POST).
- **URL** denotes the target endpoint.
- **Protocol Version** (optional) defines the HTTP version, such as HTTP/1.1 or HTTP/2.0.
- **Headers** (zero or more) define request headers.
- **Payload** contains the request body content.
