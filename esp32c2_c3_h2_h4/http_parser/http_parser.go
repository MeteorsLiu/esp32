package http_parser

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

const HTTP_PARSER_VERSION_MAJOR = 2
const HTTP_PARSER_VERSION_MINOR = 7
const HTTP_PARSER_VERSION_PATCH = 0
const HTTP_PARSER_STRICT = 1

type HttpParser struct {
	Unused [8]uint8
}

type HttpParserSettings struct {
	Unused [8]uint8
}

// llgo:type C
type HttpDataCb func(*HttpParser, *c.Char, c.SizeT) c.Int

// llgo:type C
type HttpCb func(*HttpParser) c.Int
type HttpMethod c.Int

const (
	HTTP_DELETE      HttpMethod = 0
	HTTP_GET         HttpMethod = 1
	HTTP_HEAD        HttpMethod = 2
	HTTP_POST        HttpMethod = 3
	HTTP_PUT         HttpMethod = 4
	HTTP_CONNECT     HttpMethod = 5
	HTTP_OPTIONS     HttpMethod = 6
	HTTP_TRACE       HttpMethod = 7
	HTTP_COPY        HttpMethod = 8
	HTTP_LOCK        HttpMethod = 9
	HTTP_MKCOL       HttpMethod = 10
	HTTP_MOVE        HttpMethod = 11
	HTTP_PROPFIND    HttpMethod = 12
	HTTP_PROPPATCH   HttpMethod = 13
	HTTP_SEARCH      HttpMethod = 14
	HTTP_UNLOCK      HttpMethod = 15
	HTTP_BIND        HttpMethod = 16
	HTTP_REBIND      HttpMethod = 17
	HTTP_UNBIND      HttpMethod = 18
	HTTP_ACL         HttpMethod = 19
	HTTP_REPORT      HttpMethod = 20
	HTTP_MKACTIVITY  HttpMethod = 21
	HTTP_CHECKOUT    HttpMethod = 22
	HTTP_MERGE       HttpMethod = 23
	HTTP_MSEARCH     HttpMethod = 24
	HTTP_NOTIFY      HttpMethod = 25
	HTTP_SUBSCRIBE   HttpMethod = 26
	HTTP_UNSUBSCRIBE HttpMethod = 27
	HTTP_PATCH       HttpMethod = 28
	HTTP_PURGE       HttpMethod = 29
	HTTP_MKCALENDAR  HttpMethod = 30
	HTTP_LINK        HttpMethod = 31
	HTTP_UNLINK      HttpMethod = 32
)

type HttpParserType c.Int

const (
	HTTP_REQUEST  HttpParserType = 0
	HTTP_RESPONSE HttpParserType = 1
	HTTP_BOTH     HttpParserType = 2
)

type Flags c.Int

const (
	F_CHUNKED               Flags = 1
	F_CONNECTION_KEEP_ALIVE Flags = 2
	F_CONNECTION_CLOSE      Flags = 4
	F_CONNECTION_UPGRADE    Flags = 8
	F_TRAILING              Flags = 16
	F_UPGRADE               Flags = 32
	F_SKIPBODY              Flags = 64
	F_CONTENTLENGTH         Flags = 128
)

type HttpErrno c.Int

const (
	HPE_OK                        HttpErrno = 0
	HPE_CB_message_begin          HttpErrno = 1
	HPE_CB_url                    HttpErrno = 2
	HPE_CB_header_field           HttpErrno = 3
	HPE_CB_header_value           HttpErrno = 4
	HPE_CB_headers_complete       HttpErrno = 5
	HPE_CB_body                   HttpErrno = 6
	HPE_CB_message_complete       HttpErrno = 7
	HPE_CB_status                 HttpErrno = 8
	HPE_CB_chunk_header           HttpErrno = 9
	HPE_CB_chunk_complete         HttpErrno = 10
	HPE_INVALID_EOF_STATE         HttpErrno = 11
	HPE_HEADER_OVERFLOW           HttpErrno = 12
	HPE_CLOSED_CONNECTION         HttpErrno = 13
	HPE_INVALID_VERSION           HttpErrno = 14
	HPE_INVALID_STATUS            HttpErrno = 15
	HPE_INVALID_METHOD            HttpErrno = 16
	HPE_INVALID_URL               HttpErrno = 17
	HPE_INVALID_HOST              HttpErrno = 18
	HPE_INVALID_PORT              HttpErrno = 19
	HPE_INVALID_PATH              HttpErrno = 20
	HPE_INVALID_QUERY_STRING      HttpErrno = 21
	HPE_INVALID_FRAGMENT          HttpErrno = 22
	HPE_LF_EXPECTED               HttpErrno = 23
	HPE_INVALID_HEADER_TOKEN      HttpErrno = 24
	HPE_INVALID_CONTENT_LENGTH    HttpErrno = 25
	HPE_UNEXPECTED_CONTENT_LENGTH HttpErrno = 26
	HPE_INVALID_CHUNK_SIZE        HttpErrno = 27
	HPE_INVALID_CONSTANT          HttpErrno = 28
	HPE_INVALID_INTERNAL_STATE    HttpErrno = 29
	HPE_STRICT                    HttpErrno = 30
	HPE_PAUSED                    HttpErrno = 31
	HPE_UNKNOWN                   HttpErrno = 32
)

