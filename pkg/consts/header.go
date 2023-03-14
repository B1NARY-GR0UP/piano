// Copyright 2023 BINARY Members
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package consts

const (
	HeaderDate = "Date"

	HeaderIfModifiedSince = "If-Modified-Since" // Cache
	HeaderLastModified    = "Last-Modified"
	HeaderIfNoneMatch     = "If-None-Match"
	HeaderETag            = "ETag"

	HeaderLocation = "Location" // Redirects

	HeaderTE               = "TE" // Transfer coding
	HeaderTrailer          = "Trailer"
	HeaderTrailerLower     = "trailer"
	HeaderTransferEncoding = "Transfer-Encoding"

	HeaderCookie         = "Cookie" // Controls
	HeaderExpect         = "Expect"
	HeaderMaxForwards    = "Max-Forwards"
	HeaderSetCookie      = "Set-Cookie"
	HeaderSetCookieLower = "set-cookie"

	HeaderConnection      = "Connection" // Connection management
	HeaderKeepAlive       = "Keep-Alive"
	HeaderProxyConnection = "Proxy-Connection"

	HeaderAuthorization      = "Authorization" // Authentication
	HeaderProxyAuthenticate  = "Proxy-Authenticate"
	HeaderProxyAuthorization = "Proxy-Authorization"
	HeaderWWWAuthenticate    = "WWW-Authenticate"

	HeaderAcceptRanges = "Accept-Ranges" // Range requests
	HeaderContentRange = "Content-Range"
	HeaderIfRange      = "If-Range"
	HeaderRange        = "Range"

	HeaderAllow       = "Allow" // Response context
	HeaderServer      = "Server"
	HeaderServerLower = "server"

	HeaderFrom           = "From" // Request context
	HeaderHost           = "Host"
	HeaderReferer        = "Referer"
	HeaderReferrerPolicy = "Referrer-Policy"
	HeaderUserAgent      = "User-Agent"

	HeaderContentEncoding = "Content-Encoding" // Message body information
	HeaderContentLanguage = "Content-Language"
	HeaderContentLength   = "Content-Length"
	HeaderContentLocation = "Content-Location"
	HeaderContentType     = "Content-Type"

	HeaderAccept         = "Accept" // Content negotiation
	HeaderAcceptCharset  = "Accept-Charset"
	HeaderAcceptEncoding = "Accept-Encoding"
	HeaderAcceptLanguage = "Accept-Language"
	HeaderAltSvc         = "Alt-Svc"

	HTTP10 = "HTTP/1.0" // Protocol
	HTTP11 = "HTTP/1.1"
	HTTP20 = "HTTP/2.0"
	HTTP30 = "HTTP/3.0"
)
