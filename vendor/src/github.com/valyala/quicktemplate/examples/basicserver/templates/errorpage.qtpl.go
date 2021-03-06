// This file is automatically generated by qtc from "errorpage.qtpl".
// See https://github.com/valyala/quicktemplate for details.

// Error page template. Implements BasePage methods.
//

//line vendor\src\github.com\valyala\quicktemplate\examples\basicserver\templates\errorpage.qtpl:3
package templates

//line vendor\src\github.com\valyala\quicktemplate\examples\basicserver\templates\errorpage.qtpl:3
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line vendor\src\github.com\valyala\quicktemplate\examples\basicserver\templates\errorpage.qtpl:3
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line vendor\src\github.com\valyala\quicktemplate\examples\basicserver\templates\errorpage.qtpl:4
type ErrorPage struct {
	// inherit from base page, so its' title is used in error page.
	BasePage

	// error path
	Path []byte
}

//line vendor\src\github.com\valyala\quicktemplate\examples\basicserver\templates\errorpage.qtpl:14
func (p *ErrorPage) StreamBody(qw422016 *qt422016.Writer) {
	//line vendor\src\github.com\valyala\quicktemplate\examples\basicserver\templates\errorpage.qtpl:14
	qw422016.N().S(`
	<h1>Error page</h1>
	</div>
		Unsupported path <b>`)
	//line vendor\src\github.com\valyala\quicktemplate\examples\basicserver\templates\errorpage.qtpl:17
	qw422016.E().Z(p.Path)
	//line vendor\src\github.com\valyala\quicktemplate\examples\basicserver\templates\errorpage.qtpl:17
	qw422016.N().S(`</b>.
	</div>
	Base page body: `)
	//line vendor\src\github.com\valyala\quicktemplate\examples\basicserver\templates\errorpage.qtpl:19
	p.BasePage.StreamBody(qw422016)
	//line vendor\src\github.com\valyala\quicktemplate\examples\basicserver\templates\errorpage.qtpl:19
	qw422016.N().S(`
`)
//line vendor\src\github.com\valyala\quicktemplate\examples\basicserver\templates\errorpage.qtpl:20
}

//line vendor\src\github.com\valyala\quicktemplate\examples\basicserver\templates\errorpage.qtpl:20
func (p *ErrorPage) WriteBody(qq422016 qtio422016.Writer) {
	//line vendor\src\github.com\valyala\quicktemplate\examples\basicserver\templates\errorpage.qtpl:20
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line vendor\src\github.com\valyala\quicktemplate\examples\basicserver\templates\errorpage.qtpl:20
	p.StreamBody(qw422016)
	//line vendor\src\github.com\valyala\quicktemplate\examples\basicserver\templates\errorpage.qtpl:20
	qt422016.ReleaseWriter(qw422016)
//line vendor\src\github.com\valyala\quicktemplate\examples\basicserver\templates\errorpage.qtpl:20
}

//line vendor\src\github.com\valyala\quicktemplate\examples\basicserver\templates\errorpage.qtpl:20
func (p *ErrorPage) Body() string {
	//line vendor\src\github.com\valyala\quicktemplate\examples\basicserver\templates\errorpage.qtpl:20
	qb422016 := qt422016.AcquireByteBuffer()
	//line vendor\src\github.com\valyala\quicktemplate\examples\basicserver\templates\errorpage.qtpl:20
	p.WriteBody(qb422016)
	//line vendor\src\github.com\valyala\quicktemplate\examples\basicserver\templates\errorpage.qtpl:20
	qs422016 := string(qb422016.B)
	//line vendor\src\github.com\valyala\quicktemplate\examples\basicserver\templates\errorpage.qtpl:20
	qt422016.ReleaseByteBuffer(qb422016)
	//line vendor\src\github.com\valyala\quicktemplate\examples\basicserver\templates\errorpage.qtpl:20
	return qs422016
//line vendor\src\github.com\valyala\quicktemplate\examples\basicserver\templates\errorpage.qtpl:20
}
