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

- **Scope:** Variables are globally scoped, meaning they can be referenced from anywhere within the document.
- **Identifiers:** An identifier must begin with an alphabetical character (A-Z, a-z) and may not start with numbers 
or special symbols.
- **Values:** Values can consist of any sequence of characters.
- **Value Insertion:** Variables can be inserted into request declarations by wrapping the variable name in double 
braces (`{{ }}`). For example, `{{identifier}}` will be replaced by the value assigned to `identifier`.
- **Escaping Characters:** To use special characters like `{{`, `}}`, or `=` in variable values, escape them as 
`\{{`, `\}}`, or `\=`, respectively.

### Request Declaration

**Format:**
```
method url protocol-version?
headers*
payload
```

- **Method** specifies the HTTP request method (e.g., GET, POST).
- **URL** denotes the target endpoint.
- **Protocol Version** (optional) defines the HTTP version, such as HTTP/1.1 or HTTP/2.0. If omitted, the default 
protocol is assumed to be HTTP/1.1.
- **Headers** (zero or more) define request headers as `key: value`, each on a new line. There should be a blank line 
between the headers and the payload.
- **Payload** contains the request body content, which can be in any text format (e.g., JSON, XML, plain text) and 
follows the headers.

Variables can be used in the URL, headers, and payload sections by enclosing their names in `{{ }}`, allowing for 
dynamic insertion of values at runtime.

### Example

```plaintext
# Example of a .req file with variables and request declarations

@host = api.example.com
@contentType = application/json
@token = abc123

# Request 1
POST http://{{host}}/users HTTP/1.1
Authorization: Bearer {{token}}
Content-Type: {{contentType}}

{
    "name": "John Doe",
    "email": "john.doe@example.com"
}

# Request 2
GET http://{{host}}/users/{{userId}} HTTP/1.1
Authorization: Bearer {{token}}
```