type HttpParserUrlFields c.Int

const (
	UF_SCHEMA   HttpParserUrlFields = 0
	UF_HOST     HttpParserUrlFields = 1
	UF_PORT     HttpParserUrlFields = 2
	UF_PATH     HttpParserUrlFields = 3
	UF_QUERY    HttpParserUrlFields = 4
	UF_FRAGMENT HttpParserUrlFields = 5
	UF_USERINFO HttpParserUrlFields = 6
	UF_MAX      HttpParserUrlFields = 7
)

/* Result structure for http_parser_parse_url().
 *
 * Callers should index into field_data[] with UF_* values iff field_set
 * has the relevant (1 << UF_*) bit set. As a courtesy to clients (and
 * because we probably have padding left over), we convert any port to
 * a uint16_t.
 */

type HttpParserUrl struct {
	FieldSet  c.Uint16T
	Port      c.Uint16T
	FieldData [7]struct {
		Off c.Uint16T
		Len c.Uint16T
	}
}

/* Returns the library version. Bits 16-23 contain the major version number,
 * bits 8-15 the minor version number and bits 0-7 the patch level.
 * Usage example:
 *
 *   unsigned long version = http_parser_version();
 *   unsigned major = (version >> 16) & 255;
 *   unsigned minor = (version >> 8) & 255;
 *   unsigned patch = version & 255;
 *   printf("http_parser v%u.%u.%u\n", major, minor, patch);
 */
//go:linkname HttpParserVersion C.http_parser_version
func HttpParserVersion() c.Ulong

// llgo:link (*HttpParser).HttpParserInit C.http_parser_init
func (recv_ *HttpParser) HttpParserInit(type_ HttpParserType) {
}

/* Initialize http_parser_settings members to 0
 */
// llgo:link (*HttpParserSettings).HttpParserSettingsInit C.http_parser_settings_init
func (recv_ *HttpParserSettings) HttpParserSettingsInit() {
}

/* Executes the parser. Returns number of parsed bytes. Sets
 * `parser->http_errno` on error. */
// llgo:link (*HttpParser).HttpParserExecute C.http_parser_execute
func (recv_ *HttpParser) HttpParserExecute(settings *HttpParserSettings, data *c.Char, len c.SizeT) c.SizeT {
	return 0
}

/* If http_should_keep_alive() in the on_headers_complete or
 * on_message_complete callback returns 0, then this should be
 * the last message on the connection.
 * If you are the server, respond with the "Connection: close" header.
 * If you are the client, close the connection.
 */
// llgo:link (*HttpParser).HttpShouldKeepAlive C.http_should_keep_alive
func (recv_ *HttpParser) HttpShouldKeepAlive() c.Int {
	return 0
}

/* Returns a string version of the HTTP method. */
// llgo:link HttpMethod.HttpMethodStr C.http_method_str
func (recv_ HttpMethod) HttpMethodStr() *c.Char {
	return nil
}

/* Return a string name of the given error */
// llgo:link HttpErrno.HttpErrnoName C.http_errno_name
func (recv_ HttpErrno) HttpErrnoName() *c.Char {
	return nil
}

/* Return a string description of the given error */
// llgo:link HttpErrno.HttpErrnoDescription C.http_errno_description
func (recv_ HttpErrno) HttpErrnoDescription() *c.Char {
	return nil
}

/* Initialize all http_parser_url members to 0 */
// llgo:link (*HttpParserUrl).HttpParserUrlInit C.http_parser_url_init
func (recv_ *HttpParserUrl) HttpParserUrlInit() {
}

/* Parse a URL; return nonzero on failure */
//go:linkname HttpParserParseUrl C.http_parser_parse_url
func HttpParserParseUrl(buf *c.Char, buflen c.SizeT, is_connect c.Int, u *HttpParserUrl) c.Int

/* Pause or un-pause the parser; a nonzero value pauses */
// llgo:link (*HttpParser).HttpParserPause C.http_parser_pause
func (recv_ *HttpParser) HttpParserPause(paused c.Int) {
}

/* Checks if this is the final chunk of the body. */
// llgo:link (*HttpParser).HttpBodyIsFinal C.http_body_is_final
func (recv_ *HttpParser) HttpBodyIsFinal() c.Int {
	return 0
}
